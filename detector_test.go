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
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"sync"
	"testing"
)

// ##############################
// LANGUAGE MODELS FOR ENGLISH
// ##############################

var unigramModelForEnglish = map[string]float64{
	"a": 0.01,
	"l": 0.02,
	"t": 0.03,
	"e": 0.04,
	"r": 0.05,
	// unknown unigrams
	"w": 0.0,
}

var bigramModelForEnglish = map[string]float64{
	"al": 0.11,
	"lt": 0.12,
	"te": 0.13,
	"er": 0.14,
	// unknown bigrams
	"aq": 0.0,
	"wx": 0.0,
}

var trigramModelForEnglish = map[string]float64{
	"alt": 0.19,
	"lte": 0.2,
	"ter": 0.21,
	// unknown trigrams
	"aqu": 0.0,
	"tez": 0.0,
	"wxy": 0.0,
}

var quadrigramModelForEnglish = map[string]float64{
	"alte": 0.25,
	"lter": 0.26,
	// unknown quadrigrams
	"aqua": 0.0,
	"wxyz": 0.0,
}

var fivegramModelForEnglish = map[string]float64{
	"alter": 0.29,
	// unknown fivegrams
	"aquas": 0.0,
}

// ##############################
// LANGUAGE MODELS FOR GERMAN
// ##############################

var unigramModelForGerman = map[string]float64{
	"a": 0.06,
	"l": 0.07,
	"t": 0.08,
	"e": 0.09,
	"r": 0.1,
	// unknown unigrams
	"w": 0.0,
}

var bigramModelForGerman = map[string]float64{
	"al": 0.15,
	"lt": 0.16,
	"te": 0.17,
	"er": 0.18,
	// unknown bigrams
	"wx": 0.0,
}

var trigramModelForGerman = map[string]float64{
	"alt": 0.22,
	"lte": 0.23,
	"ter": 0.24,
	// unknown trigrams
	"wxy": 0.0,
}

var quadrigramModelForGerman = map[string]float64{
	"alte": 0.27,
	"lter": 0.28,
	// unknown quadrigrams
	"wxyz": 0.0,
}

var fivegramModelForGerman = map[string]float64{
	"alter": 0.3,
}

// ##############################
// TEST DATA MODELS
// ##############################

func testDataModel(strings [][]string) testDataLanguageModel {
	ngrams := make([][]ngram, len(strings))
	for i, strs := range strings {
		n := make([]ngram, len(strs))
		for j, s := range strs {
			n[j] = newNgram(s)
		}
		ngrams[i] = n
	}
	return testDataLanguageModel{ngrams}
}

// ##############################
// DETECTORS
// ##############################

func newDetectorForEnglishAndGerman() languageDetector {
	languages := []Language{English, German}

	var unigramLanguageModels sync.Map
	unigramLanguageModels.Store(English, unigramModelForEnglish)
	unigramLanguageModels.Store(German, unigramModelForGerman)

	var bigramLanguageModels sync.Map
	bigramLanguageModels.Store(English, bigramModelForEnglish)
	bigramLanguageModels.Store(German, bigramModelForGerman)

	var trigramLanguageModels sync.Map
	trigramLanguageModels.Store(English, trigramModelForEnglish)
	trigramLanguageModels.Store(German, trigramModelForGerman)

	var quadrigramLanguageModels sync.Map
	quadrigramLanguageModels.Store(English, quadrigramModelForEnglish)
	quadrigramLanguageModels.Store(German, quadrigramModelForGerman)

	var fivegramLanguageModels sync.Map
	fivegramLanguageModels.Store(English, fivegramModelForEnglish)
	fivegramLanguageModels.Store(German, fivegramModelForGerman)

	return languageDetector{
		languages:                     languages,
		minimumRelativeDistance:       0.0,
		isLowAccuracyModeEnabled:      false,
		languagesWithUniqueCharacters: collectLanguagesWithUniqueCharacters(languages),
		oneLanguageAlphabets:          collectOneLanguageAlphabets(languages),
		unigramLanguageModels:         &unigramLanguageModels,
		bigramLanguageModels:          &bigramLanguageModels,
		trigramLanguageModels:         &trigramLanguageModels,
		quadrigramLanguageModels:      &quadrigramLanguageModels,
		fivegramLanguageModels:        &fivegramLanguageModels,
	}
}

var detectorForEnglishAndGerman = newDetectorForEnglishAndGerman()
var detectorForAllLanguages = newLanguageDetector(AllLanguages(), 0.0, true, false)

// ##############################
// TESTS
// ##############################

var delta = 0.00000000000001

func TestSplitTextIntoWords(t *testing.T) {
	testCases := []struct {
		text          string
		expectedWords []string
	}{
		{
			"this is a sentence",
			[]string{"this", "is", "a", "sentence"},
		},
		{
			"sentence",
			[]string{"sentence"},
		},
		{
			"上海大学是一个好大学 this is a sentence",
			[]string{"上", "海", "大", "学", "是", "一", "个", "好", "大", "学", "this", "is", "a", "sentence"},
		},
		{
			"Weltweit    gibt es ungefähr 6.000 Sprachen.",
			[]string{"weltweit", "gibt", "es", "ungefähr", "sprachen"},
		},
	}
	for _, testCase := range testCases {
		assert.Equal(
			t,
			testCase.expectedWords,
			splitTextIntoWords(testCase.text),
			fmt.Sprintf("unexpected tokenization for text '%s'", testCase.text),
		)
	}
}

func TestLookUpNgramProbability(t *testing.T) {
	testCases := []struct {
		language            Language
		ngram               string
		expectedProbability float64
	}{
		{English, "a", 0.01},
		{English, "lt", 0.12},
		{English, "ter", 0.21},
		{English, "alte", 0.25},
		{English, "alter", 0.29},
		{German, "t", 0.08},
		{German, "er", 0.18},
		{German, "alt", 0.22},
		{German, "lter", 0.28},
		{German, "alter", 0.3},
	}
	for _, testCase := range testCases {
		probability := detectorForEnglishAndGerman.lookUpNgramProbability(testCase.language, newNgram(testCase.ngram))
		message := fmt.Sprintf(
			"expected probability %v for language %v and ngram '%s', got %v",
			testCase.expectedProbability,
			testCase.language,
			testCase.ngram,
			probability,
		)
		assert.Equal(t, testCase.expectedProbability, probability, message)
	}

	assert.Panicsf(t, func() {
		detectorForEnglishAndGerman.lookUpNgramProbability(English, newNgram(""))
	}, "zerogram detected")
}

func TestComputeSumOfNgramProbabilities(t *testing.T) {
	testCases := []struct {
		ngramModel                 testDataLanguageModel
		expectedSumOfProbabilities float64
	}{
		{
			testDataModel([][]string{{"a"}, {"l"}, {"t"}, {"e"}, {"r"}}),
			math.Log(0.01) + math.Log(0.02) + math.Log(0.03) + math.Log(0.04) + math.Log(0.05),
		},
		{
			// back off unknown Trigram("tez") to known Bigram("te")
			testDataModel([][]string{{"alt", "al", "a"}, {"lte", "lt", "l"}, {"tez", "te", "t"}}),
			math.Log(0.19) + math.Log(0.2) + math.Log(0.13),
		},
		{
			// back off unknown Fivegram("aquas") to known Unigram("a")
			testDataModel([][]string{{"aquas", "aqua", "aqu", "aq", "a"}}),
			math.Log(0.01),
		},
	}
	for _, testCase := range testCases {
		sumOfProbabilities := detectorForEnglishAndGerman.computeSumOfNgramProbabilities(English, testCase.ngramModel)
		message := fmt.Sprintf(
			"expected sum %v for language %v and ngrams %v, got %v",
			testCase.expectedSumOfProbabilities,
			English,
			testCase.ngramModel.ngrams,
			sumOfProbabilities,
		)
		assert.InDelta(t, testCase.expectedSumOfProbabilities, sumOfProbabilities, delta, message)
	}
}

func TestComputeLanguageProbabilities(t *testing.T) {
	testCases := []struct {
		ngramModel            testDataLanguageModel
		expectedProbabilities map[Language]float64
	}{
		{
			testDataModel([][]string{{"a"}, {"l"}, {"t"}, {"e"}, {"r"}}),
			map[Language]float64{
				English: math.Log(0.01) + math.Log(0.02) + math.Log(0.03) + math.Log(0.04) + math.Log(0.05),
				German:  math.Log(0.06) + math.Log(0.07) + math.Log(0.08) + math.Log(0.09) + math.Log(0.1),
			},
		},
		{
			testDataModel([][]string{{"alt", "al", "a"}, {"lte", "lt", "l"}, {"ter", "te", "t"}, {"wxy", "wx", "w"}}),
			map[Language]float64{
				English: math.Log(0.19) + math.Log(0.2) + math.Log(0.21),
				German:  math.Log(0.22) + math.Log(0.23) + math.Log(0.24),
			},
		},
		{
			testDataModel([][]string{{"alte", "alt", "al", "a"}, {"lter", "lte", "lt", "l"}, {"wxyz", "wxy", "wx", "w"}}),
			map[Language]float64{
				English: math.Log(0.25) + math.Log(0.26),
				German:  math.Log(0.27) + math.Log(0.28),
			},
		},
	}
	languages := []Language{English, German}
	for _, testCase := range testCases {
		probabilities := detectorForEnglishAndGerman.computeLanguageProbabilities(testCase.ngramModel, languages)

		for language, probability := range probabilities {
			expectedProbability := testCase.expectedProbabilities[language]
			message := fmt.Sprintf(
				"expected probability %v for language %v, got %v",
				expectedProbability,
				language,
				probability,
			)
			assert.InDelta(t, expectedProbability, probability, delta, message)
		}
	}
}

func TestComputeLanguageConfidenceValues_LanguageDetectedByRules(t *testing.T) {
	confidenceValues := detectorForEnglishAndGerman.ComputeLanguageConfidenceValues("groß")

	assert.Equal(
		t,
		2,
		len(confidenceValues),
		fmt.Sprintf("expected 2 confidence values, got %v", len(confidenceValues)),
	)

	first, second := confidenceValues[0], confidenceValues[1]

	assert.Equal(t, German, first.Language())
	assert.Equal(t, 1.0, first.Value())

	assert.Equal(t, English, second.Language())
	assert.Equal(t, 0.0, second.Value())
}

func TestComputeLanguageConfidenceValues_KnownNgrams(t *testing.T) {
	confidenceValues := detectorForEnglishAndGerman.ComputeLanguageConfidenceValues("Alter")

	assert.Equal(
		t,
		2,
		len(confidenceValues),
		fmt.Sprintf("expected 2 confidence values, got %v", len(confidenceValues)),
	)

	first, second := confidenceValues[0], confidenceValues[1]

	assert.Equal(t, German, first.Language())
	assert.Equal(t, 0.81, roundToTwoDecimalPlaces(first.Value()))

	assert.Equal(t, English, second.Language())
	assert.Equal(t, 0.19, roundToTwoDecimalPlaces(second.Value()))
}

func TestComputeLanguageConfidenceValues_UnknownNgrams(t *testing.T) {
	confidenceValues := detectorForEnglishAndGerman.ComputeLanguageConfidenceValues("проарплап")

	assert.Equal(
		t,
		2,
		len(confidenceValues),
		fmt.Sprintf("expected 2 confidence values, got %v", len(confidenceValues)),
	)

	first, second := confidenceValues[0], confidenceValues[1]

	assert.Equal(t, English, first.Language())
	assert.Equal(t, 0.0, first.Value())

	assert.Equal(t, German, second.Language())
	assert.Equal(t, 0.0, second.Value())
}

func TestComputeLanguageConfidence_LanguageDetectedByRules(t *testing.T) {
	confidence := detectorForEnglishAndGerman.ComputeLanguageConfidence("groß", German)
	assert.Equal(t, 1.0, confidence)

	confidence = detectorForEnglishAndGerman.ComputeLanguageConfidence("groß", English)
	assert.Equal(t, 0.0, confidence)
}

func TestComputeLanguageConfidence_KnownNgrams(t *testing.T) {
	confidence := detectorForEnglishAndGerman.ComputeLanguageConfidence("Alter", German)
	assert.Equal(t, 0.81, roundToTwoDecimalPlaces(confidence))

	confidence = detectorForEnglishAndGerman.ComputeLanguageConfidence("Alter", English)
	assert.Equal(t, 0.19, roundToTwoDecimalPlaces(confidence))
}

func TestComputeLanguageConfidence_UnknownNgrams(t *testing.T) {
	confidence := detectorForEnglishAndGerman.ComputeLanguageConfidence("проарплап", German)
	assert.Equal(t, 0.0, confidence)

	confidence = detectorForEnglishAndGerman.ComputeLanguageConfidence("проарплап", English)
	assert.Equal(t, 0.0, confidence)
}

func TestComputeLanguageConfidence_UnknownLanguage(t *testing.T) {
	confidence := detectorForEnglishAndGerman.ComputeLanguageConfidence("groß", French)
	assert.Equal(t, 0.0, confidence)
}

func TestDetectLanguage(t *testing.T) {
	language, exists := detectorForEnglishAndGerman.DetectLanguageOf("Alter")
	assert.Equal(t, German, language)
	assert.True(t, exists)

	language, exists = detectorForEnglishAndGerman.DetectLanguageOf("проарплап")
	assert.Equal(t, Unknown, language)
	assert.False(t, exists)
}

func TestDetectMultipleLanguages_EmptyString(t *testing.T) {
	assert.Empty(t, detectorForAllLanguages.DetectMultipleLanguagesOf(""))
}

func TestDetectMultipleLanguages_English(t *testing.T) {
	sentence := "I'm really not sure whether multi-language detection is a good idea."

	results := detectorForAllLanguages.DetectMultipleLanguagesOf(sentence)
	assert.Equal(t, 1, len(results))

	result := results[0]
	substring := sentence[result.StartIndex():result.EndIndex()]
	assert.Equal(t, sentence, substring)
	assert.Equal(t, English, result.Language())
}

func TestDetectMultipleLanguages_EnglishGerman(t *testing.T) {
	sentence := "  He   turned around and asked: " +
		"\"Entschuldigen Sie, sprechen Sie Deutsch?\""

	results := detectorForAllLanguages.DetectMultipleLanguagesOf(sentence)
	assert.Equal(t, 2, len(results))

	firstResult := results[0]
	firstSubstring := sentence[firstResult.StartIndex():firstResult.EndIndex()]
	assert.Equal(t, "  He   turned around and asked: ", firstSubstring)
	assert.Equal(t, English, firstResult.Language())

	secondResult := results[1]
	secondSubstring := sentence[secondResult.StartIndex():secondResult.EndIndex()]
	assert.Equal(t, "\"Entschuldigen Sie, sprechen Sie Deutsch?\"", secondSubstring)
	assert.Equal(t, German, secondResult.Language())
}

func TestDetectMultipleLanguages_ChineseEnglish(t *testing.T) {
	sentence := "上海大学是一个好大学. It is such a great university."

	results := detectorForAllLanguages.DetectMultipleLanguagesOf(sentence)
	assert.Equal(t, 2, len(results))

	firstResult := results[0]
	firstSubstring := sentence[firstResult.StartIndex():firstResult.EndIndex()]
	assert.Equal(t, "上海大学是一个好大学. ", firstSubstring)
	assert.Equal(t, Chinese, firstResult.Language())

	secondResult := results[1]
	secondSubstring := sentence[secondResult.StartIndex():secondResult.EndIndex()]
	assert.Equal(t, "It is such a great university.", secondSubstring)
	assert.Equal(t, English, secondResult.Language())
}

func TestDetectMultipleLanguages_FrenchGermanEnglish(t *testing.T) {
	sentence := "Parlez-vous français? " +
		"Ich spreche Französisch nur ein bisschen. " +
		"A little bit is better than nothing."

	results := detectorForAllLanguages.DetectMultipleLanguagesOf(sentence)
	assert.Equal(t, 3, len(results))

	firstResult := results[0]
	firstSubstring := sentence[firstResult.StartIndex():firstResult.EndIndex()]
	assert.Equal(t, "Parlez-vous français? ", firstSubstring)
	assert.Equal(t, French, firstResult.Language())

	secondResult := results[1]
	secondSubstring := sentence[secondResult.StartIndex():secondResult.EndIndex()]
	assert.Equal(t, "Ich spreche Französisch nur ein bisschen. ", secondSubstring)
	assert.Equal(t, German, secondResult.Language())

	thirdResult := results[2]
	thirdSubstring := sentence[thirdResult.StartIndex():thirdResult.EndIndex()]
	assert.Equal(t, "A little bit is better than nothing.", thirdSubstring)
	assert.Equal(t, English, thirdResult.Language())
}

func TestDetectLanguageWithRules(t *testing.T) {
	testCases := []struct {
		word             string
		expectedLanguage Language
	}{
		// words with unique characters
		{"məhərrəm", Azerbaijani},
		{"substituïts", Catalan},
		{"rozdělit", Czech},
		{"tvořen", Czech},
		{"subjektů", Czech},
		{"nesufiĉecon", Esperanto},
		{"intermiksiĝis", Esperanto},
		{"monaĥinoj", Esperanto},
		{"kreitaĵoj", Esperanto},
		{"ŝpinante", Esperanto},
		{"apenaŭ", Esperanto},
		{"groß", German},
		{"σχέδια", Greek},
		{"fekvő", Hungarian},
		{"meggyűrűzni", Hungarian},
		{"ヴェダイヤモンド", Japanese},
		{"әлем", Kazakh},
		{"шаруашылығы", Kazakh},
		{"ақын", Kazakh},
		{"оның", Kazakh},
		{"шұрайлы", Kazakh},
		{"teoloģiska", Latvian},
		{"blaķene", Latvian},
		{"ceļojumiem", Latvian},
		{"numuriņu", Latvian},
		{"mergelės", Lithuanian},
		{"įrengus", Lithuanian},
		{"slegiamų", Lithuanian},
		{"припаѓа", Macedonian},
		{"ѕидови", Macedonian},
		{"ќерка", Macedonian},
		{"џамиите", Macedonian},
		{"मिळते", Marathi},
		{"үндсэн", Mongolian},
		{"дөхөж", Mongolian},
		{"zmieniły", Polish},
		{"państwowych", Polish},
		{"mniejszości", Polish},
		{"groźne", Polish},
		{"ialomiţa", Romanian},
		{"наслеђивања", Serbian},
		{"неисквареношћу", Serbian},
		{"podĺa", Slovak},
		{"pohľade", Slovak},
		{"mŕtvych", Slovak},
		{"ґрунтовому", Ukrainian},
		{"пропонує", Ukrainian},
		{"пристрої", Ukrainian},
		{"cằm", Vietnamese},
		{"thần", Vietnamese},
		{"chẳng", Vietnamese},
		{"quẩy", Vietnamese},
		{"sẵn", Vietnamese},
		{"nhẫn", Vietnamese},
		{"dắt", Vietnamese},
		{"chất", Vietnamese},
		{"đạp", Vietnamese},
		{"mặn", Vietnamese},
		{"hậu", Vietnamese},
		{"hiền", Vietnamese},
		{"lẻn", Vietnamese},
		{"biểu", Vietnamese},
		{"kẽm", Vietnamese},
		{"diễm", Vietnamese},
		{"phế", Vietnamese},
		{"việc", Vietnamese},
		{"chỉnh", Vietnamese},
		{"trĩ", Vietnamese},
		{"ravị", Vietnamese},
		{"thơ", Vietnamese},
		{"nguồn", Vietnamese},
		{"thờ", Vietnamese},
		{"sỏi", Vietnamese},
		{"tổng", Vietnamese},
		{"nhở", Vietnamese},
		{"mỗi", Vietnamese},
		{"bỡi", Vietnamese},
		{"tốt", Vietnamese},
		{"giới", Vietnamese},
		{"một", Vietnamese},
		{"hợp", Vietnamese},
		{"hưng", Vietnamese},
		{"từng", Vietnamese},
		{"của", Vietnamese},
		{"sử", Vietnamese},
		{"cũng", Vietnamese},
		{"những", Vietnamese},
		{"chức", Vietnamese},
		{"dụng", Vietnamese},
		{"thực", Vietnamese},
		{"kỳ", Vietnamese},
		{"kỷ", Vietnamese},
		{"mỹ", Vietnamese},
		{"mỵ", Vietnamese},
		{"aṣiwèrè", Yoruba},
		{"ṣaaju", Yoruba},
		{"والموضوع", Unknown},
		{"сопротивление", Unknown},
		{"house", Unknown},

		// words with unique alphabet
		{"ունենա", Armenian},
		{"জানাতে", Bengali},
		{"გარეუბან", Georgian},
		{"σταμάτησε", Greek},
		{"ઉપકરણોની", Gujarati},
		{"בתחרויות", Hebrew},
		{"びさ", Japanese},
		{"대결구도가", Korean},
		{"ਮੋਟਰਸਾਈਕਲਾਂ", Punjabi},
		{"துன்பங்களை", Tamil},
		{"కృష్ణదేవరాయలు", Telugu},
		{"ในทางหลวงหมายเลข", Thai},
	}
	for _, testCase := range testCases {
		detectedLanguage := detectorForAllLanguages.detectLanguageWithRules([]string{testCase.word})
		message := fmt.Sprintf(
			"expected %v for word '%s', got %v",
			testCase.expectedLanguage,
			testCase.word,
			detectedLanguage,
		)
		assert.Equal(t, testCase.expectedLanguage, detectedLanguage, message)
	}
}

func TestFilterLanguagesByRules(t *testing.T) {
	testCases := []struct {
		word              string
		expectedLanguages []Language
	}{
		{"والموضوع", []Language{Arabic, Persian, Urdu}},
		{"сопротивление", []Language{
			Belarusian, Bulgarian, Kazakh, Macedonian, Mongolian, Russian, Serbian, Ukrainian},
		},
		{"раскрывае", []Language{Belarusian, Kazakh, Mongolian, Russian}},
		{"этот", []Language{Belarusian, Kazakh, Mongolian, Russian}},
		{"огнём", []Language{Belarusian, Kazakh, Mongolian, Russian}},
		{"плаваща", []Language{Bulgarian, Kazakh, Mongolian, Russian}},
		{"довършат", []Language{Bulgarian, Kazakh, Mongolian, Russian}},
		{"павінен", []Language{Belarusian, Kazakh, Ukrainian}},
		{"затоплување", []Language{Macedonian, Serbian}},
		{"ректасцензија", []Language{Macedonian, Serbian}},
		{"набљудувач", []Language{Macedonian, Serbian}},
		{"aizklātā", []Language{Latvian, Maori, Yoruba}},
		{"sistēmas", []Language{Latvian, Maori, Yoruba}},
		{"palīdzi", []Language{Latvian, Maori, Yoruba}},
		{"nhẹn", []Language{Vietnamese, Yoruba}},
		{"chọn", []Language{Vietnamese, Yoruba}},
		{"prihvaćanju", []Language{Bosnian, Croatian, Polish}},
		{"nađete", []Language{Bosnian, Croatian, Vietnamese}},
		{"visão", []Language{Portuguese, Vietnamese}},
		{"wystąpią", []Language{Lithuanian, Polish}},
		{"budowę", []Language{Lithuanian, Polish}},
		{"nebūsime", []Language{Latvian, Lithuanian, Maori, Yoruba}},
		{"afişate", []Language{Azerbaijani, Romanian, Turkish}},
		{"kradzieżami", []Language{Polish, Romanian}},
		{"înviat", []Language{French, Romanian}},
		{"venerdì", []Language{Italian, Vietnamese, Yoruba}},
		{"años", []Language{Basque, Spanish}},
		{"rozohňuje", []Language{Czech, Slovak}},
		{"rtuť", []Language{Czech, Slovak}},
		{"pregătire", []Language{Romanian, Vietnamese}},
		{"jeďte", []Language{Czech, Romanian, Slovak}},
		{"minjaverðir", []Language{Icelandic, Turkish}},
		{"þagnarskyldu", []Language{Icelandic, Turkish}},
		{"nebûtu", []Language{French, Hungarian}},
		{"hashemidëve", []Language{Afrikaans, Albanian, Dutch, French}},
		{"forêt", []Language{Afrikaans, French, Portuguese, Vietnamese}},
		{"succèdent", []Language{French, Italian, Vietnamese, Yoruba}},
		{"où", []Language{French, Italian, Vietnamese, Yoruba}},
		{"tõeliseks", []Language{Estonian, Hungarian, Portuguese, Vietnamese}},
		{"viòiem", []Language{Catalan, Italian, Vietnamese, Yoruba}},
		{"contrôle", []Language{French, Portuguese, Slovak, Vietnamese}},
		{"direktør", []Language{Bokmal, Danish, Nynorsk}},
		{"vývoj", []Language{Czech, Icelandic, Slovak, Turkish, Vietnamese}},
		{"päralt", []Language{Estonian, Finnish, German, Slovak, Swedish}},
		{"labâk", []Language{French, Portuguese, Romanian, Turkish, Vietnamese}},
		{"pràctiques", []Language{Catalan, French, Italian, Portuguese, Vietnamese}},
		{"überrascht", []Language{
			Azerbaijani, Catalan, Estonian, German, Hungarian, Spanish, Turkish},
		},
		{"indebærer", []Language{Bokmal, Danish, Icelandic, Nynorsk}},
		{"måned", []Language{Bokmal, Danish, Nynorsk, Swedish}},
		{"zaručen", []Language{Bosnian, Czech, Croatian, Latvian, Lithuanian, Slovak, Slovene}},
		{"zkouškou", []Language{Bosnian, Czech, Croatian, Latvian, Lithuanian, Slovak, Slovene}},
		{"navržen", []Language{Bosnian, Czech, Croatian, Latvian, Lithuanian, Slovak, Slovene}},
		{"façonnage", []Language{
			Albanian, Azerbaijani, Basque, Catalan, French, Portuguese, Turkish},
		},
		{"höher", []Language{
			Azerbaijani, Estonian, Finnish, German, Hungarian, Icelandic, Swedish, Turkish},
		},
		{"catedráticos", []Language{
			Catalan, Czech, Icelandic, Irish, Hungarian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
		},
		{"política", []Language{
			Catalan, Czech, Icelandic, Irish, Hungarian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
		},
		{"música", []Language{
			Catalan, Czech, Icelandic, Irish, Hungarian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
		},
		{"contradicció", []Language{
			Catalan, Hungarian, Icelandic, Irish, Polish, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
		},
		{"només", []Language{
			Catalan, Czech, French, Hungarian, Icelandic, Irish,
			Italian, Portuguese, Slovak, Spanish, Vietnamese, Yoruba},
		},
		{"house", []Language{
			Afrikaans, Albanian, Azerbaijani, Basque, Bokmal, Bosnian, Catalan, Croatian, Czech, Danish, Dutch, English,
			Esperanto, Estonian, Finnish, French, Ganda, German, Hungarian, Icelandic, Indonesian, Irish, Italian,
			Latin, Latvian, Lithuanian, Malay, Maori, Nynorsk, Polish, Portuguese, Romanian, Shona, Slovak, Slovene,
			Somali, Sotho, Spanish, Swahili, Swedish, Tagalog, Tsonga, Tswana, Turkish, Vietnamese, Welsh, Xhosa,
			Yoruba, Zulu},
		},
	}
	for _, testCase := range testCases {
		filteredLanguages := detectorForAllLanguages.filterLanguagesByRules([]string{testCase.word})
		message := fmt.Sprintf("expected %v for word '%s', got %v", testCase.expectedLanguages, testCase.word, filteredLanguages)
		assert.ElementsMatch(t, testCase.expectedLanguages, filteredLanguages, message)
	}
}

func TestDetectionOfInvalidStrings(t *testing.T) {
	testCases := []string{"", " \n  \t;", "3<856%)§"}
	for _, testCase := range testCases {
		language, exists := detectorForAllLanguages.DetectLanguageOf(testCase)
		assert.Equal(t, Unknown, language)
		assert.False(t, exists)
	}
}

func TestLanguageDetectionIsDeterministic(t *testing.T) {
	testCases := []struct {
		text      string
		languages []Language
	}{
		{
			"ام وی با نیکی میناج تیزر داشت؟؟؟؟؟؟ i vote for bts ( _ ) as the _ via ( _ )",
			[]Language{English, Urdu},
		},
		{
			"Az elmúlt hétvégén 12-re emelkedett az elhunyt koronavírus-fertőzöttek száma Szlovákiában. Mindegyik szociális otthon dolgozóját letesztelik, Matovič szerint az ingázóknak még várniuk kellene a teszteléssel",
			[]Language{Hungarian, Slovak},
		},
	}
	for _, testCase := range testCases {
		detector := NewLanguageDetectorBuilder().
			FromLanguages(testCase.languages...).
			WithPreloadedLanguageModels().
			Build()
		detectedLanguages := make(map[Language]bool)
		for i := 0; i < 100; i++ {
			language, _ := detector.DetectLanguageOf(testCase.text)
			detectedLanguages[language] = true
		}
		assert.Len(
			t,
			detectedLanguages,
			1,
			fmt.Sprintf("language detector is non-deterministic for languages %v", testCase.languages),
		)
	}
}

func TestLowAccuracyModeReportsUnknownLanguageForUnigramsAndBigrams(t *testing.T) {
	detector := NewLanguageDetectorBuilder().
		FromLanguages(English, German).
		WithPreloadedLanguageModels().
		WithLowAccuracyMode().
		Build()

	languageForTrigram, exists := detector.DetectLanguageOf("bed")
	assert.NotEqual(t, Unknown, languageForTrigram)
	assert.Equal(t, true, exists)

	languageForBigram, exists := detector.DetectLanguageOf("be")
	assert.Equal(t, Unknown, languageForBigram)
	assert.Equal(t, false, exists)

	languageForUnigram, exists := detector.DetectLanguageOf("b")
	assert.Equal(t, Unknown, languageForUnigram)
	assert.Equal(t, false, exists)

	languageForEmptyString, exists := detector.DetectLanguageOf("")
	assert.Equal(t, Unknown, languageForEmptyString)
	assert.Equal(t, false, exists)
}

func BenchmarkLanguageDetection(b *testing.B) {
	detector := NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()
	sentences := []string{
		"ربما يبتعد العقرب عن بعض الذين يخيبون أمله، أو يشعر بالحاجة إلى الانتقاء، وعدم البحث عن النشاطات التي ترهق أكثر مما تسعده.",
		"Επί της ουσίας τόσο οι υφιστάμενες ενισχύσεις που οφείλονται στους κτηνοτρόφους όσο και αυτές της νέας προγραμματικής περιόδου παραμένουν στον αέρα.",
		"It has three co-chairs, one from each of a provincial health and agriculture department, and a third from the federal government.",
		"અશ્વિની ભટ્ટની નવલકથામાંથી થોડુંક માણસ જ્યારે વેદનાની પરાકાષ્ટાની સીમા વટાવી જાય પછી એક એવી પરિસ્થિતિ આવે છે જ્યારે દર્દ-વેદના નથી રહેતી, વેદના છે કે નહિ તેનો પણ કોઇ ખ્યાલ નથી રહેતો.",
		"・京都大学施設に電離圏における電子数などの状況を取得可能なイオノゾンデ受信機（斜入射観測装置）を設置することで、新たな観測手法が地震先行現象検出に資するかを検証する。",
		"ამასთანავე წანარები სათავეში უდგებიან (თუ წარმართავენ?) კახეთის გაერთიანებისა და ერთიანი სამთავროს ჩამოყალიბების პროცესს.",
		"하지만 금융 전문가들은 “전체 대출 중 부동산 PF로의 쏠림 현상이 심각한 상태에서 각종 대출 규제로 자금 여력이 부족해질 경우 연체율이 높아질 수 있는데 당국이 안이하게 대응하는 측면이 있다”고 지적했다.",
		"И потому я должен возблагодарить провидение; если бы не провидение, то сердце твое, бедный сэр Пол, все конечно разбилось бы.",
		"ส.บัญชีรายชื่อ พรรคเพื่อไทย แต่อยู่ในระหว่างการตัดสินเรื่องการเป็นสมาชิกภาพของพรรคการเมือง เพราะถูกคุมขังโดยหมายศาล ระหว่างการสมัครรับเลือกตั้ง ซึ่งขณะนี้อยู่ในระหว่างการพิจารณาของ กกต.",
		"人们必须面对：遭受严重破坏的自然生态；大自然反扑所造成的天灾人祸；人口快速成长的沈重压力；生存竞争日异严峻的社会情况；传统家庭结构逐渐瓦解的隐忧，社会价值观念混淆等问题。",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, sentence := range sentences {
			detector.DetectLanguageOf(sentence)
		}
	}
}

func roundToTwoDecimalPlaces(value float64) float64 {
	return math.Round(value*100) / 100
}
