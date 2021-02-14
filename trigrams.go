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

var trigramModels = map[Language]func() *trainingDataLanguageModel{
	Afrikaans:   loadAfrikaansTrigramModelFunc(),
	Albanian:    loadAlbanianTrigramModelFunc(),
	Arabic:      loadArabicTrigramModelFunc(),
	Armenian:    loadArmenianTrigramModelFunc(),
	Azerbaijani: loadAzerbaijaniTrigramModelFunc(),
	Basque:      loadBasqueTrigramModelFunc(),
	Belarusian:  loadBelarusianTrigramModelFunc(),
	Bengali:     loadBengaliTrigramModelFunc(),
	Bokmal:      loadBokmalTrigramModelFunc(),
	Bosnian:     loadBosnianTrigramModelFunc(),
	Bulgarian:   loadBulgarianTrigramModelFunc(),
	Catalan:     loadCatalanTrigramModelFunc(),
	Chinese:     loadChineseTrigramModelFunc(),
	Croatian:    loadCroatianTrigramModelFunc(),
	Czech:       loadCzechTrigramModelFunc(),
	Danish:      loadDanishTrigramModelFunc(),
	Dutch:       loadDutchTrigramModelFunc(),
	English:     loadEnglishTrigramModelFunc(),
	Esperanto:   loadEsperantoTrigramModelFunc(),
	Estonian:    loadEstonianTrigramModelFunc(),
	Finnish:     loadFinnishTrigramModelFunc(),
	French:      loadFrenchTrigramModelFunc(),
	Ganda:       loadGandaTrigramModelFunc(),
	Georgian:    loadGeorgianTrigramModelFunc(),
	German:      loadGermanTrigramModelFunc(),
	Greek:       loadGreekTrigramModelFunc(),
	Gujarati:    loadGujaratiTrigramModelFunc(),
	Hebrew:      loadHebrewTrigramModelFunc(),
	Hindi:       loadHindiTrigramModelFunc(),
	Hungarian:   loadHungarianTrigramModelFunc(),
	Icelandic:   loadIcelandicTrigramModelFunc(),
	Indonesian:  loadIndonesianTrigramModelFunc(),
	Irish:       loadIrishTrigramModelFunc(),
	Italian:     loadItalianTrigramModelFunc(),
	Japanese:    loadJapaneseTrigramModelFunc(),
	Kazakh:      loadKazakhTrigramModelFunc(),
	Korean:      loadKoreanTrigramModelFunc(),
	Latin:       loadLatinTrigramModelFunc(),
	Latvian:     loadLatvianTrigramModelFunc(),
	Lithuanian:  loadLithuanianTrigramModelFunc(),
	Macedonian:  loadMacedonianTrigramModelFunc(),
	Malay:       loadMalayTrigramModelFunc(),
	Maori:       loadMaoriTrigramModelFunc(),
	Marathi:     loadMarathiTrigramModelFunc(),
	Mongolian:   loadMongolianTrigramModelFunc(),
	Nynorsk:     loadNynorskTrigramModelFunc(),
	Persian:     loadPersianTrigramModelFunc(),
	Polish:      loadPolishTrigramModelFunc(),
	Portuguese:  loadPortugueseTrigramModelFunc(),
	Punjabi:     loadPunjabiTrigramModelFunc(),
	Romanian:    loadRomanianTrigramModelFunc(),
	Russian:     loadRussianTrigramModelFunc(),
	Serbian:     loadSerbianTrigramModelFunc(),
	Shona:       loadShonaTrigramModelFunc(),
	Slovak:      loadSlovakTrigramModelFunc(),
	Slovene:     loadSloveneTrigramModelFunc(),
	Somali:      loadSomaliTrigramModelFunc(),
	Sotho:       loadSothoTrigramModelFunc(),
	Spanish:     loadSpanishTrigramModelFunc(),
	Swahili:     loadSwahiliTrigramModelFunc(),
	Swedish:     loadSwedishTrigramModelFunc(),
	Tagalog:     loadTagalogTrigramModelFunc(),
	Tamil:       loadTamilTrigramModelFunc(),
	Telugu:      loadTeluguTrigramModelFunc(),
	Thai:        loadThaiTrigramModelFunc(),
	Tsonga:      loadTsongaTrigramModelFunc(),
	Tswana:      loadTswanaTrigramModelFunc(),
	Turkish:     loadTurkishTrigramModelFunc(),
	Ukrainian:   loadUkrainianTrigramModelFunc(),
	Urdu:        loadUrduTrigramModelFunc(),
	Vietnamese:  loadVietnameseTrigramModelFunc(),
	Welsh:       loadWelshTrigramModelFunc(),
	Xhosa:       loadXhosaTrigramModelFunc(),
	Yoruba:      loadYorubaTrigramModelFunc(),
	Zulu:        loadZuluTrigramModelFunc(),
}

func loadAfrikaansTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Afrikaans))
		})
		return &model
	}
}

func loadAlbanianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Albanian))
		})
		return &model
	}
}

func loadArabicTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Arabic))
		})
		return &model
	}
}

func loadArmenianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Armenian))
		})
		return &model
	}
}

func loadAzerbaijaniTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Azerbaijani))
		})
		return &model
	}
}

func loadBasqueTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Basque))
		})
		return &model
	}
}

func loadBelarusianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Belarusian))
		})
		return &model
	}
}

func loadBengaliTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bengali))
		})
		return &model
	}
}

func loadBokmalTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bokmal))
		})
		return &model
	}
}

func loadBosnianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bosnian))
		})
		return &model
	}
}

func loadBulgarianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bulgarian))
		})
		return &model
	}
}

func loadCatalanTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Catalan))
		})
		return &model
	}
}

func loadChineseTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Chinese))
		})
		return &model
	}
}

func loadCroatianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Croatian))
		})
		return &model
	}
}

func loadCzechTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Czech))
		})
		return &model
	}
}

func loadDanishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Danish))
		})
		return &model
	}
}

func loadDutchTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Dutch))
		})
		return &model
	}
}

func loadEnglishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(English))
		})
		return &model
	}
}

func loadEsperantoTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Esperanto))
		})
		return &model
	}
}

func loadEstonianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Estonian))
		})
		return &model
	}
}

func loadFinnishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Finnish))
		})
		return &model
	}
}

func loadFrenchTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(French))
		})
		return &model
	}
}

func loadGandaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Ganda))
		})
		return &model
	}
}

func loadGeorgianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Georgian))
		})
		return &model
	}
}

func loadGermanTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(German))
		})
		return &model
	}
}

func loadGreekTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Greek))
		})
		return &model
	}
}

func loadGujaratiTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Gujarati))
		})
		return &model
	}
}

func loadHebrewTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hebrew))
		})
		return &model
	}
}

func loadHindiTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hindi))
		})
		return &model
	}
}

func loadHungarianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hungarian))
		})
		return &model
	}
}

func loadIcelandicTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Icelandic))
		})
		return &model
	}
}

func loadIndonesianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Indonesian))
		})
		return &model
	}
}

func loadIrishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Irish))
		})
		return &model
	}
}

func loadItalianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Italian))
		})
		return &model
	}
}

func loadJapaneseTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Japanese))
		})
		return &model
	}
}

func loadKazakhTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Kazakh))
		})
		return &model
	}
}

func loadKoreanTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Korean))
		})
		return &model
	}
}

func loadLatinTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Latin))
		})
		return &model
	}
}

func loadLatvianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Latvian))
		})
		return &model
	}
}

func loadLithuanianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Lithuanian))
		})
		return &model
	}
}

func loadMacedonianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Macedonian))
		})
		return &model
	}
}

func loadMalayTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Malay))
		})
		return &model
	}
}

func loadMaoriTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Maori))
		})
		return &model
	}
}

func loadMarathiTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Marathi))
		})
		return &model
	}
}

func loadMongolianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Mongolian))
		})
		return &model
	}
}

func loadNynorskTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Nynorsk))
		})
		return &model
	}
}

func loadPersianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Persian))
		})
		return &model
	}
}

func loadPolishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Polish))
		})
		return &model
	}
}

func loadPortugueseTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Portuguese))
		})
		return &model
	}
}

func loadPunjabiTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Punjabi))
		})
		return &model
	}
}

func loadRomanianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Romanian))
		})
		return &model
	}
}

func loadRussianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Russian))
		})
		return &model
	}
}

func loadSerbianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Serbian))
		})
		return &model
	}
}

func loadShonaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Shona))
		})
		return &model
	}
}

func loadSlovakTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Slovak))
		})
		return &model
	}
}

func loadSloveneTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Slovene))
		})
		return &model
	}
}

func loadSomaliTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Somali))
		})
		return &model
	}
}

func loadSothoTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Sotho))
		})
		return &model
	}
}

func loadSpanishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Spanish))
		})
		return &model
	}
}

func loadSwahiliTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Swahili))
		})
		return &model
	}
}

func loadSwedishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Swedish))
		})
		return &model
	}
}

func loadTagalogTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tagalog))
		})
		return &model
	}
}

func loadTamilTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tamil))
		})
		return &model
	}
}

func loadTeluguTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Telugu))
		})
		return &model
	}
}

func loadThaiTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Thai))
		})
		return &model
	}
}

func loadTsongaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tsonga))
		})
		return &model
	}
}

func loadTswanaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tswana))
		})
		return &model
	}
}

func loadTurkishTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Turkish))
		})
		return &model
	}
}

func loadUkrainianTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Ukrainian))
		})
		return &model
	}
}

func loadUrduTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Urdu))
		})
		return &model
	}
}

func loadVietnameseTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Vietnamese))
		})
		return &model
	}
}

func loadWelshTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Welsh))
		})
		return &model
	}
}

func loadXhosaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Xhosa))
		})
		return &model
	}
}

func loadYorubaTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Yoruba))
		})
		return &model
	}
}

func loadZuluTrigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Zulu))
		})
		return &model
	}
}

func loadTrigrams(language Language) []byte {
	return loadJson(language, 3)
}
