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

import "sync"

var unigramModels = map[Language]lazyTrainingDataLanguageModel{
	Afrikaans:   afrikaansUnigramModel(),
	Albanian:    albanianUnigramModel(),
	Arabic:      arabicUnigramModel(),
	Armenian:    armenianUnigramModel(),
	Azerbaijani: azerbaijaniUnigramModel(),
	Basque:      basqueUnigramModel(),
	Belarusian:  belarusianUnigramModel(),
	Bengali:     bengaliUnigramModel(),
	Bokmal:      bokmalUnigramModel(),
	Bosnian:     bosnianUnigramModel(),
	Bulgarian:   bulgarianUnigramModel(),
	Catalan:     catalanUnigramModel(),
	Chinese:     chineseUnigramModel(),
	Croatian:    croatianUnigramModel(),
	Czech:       czechUnigramModel(),
	Danish:      danishUnigramModel(),
	Dutch:       dutchUnigramModel(),
	English:     englishUnigramModel(),
	Esperanto:   esperantoUnigramModel(),
	Estonian:    estonianUnigramModel(),
	Finnish:     finnishUnigramModel(),
	French:      frenchUnigramModel(),
	Ganda:       gandaUnigramModel(),
	Georgian:    georgianUnigramModel(),
	German:      germanUnigramModel(),
	Greek:       greekUnigramModel(),
	Gujarati:    gujaratiUnigramModel(),
	Hebrew:      hebrewUnigramModel(),
	Hindi:       hindiUnigramModel(),
	Hungarian:   hungarianUnigramModel(),
	Icelandic:   icelandicUnigramModel(),
	Indonesian:  indonesianUnigramModel(),
	Irish:       irishUnigramModel(),
	Italian:     italianUnigramModel(),
	Japanese:    japaneseUnigramModel(),
	Kazakh:      kazakhUnigramModel(),
	Korean:      koreanUnigramModel(),
	Latin:       latinUnigramModel(),
	Latvian:     latvianUnigramModel(),
	Lithuanian:  lithuanianUnigramModel(),
	Macedonian:  macedonianUnigramModel(),
	Malay:       malayUnigramModel(),
	Maori:       maoriUnigramModel(),
	Marathi:     marathiUnigramModel(),
	Mongolian:   mongolianUnigramModel(),
	Nynorsk:     nynorskUnigramModel(),
	Persian:     persianUnigramModel(),
	Polish:      polishUnigramModel(),
	Portuguese:  portugueseUnigramModel(),
	Punjabi:     punjabiUnigramModel(),
	Romanian:    romanianUnigramModel(),
	Russian:     russianUnigramModel(),
	Serbian:     serbianUnigramModel(),
	Shona:       shonaUnigramModel(),
	Slovak:      slovakUnigramModel(),
	Sinhala:     sinhalaUnigramModel(),
	Slovene:     sloveneUnigramModel(),
	Somali:      somaliUnigramModel(),
	Sotho:       sothoUnigramModel(),
	Spanish:     spanishUnigramModel(),
	Swahili:     swahiliUnigramModel(),
	Swedish:     swedishUnigramModel(),
	Tagalog:     tagalogUnigramModel(),
	Tamil:       tamilUnigramModel(),
	Telugu:      teluguUnigramModel(),
	Thai:        thaiUnigramModel(),
	Tsonga:      tsongaUnigramModel(),
	Tswana:      tswanaUnigramModel(),
	Turkish:     turkishUnigramModel(),
	Ukrainian:   ukrainianUnigramModel(),
	Urdu:        urduUnigramModel(),
	Vietnamese:  vietnameseUnigramModel(),
	Welsh:       welshUnigramModel(),
	Xhosa:       xhosaUnigramModel(),
	Yoruba:      yorubaUnigramModel(),
	Zulu:        zuluUnigramModel(),
}

func afrikaansUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Afrikaans))
		})
		return model
	}
}

func albanianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Albanian))
		})
		return model
	}
}

func arabicUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Arabic))
		})
		return model
	}
}

func armenianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Armenian))
		})
		return model
	}
}

func azerbaijaniUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Azerbaijani))
		})
		return model
	}
}

func basqueUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Basque))
		})
		return model
	}
}

func belarusianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Belarusian))
		})
		return model
	}
}

func bengaliUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bengali))
		})
		return model
	}
}

func bokmalUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bokmal))
		})
		return model
	}
}

func bosnianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bosnian))
		})
		return model
	}
}

func bulgarianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Bulgarian))
		})
		return model
	}
}

func catalanUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Catalan))
		})
		return model
	}
}

func chineseUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Chinese))
		})
		return model
	}
}

func croatianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Croatian))
		})
		return model
	}
}

func czechUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Czech))
		})
		return model
	}
}

func danishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Danish))
		})
		return model
	}
}

func dutchUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Dutch))
		})
		return model
	}
}

func englishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(English))
		})
		return model
	}
}

func esperantoUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Esperanto))
		})
		return model
	}
}

func estonianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Estonian))
		})
		return model
	}
}

func finnishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Finnish))
		})
		return model
	}
}

func frenchUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(French))
		})
		return model
	}
}

func gandaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Ganda))
		})
		return model
	}
}

func georgianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Georgian))
		})
		return model
	}
}

func germanUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(German))
		})
		return model
	}
}

func greekUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Greek))
		})
		return model
	}
}

func gujaratiUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Gujarati))
		})
		return model
	}
}

func hebrewUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hebrew))
		})
		return model
	}
}

func hindiUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hindi))
		})
		return model
	}
}

func hungarianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Hungarian))
		})
		return model
	}
}

func icelandicUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Icelandic))
		})
		return model
	}
}

func indonesianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Indonesian))
		})
		return model
	}
}

func irishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Irish))
		})
		return model
	}
}

func italianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Italian))
		})
		return model
	}
}

func japaneseUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Japanese))
		})
		return model
	}
}

func kazakhUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Kazakh))
		})
		return model
	}
}

func koreanUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Korean))
		})
		return model
	}
}

func latinUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Latin))
		})
		return model
	}
}

func latvianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Latvian))
		})
		return model
	}
}

func lithuanianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Lithuanian))
		})
		return model
	}
}

func macedonianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Macedonian))
		})
		return model
	}
}

func malayUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Malay))
		})
		return model
	}
}

func maoriUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Maori))
		})
		return model
	}
}

func marathiUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Marathi))
		})
		return model
	}
}

func mongolianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Mongolian))
		})
		return model
	}
}

func nynorskUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Nynorsk))
		})
		return model
	}
}

func persianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Persian))
		})
		return model
	}
}

func polishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Polish))
		})
		return model
	}
}

func portugueseUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Portuguese))
		})
		return model
	}
}

func punjabiUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Punjabi))
		})
		return model
	}
}

func romanianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Romanian))
		})
		return model
	}
}

func russianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Russian))
		})
		return model
	}
}

func serbianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Serbian))
		})
		return model
	}
}

func shonaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Shona))
		})
		return model
	}
}

func sinhalaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Sinhala))
		})
		return model
	}
}

func slovakUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Slovak))
		})
		return model
	}
}

func sloveneUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Slovene))
		})
		return model
	}
}

func somaliUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Somali))
		})
		return model
	}
}

func sothoUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Sotho))
		})
		return model
	}
}

func spanishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Spanish))
		})
		return model
	}
}

func swahiliUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Swahili))
		})
		return model
	}
}

func swedishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Swedish))
		})
		return model
	}
}

func tagalogUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tagalog))
		})
		return model
	}
}

func tamilUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tamil))
		})
		return model
	}
}

func teluguUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Telugu))
		})
		return model
	}
}

func thaiUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Thai))
		})
		return model
	}
}

func tsongaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tsonga))
		})
		return model
	}
}

func tswanaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Tswana))
		})
		return model
	}
}

func turkishUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Turkish))
		})
		return model
	}
}

func ukrainianUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Ukrainian))
		})
		return model
	}
}

func urduUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Urdu))
		})
		return model
	}
}

func vietnameseUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Vietnamese))
		})
		return model
	}
}

func welshUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Welsh))
		})
		return model
	}
}

func xhosaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Xhosa))
		})
		return model
	}
}

func yorubaUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Yoruba))
		})
		return model
	}
}

func zuluUnigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadUnigrams(Zulu))
		})
		return model
	}
}

func loadUnigrams(language Language) []byte {
	return loadJson(language, 1)
}
