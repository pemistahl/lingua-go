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

package lingua_test

import (
	"fmt"
	"github.com/pemistahl/lingua-go"
)

func Example_basic() {
	languages := []lingua.Language{
		lingua.English,
		lingua.French,
		lingua.German,
		lingua.Spanish,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf("languages are awesome"); exists {
		fmt.Println(language)
	}

	// Output: English
}

// By default, Lingua returns the most likely language for a given input text.
// However, there are certain words that are spelled the same in more than one
// language. The word `prologue`, for instance, is both a valid English and
// French word. Lingua would output either English or French which might be
// wrong in the given context. For cases like that, it is possible to specify a
// minimum relative distance that the logarithmized and summed up probabilities
// for each possible language have to satisfy. It can be stated as seen below.
//
// Be aware that the distance between the language probabilities is dependent on
// the length of the input text. The longer the input text, the larger the
// distance between the languages. So if you want to classify very short text
// phrases, do not set the minimum relative distance too high. Otherwise Unknown
// will be returned most of the time as in the example below. This is the return
// value for cases where language detection is not reliably possible.
func Example_minimumRelativeDistance() {
	languages := []lingua.Language{
		lingua.English,
		lingua.French,
		lingua.German,
		lingua.Spanish,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		WithMinimumRelativeDistance(0.7).
		Build()

	language, exists := detector.DetectLanguageOf("languages are awesome")

	fmt.Println(language)
	fmt.Println(exists)

	// Output:
	// Unknown
	// false
}

// Knowing about the most likely language is nice but how reliable is the
// computed likelihood? And how less likely are the other examined languages in
// comparison to the most likely one? In the example below, a slice of
// ConfidenceValue is returned containing those languages which the calling
// instance of LanguageDetector has been built from. The entries are sorted by
// their confidence value in descending order. The values that this method
// computes are part of a relative confidence metric, not of an absolute one.
// Each value is a number between 0.0 and 1.0.
//
// If the language is unambiguously identified by the rule engine, the value
// 1.0 will always be returned for this language. The other languages will
// receive a value of 0.0. If the statistics engine is additionally needed,
// the most likely language will be returned with value 0.99 and the least
// likely language will be returned with value 0.01. All other languages get
// values assigned between 0.01 and 0.99, denoting how less likely those
// languages are in comparison to the most likely language.
func Example_confidenceValues() {
	languages := []lingua.Language{
		lingua.English,
		lingua.French,
		lingua.German,
		lingua.Spanish,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	confidenceValues := detector.ComputeLanguageConfidenceValues("languages are awesome")

	for _, elem := range confidenceValues {
		fmt.Printf("%s: %.2f\n", elem.Language(), elem.Value())
	}

	// Output:
	// English: 0.99
	// French: 0.32
	// German: 0.15
	// Spanish: 0.01
}

// By default, Lingua uses lazy-loading to load only those language models on
// demand which are considered relevant by the rule-based filter engine. For web
// services, for instance, it is rather beneficial to preload all language models
// into memory to avoid unexpected latency while waiting for the service response.
// If you want to enable the eager-loading mode, you can do it as seen below.
// Multiple instances of LanguageDetector share the same language models in
// memory which are accessed asynchronously by the instances.
func Example_eagerLoading() {
	lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		WithPreloadedLanguageModels().
		Build()
}

// There might be classification tasks where you know beforehand that your language
// data is definitely not written in Latin, for instance. The detection accuracy
// can become better in such cases if you exclude certain languages from the
// decision process or just explicitly include relevant languages.
func Example_builderApi() {
	// Including all languages available in the library
	// consumes at least 2GB of memory and might
	// lead to slow runtime performance.
	lingua.NewLanguageDetectorBuilder().FromAllLanguages()

	// Include only languages that are not yet extinct
	// (= currently excludes Latin).
	lingua.NewLanguageDetectorBuilder().FromAllSpokenLanguages()

	// Include only languages written with Cyrillic script.
	lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithCyrillicScript()

	// Exclude only the Spanish language from the decision algorithm.
	lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithout(lingua.Spanish)

	// Only decide between English and German.
	lingua.NewLanguageDetectorBuilder().FromLanguages(lingua.English, lingua.German)

	// Select languages by ISO 639-1 code.
	lingua.NewLanguageDetectorBuilder().FromIsoCodes639_1(lingua.EN, lingua.DE)

	// Select languages by ISO 639-3 code.
	lingua.NewLanguageDetectorBuilder().FromIsoCodes639_3(lingua.ENG, lingua.DEU)
}
