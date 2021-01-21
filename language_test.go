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
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllLanguages(t *testing.T) {
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
			Romanian,
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
			Turkish,
			Ukrainian,
			Urdu,
			Vietnamese,
			Welsh,
			Xhosa,
			Yoruba,
			Zulu,
		},
		AllLanguages())
}

func TestAllSpokenLanguages(t *testing.T) {
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
			Romanian,
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
			Turkish,
			Ukrainian,
			Urdu,
			Vietnamese,
			Welsh,
			Xhosa,
			Yoruba,
			Zulu,
		},
		AllSpokenLanguages())
}

func TestAllLanguagesWithArabicScript(t *testing.T) {
	assert.Equal(t, []Language{Arabic, Persian, Urdu}, AllLanguagesWithArabicScript())
}

func TestAllLanguagesWithCyrillicScript(t *testing.T) {
	assert.Equal(
		t,
		[]Language{
			Belarusian, Bulgarian, Kazakh, Macedonian, Mongolian, Russian, Serbian, Ukrainian,
		},
		AllLanguagesWithCyrillicScript())
}

func TestAllLanguagesWithDevanagariScript(t *testing.T) {
	assert.Equal(t, []Language{Hindi, Marathi}, AllLanguagesWithDevanagariScript())
}

func TestAllLanguagesWithLatinScript(t *testing.T) {
	assert.Equal(
		t,
		[]Language{
			Afrikaans,
			Albanian,
			Azerbaijani,
			Basque,
			Bokmal,
			Bosnian,
			Catalan,
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
			German,
			Hungarian,
			Icelandic,
			Indonesian,
			Irish,
			Italian,
			Latin,
			Latvian,
			Lithuanian,
			Malay,
			Nynorsk,
			Polish,
			Portuguese,
			Romanian,
			Shona,
			Slovak,
			Slovene,
			Somali,
			Sotho,
			Spanish,
			Swahili,
			Swedish,
			Tagalog,
			Tsonga,
			Tswana,
			Turkish,
			Vietnamese,
			Welsh,
			Xhosa,
			Yoruba,
			Zulu,
		},
		AllLanguagesWithLatinScript())
}

func TestLanguage_MarshalJSON(t *testing.T) {
	language, err := json.Marshal(German)
	assert.Equal(t, "\"GERMAN\"", string(language))
	assert.Equal(t, nil, err)
}

func TestLanguage_UnmarshalJSON(t *testing.T) {
	var language Language
	err := json.Unmarshal([]byte("\"GERMAN\""), &language)
	assert.Equal(t, German, language)
	assert.Equal(t, nil, err)

	err = json.Unmarshal([]byte("\"GERM\""), &language)
	assert.Equal(t, fmt.Errorf("string \"GERM\" cannot be unmarshalled to an instance of type Language"), err)
}
