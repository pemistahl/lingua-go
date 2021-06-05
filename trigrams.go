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

var trigramModels = map[Language]lazyTrainingDataLanguageModel{
	Afrikaans:   afrikaansTrigramModel(),
	Albanian:    albanianTrigramModel(),
	Arabic:      arabicTrigramModel(),
	Armenian:    armenianTrigramModel(),
	Azerbaijani: azerbaijaniTrigramModel(),
	Basque:      basqueTrigramModel(),
	Belarusian:  belarusianTrigramModel(),
	Bengali:     bengaliTrigramModel(),
	Bokmal:      bokmalTrigramModel(),
	Bosnian:     bosnianTrigramModel(),
	Bulgarian:   bulgarianTrigramModel(),
	Catalan:     catalanTrigramModel(),
	Chinese:     chineseTrigramModel(),
	Croatian:    croatianTrigramModel(),
	Czech:       czechTrigramModel(),
	Danish:      danishTrigramModel(),
	Dutch:       dutchTrigramModel(),
	English:     englishTrigramModel(),
	Esperanto:   esperantoTrigramModel(),
	Estonian:    estonianTrigramModel(),
	Finnish:     finnishTrigramModel(),
	French:      frenchTrigramModel(),
	Ganda:       gandaTrigramModel(),
	Georgian:    georgianTrigramModel(),
	German:      germanTrigramModel(),
	Greek:       greekTrigramModel(),
	Gujarati:    gujaratiTrigramModel(),
	Hebrew:      hebrewTrigramModel(),
	Hindi:       hindiTrigramModel(),
	Hungarian:   hungarianTrigramModel(),
	Icelandic:   icelandicTrigramModel(),
	Indonesian:  indonesianTrigramModel(),
	Irish:       irishTrigramModel(),
	Italian:     italianTrigramModel(),
	Japanese:    japaneseTrigramModel(),
	Kazakh:      kazakhTrigramModel(),
	Korean:      koreanTrigramModel(),
	Latin:       latinTrigramModel(),
	Latvian:     latvianTrigramModel(),
	Lithuanian:  lithuanianTrigramModel(),
	Macedonian:  macedonianTrigramModel(),
	Malay:       malayTrigramModel(),
	Maori:       maoriTrigramModel(),
	Marathi:     marathiTrigramModel(),
	Mongolian:   mongolianTrigramModel(),
	Nynorsk:     nynorskTrigramModel(),
	Persian:     persianTrigramModel(),
	Polish:      polishTrigramModel(),
	Portuguese:  portugueseTrigramModel(),
	Punjabi:     punjabiTrigramModel(),
	Romanian:    romanianTrigramModel(),
	Russian:     russianTrigramModel(),
	Serbian:     serbianTrigramModel(),
	Shona:       shonaTrigramModel(),
	Slovak:      slovakTrigramModel(),
	Slovene:     sloveneTrigramModel(),
	Somali:      somaliTrigramModel(),
	Sotho:       sothoTrigramModel(),
	Spanish:     spanishTrigramModel(),
	Swahili:     swahiliTrigramModel(),
	Swedish:     swedishTrigramModel(),
	Tagalog:     tagalogTrigramModel(),
	Tamil:       tamilTrigramModel(),
	Telugu:      teluguTrigramModel(),
	Thai:        thaiTrigramModel(),
	Tsonga:      tsongaTrigramModel(),
	Tswana:      tswanaTrigramModel(),
	Turkish:     turkishTrigramModel(),
	Ukrainian:   ukrainianTrigramModel(),
	Urdu:        urduTrigramModel(),
	Vietnamese:  vietnameseTrigramModel(),
	Welsh:       welshTrigramModel(),
	Xhosa:       xhosaTrigramModel(),
	Yoruba:      yorubaTrigramModel(),
	Zulu:        zuluTrigramModel(),
}

func afrikaansTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Afrikaans))
		})
		return model
	}
}

func albanianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Albanian))
		})
		return model
	}
}

func arabicTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Arabic))
		})
		return model
	}
}

func armenianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Armenian))
		})
		return model
	}
}

func azerbaijaniTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Azerbaijani))
		})
		return model
	}
}

func basqueTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Basque))
		})
		return model
	}
}

func belarusianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Belarusian))
		})
		return model
	}
}

func bengaliTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bengali))
		})
		return model
	}
}

func bokmalTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bokmal))
		})
		return model
	}
}

func bosnianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bosnian))
		})
		return model
	}
}

func bulgarianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Bulgarian))
		})
		return model
	}
}

func catalanTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Catalan))
		})
		return model
	}
}

func chineseTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Chinese))
		})
		return model
	}
}

func croatianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Croatian))
		})
		return model
	}
}

func czechTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Czech))
		})
		return model
	}
}

func danishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Danish))
		})
		return model
	}
}

func dutchTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Dutch))
		})
		return model
	}
}

func englishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(English))
		})
		return model
	}
}

func esperantoTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Esperanto))
		})
		return model
	}
}

func estonianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Estonian))
		})
		return model
	}
}

func finnishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Finnish))
		})
		return model
	}
}

func frenchTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(French))
		})
		return model
	}
}

func gandaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Ganda))
		})
		return model
	}
}

func georgianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Georgian))
		})
		return model
	}
}

func germanTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(German))
		})
		return model
	}
}

func greekTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Greek))
		})
		return model
	}
}

func gujaratiTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Gujarati))
		})
		return model
	}
}

func hebrewTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hebrew))
		})
		return model
	}
}

func hindiTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hindi))
		})
		return model
	}
}

func hungarianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Hungarian))
		})
		return model
	}
}

func icelandicTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Icelandic))
		})
		return model
	}
}

func indonesianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Indonesian))
		})
		return model
	}
}

func irishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Irish))
		})
		return model
	}
}

func italianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Italian))
		})
		return model
	}
}

func japaneseTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Japanese))
		})
		return model
	}
}

func kazakhTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Kazakh))
		})
		return model
	}
}

func koreanTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Korean))
		})
		return model
	}
}

func latinTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Latin))
		})
		return model
	}
}

func latvianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Latvian))
		})
		return model
	}
}

func lithuanianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Lithuanian))
		})
		return model
	}
}

func macedonianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Macedonian))
		})
		return model
	}
}

func malayTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Malay))
		})
		return model
	}
}

func maoriTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Maori))
		})
		return model
	}
}

func marathiTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Marathi))
		})
		return model
	}
}

func mongolianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Mongolian))
		})
		return model
	}
}

func nynorskTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Nynorsk))
		})
		return model
	}
}

func persianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Persian))
		})
		return model
	}
}

func polishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Polish))
		})
		return model
	}
}

func portugueseTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Portuguese))
		})
		return model
	}
}

func punjabiTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Punjabi))
		})
		return model
	}
}

func romanianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Romanian))
		})
		return model
	}
}

func russianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Russian))
		})
		return model
	}
}

func serbianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Serbian))
		})
		return model
	}
}

func shonaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Shona))
		})
		return model
	}
}

func slovakTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Slovak))
		})
		return model
	}
}

func sloveneTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Slovene))
		})
		return model
	}
}

func somaliTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Somali))
		})
		return model
	}
}

func sothoTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Sotho))
		})
		return model
	}
}

func spanishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Spanish))
		})
		return model
	}
}

func swahiliTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Swahili))
		})
		return model
	}
}

func swedishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Swedish))
		})
		return model
	}
}

func tagalogTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tagalog))
		})
		return model
	}
}

func tamilTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tamil))
		})
		return model
	}
}

func teluguTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Telugu))
		})
		return model
	}
}

func thaiTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Thai))
		})
		return model
	}
}

func tsongaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tsonga))
		})
		return model
	}
}

func tswanaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Tswana))
		})
		return model
	}
}

func turkishTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Turkish))
		})
		return model
	}
}

func ukrainianTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Ukrainian))
		})
		return model
	}
}

func urduTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Urdu))
		})
		return model
	}
}

func vietnameseTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Vietnamese))
		})
		return model
	}
}

func welshTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Welsh))
		})
		return model
	}
}

func xhosaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Xhosa))
		})
		return model
	}
}

func yorubaTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Yoruba))
		})
		return model
	}
}

func zuluTrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadTrigrams(Zulu))
		})
		return model
	}
}

func loadTrigrams(language Language) []byte {
	return loadJson(language, 3)
}
