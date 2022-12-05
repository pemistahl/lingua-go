/*
 * Copyright © 2021-today Peter M. Stahl pemistahl@gmail.com
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
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
	"io"
	"math"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

//go:embed language-models
var languageModels embed.FS

var unigramModels sync.Map
var bigramModels sync.Map
var trigramModels sync.Map
var quadrigramModels sync.Map
var fivegramModels sync.Map

// LanguageDetector is the interface describing the available methods
// for detecting the language of some textual input.
type LanguageDetector interface {
	// DetectLanguageOf detects the language of the given text.
	// The boolean return value indicates whether a language can be reliably
	// detected. If this is not possible, (Unknown, false) is returned.
	DetectLanguageOf(text string) (Language, bool)

	// ComputeLanguageConfidenceValues computes confidence values for each
	// language supported by this detector for the given input text. These
	// values denote how likely it is that the given text has been written
	// in any of the languages supported by this detector.
	//
	// A slice of ConfidenceValue is returned containing those languages which
	// the calling instance of LanguageDetector has been built from. The entries
	// are sorted by their confidence value in descending order. The values that
	// this method computes are part of a relative confidence metric, not of an
	// absolute one. Each value is a number between 0.0 and 1.0.
	//
	// If the language is unambiguously identified by the rule engine, the value
	// 1.0 will always be returned for this language. The other languages will
	// receive a value of 0.0. If the statistics engine is additionally needed,
	// the most likely language will be returned with value 0.99 and the least
	// likely language will be returned with value 0.01. All other languages get
	// values assigned between 0.01 and 0.99, denoting how less likely those
	// languages are in comparison to the most likely language.
	ComputeLanguageConfidenceValues(text string) []ConfidenceValue

	// ComputeLanguageConfidence computes the confidence value for the given
	// language and input text. This value denotes how likely it is that the
	// given text has been written in the given language.
	//
	// The value that this method computes is a number between 0.0 and 1.0.
	// If the language is unambiguously identified by the rule engine, the value
	// 1.0 will always be returned. If the given language is not supported by
	// this detector instance, the value 0.0 will always be returned. Otherwise,
	// a value between 0.01 and 0.99 will be returned.
	ComputeLanguageConfidence(text string, language Language) float64
}

type languageDetector struct {
	languages                     []Language
	minimumRelativeDistance       float64
	isLowAccuracyModeEnabled      bool
	languagesWithUniqueCharacters []Language
	oneLanguageAlphabets          map[alphabet]Language
	unigramLanguageModels         *sync.Map
	bigramLanguageModels          *sync.Map
	trigramLanguageModels         *sync.Map
	quadrigramLanguageModels      *sync.Map
	fivegramLanguageModels        *sync.Map
}

func newLanguageDetector(
	languages []Language,
	minimumRelativeDistance float64,
	isEveryLanguageModelPreloaded bool,
	isLowAccuracyModeEnabled bool,
) LanguageDetector {
	detector := languageDetector{
		languages,
		minimumRelativeDistance,
		isLowAccuracyModeEnabled,
		collectLanguagesWithUniqueCharacters(languages),
		collectOneLanguageAlphabets(languages),
		&unigramModels,
		&bigramModels,
		&trigramModels,
		&quadrigramModels,
		&fivegramModels,
	}
	if isEveryLanguageModelPreloaded {
		detector.preloadLanguageModels(languages)
	}
	return detector
}

func (detector languageDetector) preloadLanguageModels(languages []Language) {
	var wg sync.WaitGroup
	for _, language := range languages {
		wg.Add(1)
		go func(language Language, wg *sync.WaitGroup) {
			defer wg.Done()
			loadLanguageModels(detector.trigramLanguageModels, language, 3)

			if !detector.isLowAccuracyModeEnabled {
				loadLanguageModels(detector.unigramLanguageModels, language, 1)
				loadLanguageModels(detector.bigramLanguageModels, language, 2)
				loadLanguageModels(detector.quadrigramLanguageModels, language, 4)
				loadLanguageModels(detector.fivegramLanguageModels, language, 5)
			}
		}(language, &wg)
	}
	wg.Wait()
}

func (detector languageDetector) DetectLanguageOf(text string) (Language, bool) {
	confidenceValues := detector.ComputeLanguageConfidenceValues(text)
	mostLikely := confidenceValues[0]
	secondMostLikely := confidenceValues[1]

	if mostLikely.Value() == secondMostLikely.Value() {
		return Unknown, false
	}
	if (mostLikely.Value() - secondMostLikely.Value()) < detector.minimumRelativeDistance {
		return Unknown, false
	}

	return mostLikely.Language(), true
}

func (detector languageDetector) ComputeLanguageConfidenceValues(text string) []ConfidenceValue {
	values := make(confidenceValueSlice, len(detector.languages))
	for i, language := range detector.languages {
		values[i] = newConfidenceValue(language, 0)
	}

	words := splitTextIntoWords(text)

	if len(words) == 0 {
		sort.Sort(values)
		return values
	}

	languageDetectedByRules := detector.detectLanguageWithRules(words)

	if languageDetectedByRules != Unknown {
		for i := range values {
			if values[i].Language() == languageDetectedByRules {
				values[i] = newConfidenceValue(languageDetectedByRules, 1)
				break
			}
		}
		sort.Sort(values)
		return values
	}

	filteredLanguages := detector.filterLanguagesByRules(words)

	if len(filteredLanguages) == 1 {
		languageDetectedByFilter := filteredLanguages[0]
		for i := range values {
			if values[i].Language() == languageDetectedByFilter {
				values[i] = newConfidenceValue(languageDetectedByFilter, 1)
				break
			}
		}
		sort.Sort(values)
		return values
	}

	characterCount := 0
	for _, word := range words {
		characterCount += utf8.RuneCountInString(word)
	}

	if detector.isLowAccuracyModeEnabled && characterCount < 3 {
		sort.Sort(values)
		return values
	}

	var ngramLengthRange []int

	if characterCount >= 120 || detector.isLowAccuracyModeEnabled {
		ngramLengthRange = []int{3}
	} else {
		ngramLengthRange = []int{1, 2, 3, 4, 5}
	}

	probabilityChannel := make(chan map[Language]float64, len(ngramLengthRange))
	unigramCountChannel := make(chan map[Language]uint32, 1)

	for _, ngramLength := range ngramLengthRange {
		go detector.lookUpLanguageModels(
			words,
			ngramLength,
			filteredLanguages,
			probabilityChannel,
			unigramCountChannel,
		)
	}

	var unigramCounts map[Language]uint32
	if slices.Contains(ngramLengthRange, 1) {
		unigramCounts = <-unigramCountChannel
	}

	probabilityMaps := getProbabilityMaps(probabilityChannel, ngramLengthRange)
	summedUpProbabilities := sumUpProbabilities(probabilityMaps, unigramCounts, filteredLanguages)

	if len(summedUpProbabilities) == 0 {
		sort.Sort(values)
		return values
	}

	highestProbability, lowestProbability := getHighestAndLowestProbability(summedUpProbabilities)
	return detector.computeConfidenceValues(values, summedUpProbabilities, highestProbability, lowestProbability)
}

func (detector languageDetector) ComputeLanguageConfidence(text string, language Language) float64 {
	confidenceValues := detector.ComputeLanguageConfidenceValues(text)
	for _, confidenceValue := range confidenceValues {
		if confidenceValue.Language() == language {
			return confidenceValue.Value()
		}
	}
	return 0
}

func getProbabilityMaps(
	probabilityChannel <-chan map[Language]float64,
	ngramLengthRange []int,
) []map[Language]float64 {
	probabilityMaps := make([]map[Language]float64, len(ngramLengthRange))
	for i := range ngramLengthRange {
		probabilityMaps[i] = <-probabilityChannel
	}
	return probabilityMaps
}

func splitTextIntoWords(text string) []string {
	return letters.FindAllString(strings.ToLower(text), -1)
}

func (detector languageDetector) detectLanguageWithRules(words []string) Language {
	totalLanguageCounts := make(map[Language]uint32)
	halfWordCount := float64(len(words)) * 0.5

	for _, word := range words {
		wordLanguageCounts := make(map[Language]uint32)

		for _, chr := range []rune(word) {
			char := string(chr)
			isMatch := false

			for alphabet, language := range detector.oneLanguageAlphabets {
				if alphabet.matches(char) {
					wordLanguageCounts[language]++
					isMatch = true
					break
				}
			}

			if !isMatch {
				if han.matches(char) {
					wordLanguageCounts[Chinese]++
				} else if japaneseCharacterSet.MatchString(char) {
					wordLanguageCounts[Japanese]++
				} else if latin.matches(char) || cyrillic.matches(char) || devanagari.matches(char) {
					for _, language := range detector.languagesWithUniqueCharacters {
						if strings.Contains(language.uniqueCharacters(), char) {
							wordLanguageCounts[language]++
						}
					}
				}
			}
		}

		if len(wordLanguageCounts) == 0 {
			totalLanguageCounts[Unknown]++
		} else if len(wordLanguageCounts) == 1 {
			var language Language
			for key := range wordLanguageCounts {
				language = key
			}
			if slices.Contains(detector.languages, language) {
				totalLanguageCounts[language]++
			} else {
				totalLanguageCounts[Unknown]++
			}
		} else {
			_, containsChinese := wordLanguageCounts[Chinese]
			_, containsJapanese := wordLanguageCounts[Japanese]
			if containsChinese && containsJapanese {
				totalLanguageCounts[Japanese]++
			} else {
				keys := maps.Keys(wordLanguageCounts)
				sort.Slice(keys, func(i, j int) bool {
					return wordLanguageCounts[keys[i]] > wordLanguageCounts[keys[j]]
				})
				mostFrequentLanguage := keys[0]
				mostFrequentLanguageCount := wordLanguageCounts[keys[0]]
				secondMostFrequentLanguageCount := wordLanguageCounts[keys[1]]

				if mostFrequentLanguageCount > secondMostFrequentLanguageCount &&
					slices.Contains(detector.languages, mostFrequentLanguage) {
					totalLanguageCounts[mostFrequentLanguage]++
				} else {
					totalLanguageCounts[Unknown]++
				}
			}
		}
	}
	var unknownLanguageCount float64 = 0
	if value, exists := totalLanguageCounts[Unknown]; exists {
		unknownLanguageCount = float64(value)
	}
	if unknownLanguageCount < halfWordCount {
		delete(totalLanguageCounts, Unknown)
	}
	if len(totalLanguageCounts) == 0 {
		return Unknown
	}
	if len(totalLanguageCounts) == 1 {
		for language := range totalLanguageCounts {
			return language
		}
	}
	if len(totalLanguageCounts) == 2 {
		_, containsChinese := totalLanguageCounts[Chinese]
		_, containsJapanese := totalLanguageCounts[Japanese]
		if containsChinese && containsJapanese {
			return Japanese
		}
	}
	sortedLanguages := maps.Keys(totalLanguageCounts)
	sort.Slice(sortedLanguages, func(i, j int) bool {
		return totalLanguageCounts[sortedLanguages[i]] > totalLanguageCounts[sortedLanguages[j]]
	})
	mostFrequentLanguage := sortedLanguages[0]
	mostFrequentLanguageCount := totalLanguageCounts[sortedLanguages[0]]
	secondMostFrequentLanguageCount := totalLanguageCounts[sortedLanguages[1]]

	if mostFrequentLanguageCount == secondMostFrequentLanguageCount {
		return Unknown
	}

	return mostFrequentLanguage
}

func (detector languageDetector) filterLanguagesByRules(words []string) []Language {
	detectedAlphabets := make(map[alphabet]uint32)
	halfWordCount := float64(len(words)) * 0.5

	for _, word := range words {
		for _, alphabet := range allAlphabets() {
			if alphabet.matches(word) {
				detectedAlphabets[alphabet]++
				break
			}
		}
	}

	if len(detectedAlphabets) == 0 {
		return detector.languages
	}

	if len(detectedAlphabets) > 1 {
		distinctAlphabetCounts := make(map[uint32]struct{})
		for _, count := range detectedAlphabets {
			distinctAlphabetCounts[count] = struct{}{}
		}
		if len(distinctAlphabetCounts) == 1 {
			return detector.languages
		}
	}

	sortedAlphabets := maps.Keys(detectedAlphabets)
	sort.Slice(sortedAlphabets, func(i, j int) bool {
		return detectedAlphabets[sortedAlphabets[i]] > detectedAlphabets[sortedAlphabets[j]]
	})

	mostFrequentAlphabet := sortedAlphabets[0]
	var filteredLanguages []Language

	for _, language := range detector.languages {
		if slices.Contains(language.alphabets(), mostFrequentAlphabet) {
			filteredLanguages = append(filteredLanguages, language)
		}
	}

	languageCounts := make(map[Language]uint32)

	for characters, languages := range charsToLanguagesMapping {
		var relevantLanguages []Language
		for _, language := range languages {
			if slices.Contains(filteredLanguages, language) {
				relevantLanguages = append(relevantLanguages, language)
			}
		}
		for _, word := range words {
			for _, character := range []rune(characters) {
				if strings.ContainsRune(word, character) {
					for _, language := range relevantLanguages {
						languageCounts[language]++
					}
				}
			}
		}
	}

	var languageSubset []Language

	for language, count := range languageCounts {
		if float64(count) >= halfWordCount {
			languageSubset = append(languageSubset, language)
		}
	}

	if len(languageSubset) > 0 {
		return languageSubset
	}

	return filteredLanguages
}

func (detector languageDetector) lookUpLanguageModels(
	words []string,
	ngramLength int,
	filteredLanguages []Language,
	probabilityChannel chan<- map[Language]float64,
	unigramCountChannel chan<- map[Language]uint32,
) {
	testDataModel := newTestDataLanguageModel(words, ngramLength)
	probabilities := detector.computeLanguageProbabilities(testDataModel, filteredLanguages)
	probabilityChannel <- probabilities

	if ngramLength == 1 {
		intersectedLanguages := make([]Language, len(filteredLanguages))

		if len(probabilities) > 0 {
			for i, language := range filteredLanguages {
				if _, exists := probabilities[language]; exists {
					intersectedLanguages[i] = language
				}
			}
		} else {
			copy(intersectedLanguages, filteredLanguages)
		}

		detector.countUnigrams(unigramCountChannel, testDataModel, intersectedLanguages)
	}
}

func (detector languageDetector) computeLanguageProbabilities(
	model testDataLanguageModel,
	filteredLanguages []Language,
) map[Language]float64 {
	probabilities := make(map[Language]float64)
	for _, language := range filteredLanguages {
		sum := detector.computeSumOfNgramProbabilities(language, model.ngrams)
		if sum < 0 {
			probabilities[language] = sum
		}
	}
	return probabilities
}

func getHighestAndLowestProbability(probabilities map[Language]float64) (float64, float64) {
	probabilityCount := len(probabilities)
	keys := maps.Keys(probabilities)
	sort.Slice(keys, func(i, j int) bool {
		return probabilities[keys[i]] > probabilities[keys[j]]
	})
	return probabilities[keys[0]], probabilities[keys[probabilityCount-1]]
}

func (detector languageDetector) computeConfidenceValues(
	confidenceValues confidenceValueSlice,
	probabilities map[Language]float64,
	highestProbability float64,
	lowestProbability float64,
) []ConfidenceValue {
	denominator := highestProbability - lowestProbability
	for language, probability := range probabilities {
		// Apply min-max normalization
		normalizedProbability := 0.98*(probability-lowestProbability)/denominator + 0.01
		for i := range confidenceValues {
			if confidenceValues[i].Language() == language {
				confidenceValues[i] = newConfidenceValue(language, normalizedProbability)
				break
			}
		}
	}
	sort.Sort(confidenceValues)
	return confidenceValues
}

func (detector languageDetector) computeSumOfNgramProbabilities(language Language, ngrams map[ngram]struct{}) float64 {
	sum := 0.0
	for ngram := range ngrams {
		for _, elem := range ngram.rangeOfLowerOrderNgrams() {
			probability := detector.lookUpNgramProbability(language, elem)
			if probability > 0 {
				sum += math.Log(probability)
				break
			}
		}
	}
	return sum
}

func (detector languageDetector) lookUpNgramProbability(language Language, ngrm ngram) float64 {
	ngramLength := utf8.RuneCountInString(ngrm.value)
	var models map[string]float64

	switch ngramLength {
	case 5:
		models = loadLanguageModels(detector.fivegramLanguageModels, language, ngramLength)
	case 4:
		models = loadLanguageModels(detector.quadrigramLanguageModels, language, ngramLength)
	case 3:
		models = loadLanguageModels(detector.trigramLanguageModels, language, ngramLength)
	case 2:
		models = loadLanguageModels(detector.bigramLanguageModels, language, ngramLength)
	case 1:
		models = loadLanguageModels(detector.unigramLanguageModels, language, ngramLength)
	case 0:
		panic("zerogram detected")
	default:
		panic(fmt.Sprintf("unsupported ngram length detected: %v", ngramLength))
	}

	if frequency, exists := models[ngrm.value]; exists {
		return frequency
	}
	return 0
}

func (detector languageDetector) countUnigrams(
	unigramCountChannel chan<- map[Language]uint32,
	unigramModel testDataLanguageModel,
	filteredLanguages []Language,
) {
	unigramCounts := make(map[Language]uint32)
	for _, language := range filteredLanguages {
		for unigram := range unigramModel.ngrams {
			if detector.lookUpNgramProbability(language, unigram) > 0 {
				unigramCounts[language]++
			}
		}
	}
	unigramCountChannel <- unigramCounts
}

func sumUpProbabilities(
	probabilityMaps []map[Language]float64,
	unigramCounts map[Language]uint32,
	filteredLanguages []Language,
) map[Language]float64 {
	summedUpProbabilities := make(map[Language]float64)
	hasUnigramCounts := unigramCounts != nil
	for _, language := range filteredLanguages {
		sum := 0.0
		for _, probabilities := range probabilityMaps {
			if probability, exists := probabilities[language]; exists {
				sum += probability
			}
		}
		if hasUnigramCounts {
			if unigramCount, exists := unigramCounts[language]; exists {
				sum /= float64(unigramCount)
			}
		}
		if sum != 0 {
			summedUpProbabilities[language] = sum
		}
	}
	return summedUpProbabilities
}

func loadLanguageModels(
	languageModels *sync.Map,
	language Language,
	ngramLength int,
) map[string]float64 {
	existingModels, exists := languageModels.Load(language)
	if exists {
		return existingModels.(map[string]float64)
	}

	protobufData := loadProtobufData(language, ngramLength)
	if protobufData == nil {
		return nil
	}

	model := SerializableLanguageModel{}
	if err := proto.Unmarshal(protobufData, &model); err != nil {
		panic(err.Error())
	}

	modelMap := make(map[string]float64, model.TotalNgrams)
	for _, ngramSet := range model.NgramSets {
		for _, ngrm := range ngramSet.Ngrams {
			modelMap[ngrm] = ngramSet.Probability
		}
	}

	languageModels.Store(language, modelMap)
	return modelMap
}

func loadProtobufData(language Language, ngramLength int) []byte {
	ngramName := getNgramNameByLength(ngramLength)
	isoCode := strings.ToLower(language.IsoCode639_1().String())
	zipFilePath := fmt.Sprintf("language-models/%s/%ss.pb.bin.zip", isoCode, ngramName)
	zipFileBytes, err := languageModels.ReadFile(zipFilePath)
	if err != nil {
		return nil
	}
	zipFile, _ := zip.NewReader(bytes.NewReader(zipFileBytes), int64(len(zipFileBytes)))
	protobufFile, _ := zipFile.File[0].Open()
	defer protobufFile.Close()
	protobufFileContent, _ := io.ReadAll(protobufFile)
	return protobufFileContent
}

func collectLanguagesWithUniqueCharacters(languages []Language) []Language {
	var languagesWithUniqueCharacters []Language
	for _, language := range languages {
		if len(language.uniqueCharacters()) > 0 {
			languagesWithUniqueCharacters = append(languagesWithUniqueCharacters, language)
		}
	}
	return languagesWithUniqueCharacters
}

func collectOneLanguageAlphabets(languages []Language) map[alphabet]Language {
	oneLanguageAlphabets := make(map[alphabet]Language)
	for alphabet, language := range allAlphabetsSupportingSingleLanguage() {
		if slices.Contains(languages, language) {
			oneLanguageAlphabets[alphabet] = language
		}
	}
	return oneLanguageAlphabets
}
