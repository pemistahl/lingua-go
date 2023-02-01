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

package lingua

import (
	"archive/zip"
	"bufio"
	"errors"
	"fmt"
	"github.com/pemistahl/lingua-go/serialization"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// CreateAndWriteLanguageModelFiles creates language model files for
// accuracy report generation and writes them to a directory.
//
// `inputFilePath` is the path to a txt file used for language model creation.
// The assumed enconding of the txt file is UTF-8.
//
// `outputDirectoryPath` is the path to an existing directory where the language
// model files are to be written.
//
// `language` is the language for which to create the language models.
//
// `charClass` is a regex character class such as `\\p{L}` to restrict the
// set of characters that the language models are built from.
//
// An error is returned if:
//
// - the input file path is not absolute or does not point to an existing txt file
//
// - the input file's encoding is not UTF-8
//
// - the output directory path is not absolute or does not point to an existing directory
//
// Panics if the character class cannot be compiled to a valid regular expression.
func CreateAndWriteLanguageModelFiles(
	inputFilePath string,
	outputDirectoryPath string,
	language Language,
	charClass string,
) error {
	err := checkInputFilePath(inputFilePath)
	if err != nil {
		return err
	}
	err = checkOutputDirectoryPath(outputDirectoryPath)
	if err != nil {
		return err
	}

	unigramModel, err := createLanguageModel(
		inputFilePath,
		language,
		1,
		charClass,
		map[ngram]uint32{},
	)
	if err != nil {
		return err
	}

	bigramModel, err := createLanguageModel(
		inputFilePath,
		language,
		2,
		charClass,
		unigramModel.absoluteFrequencies,
	)
	if err != nil {
		return err
	}

	trigramModel, err := createLanguageModel(
		inputFilePath,
		language,
		3,
		charClass,
		bigramModel.absoluteFrequencies,
	)
	if err != nil {
		return err
	}

	quadrigramModel, err := createLanguageModel(
		inputFilePath,
		language,
		4,
		charClass,
		trigramModel.absoluteFrequencies,
	)
	if err != nil {
		return err
	}

	fivegramModel, err := createLanguageModel(
		inputFilePath,
		language,
		5,
		charClass,
		quadrigramModel.absoluteFrequencies,
	)
	if err != nil {
		return err
	}

	err = writeCompressedLanguageModel(
		unigramModel,
		1,
		outputDirectoryPath,
		"unigrams.pb.bin",
	)
	if err != nil {
		return err
	}

	err = writeCompressedLanguageModel(
		bigramModel,
		2,
		outputDirectoryPath,
		"bigrams.pb.bin",
	)
	if err != nil {
		return err
	}

	err = writeCompressedLanguageModel(
		trigramModel,
		3,
		outputDirectoryPath,
		"trigrams.pb.bin",
	)
	if err != nil {
		return err
	}

	err = writeCompressedLanguageModel(
		quadrigramModel,
		4,
		outputDirectoryPath,
		"quadrigrams.pb.bin",
	)
	if err != nil {
		return err
	}

	err = writeCompressedLanguageModel(
		fivegramModel,
		5,
		outputDirectoryPath,
		"fivegrams.pb.bin",
	)
	if err != nil {
		return err
	}

	return nil
}

// CreateAndWriteTestDataFiles creates test data files for accuracy report
// generation and writes them to a directory.
//
// `inputFilePath` is the path to a txt file used for test data creation.
// The assumed enconding of the txt file is UTF-8.
//
// `outputDirectoryPath` is the path to an existing directory where the test
// data files are to be written.
//
// `charClass` is a regex character class such as `\\p{L}` to restrict the
// set of characters that the language models are built from.
//
// `maximumLines` is the maximum number of lines each test data file should have.
//
// An error is returned if:
//
// - the input file path is not absolute or does not point to an existing txt file
//
// - the input file's encoding is not UTF-8
//
// - the output directory path is not absolute or does not point to an existing directory
//
// Panics if the character class cannot be compiled to a valid regular expression.
func CreateAndWriteTestDataFiles(
	inputFilePath string,
	outputDirectoryPath string,
	charClass string,
	maximumLines int,
) error {
	err := checkInputFilePath(inputFilePath)
	if err != nil {
		return err
	}
	err = checkOutputDirectoryPath(outputDirectoryPath)
	if err != nil {
		return err
	}

	err = createAndWriteSentencesFile(
		inputFilePath,
		outputDirectoryPath,
		maximumLines,
	)
	if err != nil {
		return err
	}

	singleWords, err := createAndWriteSingleWordsFile(
		inputFilePath,
		outputDirectoryPath,
		charClass,
		maximumLines,
	)
	if err != nil {
		return err
	}

	err = createAndWriteWordPairsFile(
		singleWords,
		outputDirectoryPath,
		maximumLines,
	)
	if err != nil {
		return err
	}

	return nil
}

func createAndWriteSentencesFile(
	inputFilePath string,
	outputDirectoryPath string,
	maximumLines int,
) error {
	sentencesFilePath := filepath.Join(outputDirectoryPath, "sentences.txt")
	if _, err := os.Stat(sentencesFilePath); errors.Is(err, os.ErrExist) {
		err = os.Remove(sentencesFilePath)
		if err != nil {
			return err
		}
	}
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var inputLines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			inputLines = append(inputLines, line)
		}
	}

	sentencesFile, err := os.Create(sentencesFilePath)
	if err != nil {
		return err
	}
	defer sentencesFile.Close()

	lineCounter := 0

	for _, line := range inputLines {
		normalizedWhitespace := multipleWhitespace.ReplaceAllString(line, " ")
		removedQuotes := strings.ReplaceAll(normalizedWhitespace, "\"", "")

		if lineCounter < maximumLines {
			_, err = sentencesFile.WriteString(fmt.Sprintf("%s\n", removedQuotes))
			if err != nil {
				return err
			}
			lineCounter++
		} else {
			break
		}
	}

	return nil
}

func createAndWriteSingleWordsFile(
	inputFilePath string,
	outputDirectoryPath string,
	charClass string,
	maximumLines int,
) ([]string, error) {
	singleWordsFilePath := filepath.Join(outputDirectoryPath, "single-words.txt")
	wordRegex, err := regexp.Compile(fmt.Sprintf("[%s]{5,}", charClass))
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(singleWordsFilePath); errors.Is(err, os.ErrExist) {
		err = os.Remove(singleWordsFilePath)
		if err != nil {
			return nil, err
		}
	}

	var words []string

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	var inputLines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			inputLines = append(inputLines, line)
		}
	}

	singleWordsFile, err := os.Create(singleWordsFilePath)
	if err != nil {
		return nil, err
	}
	defer singleWordsFile.Close()

	lineCounter := 0

	for _, line := range inputLines {
		removedPunctuation := punctuation.ReplaceAllString(line, "")
		removedNumbers := numbers.ReplaceAllString(removedPunctuation, "")
		normalizedWhitespace := multipleWhitespace.ReplaceAllString(removedNumbers, " ")
		removedQuotes := strings.ReplaceAll(normalizedWhitespace, "\"", "")

		for _, word := range strings.Split(removedQuotes, " ") {
			word = strings.ToLower(strings.TrimSpace(word))
			if wordRegex.MatchString(word) {
				words = append(words, word)
			}
		}
	}

	for _, word := range words {
		if lineCounter < maximumLines {
			_, err = singleWordsFile.WriteString(fmt.Sprintf("%s\n", word))
			if err != nil {
				return nil, err
			}
			lineCounter++
		} else {
			break
		}
	}

	return words, nil
}

func createAndWriteWordPairsFile(
	words []string,
	outputDirectoryPath string,
	maximumLines int,
) error {
	wordPairsFilePath := filepath.Join(outputDirectoryPath, "word-pairs.txt")

	if _, err := os.Stat(wordPairsFilePath); errors.Is(err, os.ErrExist) {
		err = os.Remove(wordPairsFilePath)
		if err != nil {
			return err
		}
	}

	var wordPairs []string

	for i := 0; i <= len(words)-2; i++ {
		if i%2 == 0 {
			wordPairs = append(wordPairs, strings.Join(words[i:i+2], " "))
		}
	}

	wordPairsFile, err := os.Create(wordPairsFilePath)
	if err != nil {
		return err
	}
	defer wordPairsFile.Close()

	lineCounter := 0

	for _, wordPair := range wordPairs {
		if lineCounter < maximumLines {
			_, err = wordPairsFile.WriteString(fmt.Sprintf("%s\n", wordPair))
			if err != nil {
				return err
			}
			lineCounter++
		} else {
			break
		}
	}

	return nil
}

func createLanguageModel(
	inputFilePath string,
	language Language,
	ngramLength int,
	charClass string,
	lowerNgramAbsoluteFrequencies map[ngram]uint32,
) (trainingDataLanguageModel, error) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return trainingDataLanguageModel{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) > 0 {
			lines = append(lines, line)
		}
	}
	return newTrainingDataLanguageModel(
		lines,
		language,
		ngramLength,
		charClass,
		lowerNgramAbsoluteFrequencies,
	), nil
}

func writeCompressedLanguageModel(
	model trainingDataLanguageModel,
	ngramLength int,
	outputDirectoryPath string,
	fileName string,
) error {
	languageName := strings.ToUpper(model.language.String())
	languageEnumValue := serialization.SerializableLanguage_value[languageName]
	serializableLanguage := serialization.SerializableLanguage(languageEnumValue)

	probabilitiesToNgrams := make(map[float64][]string)
	for ngrm, probability := range model.relativeFrequencies {
		probabilitiesToNgrams[probability] = append(probabilitiesToNgrams[probability], ngrm.value)
	}

	var ngramSets []*serialization.SerializableNgramSet
	for probability, ngrams := range probabilitiesToNgrams {
		sort.Strings(ngrams)
		ngramSet := serialization.SerializableNgramSet{
			Probability: probability,
			Ngrams:      ngrams,
		}
		ngramSets = append(ngramSets, &ngramSet)
	}

	serializableModel := serialization.SerializableLanguageModel{
		Language:    serializableLanguage,
		NgramLength: uint32(ngramLength),
		TotalNgrams: uint32(len(model.relativeFrequencies)),
		NgramSets:   ngramSets,
	}

	serializedModel, err := proto.Marshal(&serializableModel)
	if err != nil {
		return err
	}

	zipFileName := fmt.Sprintf("%s.zip", fileName)
	zipFilePath := filepath.Join(outputDirectoryPath, zipFileName)
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	protobufFileWriter, err := zipWriter.Create(fileName)
	if err != nil {
		return err
	}
	_, err = protobufFileWriter.Write(serializedModel)
	if err != nil {
		return err
	}
	return nil
}

func checkInputFilePath(inputFilePath string) error {
	if !filepath.IsAbs(inputFilePath) {
		return fmt.Errorf("input file path '%s' is not absolute", inputFilePath)
	}

	fileInfo, err := os.Stat(inputFilePath)

	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("input file '%s' does not exist", inputFilePath)
	}
	if !fileInfo.Mode().IsRegular() {
		return fmt.Errorf("input file path '%s' does not represent a regular file", inputFilePath)
	}

	return nil
}

func checkOutputDirectoryPath(outputDirectoryPath string) error {
	if !filepath.IsAbs(outputDirectoryPath) {
		return fmt.Errorf("output directory path '%s' is not absolute", outputDirectoryPath)
	}

	fileInfo, err := os.Stat(outputDirectoryPath)

	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("output directory '%s' does not exist", outputDirectoryPath)
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("output directory path '%s' does not exist", outputDirectoryPath)
	}

	return nil
}
