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

const missingLanguageMessage = "LanguageDetector needs at least 2 languages to choose from"

type UnconfiguredLanguageDetectorBuilder interface {
	FromAllLanguages() LanguageDetectorBuilder
	FromAllSpokenLanguages() LanguageDetectorBuilder
	FromAllLanguagesWithArabicScript() LanguageDetectorBuilder
	FromAllLanguagesWithCyrillicScript() LanguageDetectorBuilder
	FromAllLanguagesWithDevanagariScript() LanguageDetectorBuilder
	FromAllLanguagesWithLatinScript() LanguageDetectorBuilder
	FromAllLanguagesWithout(languages []Language) LanguageDetectorBuilder
	FromLanguages(languages []Language) LanguageDetectorBuilder
	FromIsoCodes639_1(isoCodes []IsoCode639_1) LanguageDetectorBuilder
	FromIsoCodes639_3(isoCodes []IsoCode639_3) LanguageDetectorBuilder
}

type LanguageDetectorBuilder interface {
	getLanguages() []Language
	getMinimumRelativeDistance() float64
	WithMinimumRelativeDistance(distance float64) LanguageDetectorBuilder
	Build() LanguageDetector
}

type languageDetectorBuilder struct {
	languages               []Language
	minimumRelativeDistance float64
}

func NewLanguageDetectorBuilder() UnconfiguredLanguageDetectorBuilder {
	return &languageDetectorBuilder{}
}

func (builder *languageDetectorBuilder) FromAllLanguages() LanguageDetectorBuilder {
	return builder.from(AllLanguages())
}

func (builder *languageDetectorBuilder) FromAllSpokenLanguages() LanguageDetectorBuilder {
	return builder.from(AllSpokenLanguages())
}

func (builder *languageDetectorBuilder) FromAllLanguagesWithArabicScript() LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithArabicScript())
}

func (builder *languageDetectorBuilder) FromAllLanguagesWithCyrillicScript() LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithCyrillicScript())
}

func (builder *languageDetectorBuilder) FromAllLanguagesWithDevanagariScript() LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithDevanagariScript())
}

func (builder *languageDetectorBuilder) FromAllLanguagesWithLatinScript() LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithLatinScript())
}

func (builder *languageDetectorBuilder) FromAllLanguagesWithout(languages []Language) LanguageDetectorBuilder {
	languagesToLoad := AllLanguages()
	for _, languageToRemove := range languages {
		for i, currentLanguage := range languagesToLoad {
			if currentLanguage == languageToRemove {
				languagesToLoad = append(languagesToLoad[:i], languagesToLoad[i+1:]...)
				break
			}
		}
	}
	if len(languagesToLoad) < 2 {
		panic(missingLanguageMessage)
	}
	return builder.from(languagesToLoad)
}

func (builder *languageDetectorBuilder) FromLanguages(languages []Language) LanguageDetectorBuilder {
	if len(languages) < 2 {
		panic(missingLanguageMessage)
	}
	return builder.from(languages)
}

func (builder *languageDetectorBuilder) FromIsoCodes639_1(isoCodes []IsoCode639_1) LanguageDetectorBuilder {
	if len(isoCodes) < 2 {
		panic(missingLanguageMessage)
	}
	var languages []Language
	for _, isoCode := range isoCodes {
		languages = append(languages, GetLanguageFromIsoCode639_1(isoCode))
	}
	return builder.from(languages)
}

func (builder *languageDetectorBuilder) FromIsoCodes639_3(isoCodes []IsoCode639_3) LanguageDetectorBuilder {
	if len(isoCodes) < 2 {
		panic(missingLanguageMessage)
	}
	var languages []Language
	for _, isoCode := range isoCodes {
		languages = append(languages, GetLanguageFromIsoCode639_3(isoCode))
	}
	return builder.from(languages)
}

func (builder *languageDetectorBuilder) WithMinimumRelativeDistance(distance float64) LanguageDetectorBuilder {
	if distance < 0.0 || distance > 0.99 {
		panic("Minimum relative distance must lie in between 0.0 and 0.99")
	}
	builder.minimumRelativeDistance = distance
	return builder
}

func (builder *languageDetectorBuilder) Build() LanguageDetector {
	return LanguageDetector{}
}

func (builder *languageDetectorBuilder) getLanguages() []Language {
	return builder.languages
}

func (builder *languageDetectorBuilder) getMinimumRelativeDistance() float64 {
	return builder.minimumRelativeDistance
}

func (builder *languageDetectorBuilder) from(languages []Language) LanguageDetectorBuilder {
	builder.languages = languages
	builder.minimumRelativeDistance = 0.0
	return builder
}
