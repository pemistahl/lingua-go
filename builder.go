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

type LanguageDetectorBuilder struct {
	languages               []Language
	minimumRelativeDistance float64
}

func NewLanguageDetectorBuilder() *LanguageDetectorBuilder {
	return &LanguageDetectorBuilder{}
}

func (builder *LanguageDetectorBuilder) FromAllLanguages() *LanguageDetectorBuilder {
	return builder.from(AllLanguages())
}

func (builder *LanguageDetectorBuilder) FromAllSpokenLanguages() *LanguageDetectorBuilder {
	return builder.from(AllSpokenLanguages())
}

func (builder *LanguageDetectorBuilder) FromAllLanguagesWithArabicScript() *LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithArabicScript())
}

func (builder *LanguageDetectorBuilder) FromAllLanguagesWithCyrillicScript() *LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithCyrillicScript())
}

func (builder *LanguageDetectorBuilder) FromAllLanguagesWithDevanagariScript() *LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithDevanagariScript())
}

func (builder *LanguageDetectorBuilder) FromAllLanguagesWithLatinScript() *LanguageDetectorBuilder {
	return builder.from(AllLanguagesWithLatinScript())
}

func (builder *LanguageDetectorBuilder) FromAllLanguagesWithout(languages []Language) *LanguageDetectorBuilder {
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

func (builder *LanguageDetectorBuilder) FromLanguages(languages []Language) *LanguageDetectorBuilder {
	if len(languages) < 2 {
		panic(missingLanguageMessage)
	}
	return builder.from(languages)
}

func (builder *LanguageDetectorBuilder) FromIsoCodes639_1(isoCodes []IsoCode639_1) *LanguageDetectorBuilder {
	if len(isoCodes) < 2 {
		panic(missingLanguageMessage)
	}
	var languages []Language
	for _, isoCode := range isoCodes {
		languages = append(languages, GetLanguageFromIsoCode639_1(isoCode))
	}
	return builder.from(languages)
}

func (builder *LanguageDetectorBuilder) FromIsoCodes639_3(isoCodes []IsoCode639_3) *LanguageDetectorBuilder {
	if len(isoCodes) < 2 {
		panic(missingLanguageMessage)
	}
	var languages []Language
	for _, isoCode := range isoCodes {
		languages = append(languages, GetLanguageFromIsoCode639_3(isoCode))
	}
	return builder.from(languages)
}

func (builder *LanguageDetectorBuilder) WithMinimumRelativeDistance(distance float64) *LanguageDetectorBuilder {
	if distance < 0.0 || distance > 0.99 {
		panic("Minimum relative distance must lie in between 0.0 and 0.99")
	}
	builder.minimumRelativeDistance = distance
	return builder
}

func (builder *LanguageDetectorBuilder) Build() LanguageDetector {
	if len(builder.languages) == 0 {
		panic("LanguageDetector cannot be built as no languages have been specified")
	}
	return LanguageDetector{}
}

func (builder *LanguageDetectorBuilder) from(languages []Language) *LanguageDetectorBuilder {
	builder.languages = languages
	builder.minimumRelativeDistance = 0.0
	return builder
}
