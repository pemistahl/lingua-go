/*
 * Copyright © 2021-present Peter M. Stahl pemistahl@gmail.com
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
	"github.com/abadojack/whatlanggo"
	"github.com/jmhodges/gocld3/cld3"
	"github.com/pemistahl/lingua-go"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

var zero = big.NewRat(0, 1)

type detectorStatistics struct {
	singleWordStatistic statistic
	wordPairStatistic   statistic
	sentenceStatistic   statistic
	averageAccuracies   map[lingua.Language]*big.Rat
}

type statistic struct {
	languageCounts     map[lingua.Language]int
	languageAccuracies map[lingua.Language]*big.Rat
	entityCount        int
	entityLengthCount  int
}

func newDetectorStatistics() detectorStatistics {
	return detectorStatistics{
		singleWordStatistic: newStatistic(),
		wordPairStatistic:   newStatistic(),
		sentenceStatistic:   newStatistic(),
		averageAccuracies:   make(map[lingua.Language]*big.Rat),
	}
}

func newStatistic() statistic {
	return statistic{
		languageCounts:     make(map[lingua.Language]int),
		languageAccuracies: make(map[lingua.Language]*big.Rat),
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
	sum := big.NewRat(0, 1)
	sum = sum.Add(sum, singleWordAccuracy)
	sum = sum.Add(sum, wordPairAccuracy)
	sum = sum.Add(sum, sentenceAccuracy)
	averageAccuracy := sum.Quo(sum, big.NewRat(3, 1))
	ds.averageAccuracies[language] = averageAccuracy

	if averageAccuracy.Cmp(zero) == 0 {
		return ""
	}

	averageFloatAccuracy, _ := averageAccuracy.Float64()

	return fmt.Sprintf(
		"##### %s #####\n\n>>> Accuracy on average: %.2f%%\n\n%s\n%s\n%s\n",
		language,
		averageFloatAccuracy*100,
		singleWordReport,
		wordPairReport,
		sentenceReport,
	)
}

func (ds *detectorStatistics) createAggregatedReportRow(language lingua.Language) string {
	var averageAccuracyColumn string
	accuracy, exists := ds.averageAccuracies[language]
	if exists && accuracy.Cmp(zero) == 1 {
		floatAccuracy, _ := accuracy.Float64()
		averageAccuracyColumn = fmt.Sprintf("%.0f", floatAccuracy*100)
	} else {
		averageAccuracyColumn = "NaN"
	}

	var singleWordsAccuracyColumn string
	accuracy, exists = ds.singleWordStatistic.languageAccuracies[language]
	if exists && accuracy.Cmp(zero) == 1 {
		floatAccuracy, _ := accuracy.Float64()
		singleWordsAccuracyColumn = fmt.Sprintf("%.0f", floatAccuracy*100)
	} else {
		singleWordsAccuracyColumn = "NaN"
	}

	var wordPairsAccuracyColumn string
	accuracy, exists = ds.wordPairStatistic.languageAccuracies[language]
	if exists && accuracy.Cmp(zero) == 1 {
		floatAccuracy, _ := accuracy.Float64()
		wordPairsAccuracyColumn = fmt.Sprintf("%.0f", floatAccuracy*100)
	} else {
		wordPairsAccuracyColumn = "NaN"
	}

	var sentencesAccuracyColumn string
	accuracy, exists = ds.sentenceStatistic.languageAccuracies[language]
	if exists && accuracy.Cmp(zero) == 1 {
		floatAccuracy, _ := accuracy.Float64()
		sentencesAccuracyColumn = fmt.Sprintf("%.0f", floatAccuracy*100)
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
		s.languageAccuracies[language] = big.NewRat(int64(count), int64(sumOfCounts))
	}
}

func (s *statistic) createReportData(
	language lingua.Language,
	description string,
) (*big.Rat, string) {
	accuracy, exists := s.languageAccuracies[language]
	if !exists {
		accuracy = zero
	}
	averageLength := int(math.Round(float64(s.entityLengthCount) / float64(s.entityCount)))
	floatAccuracy, _ := accuracy.Float64()
	report := fmt.Sprintf(
		">> Detection of %d %s (average length: %d chars)\nAccuracy: %.2f%%\nErroneously classified as %s\n",
		s.entityCount,
		description,
		averageLength,
		floatAccuracy*100,
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
		if firstAccuracy.Cmp(secondAccuracy) == 0 {
			return firstLanguage < secondLanguage
		}
		return firstAccuracy.Cmp(secondAccuracy) == 1
	})
	var reports []string
	for _, currentLanguage := range languages {
		accuracy, _ := s.languageAccuracies[currentLanguage].Float64()
		report := fmt.Sprintf("%s: %.2f%%", currentLanguage, accuracy*100)
		reports = append(reports, report)
	}
	return strings.Join(reports, ", ")
}

func main() {
	start := time.Now()

	linguaDetectorWithHighAccuracy := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()

	linguaDetectorWithLowAccuracy := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithLowAccuracyMode().
		WithPreloadedLanguageModels().
		Build()

	cld3Detector, _ := cld3.NewLanguageIdentifier(0, 512)
	defer cld3.FreeLanguageIdentifier(cld3Detector)

	testDataDirectory, _ := filepath.Abs("language-testdata")
	accuracyReportsDirectory, _ := filepath.Abs("accuracy-reports")
	linguaHighAccuracyReportsDirectory := filepath.Join(accuracyReportsDirectory, "lingua-high-accuracy")
	linguaLowAccuracyReportsDirectory := filepath.Join(accuracyReportsDirectory, "lingua-low-accuracy")
	cld3ReportsDirectory := filepath.Join(accuracyReportsDirectory, "cld3")
	whatlangReportsDirectory := filepath.Join(accuracyReportsDirectory, "whatlang")

	err := os.MkdirAll(linguaHighAccuracyReportsDirectory, os.ModePerm)
	if err != nil {
		panic("Lingua reports directory could not be created")
	}

	err = os.MkdirAll(linguaLowAccuracyReportsDirectory, os.ModePerm)
	if err != nil {
		panic("Lingua reports directory could not be created")
	}

	err = os.MkdirAll(cld3ReportsDirectory, os.ModePerm)
	if err != nil {
		panic("CLD3 reports directory could not be created")
	}

	err = os.MkdirAll(whatlangReportsDirectory, os.ModePerm)
	if err != nil {
		panic("Whatlang reports directory could not be created")
	}

	aggregatedReportFilePath := filepath.Join(accuracyReportsDirectory, "aggregated-accuracy-values.csv")
	aggregatedReportFile, err := os.Create(aggregatedReportFilePath)
	if err != nil {
		panic("CSV file could not be created")
	}
	defer aggregatedReportFile.Close()

	aggregatedReportColumns := []string{
		"language",
		"average-whatlang",
		"single-words-whatlang",
		"word-pairs-whatlang",
		"sentences-whatlang",
		"average-cld3",
		"single-words-cld3",
		"word-pairs-cld3",
		"sentences-cld3",
		"average-lingua-low",
		"single-words-lingua-low",
		"word-pairs-lingua-low",
		"sentences-lingua-low",
		"average-lingua-high",
		"single-words-lingua-high",
		"word-pairs-lingua-high",
		"sentences-lingua-high\n",
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

		linguaHighAccuracyStatistics := newDetectorStatistics()
		linguaLowAccuracyStatistics := newDetectorStatistics()
		cld3Statistics := newDetectorStatistics()
		whatlangStatistics := newDetectorStatistics()

		for _, singleWord := range singleWords {
			linguaLanguageInHighAccuracyMode, _ := linguaDetectorWithHighAccuracy.DetectLanguageOf(singleWord)
			linguaHighAccuracyStatistics.addSingleWordCounts(linguaLanguageInHighAccuracyMode, singleWord)

			linguaLanguageInLowAccuracyMode, _ := linguaDetectorWithLowAccuracy.DetectLanguageOf(singleWord)
			linguaLowAccuracyStatistics.addSingleWordCounts(linguaLanguageInLowAccuracyMode, singleWord)

			cld3Language := mapCld3ToLingua(cld3Detector.FindLanguage(singleWord).Language)
			cld3Statistics.addSingleWordCounts(cld3Language, singleWord)

			whatlangLanguage := mapWhatlangToLingua(whatlanggo.DetectLang(singleWord))
			whatlangStatistics.addSingleWordCounts(whatlangLanguage, singleWord)
		}

		for _, wordPair := range wordPairs {
			linguaLanguageInHighAccuracyMode, _ := linguaDetectorWithHighAccuracy.DetectLanguageOf(wordPair)
			linguaHighAccuracyStatistics.addWordPairCounts(linguaLanguageInHighAccuracyMode, wordPair)

			linguaLanguageInLowAccuracyMode, _ := linguaDetectorWithLowAccuracy.DetectLanguageOf(wordPair)
			linguaLowAccuracyStatistics.addWordPairCounts(linguaLanguageInLowAccuracyMode, wordPair)

			cld3Language := mapCld3ToLingua(cld3Detector.FindLanguage(wordPair).Language)
			cld3Statistics.addWordPairCounts(cld3Language, wordPair)

			whatlangLanguage := mapWhatlangToLingua(whatlanggo.DetectLang(wordPair))
			whatlangStatistics.addWordPairCounts(whatlangLanguage, wordPair)
		}

		for _, sentence := range sentences {
			linguaLanguageInHighAccuracyMode, _ := linguaDetectorWithHighAccuracy.DetectLanguageOf(sentence)
			linguaHighAccuracyStatistics.addSentenceCounts(linguaLanguageInHighAccuracyMode, sentence)

			linguaLanguageInLowAccuracyMode, _ := linguaDetectorWithLowAccuracy.DetectLanguageOf(sentence)
			linguaLowAccuracyStatistics.addSentenceCounts(linguaLanguageInLowAccuracyMode, sentence)

			cld3Language := mapCld3ToLingua(cld3Detector.FindLanguage(sentence).Language)
			cld3Statistics.addSentenceCounts(cld3Language, sentence)

			whatlangLanguage := mapWhatlangToLingua(whatlanggo.DetectLang(sentence))
			whatlangStatistics.addSentenceCounts(whatlangLanguage, sentence)
		}

		linguaHighAccuracyStatistics.computeAccuracyValues()
		linguaLowAccuracyStatistics.computeAccuracyValues()
		cld3Statistics.computeAccuracyValues()
		whatlangStatistics.computeAccuracyValues()

		linguaHighAccuracyReport := linguaHighAccuracyStatistics.createReportData(language)
		linguaLowAccuracyReport := linguaLowAccuracyStatistics.createReportData(language)
		cld3Report := cld3Statistics.createReportData(language)
		whatlangReport := whatlangStatistics.createReportData(language)

		linguaHighAccuracyAggregatedReportRow := linguaHighAccuracyStatistics.createAggregatedReportRow(language)
		linguaLowAccuracyAggregatedReportRow := linguaLowAccuracyStatistics.createAggregatedReportRow(language)
		cld3AggregatedReportRow := cld3Statistics.createAggregatedReportRow(language)
		whatlangAggregatedReportRow := whatlangStatistics.createAggregatedReportRow(language)
		totalAggregatedReportRow := fmt.Sprintf(
			"%s,%s,%s,%s,%s\n",
			language,
			whatlangAggregatedReportRow,
			cld3AggregatedReportRow,
			linguaLowAccuracyAggregatedReportRow,
			linguaHighAccuracyAggregatedReportRow,
		)

		_, err = aggregatedReportFile.WriteString(totalAggregatedReportRow)
		if err != nil {
			panic("CSV data row could not be written")
		}

		reportFileName := fmt.Sprintf("%s.txt", language)
		linguaHighAccuracyReportsFilePath := filepath.Join(linguaHighAccuracyReportsDirectory, reportFileName)
		linguaLowAccuracyReportsFilePath := filepath.Join(linguaLowAccuracyReportsDirectory, reportFileName)
		cld3ReportsFilePath := filepath.Join(cld3ReportsDirectory, reportFileName)
		whatlangReportsFilePath := filepath.Join(whatlangReportsDirectory, reportFileName)

		if len(linguaHighAccuracyReport) > 0 {
			linguaReportsFile, err := os.Create(linguaHighAccuracyReportsFilePath)
			if err != nil {
				panic("Lingua reports file could not be created")
			}

			_, err = linguaReportsFile.WriteString(linguaHighAccuracyReport)
			if err != nil {
				panic("Lingua reports file could not be written")
			}
			linguaReportsFile.Close()
		}

		if len(linguaLowAccuracyReport) > 0 {
			linguaReportsFile, err := os.Create(linguaLowAccuracyReportsFilePath)
			if err != nil {
				panic("Lingua reports file could not be created")
			}

			_, err = linguaReportsFile.WriteString(linguaLowAccuracyReport)
			if err != nil {
				panic("Lingua reports file could not be written")
			}
			linguaReportsFile.Close()
		}

		if len(cld3Report) > 0 {
			cld3ReportsFile, err := os.Create(cld3ReportsFilePath)
			if err != nil {
				panic("CLD3 reports file could not be created")
			}

			_, err = cld3ReportsFile.WriteString(cld3Report)
			if err != nil {
				panic("CLD3 reports file could not be written")
			}
			cld3ReportsFile.Close()
		}

		if len(whatlangReport) > 0 {
			whatlangReportsFile, err := os.Create(whatlangReportsFilePath)
			if err != nil {
				panic("Whatlang reports file could not be created")
			}

			_, err = whatlangReportsFile.WriteString(whatlangReport)
			if err != nil {
				panic("Whatlang reports file could not be written")
			}
			whatlangReportsFile.Close()
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

func mapCld3ToLingua(isoCode string) lingua.Language {
	for _, language := range lingua.AllLanguages() {
		linguaIsoCode := strings.ToLower(language.IsoCode639_1().String())
		if linguaIsoCode == isoCode {
			return language
		}
	}
	return lingua.Unknown
}

func mapWhatlangToLingua(language whatlanggo.Lang) lingua.Language {
	switch language {
	case whatlanggo.Afr:
		return lingua.Afrikaans
	case whatlanggo.Arb:
		return lingua.Arabic
	case whatlanggo.Azj:
		return lingua.Azerbaijani
	case whatlanggo.Bel:
		return lingua.Belarusian
	case whatlanggo.Ben:
		return lingua.Bengali
	case whatlanggo.Bul:
		return lingua.Bulgarian
	case whatlanggo.Ces:
		return lingua.Czech
	case whatlanggo.Cmn:
		return lingua.Chinese
	case whatlanggo.Dan:
		return lingua.Danish
	case whatlanggo.Deu:
		return lingua.German
	case whatlanggo.Ell:
		return lingua.Greek
	case whatlanggo.Eng:
		return lingua.English
	case whatlanggo.Epo:
		return lingua.Esperanto
	case whatlanggo.Est:
		return lingua.Estonian
	case whatlanggo.Fin:
		return lingua.Finnish
	case whatlanggo.Fra:
		return lingua.French
	case whatlanggo.Guj:
		return lingua.Gujarati
	case whatlanggo.Heb:
		return lingua.Hebrew
	case whatlanggo.Hin:
		return lingua.Hindi
	case whatlanggo.Hrv:
		return lingua.Croatian
	case whatlanggo.Hun:
		return lingua.Hungarian
	case whatlanggo.Ind:
		return lingua.Indonesian
	case whatlanggo.Ita:
		return lingua.Italian
	case whatlanggo.Jpn:
		return lingua.Japanese
	case whatlanggo.Kat:
		return lingua.Georgian
	case whatlanggo.Kor:
		return lingua.Korean
	case whatlanggo.Lav:
		return lingua.Latvian
	case whatlanggo.Lit:
		return lingua.Lithuanian
	case whatlanggo.Mar:
		return lingua.Marathi
	case whatlanggo.Mkd:
		return lingua.Macedonian
	case whatlanggo.Nld:
		return lingua.Dutch
	case whatlanggo.Nno:
		return lingua.Nynorsk
	case whatlanggo.Nob:
		return lingua.Bokmal
	case whatlanggo.Pan:
		return lingua.Punjabi
	case whatlanggo.Pes:
		return lingua.Persian
	case whatlanggo.Pol:
		return lingua.Polish
	case whatlanggo.Por:
		return lingua.Portuguese
	case whatlanggo.Ron:
		return lingua.Romanian
	case whatlanggo.Rus:
		return lingua.Russian
	case whatlanggo.Slv:
		return lingua.Slovene
	case whatlanggo.Sna:
		return lingua.Shona
	case whatlanggo.Som:
		return lingua.Somali
	case whatlanggo.Spa:
		return lingua.Spanish
	case whatlanggo.Srp:
		return lingua.Serbian
	case whatlanggo.Swe:
		return lingua.Swedish
	case whatlanggo.Tam:
		return lingua.Tamil
	case whatlanggo.Tel:
		return lingua.Telugu
	case whatlanggo.Tgl:
		return lingua.Tagalog
	case whatlanggo.Tha:
		return lingua.Thai
	case whatlanggo.Tur:
		return lingua.Turkish
	case whatlanggo.Ukr:
		return lingua.Ukrainian
	case whatlanggo.Urd:
		return lingua.Urdu
	case whatlanggo.Vie:
		return lingua.Vietnamese
	case whatlanggo.Yor:
		return lingua.Yoruba
	case whatlanggo.Zul:
		return lingua.Zulu
	case whatlanggo.Mal:
		return lingua.Malayalam
	default:
		return lingua.Unknown
	}
}
