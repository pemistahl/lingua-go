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
	assert.ElementsMatch(t, AllLanguages(), builder.getLanguages())
	assert.Equal(t, 0.0, builder.getMinimumRelativeDistance())

	builder.WithMinimumRelativeDistance(0.2)
	assert.Equal(t, 0.2, builder.getMinimumRelativeDistance())
}

func TestLanguageDetectorBuilder_FromAllSpokenLanguages(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllSpokenLanguages()
	assert.ElementsMatch(t, AllSpokenLanguages(), builder.getLanguages())
	assert.Equal(t, 0.0, builder.getMinimumRelativeDistance())

	builder.WithMinimumRelativeDistance(0.2)
	assert.Equal(t, 0.2, builder.getMinimumRelativeDistance())
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithArabicScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithArabicScript()
	assert.ElementsMatch(t, AllLanguagesWithArabicScript(), builder.getLanguages())
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithCyrillicScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithCyrillicScript()
	assert.ElementsMatch(t, AllLanguagesWithCyrillicScript(), builder.getLanguages())
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithDevanagariScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithDevanagariScript()
	assert.ElementsMatch(t, AllLanguagesWithDevanagariScript(), builder.getLanguages())
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithLatinScript(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithLatinScript()
	assert.ElementsMatch(t, AllLanguagesWithLatinScript(), builder.getLanguages())
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithout(t *testing.T) {
	builder := NewLanguageDetectorBuilder().FromAllLanguagesWithout(Turkish, Romanian)
	assert.ElementsMatch(
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
			Maori,
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
		builder.getLanguages(),
	)
}

func TestLanguageDetectorBuilder_FromAllLanguagesWithout_Panics(t *testing.T) {
	assert.PanicsWithValue(
		t,
		missingLanguageMessage,
		func() {
			NewLanguageDetectorBuilder().FromAllLanguagesWithout(AllLanguages()[1:]...)
		},
	)
}

func TestLanguageDetectorBuilder_FromLanguages(t *testing.T) {
	testCases := []struct {
		languages         []Language
		expectedLanguages []Language
	}{
		{
			[]Language{German, English},
			[]Language{German, English},
		},
		{
			[]Language{German, English, Unknown},
			[]Language{German, English},
		},
	}
	for _, testCase := range testCases {
		builder := NewLanguageDetectorBuilder().FromLanguages(testCase.languages...)
		assert.ElementsMatch(t, testCase.expectedLanguages, builder.getLanguages())
	}
}

func TestLanguageDetectorBuilder_FromLanguages_Panics(t *testing.T) {
	testCases := []struct {
		languages []Language
	}{
		{[]Language{German}},
		{[]Language{German, Unknown}},
	}
	for _, testCase := range testCases {
		assert.PanicsWithValue(
			t,
			missingLanguageMessage,
			func() {
				NewLanguageDetectorBuilder().FromLanguages(testCase.languages...)
			},
		)
	}
}

func TestLanguageDetectorBuilder_FromIsoCodes639_1(t *testing.T) {
	testCases := []struct {
		isoCodes          []IsoCode639_1
		expectedLanguages []Language
	}{
		{
			[]IsoCode639_1{DE, EN},
			[]Language{German, English},
		},
		{
			[]IsoCode639_1{DE, EN, UnknownIsoCode639_1},
			[]Language{German, English},
		},
	}
	for _, testCase := range testCases {
		builder := NewLanguageDetectorBuilder().FromIsoCodes639_1(testCase.isoCodes...)
		assert.ElementsMatch(t, testCase.expectedLanguages, builder.getLanguages())
	}
}

func TestLanguageDetectorBuilder_FromIsoCodes639_1_Panics(t *testing.T) {
	testCases := []struct {
		isoCodes []IsoCode639_1
	}{
		{[]IsoCode639_1{DE}},
		{[]IsoCode639_1{DE, UnknownIsoCode639_1}},
	}
	for _, testCase := range testCases {
		assert.PanicsWithValue(
			t,
			missingLanguageMessage,
			func() {
				NewLanguageDetectorBuilder().FromIsoCodes639_1(testCase.isoCodes...)
			},
		)
	}
}

func TestLanguageDetectorBuilder_FromIsoCodes639_3(t *testing.T) {
	testCases := []struct {
		isoCodes          []IsoCode639_3
		expectedLanguages []Language
	}{
		{
			[]IsoCode639_3{DEU, ENG},
			[]Language{German, English},
		},
		{
			[]IsoCode639_3{DEU, ENG, UnknownIsoCode639_3},
			[]Language{German, English},
		},
	}
	for _, testCase := range testCases {
		builder := NewLanguageDetectorBuilder().FromIsoCodes639_3(testCase.isoCodes...)
		assert.ElementsMatch(t, testCase.expectedLanguages, builder.getLanguages())
	}
}

func TestLanguageDetectorBuilder_FromIsoCodes639_3_Panics(t *testing.T) {
	testCases := []struct {
		isoCodes []IsoCode639_3
	}{
		{[]IsoCode639_3{DEU}},
		{[]IsoCode639_3{DEU, UnknownIsoCode639_3}},
	}
	for _, testCase := range testCases {
		assert.PanicsWithValue(
			t,
			missingLanguageMessage,
			func() {
				NewLanguageDetectorBuilder().FromIsoCodes639_3(testCase.isoCodes...)
			},
		)
	}
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
