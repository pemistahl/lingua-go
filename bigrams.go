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

var bigramModels = map[Language]func() *trainingDataLanguageModel{
	Afrikaans:   loadAfrikaansBigramModelFunc(),
	Albanian:    loadAlbanianBigramModelFunc(),
	Arabic:      loadArabicBigramModelFunc(),
	Armenian:    loadArmenianBigramModelFunc(),
	Azerbaijani: loadAzerbaijaniBigramModelFunc(),
	Basque:      loadBasqueBigramModelFunc(),
	Belarusian:  loadBelarusianBigramModelFunc(),
	Bengali:     loadBengaliBigramModelFunc(),
	Bokmal:      loadBokmalBigramModelFunc(),
	Bosnian:     loadBosnianBigramModelFunc(),
	Bulgarian:   loadBulgarianBigramModelFunc(),
	Catalan:     loadCatalanBigramModelFunc(),
	Chinese:     loadChineseBigramModelFunc(),
	Croatian:    loadCroatianBigramModelFunc(),
	Czech:       loadCzechBigramModelFunc(),
	Danish:      loadDanishBigramModelFunc(),
	Dutch:       loadDutchBigramModelFunc(),
	English:     loadEnglishBigramModelFunc(),
	Esperanto:   loadEsperantoBigramModelFunc(),
	Estonian:    loadEstonianBigramModelFunc(),
	Finnish:     loadFinnishBigramModelFunc(),
	French:      loadFrenchBigramModelFunc(),
	Ganda:       loadGandaBigramModelFunc(),
	Georgian:    loadGeorgianBigramModelFunc(),
	German:      loadGermanBigramModelFunc(),
	Greek:       loadGreekBigramModelFunc(),
	Gujarati:    loadGujaratiBigramModelFunc(),
	Hebrew:      loadHebrewBigramModelFunc(),
	Hindi:       loadHindiBigramModelFunc(),
	Hungarian:   loadHungarianBigramModelFunc(),
	Icelandic:   loadIcelandicBigramModelFunc(),
	Indonesian:  loadIndonesianBigramModelFunc(),
	Irish:       loadIrishBigramModelFunc(),
	Italian:     loadItalianBigramModelFunc(),
	Japanese:    loadJapaneseBigramModelFunc(),
	Kazakh:      loadKazakhBigramModelFunc(),
	Korean:      loadKoreanBigramModelFunc(),
	Latin:       loadLatinBigramModelFunc(),
	Latvian:     loadLatvianBigramModelFunc(),
	Lithuanian:  loadLithuanianBigramModelFunc(),
	Macedonian:  loadMacedonianBigramModelFunc(),
	Malay:       loadMalayBigramModelFunc(),
	Maori:       loadMaoriBigramModelFunc(),
	Marathi:     loadMarathiBigramModelFunc(),
	Mongolian:   loadMongolianBigramModelFunc(),
	Nynorsk:     loadNynorskBigramModelFunc(),
	Persian:     loadPersianBigramModelFunc(),
	Polish:      loadPolishBigramModelFunc(),
	Portuguese:  loadPortugueseBigramModelFunc(),
	Punjabi:     loadPunjabiBigramModelFunc(),
	Romanian:    loadRomanianBigramModelFunc(),
	Russian:     loadRussianBigramModelFunc(),
	Serbian:     loadSerbianBigramModelFunc(),
	Shona:       loadShonaBigramModelFunc(),
	Slovak:      loadSlovakBigramModelFunc(),
	Slovene:     loadSloveneBigramModelFunc(),
	Somali:      loadSomaliBigramModelFunc(),
	Sotho:       loadSothoBigramModelFunc(),
	Spanish:     loadSpanishBigramModelFunc(),
	Swahili:     loadSwahiliBigramModelFunc(),
	Swedish:     loadSwedishBigramModelFunc(),
	Tagalog:     loadTagalogBigramModelFunc(),
	Tamil:       loadTamilBigramModelFunc(),
	Telugu:      loadTeluguBigramModelFunc(),
	Thai:        loadThaiBigramModelFunc(),
	Tsonga:      loadTsongaBigramModelFunc(),
	Tswana:      loadTswanaBigramModelFunc(),
	Turkish:     loadTurkishBigramModelFunc(),
	Ukrainian:   loadUkrainianBigramModelFunc(),
	Urdu:        loadUrduBigramModelFunc(),
	Vietnamese:  loadVietnameseBigramModelFunc(),
	Welsh:       loadWelshBigramModelFunc(),
	Xhosa:       loadXhosaBigramModelFunc(),
	Yoruba:      loadYorubaBigramModelFunc(),
	Zulu:        loadZuluBigramModelFunc(),
}

func loadAfrikaansBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Afrikaans))
		})
		return &model
	}
}

func loadAlbanianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Albanian))
		})
		return &model
	}
}

func loadArabicBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Arabic))
		})
		return &model
	}
}

func loadArmenianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Armenian))
		})
		return &model
	}
}

func loadAzerbaijaniBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Azerbaijani))
		})
		return &model
	}
}

func loadBasqueBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Basque))
		})
		return &model
	}
}

func loadBelarusianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Belarusian))
		})
		return &model
	}
}

func loadBengaliBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bengali))
		})
		return &model
	}
}

func loadBokmalBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bokmal))
		})
		return &model
	}
}

func loadBosnianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bosnian))
		})
		return &model
	}
}

func loadBulgarianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bulgarian))
		})
		return &model
	}
}

func loadCatalanBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Catalan))
		})
		return &model
	}
}

func loadChineseBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Chinese))
		})
		return &model
	}
}

func loadCroatianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Croatian))
		})
		return &model
	}
}

func loadCzechBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Czech))
		})
		return &model
	}
}

func loadDanishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Danish))
		})
		return &model
	}
}

func loadDutchBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Dutch))
		})
		return &model
	}
}

func loadEnglishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(English))
		})
		return &model
	}
}

func loadEsperantoBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Esperanto))
		})
		return &model
	}
}

func loadEstonianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Estonian))
		})
		return &model
	}
}

func loadFinnishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Finnish))
		})
		return &model
	}
}

func loadFrenchBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(French))
		})
		return &model
	}
}

func loadGandaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Ganda))
		})
		return &model
	}
}

func loadGeorgianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Georgian))
		})
		return &model
	}
}

func loadGermanBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(German))
		})
		return &model
	}
}

func loadGreekBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Greek))
		})
		return &model
	}
}

func loadGujaratiBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Gujarati))
		})
		return &model
	}
}

func loadHebrewBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hebrew))
		})
		return &model
	}
}

func loadHindiBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hindi))
		})
		return &model
	}
}

func loadHungarianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hungarian))
		})
		return &model
	}
}

func loadIcelandicBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Icelandic))
		})
		return &model
	}
}

func loadIndonesianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Indonesian))
		})
		return &model
	}
}

func loadIrishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Irish))
		})
		return &model
	}
}

func loadItalianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Italian))
		})
		return &model
	}
}

func loadJapaneseBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Japanese))
		})
		return &model
	}
}

func loadKazakhBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Kazakh))
		})
		return &model
	}
}

func loadKoreanBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Korean))
		})
		return &model
	}
}

func loadLatinBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Latin))
		})
		return &model
	}
}

func loadLatvianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Latvian))
		})
		return &model
	}
}

func loadLithuanianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Lithuanian))
		})
		return &model
	}
}

func loadMacedonianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Macedonian))
		})
		return &model
	}
}

func loadMalayBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Malay))
		})
		return &model
	}
}

func loadMaoriBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Maori))
		})
		return &model
	}
}

func loadMarathiBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Marathi))
		})
		return &model
	}
}

func loadMongolianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Mongolian))
		})
		return &model
	}
}

func loadNynorskBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Nynorsk))
		})
		return &model
	}
}

func loadPersianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Persian))
		})
		return &model
	}
}

func loadPolishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Polish))
		})
		return &model
	}
}

func loadPortugueseBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Portuguese))
		})
		return &model
	}
}

func loadPunjabiBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Punjabi))
		})
		return &model
	}
}

func loadRomanianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Romanian))
		})
		return &model
	}
}

func loadRussianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Russian))
		})
		return &model
	}
}

func loadSerbianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Serbian))
		})
		return &model
	}
}

func loadShonaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Shona))
		})
		return &model
	}
}

func loadSlovakBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Slovak))
		})
		return &model
	}
}

func loadSloveneBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Slovene))
		})
		return &model
	}
}

func loadSomaliBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Somali))
		})
		return &model
	}
}

func loadSothoBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Sotho))
		})
		return &model
	}
}

func loadSpanishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Spanish))
		})
		return &model
	}
}

func loadSwahiliBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Swahili))
		})
		return &model
	}
}

func loadSwedishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Swedish))
		})
		return &model
	}
}

func loadTagalogBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tagalog))
		})
		return &model
	}
}

func loadTamilBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tamil))
		})
		return &model
	}
}

func loadTeluguBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Telugu))
		})
		return &model
	}
}

func loadThaiBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Thai))
		})
		return &model
	}
}

func loadTsongaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tsonga))
		})
		return &model
	}
}

func loadTswanaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tswana))
		})
		return &model
	}
}

func loadTurkishBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Turkish))
		})
		return &model
	}
}

func loadUkrainianBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Ukrainian))
		})
		return &model
	}
}

func loadUrduBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Urdu))
		})
		return &model
	}
}

func loadVietnameseBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Vietnamese))
		})
		return &model
	}
}

func loadWelshBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Welsh))
		})
		return &model
	}
}

func loadXhosaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Xhosa))
		})
		return &model
	}
}

func loadYorubaBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Yoruba))
		})
		return &model
	}
}

func loadZuluBigramModelFunc() func() *trainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Zulu))
		})
		return &model
	}
}

func loadBigrams(language Language) []byte {
	return loadJson(language, 2)
}
