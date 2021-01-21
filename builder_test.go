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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguageDetectorBuilder_FromAllLanguages(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguages()
	assert.Equal(t, AllLanguages(), builder.languages)
	assert.Equal(t, 0.0, builder.minimumRelativeDistance)

	builder.WithMinimumRelativeDistance(0.2)
	assert.Equal(t, 0.2, builder.minimumRelativeDistance)
}

func TestLanguageDetectorBuilder_FromAllSpokenLanguages(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllSpokenLanguages()
	assert.Equal(t, AllSpokenLanguages(), builder.languages)
	assert.Equal(t, 0.0, builder.minimumRelativeDistance)

	builder.WithMinimumRelativeDistance(0.2)
	assert.Equal(t, 0.2, builder.minimumRelativeDistance)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithArabicScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithArabicScript()
	assert.Equal(t, AllLanguagesWithArabicScript(), builder.languages)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithCyrillicScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithCyrillicScript()
	assert.Equal(t, AllLanguagesWithCyrillicScript(), builder.languages)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithDevanagariScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithDevanagariScript()
	assert.Equal(t, AllLanguagesWithDevanagariScript(), builder.languages)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithLatinScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithLatinScript()
	assert.Equal(t, AllLanguagesWithLatinScript(), builder.languages)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithout(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithout([]Language{Turkish, Romanian})
	assert.Equal(
		t,
		[]Language{
			Afrikaans,
			Albanian,
			Arabic,
			Armenian,
			Azerbaijani,
			Basque,
			Belarusian,
			Bengali,
			Bokmal,
			Bosnian,
			Bulgarian,
			Catalan,
			Chinese,
			Croatian,
			Czech,
			Danish,
			Dutch,
			English,
			Esperanto,
			Estonian,
			Finnish,
			French,
			Ganda,
			Georgian,
			German,
			Greek,
			Gujarati,
			Hebrew,
			Hindi,
			Hungarian,
			Icelandic,
			Indonesian,
			Irish,
			Italian,
			Japanese,
			Kazakh,
			Korean,
			Latin,
			Latvian,
			Lithuanian,
			Macedonian,
			Malay,
			Marathi,
			Mongolian,
			Nynorsk,
			Persian,
			Polish,
			Portuguese,
			Punjabi,
			Russian,
			Serbian,
			Shona,
			Slovak,
			Slovene,
			Somali,
			Sotho,
			Spanish,
			Swahili,
			Swedish,
			Tagalog,
			Tamil,
			Telugu,
			Thai,
			Tsonga,
			Tswana,
			Ukrainian,
			Urdu,
			Vietnamese,
			Welsh,
			Xhosa,
			Yoruba,
			Zulu,
		},
		builder.languages,
	)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithout_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"LanguageDetector needs at least 2 languages to choose from",
		func() {
			NewLanguageDetectorBuilder().FromAllLanguagesWithout(AllLanguages()[1:])
		},
	)
}

func TestLanguageDetectorBuilder_FromLanguages(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromLanguages([]Language{German, English})
	assert.Equal(t, []Language{German, English}, builder.languages)
}

func TestLanguageDetectorBuilder_FromLanguages_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"LanguageDetector needs at least 2 languages to choose from",
		func() {
			NewLanguageDetectorBuilder().FromLanguages([]Language{German})
		},
	)
}

func TestLanguageDetectorBuilder_FromIsoCodes639_1(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromIsoCodes639_1([]IsoCode639_1{DE, SV})
	assert.Equal(t, []Language{German, Swedish}, builder.languages)
}

func TestLanguageDetectorBuilder_FromIsoCodes639_1_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"LanguageDetector needs at least 2 languages to choose from",
		func() {
			NewLanguageDetectorBuilder().FromIsoCodes639_1([]IsoCode639_1{DE})
		},
	)
}

func TestLanguageDetectorBuilder_FromIsoCodes639_3(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromIsoCodes639_3([]IsoCode639_3{DEU, SWE})
	assert.Equal(t, []Language{German, Swedish}, builder.languages)
}

func TestLanguageDetectorBuilder_FromIsoCodes639_3_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"LanguageDetector needs at least 2 languages to choose from",
		func() {
			NewLanguageDetectorBuilder().FromIsoCodes639_3([]IsoCode639_3{DEU})
		},
	)
}

func TestLanguageDetectorBuilder_WithMinimumRelativeDistance_Panics_1(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"Minimum relative distance must lie in between 0.0 and 0.99",
		func() {
			NewLanguageDetectorBuilder().FromAllLanguages().WithMinimumRelativeDistance(-2.3)
		},
	)
}

func TestLanguageDetectorBuilder_WithMinimumRelativeDistance_Panics_2(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"Minimum relative distance must lie in between 0.0 and 0.99",
		func() {
			NewLanguageDetectorBuilder().FromAllLanguages().WithMinimumRelativeDistance(1.7)
		},
	)
}

func TestLanguageDetectorBuilder_Build_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		"LanguageDetector cannot be built as no languages have been specified",
		func() {
			NewLanguageDetectorBuilder().Build()
		},
	)
}
