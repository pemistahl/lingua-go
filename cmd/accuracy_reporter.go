// ### +build ignore

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

package main

import (
	"fmt"
	"github.com/pemistahl/lingua-go"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

type detectorStatistics struct {
	singleWordStatistic statistic
	wordPairStatistic   statistic
	sentenceStatistic   statistic
	averageAccuracies   map[lingua.Language]float64
}

type statistic struct {
	languageCounts     map[lingua.Language]int
	languageAccuracies map[lingua.Language]float64
	entityCount        int
	entityLengthCount  int
}

func newDetectorStatistics() detectorStatistics {
	return detectorStatistics{
		singleWordStatistic: newStatistic(),
		wordPairStatistic:   newStatistic(),
		sentenceStatistic:   newStatistic(),
		averageAccuracies:   make(map[lingua.Language]float64),
	}
}

func newStatistic() statistic {
	return statistic{
		languageCounts:     make(map[lingua.Language]int),
		languageAccuracies: make(map[lingua.Language]float64),
		entityCount:        0,
		entityLengthCount:  0,
	}
}

func (ds *detectorStatistics) addSingleWordCounts(language lingua.Language, singleWord string) {
	ds.singleWordStatistic.addLanguageCount(language)
	ds.singleWordStatistic.addEntityCount()
	ds.singleWordStatistic.addEntityLengthCount(singleWord)
}

func (ds *detectorStatistics) addWordPairCounts(language lingua.Language, wordPair string) {
	ds.wordPairStatistic.addLanguageCount(language)
	ds.wordPairStatistic.addEntityCount()
	ds.wordPairStatistic.addEntityLengthCount(wordPair)
}

func (ds *detectorStatistics) addSentenceCounts(language lingua.Language, sentence string) {
	ds.sentenceStatistic.addLanguageCount(language)
	ds.sentenceStatistic.addEntityCount()
	ds.sentenceStatistic.addEntityLengthCount(sentence)
}

func (ds *detectorStatistics) computeAccuracyValues() {
	ds.singleWordStatistic.mapCountsToAccuracyValues()
	ds.wordPairStatistic.mapCountsToAccuracyValues()
	ds.sentenceStatistic.mapCountsToAccuracyValues()
}

func (ds *detectorStatistics) createReportData(language lingua.Language) string {
	singleWordAccuracy, singleWordReport := ds.singleWordStatistic.createReportData(language, "single words")
	wordPairAccuracy, wordPairReport := ds.wordPairStatistic.createReportData(language, "word pairs")
	sentenceAccuracy, sentenceReport := ds.sentenceStatistic.createReportData(language, "sentences")
	averageAccuracy := (singleWordAccuracy + wordPairAccuracy + sentenceAccuracy) / 3
	ds.averageAccuracies[language] = averageAccuracy

	if averageAccuracy == 0 {
		return ""
	}

	return fmt.Sprintf(
		"##### %s #####\n\n>>> Accuracy on average: %.2f%%\n\n%s\n%s\n%s\n",
		language,
		averageAccuracy*100,
		singleWordReport,
		wordPairReport,
		sentenceReport,
	)
}

func (ds *detectorStatistics) createAggregatedReportRow(language lingua.Language) string {
	var averageAccuracyColumn string
	accuracy, exists := ds.averageAccuracies[language]
	if exists && accuracy > 0 {
		averageAccuracyColumn = fmt.Sprintf("%.0f", accuracy*100)
	} else {
		averageAccuracyColumn = "NaN"
	}

	var singleWordsAccuracyColumn string
	accuracy, exists = ds.singleWordStatistic.languageAccuracies[language]
	if exists && accuracy > 0 {
		singleWordsAccuracyColumn = fmt.Sprintf("%.0f", accuracy*100)
	} else {
		singleWordsAccuracyColumn = "NaN"
	}

	var wordPairsAccuracyColumn string
	accuracy, exists = ds.wordPairStatistic.languageAccuracies[language]
	if exists && accuracy > 0 {
		wordPairsAccuracyColumn = fmt.Sprintf("%.0f", accuracy*100)
	} else {
		wordPairsAccuracyColumn = "NaN"
	}

	var sentencesAccuracyColumn string
	accuracy, exists = ds.sentenceStatistic.languageAccuracies[language]
	if exists && accuracy > 0 {
		sentencesAccuracyColumn = fmt.Sprintf("%.0f", accuracy*100)
	} else {
		sentencesAccuracyColumn = "NaN"
	}

	return fmt.Sprintf(
		"%s,%s,%s,%s",
		averageAccuracyColumn,
		singleWordsAccuracyColumn,
		wordPairsAccuracyColumn,
		sentencesAccuracyColumn,
	)
}

func (s *statistic) addLanguageCount(language lingua.Language) {
	if _, exists := s.languageCounts[language]; !exists {
		s.languageCounts[language] = 0
	}
	s.languageCounts[language] += 1
}

func (s *statistic) addEntityCount() {
	s.entityCount += 1
}

func (s *statistic) addEntityLengthCount(entity string) {
	s.entityLengthCount += utf8.RuneCountInString(entity)
}

func (s *statistic) mapCountsToAccuracyValues() {
	sumOfCounts := 0
	for _, count := range s.languageCounts {
		sumOfCounts += count
	}
	for language, count := range s.languageCounts {
		s.languageAccuracies[language] = float64(count) / float64(sumOfCounts)
	}
}

func (s *statistic) createReportData(
	language lingua.Language,
	description string,
) (float64, string) {
	accuracy, exists := s.languageAccuracies[language]
	if !exists {
		accuracy = 0.0
	}
	averageLength := int(math.Round(float64(s.entityLengthCount) / float64(s.entityCount)))
	report := fmt.Sprintf(
		">> Detection of %d %s (average length: %d chars)\nAccuracy: %.2f%%\nErroneously classified as %s\n",
		s.entityCount,
		description,
		averageLength,
		accuracy*100,
		s.formatLanguageAccuracies(language),
	)
	return accuracy, report
}

func (s *statistic) formatLanguageAccuracies(language lingua.Language) string {
	var languages []lingua.Language
	for currentLanguage := range s.languageAccuracies {
		if currentLanguage != language {
			languages = append(languages, currentLanguage)
		}
	}
	sort.Slice(languages, func(i, j int) bool {
		firstLanguage, secondLanguage := languages[i], languages[j]
		firstAccuracy, secondAccuracy := s.languageAccuracies[firstLanguage], s.languageAccuracies[secondLanguage]
		if firstAccuracy == secondAccuracy {
			return firstLanguage < secondLanguage
		}
		return firstAccuracy > secondAccuracy
	})
	var reports []string
	for _, currentLanguage := range languages {
		report := fmt.Sprintf("%s: %.2f%%", currentLanguage, s.languageAccuracies[currentLanguage]*100)
		reports = append(reports, report)
	}
	return strings.Join(reports, ", ")
}

func main() {
	start := time.Now()

	linguaDetector := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()

	workingDirectory, _ := os.Getwd()
	testDataDirectory := filepath.Join(workingDirectory, "language-testdata")
	accuracyReportsDirectory := filepath.Join(workingDirectory, "accuracy-reports")
	linguaReportsDirectory := filepath.Join(accuracyReportsDirectory, "lingua")

	err := os.MkdirAll(linguaReportsDirectory, os.ModePerm)
	if err != nil {
		panic("Lingua reports directory could not be created")
	}

	aggregatedReportFilePath := filepath.Join(accuracyReportsDirectory, "aggregated-accuracy-values.csv")
	aggregatedReportFile, err := os.Create(aggregatedReportFilePath)
	if err != nil {
		panic("CSV file could not be created")
	}
	defer aggregatedReportFile.Close()

	aggregatedReportColumns := []string{
		"language",
		"average-lingua",
		"single-words-lingua",
		"word-pairs-lingua",
		"sentences-lingua\n",
	}

	_, err = aggregatedReportFile.WriteString(strings.Join(aggregatedReportColumns, ","))
	if err != nil {
		panic("CSV header row could not be written")
	}

	languages := lingua.AllLanguages()
	totalLanguageCount := len(languages)

	for idx, language := range languages {
		fmt.Printf("Writing reports for %v... (%d/%d)\n", language, idx+1, totalLanguageCount)

		singleWords := getFileContent(testDataDirectory, "single-words", language)
		wordPairs := getFileContent(testDataDirectory, "word-pairs", language)
		sentences := getFileContent(testDataDirectory, "sentences", language)

		linguaStatistics := newDetectorStatistics()

		for _, singleWord := range singleWords {
			linguaLanguage, _ := linguaDetector.DetectLanguageOf(singleWord)
			linguaStatistics.addSingleWordCounts(linguaLanguage, singleWord)
		}

		for _, wordPair := range wordPairs {
			linguaLanguage, _ := linguaDetector.DetectLanguageOf(wordPair)
			linguaStatistics.addWordPairCounts(linguaLanguage, wordPair)
		}

		for _, sentence := range sentences {
			linguaLanguage, _ := linguaDetector.DetectLanguageOf(sentence)
			linguaStatistics.addSentenceCounts(linguaLanguage, sentence)
		}

		linguaStatistics.computeAccuracyValues()

		linguaReport := linguaStatistics.createReportData(language)

		linguaAggregatedReportRow := linguaStatistics.createAggregatedReportRow(language)
		totalAggregatedReportRow := fmt.Sprintf("%s,%s\n", language, linguaAggregatedReportRow)

		_, err = aggregatedReportFile.WriteString(totalAggregatedReportRow)
		if err != nil {
			panic("CSV data row could not be written")
		}

		reportFileName := fmt.Sprintf("%s.txt", language)
		linguaReportsFilePath := filepath.Join(linguaReportsDirectory, reportFileName)

		if len(linguaReport) > 0 {
			linguaReportsFile, err := os.Create(linguaReportsFilePath)
			if err != nil {
				panic("Lingua reports file could not be created")
			}
			defer linguaReportsFile.Close()

			_, err = linguaReportsFile.WriteString(linguaReport)
			if err != nil {
				panic("Lingua reports file could not be written")
			}
		}

		fmt.Println("Done\n")
	}

	elapsed := time.Since(start)
	fmt.Printf("All accuracy reports successfully written in %.0f seconds\n", elapsed.Seconds())
}

func getFileContent(testDataDirectory, subdirectory string, language lingua.Language) []string {
	testDataFileName := fmt.Sprintf("%s.txt", strings.ToLower(language.IsoCode639_1().String()))
	testDataFilePath := filepath.Join(testDataDirectory, subdirectory, testDataFileName)
	testData, err := os.ReadFile(testDataFilePath)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(testData), "\n")
	var filteredLines []string

	for _, line := range lines {
		if utf8.RuneCountInString(strings.TrimSpace(line)) > 0 {
			filteredLines = append(filteredLines, line)
		}
	}
	return filteredLines
}
