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

var quadrigramModels = map[Language]lazyTrainingDataLanguageModel{
	Afrikaans:   afrikaansQuadrigramModel(),
	Albanian:    albanianQuadrigramModel(),
	Arabic:      arabicQuadrigramModel(),
	Armenian:    armenianQuadrigramModel(),
	Azerbaijani: azerbaijaniQuadrigramModel(),
	Basque:      basqueQuadrigramModel(),
	Belarusian:  belarusianQuadrigramModel(),
	Bengali:     bengaliQuadrigramModel(),
	Bokmal:      bokmalQuadrigramModel(),
	Bosnian:     bosnianQuadrigramModel(),
	Bulgarian:   bulgarianQuadrigramModel(),
	Catalan:     catalanQuadrigramModel(),
	Chinese:     chineseQuadrigramModel(),
	Croatian:    croatianQuadrigramModel(),
	Czech:       czechQuadrigramModel(),
	Danish:      danishQuadrigramModel(),
	Dutch:       dutchQuadrigramModel(),
	English:     englishQuadrigramModel(),
	Esperanto:   esperantoQuadrigramModel(),
	Estonian:    estonianQuadrigramModel(),
	Finnish:     finnishQuadrigramModel(),
	French:      frenchQuadrigramModel(),
	Ganda:       gandaQuadrigramModel(),
	Georgian:    georgianQuadrigramModel(),
	German:      germanQuadrigramModel(),
	Greek:       greekQuadrigramModel(),
	Gujarati:    gujaratiQuadrigramModel(),
	Hebrew:      hebrewQuadrigramModel(),
	Hindi:       hindiQuadrigramModel(),
	Hungarian:   hungarianQuadrigramModel(),
	Icelandic:   icelandicQuadrigramModel(),
	Indonesian:  indonesianQuadrigramModel(),
	Irish:       irishQuadrigramModel(),
	Italian:     italianQuadrigramModel(),
	Japanese:    japaneseQuadrigramModel(),
	Kazakh:      kazakhQuadrigramModel(),
	Korean:      koreanQuadrigramModel(),
	Latin:       latinQuadrigramModel(),
	Latvian:     latvianQuadrigramModel(),
	Lithuanian:  lithuanianQuadrigramModel(),
	Macedonian:  macedonianQuadrigramModel(),
	Malay:       malayQuadrigramModel(),
	Maori:       maoriQuadrigramModel(),
	Marathi:     marathiQuadrigramModel(),
	Mongolian:   mongolianQuadrigramModel(),
	Nynorsk:     nynorskQuadrigramModel(),
	Persian:     persianQuadrigramModel(),
	Polish:      polishQuadrigramModel(),
	Portuguese:  portugueseQuadrigramModel(),
	Punjabi:     punjabiQuadrigramModel(),
	Romanian:    romanianQuadrigramModel(),
	Russian:     russianQuadrigramModel(),
	Serbian:     serbianQuadrigramModel(),
	Shona:       shonaQuadrigramModel(),
	Sinhala:     sinhalaQuadrigramModel(),
	Slovak:      slovakQuadrigramModel(),
	Slovene:     sloveneQuadrigramModel(),
	Somali:      somaliQuadrigramModel(),
	Sotho:       sothoQuadrigramModel(),
	Spanish:     spanishQuadrigramModel(),
	Swahili:     swahiliQuadrigramModel(),
	Swedish:     swedishQuadrigramModel(),
	Tagalog:     tagalogQuadrigramModel(),
	Tamil:       tamilQuadrigramModel(),
	Telugu:      teluguQuadrigramModel(),
	Thai:        thaiQuadrigramModel(),
	Tsonga:      tsongaQuadrigramModel(),
	Tswana:      tswanaQuadrigramModel(),
	Turkish:     turkishQuadrigramModel(),
	Ukrainian:   ukrainianQuadrigramModel(),
	Urdu:        urduQuadrigramModel(),
	Vietnamese:  vietnameseQuadrigramModel(),
	Welsh:       welshQuadrigramModel(),
	Xhosa:       xhosaQuadrigramModel(),
	Yoruba:      yorubaQuadrigramModel(),
	Zulu:        zuluQuadrigramModel(),
}

func afrikaansQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Afrikaans))
		})
		return model
	}
}

func albanianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Albanian))
		})
		return model
	}
}

func arabicQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Arabic))
		})
		return model
	}
}

func armenianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Armenian))
		})
		return model
	}
}

func azerbaijaniQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Azerbaijani))
		})
		return model
	}
}

func basqueQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Basque))
		})
		return model
	}
}

func belarusianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Belarusian))
		})
		return model
	}
}

func bengaliQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bengali))
		})
		return model
	}
}

func bokmalQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bokmal))
		})
		return model
	}
}

func bosnianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bosnian))
		})
		return model
	}
}

func bulgarianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Bulgarian))
		})
		return model
	}
}

func catalanQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Catalan))
		})
		return model
	}
}

func chineseQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Chinese))
		})
		return model
	}
}

func croatianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Croatian))
		})
		return model
	}
}

func czechQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Czech))
		})
		return model
	}
}

func danishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Danish))
		})
		return model
	}
}

func dutchQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Dutch))
		})
		return model
	}
}

func englishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(English))
		})
		return model
	}
}

func esperantoQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Esperanto))
		})
		return model
	}
}

func estonianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Estonian))
		})
		return model
	}
}

func finnishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Finnish))
		})
		return model
	}
}

func frenchQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(French))
		})
		return model
	}
}

func gandaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Ganda))
		})
		return model
	}
}

func georgianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Georgian))
		})
		return model
	}
}

func germanQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(German))
		})
		return model
	}
}

func greekQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Greek))
		})
		return model
	}
}

func gujaratiQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Gujarati))
		})
		return model
	}
}

func hebrewQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hebrew))
		})
		return model
	}
}

func hindiQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hindi))
		})
		return model
	}
}

func hungarianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Hungarian))
		})
		return model
	}
}

func icelandicQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Icelandic))
		})
		return model
	}
}

func indonesianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Indonesian))
		})
		return model
	}
}

func irishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Irish))
		})
		return model
	}
}

func italianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Italian))
		})
		return model
	}
}

func japaneseQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Japanese))
		})
		return model
	}
}

func kazakhQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Kazakh))
		})
		return model
	}
}

func koreanQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Korean))
		})
		return model
	}
}

func latinQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Latin))
		})
		return model
	}
}

func latvianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Latvian))
		})
		return model
	}
}

func lithuanianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Lithuanian))
		})
		return model
	}
}

func macedonianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Macedonian))
		})
		return model
	}
}

func malayQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Malay))
		})
		return model
	}
}

func maoriQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Maori))
		})
		return model
	}
}

func marathiQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Marathi))
		})
		return model
	}
}

func mongolianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Mongolian))
		})
		return model
	}
}

func nynorskQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Nynorsk))
		})
		return model
	}
}

func persianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Persian))
		})
		return model
	}
}

func polishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Polish))
		})
		return model
	}
}

func portugueseQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Portuguese))
		})
		return model
	}
}

func punjabiQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Punjabi))
		})
		return model
	}
}

func romanianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Romanian))
		})
		return model
	}
}

func russianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Russian))
		})
		return model
	}
}

func serbianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Serbian))
		})
		return model
	}
}

func shonaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Shona))
		})
		return model
	}
}

func sinhalaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Sinhala))
		})
		return model
	}
}

func slovakQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Slovak))
		})
		return model
	}
}

func sloveneQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Slovene))
		})
		return model
	}
}

func somaliQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Somali))
		})
		return model
	}
}

func sothoQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Sotho))
		})
		return model
	}
}

func spanishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Spanish))
		})
		return model
	}
}

func swahiliQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Swahili))
		})
		return model
	}
}

func swedishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Swedish))
		})
		return model
	}
}

func tagalogQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tagalog))
		})
		return model
	}
}

func tamilQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tamil))
		})
		return model
	}
}

func teluguQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Telugu))
		})
		return model
	}
}

func thaiQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Thai))
		})
		return model
	}
}

func tsongaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tsonga))
		})
		return model
	}
}

func tswanaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Tswana))
		})
		return model
	}
}

func turkishQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Turkish))
		})
		return model
	}
}

func ukrainianQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Ukrainian))
		})
		return model
	}
}

func urduQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Urdu))
		})
		return model
	}
}

func vietnameseQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Vietnamese))
		})
		return model
	}
}

func welshQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Welsh))
		})
		return model
	}
}

func xhosaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Xhosa))
		})
		return model
	}
}

func yorubaQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Yoruba))
		})
		return model
	}
}

func zuluQuadrigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadQuadrigrams(Zulu))
		})
		return model
	}
}

func loadQuadrigrams(language Language) []byte {
	return loadJson(language, 4)
}
