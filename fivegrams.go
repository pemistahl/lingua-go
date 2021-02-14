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

var fivegramModels = map[Language]func() *trainingDataLanguageModel{
	Afrikaans:   loadAfrikaansFivegramModelFunc(),
	Albanian:    loadAlbanianFivegramModelFunc(),
	Arabic:      loadArabicFivegramModelFunc(),
	Armenian:    loadArmenianFivegramModelFunc(),
	Azerbaijani: loadAzerbaijaniFivegramModelFunc(),
	Basque:      loadBasqueFivegramModelFunc(),
	Belarusian:  loadBelarusianFivegramModelFunc(),
	Bengali:     loadBengaliFivegramModelFunc(),
	Bokmal:      loadBokmalFivegramModelFunc(),
	Bosnian:     loadBosnianFivegramModelFunc(),
	Bulgarian:   loadBulgarianFivegramModelFunc(),
	Catalan:     loadCatalanFivegramModelFunc(),
	Chinese:     loadChineseFivegramModelFunc(),
	Croatian:    loadCroatianFivegramModelFunc(),
	Czech:       loadCzechFivegramModelFunc(),
	Danish:      loadDanishFivegramModelFunc(),
	Dutch:       loadDutchFivegramModelFunc(),
	English:     loadEnglishFivegramModelFunc(),
	Esperanto:   loadEsperantoFivegramModelFunc(),
	Estonian:    loadEstonianFivegramModelFunc(),
	Finnish:     loadFinnishFivegramModelFunc(),
	French:      loadFrenchFivegramModelFunc(),
	Ganda:       loadGandaFivegramModelFunc(),
	Georgian:    loadGeorgianFivegramModelFunc(),
	German:      loadGermanFivegramModelFunc(),
	Greek:       loadGreekFivegramModelFunc(),
	Gujarati:    loadGujaratiFivegramModelFunc(),
	Hebrew:      loadHebrewFivegramModelFunc(),
	Hindi:       loadHindiFivegramModelFunc(),
	Hungarian:   loadHungarianFivegramModelFunc(),
	Icelandic:   loadIcelandicFivegramModelFunc(),
	Indonesian:  loadIndonesianFivegramModelFunc(),
	Irish:       loadIrishFivegramModelFunc(),
	Italian:     loadItalianFivegramModelFunc(),
	Japanese:    loadJapaneseFivegramModelFunc(),
	Kazakh:      loadKazakhFivegramModelFunc(),
	Korean:      loadKoreanFivegramModelFunc(),
	Latin:       loadLatinFivegramModelFunc(),
	Latvian:     loadLatvianFivegramModelFunc(),
	Lithuanian:  loadLithuanianFivegramModelFunc(),
	Macedonian:  loadMacedonianFivegramModelFunc(),
	Malay:       loadMalayFivegramModelFunc(),
	Maori:       loadMaoriFivegramModelFunc(),
	Marathi:     loadMarathiFivegramModelFunc(),
	Mongolian:   loadMongolianFivegramModelFunc(),
	Nynorsk:     loadNynorskFivegramModelFunc(),
	Persian:     loadPersianFivegramModelFunc(),
	Polish:      loadPolishFivegramModelFunc(),
	Portuguese:  loadPortugueseFivegramModelFunc(),
	Punjabi:     loadPunjabiFivegramModelFunc(),
	Romanian:    loadRomanianFivegramModelFunc(),
	Russian:     loadRussianFivegramModelFunc(),
	Serbian:     loadSerbianFivegramModelFunc(),
	Shona:       loadShonaFivegramModelFunc(),
	Slovak:      loadSlovakFivegramModelFunc(),
	Slovene:     loadSloveneFivegramModelFunc(),
	Somali:      loadSomaliFivegramModelFunc(),
	Sotho:       loadSothoFivegramModelFunc(),
	Spanish:     loadSpanishFivegramModelFunc(),
	Swahili:     loadSwahiliFivegramModelFunc(),
	Swedish:     loadSwedishFivegramModelFunc(),
	Tagalog:     loadTagalogFivegramModelFunc(),
	Tamil:       loadTamilFivegramModelFunc(),
	Telugu:      loadTeluguFivegramModelFunc(),
	Thai:        loadThaiFivegramModelFunc(),
	Tsonga:      loadTsongaFivegramModelFunc(),
	Tswana:      loadTswanaFivegramModelFunc(),
	Turkish:     loadTurkishFivegramModelFunc(),
	Ukrainian:   loadUkrainianFivegramModelFunc(),
	Urdu:        loadUrduFivegramModelFunc(),
	Vietnamese:  loadVietnameseFivegramModelFunc(),
	Welsh:       loadWelshFivegramModelFunc(),
	Xhosa:       loadXhosaFivegramModelFunc(),
	Yoruba:      loadYorubaFivegramModelFunc(),
	Zulu:        loadZuluFivegramModelFunc(),
}

func loadAfrikaansFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Afrikaans))
		})
		return &model
	}
}

func loadAlbanianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Albanian))
		})
		return &model
	}
}

func loadArabicFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Arabic))
		})
		return &model
	}
}

func loadArmenianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Armenian))
		})
		return &model
	}
}

func loadAzerbaijaniFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Azerbaijani))
		})
		return &model
	}
}

func loadBasqueFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Basque))
		})
		return &model
	}
}

func loadBelarusianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Belarusian))
		})
		return &model
	}
}

func loadBengaliFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bengali))
		})
		return &model
	}
}

func loadBokmalFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bokmal))
		})
		return &model
	}
}

func loadBosnianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bosnian))
		})
		return &model
	}
}

func loadBulgarianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bulgarian))
		})
		return &model
	}
}

func loadCatalanFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Catalan))
		})
		return &model
	}
}

func loadChineseFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Chinese))
		})
		return &model
	}
}

func loadCroatianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Croatian))
		})
		return &model
	}
}

func loadCzechFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Czech))
		})
		return &model
	}
}

func loadDanishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Danish))
		})
		return &model
	}
}

func loadDutchFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Dutch))
		})
		return &model
	}
}

func loadEnglishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(English))
		})
		return &model
	}
}

func loadEsperantoFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Esperanto))
		})
		return &model
	}
}

func loadEstonianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Estonian))
		})
		return &model
	}
}

func loadFinnishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Finnish))
		})
		return &model
	}
}

func loadFrenchFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(French))
		})
		return &model
	}
}

func loadGandaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Ganda))
		})
		return &model
	}
}

func loadGeorgianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Georgian))
		})
		return &model
	}
}

func loadGermanFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(German))
		})
		return &model
	}
}

func loadGreekFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Greek))
		})
		return &model
	}
}

func loadGujaratiFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Gujarati))
		})
		return &model
	}
}

func loadHebrewFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hebrew))
		})
		return &model
	}
}

func loadHindiFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hindi))
		})
		return &model
	}
}

func loadHungarianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hungarian))
		})
		return &model
	}
}

func loadIcelandicFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Icelandic))
		})
		return &model
	}
}

func loadIndonesianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Indonesian))
		})
		return &model
	}
}

func loadIrishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Irish))
		})
		return &model
	}
}

func loadItalianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Italian))
		})
		return &model
	}
}

func loadJapaneseFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Japanese))
		})
		return &model
	}
}

func loadKazakhFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Kazakh))
		})
		return &model
	}
}

func loadKoreanFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Korean))
		})
		return &model
	}
}

func loadLatinFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Latin))
		})
		return &model
	}
}

func loadLatvianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Latvian))
		})
		return &model
	}
}

func loadLithuanianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Lithuanian))
		})
		return &model
	}
}

func loadMacedonianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Macedonian))
		})
		return &model
	}
}

func loadMalayFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Malay))
		})
		return &model
	}
}

func loadMaoriFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Maori))
		})
		return &model
	}
}

func loadMarathiFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Marathi))
		})
		return &model
	}
}

func loadMongolianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Mongolian))
		})
		return &model
	}
}

func loadNynorskFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Nynorsk))
		})
		return &model
	}
}

func loadPersianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Persian))
		})
		return &model
	}
}

func loadPolishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Polish))
		})
		return &model
	}
}

func loadPortugueseFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Portuguese))
		})
		return &model
	}
}

func loadPunjabiFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Punjabi))
		})
		return &model
	}
}

func loadRomanianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Romanian))
		})
		return &model
	}
}

func loadRussianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Russian))
		})
		return &model
	}
}

func loadSerbianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Serbian))
		})
		return &model
	}
}

func loadShonaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Shona))
		})
		return &model
	}
}

func loadSlovakFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Slovak))
		})
		return &model
	}
}

func loadSloveneFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Slovene))
		})
		return &model
	}
}

func loadSomaliFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Somali))
		})
		return &model
	}
}

func loadSothoFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Sotho))
		})
		return &model
	}
}

func loadSpanishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Spanish))
		})
		return &model
	}
}

func loadSwahiliFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Swahili))
		})
		return &model
	}
}

func loadSwedishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Swedish))
		})
		return &model
	}
}

func loadTagalogFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tagalog))
		})
		return &model
	}
}

func loadTamilFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tamil))
		})
		return &model
	}
}

func loadTeluguFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Telugu))
		})
		return &model
	}
}

func loadThaiFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Thai))
		})
		return &model
	}
}

func loadTsongaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tsonga))
		})
		return &model
	}
}

func loadTswanaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tswana))
		})
		return &model
	}
}

func loadTurkishFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Turkish))
		})
		return &model
	}
}

func loadUkrainianFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Ukrainian))
		})
		return &model
	}
}

func loadUrduFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Urdu))
		})
		return &model
	}
}

func loadVietnameseFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Vietnamese))
		})
		return &model
	}
}

func loadWelshFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Welsh))
		})
		return &model
	}
}

func loadXhosaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Xhosa))
		})
		return &model
	}
}

func loadYorubaFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Yoruba))
		})
		return &model
	}
}

func loadZuluFivegramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Zulu))
		})
		return &model
	}
}

func loadFivegrams(language Language) []byte {
	return loadJson(language, 5)
}
