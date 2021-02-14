/*
 * Copyright © 2021 Peter M. Stahl pemistahl@gmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either expressed or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lingua

import (
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strings"
)

type jsonLanguageModel struct {
	Language Language          `json:"language"`
	Ngrams   map[string]string `json:"ngrams"`
}

type trainingDataLanguageModel struct {
	language                Language
	absoluteFrequencies     map[ngram]uint32
	relativeFrequencies     map[ngram]*big.Rat
	jsonRelativeFrequencies map[ngram]float64
}

type testDataLanguageModel struct {
	ngrams map[ngram]bool
}

func newTrainingDataLanguageModel(
	text []string,
	language Language,
	ngramLength int,
	charClass string,
	lowerNgramAbsoluteFrequencies map[ngram]uint32,
) trainingDataLanguageModel {
	absoluteFrequencies := computeAbsoluteFrequencies(text, ngramLength, charClass)
	relativeFrequencies := computeRelativeFrequencies(ngramLength, absoluteFrequencies, lowerNgramAbsoluteFrequencies)

	return trainingDataLanguageModel{
		language:                language,
		absoluteFrequencies:     absoluteFrequencies,
		relativeFrequencies:     relativeFrequencies,
		jsonRelativeFrequencies: nil,
	}
}

func newTrainingDataLanguageModelFromJson(jsonData []byte) trainingDataLanguageModel {
	var jsonModel jsonLanguageModel
	err := json.Unmarshal(jsonData, &jsonModel)
	if err != nil {
		panic(err.Error())
	}
	jsonRelativeFrequencies := make(map[ngram]float64)
	for rat, ngrams := range jsonModel.Ngrams {
		r := new(big.Rat)
		r.SetString(rat)
		f, _ := r.Float64()
		for _, ngram := range strings.Split(ngrams, " ") {
			jsonRelativeFrequencies[newNgram(ngram)] = f
		}
	}
	return trainingDataLanguageModel{
		language:                jsonModel.Language,
		absoluteFrequencies:     nil,
		relativeFrequencies:     nil,
		jsonRelativeFrequencies: jsonRelativeFrequencies,
	}
}

func (model trainingDataLanguageModel) toJson() []byte {
	ratsToNgrams := make(map[string]ngramSlice)
	for ngram, rat := range model.relativeFrequencies {
		r := rat.String()
		ratsToNgrams[r] = append(ratsToNgrams[r], ngram)
	}
	ratsToJoinedNgrams := make(map[string]string)
	for rat, ngrams := range ratsToNgrams {
		sort.Sort(ngrams)
		var ngramValues []string
		for _, ngram := range ngrams {
			ngramValues = append(ngramValues, ngram.value)
		}
		ratsToJoinedNgrams[rat] = strings.Join(ngramValues, " ")
	}
	jsonModel := jsonLanguageModel{
		Language: model.language,
		Ngrams:   ratsToJoinedNgrams,
	}
	serializedJsonModel, err := json.Marshal(jsonModel)
	if err != nil {
		panic(err.Error())
	}
	return serializedJsonModel
}

func newTestDataLanguageModel(text string, ngramLength int) testDataLanguageModel {
	if ngramLength > maxNgramLength {
		panic(fmt.Sprintf("ngram length %v is greater than %v", ngramLength, maxNgramLength))
	}
	ngrams := make(map[ngram]bool)
	chars := []rune(text)
	charsCount := len(chars)
	if charsCount >= ngramLength {
		for i := 0; i <= charsCount-ngramLength; i++ {
			slice := string(chars[i : i+ngramLength])
			if letter.MatchString(slice) {
				ngrams[newNgram(slice)] = true
			}
		}
	}
	return testDataLanguageModel{ngrams: ngrams}
}

func computeAbsoluteFrequencies(
	text []string,
	ngramLength int,
	charClass string,
) map[ngram]uint32 {
	absoluteFrequencies := make(map[ngram]uint32)
	regex, err := regexp.Compile(fmt.Sprintf("^[%v]+$", charClass))
	if err != nil {
		panic(fmt.Sprintf("The character class '%v' cannot be compiled to a valid regular expression", charClass))
	}
	for _, line := range text {
		chars := []rune(strings.ToLower(line))
		for i := 0; i <= len(chars)-ngramLength; i++ {
			slice := string(chars[i : i+ngramLength])
			if regex.MatchString(slice) {
				absoluteFrequencies[newNgram(slice)]++
			}
		}
	}
	return absoluteFrequencies
}

func computeRelativeFrequencies(
	ngramLength int,
	absoluteFrequencies map[ngram]uint32,
	lowerNgramAbsoluteFrequencies map[ngram]uint32,
) map[ngram]*big.Rat {
	ngramProbabilities := make(map[ngram]*big.Rat)
	var totalNgramFrequency uint32
	for _, frequency := range absoluteFrequencies {
		totalNgramFrequency += frequency
	}
	for ngram, frequency := range absoluteFrequencies {
		var denominator uint32
		if ngramLength == 1 || len(lowerNgramAbsoluteFrequencies) == 0 {
			denominator = totalNgramFrequency
		} else {
			chars := []rune(ngram.value)
			slice := string(chars[0 : ngramLength-1])
			denominator = lowerNgramAbsoluteFrequencies[newNgram(slice)]
		}
		ngramProbabilities[ngram] = big.NewRat(int64(frequency), int64(denominator))
	}
	return ngramProbabilities
}
