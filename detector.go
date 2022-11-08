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
	"fmt"
	"golang.org/x/exp/slices"
	"math"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

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
	// language considered possible for the given input text.
	//
	// A slice of ConfidenceValue is returned containing all possible languages
	// sorted by their confidence value in descending order. The values that
	// this method computes are part of a relative confidence metric, not of an
	// absolute one. Each value is a number between 0.0 and 1.0. The most likely
	// language is always returned with value 1.0. All other languages get values
	// assigned which are lower than 1.0, denoting how less likely those languages
	// are in comparison to the most likely language.
	//
	// The slice returned by this method does not necessarily contain all
	// languages which the calling instance of LanguageDetector was built from.
	// If the rule-based engine decides that a specific language is truly
	// impossible, then it will not be part of the returned slice. Likewise,
	// if no ngram probabilities can be found within the detector's languages
	// for the given input text, the returned slice will be empty.
	// The confidence value for each language not being part of the returned
	// slice is assumed to be 0.0.
	ComputeLanguageConfidenceValues(text string) []ConfidenceValue
}

type languageDetector struct {
	languages                     []Language
	minimumRelativeDistance       float64
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
) LanguageDetector {
	detector := languageDetector{
		languages,
		minimumRelativeDistance,
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
			detector.loadLanguageModels(detector.unigramLanguageModels, language, 1)
			detector.loadLanguageModels(detector.bigramLanguageModels, language, 2)
			detector.loadLanguageModels(detector.trigramLanguageModels, language, 3)
			detector.loadLanguageModels(detector.quadrigramLanguageModels, language, 4)
			detector.loadLanguageModels(detector.fivegramLanguageModels, language, 5)
		}(language, &wg)
	}
	wg.Wait()
}

func (detector languageDetector) DetectLanguageOf(text string) (Language, bool) {
	confidenceValues := detector.ComputeLanguageConfidenceValues(text)

	if len(confidenceValues) == 0 {
		return Unknown, false
	}

	mostLikely := confidenceValues[0]

	if len(confidenceValues) == 1 {
		return mostLikely.Language(), true
	}

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
	var values []ConfidenceValue
	cleanedUpText := detector.cleanUpInputText(text)

	if len(cleanedUpText) == 0 || noLetter.MatchString(cleanedUpText) {
		return values
	}

	words := detector.splitTextIntoWords(cleanedUpText)
	languageDetectedByRules := detector.detectLanguageWithRules(words)

	if languageDetectedByRules != Unknown {
		values = append(values, newConfidenceValue(languageDetectedByRules, 1.0))
		return values
	}

	filteredLanguages := detector.filterLanguagesByRules(words)

	if len(filteredLanguages) == 1 {
		values = append(values, newConfidenceValue(filteredLanguages[0], 1.0))
		return values
	}

	characterCount := utf8.RuneCountInString(cleanedUpText)
	var ngramLengthRange []int

	if characterCount >= 120 {
		ngramLengthRange = []int{3}
	} else {
		ngramLengthRange = []int{1, 2, 3, 4, 5}
	}

	probabilityChannel := make(chan map[Language]float64, len(ngramLengthRange))
	unigramCountChannel := make(chan map[Language]uint32, 1)

	for _, ngramLength := range ngramLengthRange {
		go detector.lookUpLanguageModels(
			cleanedUpText,
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

	probabilityMaps := detector.getProbabilityMaps(probabilityChannel, ngramLengthRange)
	summedUpProbabilities := detector.sumUpProbabilities(probabilityMaps, unigramCounts, filteredLanguages)

	if len(summedUpProbabilities) == 0 {
		return values
	}

	highestProbability := detector.getHighestProbability(summedUpProbabilities)
	return detector.computeConfidenceValues(summedUpProbabilities, highestProbability)
}

func (detector languageDetector) getProbabilityMaps(
	probabilityChannel <-chan map[Language]float64,
	ngramLengthRange []int,
) []map[Language]float64 {
	var probabilityMaps []map[Language]float64
	for range ngramLengthRange {
		probabilityMaps = append(probabilityMaps, <-probabilityChannel)
	}
	return probabilityMaps
}

func (detector languageDetector) cleanUpInputText(text string) string {
	trimmed := strings.ToLower(strings.TrimSpace(text))
	withoutPunctuation := punctuation.ReplaceAllString(trimmed, "")
	withoutNumbers := numbers.ReplaceAllString(withoutPunctuation, "")
	normalizedWhitespace := multipleWhitespace.ReplaceAllString(withoutNumbers, " ")
	return normalizedWhitespace
}

func (detector languageDetector) splitTextIntoWords(text string) []string {
	var normalizedTextBuilder []string
	for _, chr := range text {
		char := string(chr)
		normalizedTextBuilder = append(normalizedTextBuilder, char)
		if detector.isLogogram(char) {
			normalizedTextBuilder = append(normalizedTextBuilder, " ")
		}
	}
	normalizedText := strings.Join(normalizedTextBuilder, "")
	if strings.Contains(normalizedText, " ") {
		substrings := strings.Split(normalizedText, " ")
		var filteredSubstrings []string
		for _, substring := range substrings {
			if len(substring) > 0 {
				filteredSubstrings = append(filteredSubstrings, substring)
			}
		}
		return filteredSubstrings
	}
	return []string{normalizedText}
}

func (detector languageDetector) isLogogram(char string) bool {
	if strings.TrimSpace(char) == "" {
		return false
	}
	for _, language := range languagesSupportingLogograms {
		for _, alphabet := range language.alphabets() {
			if alphabet.matches(char) {
				return true
			}
		}
	}
	return false
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
			if containsLanguage(detector.languages, language) {
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
				keys := make([]Language, 0, len(wordLanguageCounts))
				for key := range wordLanguageCounts {
					keys = append(keys, key)
				}
				sort.Slice(keys, func(i, j int) bool {
					return wordLanguageCounts[keys[i]] > wordLanguageCounts[keys[j]]
				})
				mostFrequentLanguage := keys[0]
				mostFrequentLanguageCount := wordLanguageCounts[keys[0]]
				secondMostFrequentLanguageCount := wordLanguageCounts[keys[1]]

				if mostFrequentLanguageCount > secondMostFrequentLanguageCount &&
					containsLanguage(detector.languages, mostFrequentLanguage) {
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
	sortedLanguages := make([]Language, 0, len(totalLanguageCounts))
	for language := range totalLanguageCounts {
		sortedLanguages = append(sortedLanguages, language)
	}
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
		distinctAlphabetCounts := make(map[uint32]bool)
		for _, count := range detectedAlphabets {
			distinctAlphabetCounts[count] = true
		}
		if len(distinctAlphabetCounts) == 1 {
			return detector.languages
		}
	}

	sortedAlphabets := make([]alphabet, 0, len(detectedAlphabets))
	for alphabet := range detectedAlphabets {
		sortedAlphabets = append(sortedAlphabets, alphabet)
	}
	sort.Slice(sortedAlphabets, func(i, j int) bool {
		return detectedAlphabets[sortedAlphabets[i]] > detectedAlphabets[sortedAlphabets[j]]
	})

	mostFrequentAlphabet := sortedAlphabets[0]
	var filteredLanguages []Language

	for _, language := range detector.languages {
		if containsAlphabet(language.alphabets(), mostFrequentAlphabet) {
			filteredLanguages = append(filteredLanguages, language)
		}
	}

	languageCounts := make(map[Language]uint32)

	for characters, languages := range charsToLanguagesMapping {
		var relevantLanguages []Language
		for _, language := range languages {
			if containsLanguage(filteredLanguages, language) {
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
	text string,
	ngramLength int,
	filteredLanguages []Language,
	probabilityChannel chan<- map[Language]float64,
	unigramCountChannel chan<- map[Language]uint32,
) {
	testDataModel := newTestDataLanguageModel(text, ngramLength)
	probabilities := detector.computeLanguageProbabilities(testDataModel, filteredLanguages)
	probabilityChannel <- probabilities

	if ngramLength == 1 {
		var languages []Language
		for language := range probabilities {
			languages = append(languages, language)
		}

		var intersectedLanguages []Language
		if len(languages) > 0 {
			for _, language := range filteredLanguages {
				if containsLanguage(languages, language) {
					intersectedLanguages = append(intersectedLanguages, language)
				}
			}
		} else {
			intersectedLanguages = filteredLanguages
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

func (detector languageDetector) getHighestProbability(probabilities map[Language]float64) float64 {
	keys := make([]Language, 0, len(probabilities))
	for key := range probabilities {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return probabilities[keys[i]] > probabilities[keys[j]]
	})
	return probabilities[keys[0]]
}

func (detector languageDetector) computeConfidenceValues(
	probabilities map[Language]float64,
	highestProbability float64,
) []ConfidenceValue {
	var confidenceValues []ConfidenceValue
	for language, probability := range probabilities {
		value := newConfidenceValue(language, highestProbability/probability)
		confidenceValues = append(confidenceValues, value)
	}
	sort.Slice(confidenceValues, func(i, j int) bool {
		first, second := confidenceValues[i], confidenceValues[j]
		if first.Value() == second.Value() {
			return first.Language() < second.Language()
		}
		return first.Value() > second.Value()
	})
	return confidenceValues
}

func (detector languageDetector) computeSumOfNgramProbabilities(language Language, ngrams map[ngram]bool) float64 {
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
		models = detector.loadLanguageModels(detector.fivegramLanguageModels, language, ngramLength)
	case 4:
		models = detector.loadLanguageModels(detector.quadrigramLanguageModels, language, ngramLength)
	case 3:
		models = detector.loadLanguageModels(detector.trigramLanguageModels, language, ngramLength)
	case 2:
		models = detector.loadLanguageModels(detector.bigramLanguageModels, language, ngramLength)
	case 1:
		models = detector.loadLanguageModels(detector.unigramLanguageModels, language, ngramLength)
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

func (detector languageDetector) sumUpProbabilities(
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

func (detector languageDetector) loadLanguageModels(
	languageModels *sync.Map,
	language Language,
	ngramLength int,
) map[string]float64 {
	existingModels, exists := languageModels.Load(language)
	if exists {
		return existingModels.(map[string]float64)
	}
	jsonData := loadJson(language, ngramLength)
	loadedModels := newTrainingDataLanguageModelFromJson(jsonData)
	languageModels.Store(language, loadedModels)
	return loadedModels
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
		if containsLanguage(languages, language) {
			oneLanguageAlphabets[alphabet] = language
		}
	}
	return oneLanguageAlphabets
}
