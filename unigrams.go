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

var unigramModels = map[Language]func() *trainingDataLanguageModel{
	Afrikaans:   loadAfrikaansUnigramModelFunc(),
	Albanian:    loadAlbanianUnigramModelFunc(),
	Arabic:      loadArabicUnigramModelFunc(),
	Armenian:    loadArmenianUnigramModelFunc(),
	Azerbaijani: loadAzerbaijaniUnigramModelFunc(),
	Basque:      loadBasqueUnigramModelFunc(),
	Belarusian:  loadBelarusianUnigramModelFunc(),
	Bengali:     loadBengaliUnigramModelFunc(),
	Bokmal:      loadBokmalUnigramModelFunc(),
	Bosnian:     loadBosnianUnigramModelFunc(),
	Bulgarian:   loadBulgarianUnigramModelFunc(),
	Catalan:     loadCatalanUnigramModelFunc(),
	Chinese:     loadChineseUnigramModelFunc(),
	Croatian:    loadCroatianUnigramModelFunc(),
	Czech:       loadCzechUnigramModelFunc(),
	Danish:      loadDanishUnigramModelFunc(),
	Dutch:       loadDutchUnigramModelFunc(),
	English:     loadEnglishUnigramModelFunc(),
	Esperanto:   loadEsperantoUnigramModelFunc(),
	Estonian:    loadEstonianUnigramModelFunc(),
	Finnish:     loadFinnishUnigramModelFunc(),
	French:      loadFrenchUnigramModelFunc(),
	Ganda:       loadGandaUnigramModelFunc(),
	Georgian:    loadGeorgianUnigramModelFunc(),
	German:      loadGermanUnigramModelFunc(),
	Greek:       loadGreekUnigramModelFunc(),
	Gujarati:    loadGujaratiUnigramModelFunc(),
	Hebrew:      loadHebrewUnigramModelFunc(),
	Hindi:       loadHindiUnigramModelFunc(),
	Hungarian:   loadHungarianUnigramModelFunc(),
	Icelandic:   loadIcelandicUnigramModelFunc(),
	Indonesian:  loadIndonesianUnigramModelFunc(),
	Irish:       loadIrishUnigramModelFunc(),
	Italian:     loadItalianUnigramModelFunc(),
	Japanese:    loadJapaneseUnigramModelFunc(),
	Kazakh:      loadKazakhUnigramModelFunc(),
	Korean:      loadKoreanUnigramModelFunc(),
	Latin:       loadLatinUnigramModelFunc(),
	Latvian:     loadLatvianUnigramModelFunc(),
	Lithuanian:  loadLithuanianUnigramModelFunc(),
	Macedonian:  loadMacedonianUnigramModelFunc(),
	Malay:       loadMalayUnigramModelFunc(),
	Maori:       loadMaoriUnigramModelFunc(),
	Marathi:     loadMarathiUnigramModelFunc(),
	Mongolian:   loadMongolianUnigramModelFunc(),
	Nynorsk:     loadNynorskUnigramModelFunc(),
	Persian:     loadPersianUnigramModelFunc(),
	Polish:      loadPolishUnigramModelFunc(),
	Portuguese:  loadPortugueseUnigramModelFunc(),
	Punjabi:     loadPunjabiUnigramModelFunc(),
	Romanian:    loadRomanianUnigramModelFunc(),
	Russian:     loadRussianUnigramModelFunc(),
	Serbian:     loadSerbianUnigramModelFunc(),
	Shona:       loadShonaUnigramModelFunc(),
	Slovak:      loadSlovakUnigramModelFunc(),
	Slovene:     loadSloveneUnigramModelFunc(),
	Somali:      loadSomaliUnigramModelFunc(),
	Sotho:       loadSothoUnigramModelFunc(),
	Spanish:     loadSpanishUnigramModelFunc(),
	Swahili:     loadSwahiliUnigramModelFunc(),
	Swedish:     loadSwedishUnigramModelFunc(),
	Tagalog:     loadTagalogUnigramModelFunc(),
	Tamil:       loadTamilUnigramModelFunc(),
	Telugu:      loadTeluguUnigramModelFunc(),
	Thai:        loadThaiUnigramModelFunc(),
	Tsonga:      loadTsongaUnigramModelFunc(),
	Tswana:      loadTswanaUnigramModelFunc(),
	Turkish:     loadTurkishUnigramModelFunc(),
	Ukrainian:   loadUkrainianUnigramModelFunc(),
	Urdu:        loadUrduUnigramModelFunc(),
	Vietnamese:  loadVietnameseUnigramModelFunc(),
	Welsh:       loadWelshUnigramModelFunc(),
	Xhosa:       loadXhosaUnigramModelFunc(),
	Yoruba:      loadYorubaUnigramModelFunc(),
	Zulu:        loadZuluUnigramModelFunc(),
}

func loadAfrikaansUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Afrikaans))
		})
		return &model
	}
}

func loadAlbanianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Albanian))
		})
		return &model
	}
}

func loadArabicUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Arabic))
		})
		return &model
	}
}

func loadArmenianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Armenian))
		})
		return &model
	}
}

func loadAzerbaijaniUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Azerbaijani))
		})
		return &model
	}
}

func loadBasqueUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Basque))
		})
		return &model
	}
}

func loadBelarusianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Belarusian))
		})
		return &model
	}
}

func loadBengaliUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bengali))
		})
		return &model
	}
}

func loadBokmalUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bokmal))
		})
		return &model
	}
}

func loadBosnianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bosnian))
		})
		return &model
	}
}

func loadBulgarianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bulgarian))
		})
		return &model
	}
}

func loadCatalanUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Catalan))
		})
		return &model
	}
}

func loadChineseUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Chinese))
		})
		return &model
	}
}

func loadCroatianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Croatian))
		})
		return &model
	}
}

func loadCzechUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Czech))
		})
		return &model
	}
}

func loadDanishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Danish))
		})
		return &model
	}
}

func loadDutchUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Dutch))
		})
		return &model
	}
}

func loadEnglishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(English))
		})
		return &model
	}
}

func loadEsperantoUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Esperanto))
		})
		return &model
	}
}

func loadEstonianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Estonian))
		})
		return &model
	}
}

func loadFinnishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Finnish))
		})
		return &model
	}
}

func loadFrenchUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(French))
		})
		return &model
	}
}

func loadGandaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Ganda))
		})
		return &model
	}
}

func loadGeorgianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Georgian))
		})
		return &model
	}
}

func loadGermanUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(German))
		})
		return &model
	}
}

func loadGreekUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Greek))
		})
		return &model
	}
}

func loadGujaratiUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Gujarati))
		})
		return &model
	}
}

func loadHebrewUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hebrew))
		})
		return &model
	}
}

func loadHindiUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hindi))
		})
		return &model
	}
}

func loadHungarianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hungarian))
		})
		return &model
	}
}

func loadIcelandicUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Icelandic))
		})
		return &model
	}
}

func loadIndonesianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Indonesian))
		})
		return &model
	}
}

func loadIrishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Irish))
		})
		return &model
	}
}

func loadItalianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Italian))
		})
		return &model
	}
}

func loadJapaneseUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Japanese))
		})
		return &model
	}
}

func loadKazakhUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Kazakh))
		})
		return &model
	}
}

func loadKoreanUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Korean))
		})
		return &model
	}
}

func loadLatinUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Latin))
		})
		return &model
	}
}

func loadLatvianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Latvian))
		})
		return &model
	}
}

func loadLithuanianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Lithuanian))
		})
		return &model
	}
}

func loadMacedonianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Macedonian))
		})
		return &model
	}
}

func loadMalayUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Malay))
		})
		return &model
	}
}

func loadMaoriUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Maori))
		})
		return &model
	}
}

func loadMarathiUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Marathi))
		})
		return &model
	}
}

func loadMongolianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Mongolian))
		})
		return &model
	}
}

func loadNynorskUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Nynorsk))
		})
		return &model
	}
}

func loadPersianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Persian))
		})
		return &model
	}
}

func loadPolishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Polish))
		})
		return &model
	}
}

func loadPortugueseUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Portuguese))
		})
		return &model
	}
}

func loadPunjabiUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Punjabi))
		})
		return &model
	}
}

func loadRomanianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Romanian))
		})
		return &model
	}
}

func loadRussianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Russian))
		})
		return &model
	}
}

func loadSerbianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Serbian))
		})
		return &model
	}
}

func loadShonaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Shona))
		})
		return &model
	}
}

func loadSlovakUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Slovak))
		})
		return &model
	}
}

func loadSloveneUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Slovene))
		})
		return &model
	}
}

func loadSomaliUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Somali))
		})
		return &model
	}
}

func loadSothoUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Sotho))
		})
		return &model
	}
}

func loadSpanishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Spanish))
		})
		return &model
	}
}

func loadSwahiliUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Swahili))
		})
		return &model
	}
}

func loadSwedishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Swedish))
		})
		return &model
	}
}

func loadTagalogUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tagalog))
		})
		return &model
	}
}

func loadTamilUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tamil))
		})
		return &model
	}
}

func loadTeluguUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Telugu))
		})
		return &model
	}
}

func loadThaiUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Thai))
		})
		return &model
	}
}

func loadTsongaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tsonga))
		})
		return &model
	}
}

func loadTswanaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tswana))
		})
		return &model
	}
}

func loadTurkishUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Turkish))
		})
		return &model
	}
}

func loadUkrainianUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Ukrainian))
		})
		return &model
	}
}

func loadUrduUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Urdu))
		})
		return &model
	}
}

func loadVietnameseUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Vietnamese))
		})
		return &model
	}
}

func loadWelshUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Welsh))
		})
		return &model
	}
}

func loadXhosaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Xhosa))
		})
		return &model
	}
}

func loadYorubaUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Yoruba))
		})
		return &model
	}
}

func loadZuluUnigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Zulu))
		})
		return &model
	}
}

func loadUnigrams(language Language) []byte {
	return loadJson(language, 1)
}
