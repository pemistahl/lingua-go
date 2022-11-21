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
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
	"path/filepath"
	"testing"
)

const content = `These sentences are intended for testing purposes.
Do not use them in production!
By the way, they consist of 23 words in total.`

const corpus = `There are many attributes associated with good software.
Some of these can be mutually contradictory, and different customers and participants may have different priorities.
Weinberg provides an example of how different goals can have a dramatic effect on both effort required and efficiency.
Furthermore, he notes that programmers will generally aim to achieve any explicit goals which may be set, probably at the expense of any other quality attributes.
Sommerville has identified four generalised attributes which are not concerned with what a program does, but how well the program does it:
Maintainability, Dependability, Efficiency, Usability`

const expectedSentencesFileContent = `There are many attributes associated with good software.
Some of these can be mutually contradictory, and different customers and participants may have different priorities.
Weinberg provides an example of how different goals can have a dramatic effect on both effort required and efficiency.
Furthermore, he notes that programmers will generally aim to achieve any explicit goals which may be set, probably at the expense of any other quality attributes.
`

const expectedSingleWordsFileContent = `there
attributes
associated
software
`

const expectedWordPairsFileContent = `there attributes
associated software
these mutually
contradictory different
`

var expectedUnigramModel = SerializableLanguageModel{
	Language:    SerializableLanguage_ENGLISH,
	NgramLength: 1,
	TotalNgrams: 20,
	NgramSets: []*SerializableNgramSet{
		{
			Probability: 0.01,
			Ngrams:      []string{"b", "g", "l", "m"},
		},
		{
			Probability: 0.02,
			Ngrams:      []string{"f", "w"},
		},
		{
			Probability: 0.03,
			Ngrams:      []string{"a", "c", "p", "u", "y"},
		},
		{
			Probability: 0.04,
			Ngrams:      []string{"h"},
		},
		{
			Probability: 0.05,
			Ngrams:      []string{"d", "r"},
		},
		{
			Probability: 0.06,
			Ngrams:      []string{"i"},
		},
		{
			Probability: 0.1,
			Ngrams:      []string{"n", "o", "s"},
		},
		{
			Probability: 0.13,
			Ngrams:      []string{"t"},
		},
		{
			Probability: 0.14,
			Ngrams:      []string{"e"},
		},
	},
}

var expectedBigramModel = SerializableLanguageModel{
	Language:    SerializableLanguage_ENGLISH,
	NgramLength: 2,
	TotalNgrams: 53,
	NgramSets: []*SerializableNgramSet{
		{
			Probability: 1.0 / 14,
			Ngrams:      []string{"ed", "em", "ey"},
		},
		{
			Probability: 1.0 / 13,
			Ngrams:      []string{"ta", "to"},
		},
		{
			Probability: 0.1,
			Ngrams:      []string{"nc", "nd", "ng", "no", "ns", "od", "of", "os", "si"},
		},
		{
			Probability: 2.0 / 13,
			Ngrams:      []string{"ti"},
		},
		{
			Probability: 1.0 / 6,
			Ngrams:      []string{"io", "is"},
		},
		{
			Probability: 0.2,
			Ngrams:      []string{"de", "do", "ds", "du", "nt", "on", "or", "ot", "rd", "re", "ro", "rp", "st"},
		},
		{
			Probability: 3.0 / 14,
			Ngrams:      []string{"en"},
		},
		{
			Probability: 3.0 / 13,
			Ngrams:      []string{"te"},
		},
		{
			Probability: 2.0 / 7,
			Ngrams:      []string{"es"},
		},
		{
			Probability: 4.0 / 13,
			Ngrams:      []string{"th"},
		},
		{
			Probability: 1.0 / 3,
			Ngrams:      []string{"al", "ar", "ay", "ce", "co", "ct", "po", "pr", "pu", "uc", "ur", "us"},
		},
		{
			Probability: 0.4,
			Ngrams:      []string{"se"},
		},
		{
			Probability: 0.5,
			Ngrams:      []string{"fo", "wa", "wo"},
		},
		{
			Probability: 2.0 / 3,
			Ngrams:      []string{"in"},
		},
		{
			Probability: 1,
			Ngrams:      []string{"by", "he"},
		},
	},
}

var expectedTrigramModel = SerializableLanguageModel{
	Language:    SerializableLanguage_ENGLISH,
	NgramLength: 3,
	TotalNgrams: 51,
	NgramSets: []*SerializableNgramSet{
		{
			Probability: 0.25,
			Ngrams:      []string{"ese", "est", "hem", "hes", "hey", "ing", "int", "sen", "ses"},
		},
		{
			Probability: 1.0 / 3,
			Ngrams:      []string{"enc", "end", "ent", "tes"},
		},
		{
			Probability: 0.5,
			Ngrams:      []string{"ons", "ord", "ota", "sti", "tin", "tio"},
		},
		{
			Probability: 2.0 / 3,
			Ngrams:      []string{"ten"},
		},
		{
			Probability: 1,
			Ngrams:      []string{"are", "ces", "con", "cti", "ded", "duc", "for", "ion", "ist", "nce", "nde", "not", "nsi", "nte", "odu", "ose", "pos", "pro", "pur", "rds", "rod", "rpo", "sis", "tal", "the", "tot", "uct", "urp", "use", "way", "wor"},
		},
	},
}

var expectedQuadrigramModel = SerializableLanguageModel{
	Language:    SerializableLanguage_ENGLISH,
	NgramLength: 4,
	TotalNgrams: 38,
	NgramSets: []*SerializableNgramSet{
		{
			Probability: 0.25,
			Ngrams:      []string{"them", "thes", "they"},
		},
		{
			Probability: 0.5,
			Ngrams:      []string{"tenc", "tend"},
		},
		{
			Probability: 1,
			Ngrams:      []string{"cons", "ctio", "duct", "ence", "ende", "ente", "esti", "hese", "inte", "nces", "nded", "nsis", "nten", "oduc", "onsi", "ords", "oses", "otal", "pose", "prod", "purp", "rodu", "rpos", "sent", "sist", "stin", "test", "ting", "tion", "tota", "ucti", "urpo", "word"},
		},
	},
}

var expectedFivegramModel = SerializableLanguageModel{
	Language:    SerializableLanguage_ENGLISH,
	NgramLength: 5,
	TotalNgrams: 28,
	NgramSets: []*SerializableNgramSet{
		{
			Probability: 0.5,
			Ngrams:      []string{"ntenc", "ntend"},
		},
		{
			Probability: 1,
			Ngrams:      []string{"consi", "ction", "ducti", "ences", "ended", "enten", "estin", "inten", "nsist", "oduct", "onsis", "poses", "produ", "purpo", "roduc", "rpose", "sente", "sting", "tence", "tende", "testi", "these", "total", "uctio", "urpos", "words"},
		},
	},
}

func TestCreateAndWriteLanguageModelFiles(t *testing.T) {
	inputFilePath := createTempInputFile(content)
	outputDirectoryPath, _ := os.MkdirTemp("", "linguaOutputDirectory")

	err := CreateAndWriteLanguageModelFiles(inputFilePath, outputDirectoryPath, English, "\\p{L}")

	assert.Nil(t, err, "language model files could not be created")

	files, _ := os.ReadDir(outputDirectoryPath)

	assert.Equal(t, 5, len(files), "number of language model files is not correct")

	unigramsFile := files[4]
	bigramsFile := files[0]
	trigramsFile := files[3]
	quadrigramsFile := files[2]
	fivegramsFile := files[1]

	assert.Equal(t, "unigrams.pb.bin.zip", unigramsFile.Name())
	assert.Equal(t, "bigrams.pb.bin.zip", bigramsFile.Name())
	assert.Equal(t, "trigrams.pb.bin.zip", trigramsFile.Name())
	assert.Equal(t, "quadrigrams.pb.bin.zip", quadrigramsFile.Name())
	assert.Equal(t, "fivegrams.pb.bin.zip", fivegramsFile.Name())

	assertLanguageModelFileContent(t, outputDirectoryPath, unigramsFile.Name(), "unigrams.pb.bin", &expectedUnigramModel)
	assertLanguageModelFileContent(t, outputDirectoryPath, bigramsFile.Name(), "bigrams.pb.bin", &expectedBigramModel)
	assertLanguageModelFileContent(t, outputDirectoryPath, trigramsFile.Name(), "trigrams.pb.bin", &expectedTrigramModel)
	assertLanguageModelFileContent(t, outputDirectoryPath, quadrigramsFile.Name(), "quadrigrams.pb.bin", &expectedQuadrigramModel)
	assertLanguageModelFileContent(t, outputDirectoryPath, fivegramsFile.Name(), "fivegrams.pb.bin", &expectedFivegramModel)

	cleanUp(inputFilePath, outputDirectoryPath)
}

func TestCreateAndWriteTestDataFiles(t *testing.T) {
	inputFilePath := createTempInputFile(corpus)
	outputDirectoryPath, _ := os.MkdirTemp("", "linguaOutputDirectory")

	err := CreateAndWriteTestDataFiles(inputFilePath, outputDirectoryPath, "\\p{L}", 4)

	assert.Nil(t, err, "test data files could not be created")

	files, _ := os.ReadDir(outputDirectoryPath)

	assert.Equal(t, 3, len(files), "number of test data files is not correct")

	assertTestDataFileContent(
		t,
		files[0],
		outputDirectoryPath,
		"sentences.txt",
		expectedSentencesFileContent,
	)

	assertTestDataFileContent(
		t,
		files[1],
		outputDirectoryPath,
		"single-words.txt",
		expectedSingleWordsFileContent,
	)

	assertTestDataFileContent(
		t,
		files[2],
		outputDirectoryPath,
		"word-pairs.txt",
		expectedWordPairsFileContent,
	)

	cleanUp(inputFilePath, outputDirectoryPath)
}

func createTempInputFile(content string) string {
	inputFile, _ := os.CreateTemp("", "text.*.txt")
	defer inputFile.Close()
	_, _ = inputFile.WriteString(content)
	return inputFile.Name()
}

func assertLanguageModelFileContent(
	t *testing.T,
	outputDirectoryPath string,
	zipFileName string,
	expectedFileName string,
	expectedModel *SerializableLanguageModel,
) {
	zipFilePath := filepath.Join(outputDirectoryPath, zipFileName)
	zipFile, _ := zip.OpenReader(zipFilePath)
	defer zipFile.Close()

	assert.Equal(t, 1, len(zipFile.File), "zip archive does not contain exactly one file")

	protobufFile := zipFile.File[0]

	assert.Equal(t, expectedFileName, protobufFile.Name)

	protobufFileReader, _ := protobufFile.Open()
	defer protobufFileReader.Close()
	protobufFileContent, _ := io.ReadAll(protobufFileReader)

	actualModel := SerializableLanguageModel{}
	_ = proto.Unmarshal(protobufFileContent, &actualModel)

	expectedNgrams := make(map[float64][]string)
	for _, ngramSet := range expectedModel.NgramSets {
		expectedNgrams[ngramSet.Probability] = ngramSet.Ngrams
	}

	actualNgrams := make(map[float64][]string)
	for _, ngramSet := range actualModel.NgramSets {
		actualNgrams[ngramSet.Probability] = ngramSet.Ngrams
	}

	assert.Equal(t, expectedModel.Language, actualModel.Language)
	assert.Equal(t, expectedModel.NgramLength, actualModel.NgramLength)
	assert.Equal(t, expectedModel.TotalNgrams, actualModel.TotalNgrams)
	assert.Equal(t, expectedNgrams, actualNgrams)
}

func assertTestDataFileContent(
	t *testing.T,
	file os.DirEntry,
	outputDirectoryPath,
	expectedFileName,
	expectedFileContent string,
) {
	assert.False(t, file.IsDir())
	assert.Equal(t, expectedFileName, file.Name())

	filePath := filepath.Join(outputDirectoryPath, file.Name())
	testDataFileContent, _ := os.ReadFile(filePath)

	assert.Equal(t, expectedFileContent, string(testDataFileContent))
}

func cleanUp(inputFilePath, outputDirectoryPath string) {
	err := os.Remove(inputFilePath)
	if err != nil {
		panic(fmt.Errorf("temporary file could not be removed: %v", err))
	}
	err = os.RemoveAll(outputDirectoryPath)
	if err != nil {
		panic(fmt.Errorf("temporary directory could not be removed: %v", err))
	}
}
