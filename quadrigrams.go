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
	"sync"
)

var quadrigramModels = map[Language]func() *trainingDataLanguageModel{
	Afrikaans:   loadAfrikaansQuadrigramModelFunc(),
	Albanian:    loadAlbanianQuadrigramModelFunc(),
	Arabic:      loadArabicQuadrigramModelFunc(),
	Armenian:    loadArmenianQuadrigramModelFunc(),
	Azerbaijani: loadAzerbaijaniQuadrigramModelFunc(),
	Basque:      loadBasqueQuadrigramModelFunc(),
	Belarusian:  loadBelarusianQuadrigramModelFunc(),
	Bengali:     loadBengaliQuadrigramModelFunc(),
	Bokmal:      loadBokmalQuadrigramModelFunc(),
	Bosnian:     loadBosnianQuadrigramModelFunc(),
	Bulgarian:   loadBulgarianQuadrigramModelFunc(),
	Catalan:     loadCatalanQuadrigramModelFunc(),
	Chinese:     loadChineseQuadrigramModelFunc(),
	Croatian:    loadCroatianQuadrigramModelFunc(),
	Czech:       loadCzechQuadrigramModelFunc(),
	Danish:      loadDanishQuadrigramModelFunc(),
	Dutch:       loadDutchQuadrigramModelFunc(),
	English:     loadEnglishQuadrigramModelFunc(),
	Esperanto:   loadEsperantoQuadrigramModelFunc(),
	Estonian:    loadEstonianQuadrigramModelFunc(),
	Finnish:     loadFinnishQuadrigramModelFunc(),
	French:      loadFrenchQuadrigramModelFunc(),
	Ganda:       loadGandaQuadrigramModelFunc(),
	Georgian:    loadGeorgianQuadrigramModelFunc(),
	German:      loadGermanQuadrigramModelFunc(),
	Greek:       loadGreekQuadrigramModelFunc(),
	Gujarati:    loadGujaratiQuadrigramModelFunc(),
	Hebrew:      loadHebrewQuadrigramModelFunc(),
	Hindi:       loadHindiQuadrigramModelFunc(),
	Hungarian:   loadHungarianQuadrigramModelFunc(),
	Icelandic:   loadIcelandicQuadrigramModelFunc(),
	Indonesian:  loadIndonesianQuadrigramModelFunc(),
	Irish:       loadIrishQuadrigramModelFunc(),
	Italian:     loadItalianQuadrigramModelFunc(),
	Japanese:    loadJapaneseQuadrigramModelFunc(),
	Kazakh:      loadKazakhQuadrigramModelFunc(),
	Korean:      loadKoreanQuadrigramModelFunc(),
	Latin:       loadLatinQuadrigramModelFunc(),
	Latvian:     loadLatvianQuadrigramModelFunc(),
	Lithuanian:  loadLithuanianQuadrigramModelFunc(),
	Macedonian:  loadMacedonianQuadrigramModelFunc(),
	Malay:       loadMalayQuadrigramModelFunc(),
	Maori:       loadMaoriQuadrigramModelFunc(),
	Marathi:     loadMarathiQuadrigramModelFunc(),
	Mongolian:   loadMongolianQuadrigramModelFunc(),
	Nynorsk:     loadNynorskQuadrigramModelFunc(),
	Persian:     loadPersianQuadrigramModelFunc(),
	Polish:      loadPolishQuadrigramModelFunc(),
	Portuguese:  loadPortugueseQuadrigramModelFunc(),
	Punjabi:     loadPunjabiQuadrigramModelFunc(),
	Romanian:    loadRomanianQuadrigramModelFunc(),
	Russian:     loadRussianQuadrigramModelFunc(),
	Serbian:     loadSerbianQuadrigramModelFunc(),
	Shona:       loadShonaQuadrigramModelFunc(),
	Slovak:      loadSlovakQuadrigramModelFunc(),
	Slovene:     loadSloveneQuadrigramModelFunc(),
	Somali:      loadSomaliQuadrigramModelFunc(),
	Sotho:       loadSothoQuadrigramModelFunc(),
	Spanish:     loadSpanishQuadrigramModelFunc(),
	Swahili:     loadSwahiliQuadrigramModelFunc(),
	Swedish:     loadSwedishQuadrigramModelFunc(),
	Tagalog:     loadTagalogQuadrigramModelFunc(),
	Tamil:       loadTamilQuadrigramModelFunc(),
	Telugu:      loadTeluguQuadrigramModelFunc(),
	Thai:        loadThaiQuadrigramModelFunc(),
	Tsonga:      loadTsongaQuadrigramModelFunc(),
	Tswana:      loadTswanaQuadrigramModelFunc(),
	Turkish:     loadTurkishQuadrigramModelFunc(),
	Ukrainian:   loadUkrainianQuadrigramModelFunc(),
	Urdu:        loadUrduQuadrigramModelFunc(),
	Vietnamese:  loadVietnameseQuadrigramModelFunc(),
	Welsh:       loadWelshQuadrigramModelFunc(),
	Xhosa:       loadXhosaQuadrigramModelFunc(),
	Yoruba:      loadYorubaQuadrigramModelFunc(),
	Zulu:        loadZuluQuadrigramModelFunc(),
}

func loadAfrikaansQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Afrikaans))
		})
		return &model
	}
}

func loadAlbanianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Albanian))
		})
		return &model
	}
}

func loadArabicQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Arabic))
		})
		return &model
	}
}

func loadArmenianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Armenian))
		})
		return &model
	}
}

func loadAzerbaijaniQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Azerbaijani))
		})
		return &model
	}
}

func loadBasqueQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Basque))
		})
		return &model
	}
}

func loadBelarusianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Belarusian))
		})
		return &model
	}
}

func loadBengaliQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bengali))
		})
		return &model
	}
}

func loadBokmalQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bokmal))
		})
		return &model
	}
}

func loadBosnianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bosnian))
		})
		return &model
	}
}

func loadBulgarianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bulgarian))
		})
		return &model
	}
}

func loadCatalanQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Catalan))
		})
		return &model
	}
}

func loadChineseQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Chinese))
		})
		return &model
	}
}

func loadCroatianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Croatian))
		})
		return &model
	}
}

func loadCzechQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Czech))
		})
		return &model
	}
}

func loadDanishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Danish))
		})
		return &model
	}
}

func loadDutchQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Dutch))
		})
		return &model
	}
}

func loadEnglishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(English))
		})
		return &model
	}
}

func loadEsperantoQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Esperanto))
		})
		return &model
	}
}

func loadEstonianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Estonian))
		})
		return &model
	}
}

func loadFinnishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Finnish))
		})
		return &model
	}
}

func loadFrenchQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(French))
		})
		return &model
	}
}

func loadGandaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Ganda))
		})
		return &model
	}
}

func loadGeorgianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Georgian))
		})
		return &model
	}
}

func loadGermanQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(German))
		})
		return &model
	}
}

func loadGreekQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Greek))
		})
		return &model
	}
}

func loadGujaratiQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Gujarati))
		})
		return &model
	}
}

func loadHebrewQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hebrew))
		})
		return &model
	}
}

func loadHindiQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hindi))
		})
		return &model
	}
}

func loadHungarianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hungarian))
		})
		return &model
	}
}

func loadIcelandicQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Icelandic))
		})
		return &model
	}
}

func loadIndonesianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Indonesian))
		})
		return &model
	}
}

func loadIrishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Irish))
		})
		return &model
	}
}

func loadItalianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Italian))
		})
		return &model
	}
}

func loadJapaneseQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Japanese))
		})
		return &model
	}
}

func loadKazakhQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Kazakh))
		})
		return &model
	}
}

func loadKoreanQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Korean))
		})
		return &model
	}
}

func loadLatinQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Latin))
		})
		return &model
	}
}

func loadLatvianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Latvian))
		})
		return &model
	}
}

func loadLithuanianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Lithuanian))
		})
		return &model
	}
}

func loadMacedonianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Macedonian))
		})
		return &model
	}
}

func loadMalayQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Malay))
		})
		return &model
	}
}

func loadMaoriQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Maori))
		})
		return &model
	}
}

func loadMarathiQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Marathi))
		})
		return &model
	}
}

func loadMongolianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Mongolian))
		})
		return &model
	}
}

func loadNynorskQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Nynorsk))
		})
		return &model
	}
}

func loadPersianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Persian))
		})
		return &model
	}
}

func loadPolishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Polish))
		})
		return &model
	}
}

func loadPortugueseQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Portuguese))
		})
		return &model
	}
}

func loadPunjabiQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Punjabi))
		})
		return &model
	}
}

func loadRomanianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Romanian))
		})
		return &model
	}
}

func loadRussianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Russian))
		})
		return &model
	}
}

func loadSerbianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Serbian))
		})
		return &model
	}
}

func loadShonaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Shona))
		})
		return &model
	}
}

func loadSlovakQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Slovak))
		})
		return &model
	}
}

func loadSloveneQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Slovene))
		})
		return &model
	}
}

func loadSomaliQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Somali))
		})
		return &model
	}
}

func loadSothoQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Sotho))
		})
		return &model
	}
}

func loadSpanishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Spanish))
		})
		return &model
	}
}

func loadSwahiliQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Swahili))
		})
		return &model
	}
}

func loadSwedishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Swedish))
		})
		return &model
	}
}

func loadTagalogQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tagalog))
		})
		return &model
	}
}

func loadTamilQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tamil))
		})
		return &model
	}
}

func loadTeluguQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Telugu))
		})
		return &model
	}
}

func loadThaiQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Thai))
		})
		return &model
	}
}

func loadTsongaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tsonga))
		})
		return &model
	}
}

func loadTswanaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tswana))
		})
		return &model
	}
}

func loadTurkishQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Turkish))
		})
		return &model
	}
}

func loadUkrainianQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Ukrainian))
		})
		return &model
	}
}

func loadUrduQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Urdu))
		})
		return &model
	}
}

func loadVietnameseQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Vietnamese))
		})
		return &model
	}
}

func loadWelshQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Welsh))
		})
		return &model
	}
}

func loadXhosaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Xhosa))
		})
		return &model
	}
}

func loadYorubaQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Yoruba))
		})
		return &model
	}
}

func loadZuluQuadrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Zulu))
		})
		return &model
	}
}

func loadQuadrigrams(language Language) []byte {
	return loadJson(language, 4)
}
