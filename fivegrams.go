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

var fivegramModels = map[Language]lazyTrainingDataLanguageModel{
	Afrikaans:   afrikaansFivegramModel(),
	Albanian:    albanianFivegramModel(),
	Arabic:      arabicFivegramModel(),
	Armenian:    armenianFivegramModel(),
	Azerbaijani: azerbaijaniFivegramModel(),
	Basque:      basqueFivegramModel(),
	Belarusian:  belarusianFivegramModel(),
	Bengali:     bengaliFivegramModel(),
	Bokmal:      bokmalFivegramModel(),
	Bosnian:     bosnianFivegramModel(),
	Bulgarian:   bulgarianFivegramModel(),
	Catalan:     catalanFivegramModel(),
	Chinese:     chineseFivegramModel(),
	Croatian:    croatianFivegramModel(),
	Czech:       czechFivegramModel(),
	Danish:      danishFivegramModel(),
	Dutch:       dutchFivegramModel(),
	English:     englishFivegramModel(),
	Esperanto:   esperantoFivegramModel(),
	Estonian:    estonianFivegramModel(),
	Finnish:     finnishFivegramModel(),
	French:      frenchFivegramModel(),
	Ganda:       gandaFivegramModel(),
	Georgian:    georgianFivegramModel(),
	German:      germanFivegramModel(),
	Greek:       greekFivegramModel(),
	Gujarati:    gujaratiFivegramModel(),
	Hebrew:      hebrewFivegramModel(),
	Hindi:       hindiFivegramModel(),
	Hungarian:   hungarianFivegramModel(),
	Icelandic:   icelandicFivegramModel(),
	Indonesian:  indonesianFivegramModel(),
	Irish:       irishFivegramModel(),
	Italian:     italianFivegramModel(),
	Japanese:    japaneseFivegramModel(),
	Kazakh:      kazakhFivegramModel(),
	Korean:      koreanFivegramModel(),
	Latin:       latinFivegramModel(),
	Latvian:     latvianFivegramModel(),
	Lithuanian:  lithuanianFivegramModel(),
	Macedonian:  macedonianFivegramModel(),
	Malay:       malayFivegramModel(),
	Maori:       maoriFivegramModel(),
	Marathi:     marathiFivegramModel(),
	Mongolian:   mongolianFivegramModel(),
	Nynorsk:     nynorskFivegramModel(),
	Persian:     persianFivegramModel(),
	Polish:      polishFivegramModel(),
	Portuguese:  portugueseFivegramModel(),
	Punjabi:     punjabiFivegramModel(),
	Romanian:    romanianFivegramModel(),
	Russian:     russianFivegramModel(),
	Serbian:     serbianFivegramModel(),
	Shona:       shonaFivegramModel(),
	Sinhala:     sinhalaFivegramModel(),
	Slovak:      slovakFivegramModel(),
	Slovene:     sloveneFivegramModel(),
	Somali:      somaliFivegramModel(),
	Sotho:       sothoFivegramModel(),
	Spanish:     spanishFivegramModel(),
	Swahili:     swahiliFivegramModel(),
	Swedish:     swedishFivegramModel(),
	Tagalog:     tagalogFivegramModel(),
	Tamil:       tamilFivegramModel(),
	Telugu:      teluguFivegramModel(),
	Thai:        thaiFivegramModel(),
	Tsonga:      tsongaFivegramModel(),
	Tswana:      tswanaFivegramModel(),
	Turkish:     turkishFivegramModel(),
	Ukrainian:   ukrainianFivegramModel(),
	Urdu:        urduFivegramModel(),
	Vietnamese:  vietnameseFivegramModel(),
	Welsh:       welshFivegramModel(),
	Xhosa:       xhosaFivegramModel(),
	Yoruba:      yorubaFivegramModel(),
	Zulu:        zuluFivegramModel(),
}

func afrikaansFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Afrikaans))
		})
		return model
	}
}

func albanianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Albanian))
		})
		return model
	}
}

func arabicFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Arabic))
		})
		return model
	}
}

func armenianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Armenian))
		})
		return model
	}
}

func azerbaijaniFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Azerbaijani))
		})
		return model
	}
}

func basqueFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Basque))
		})
		return model
	}
}

func belarusianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Belarusian))
		})
		return model
	}
}

func bengaliFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bengali))
		})
		return model
	}
}

func bokmalFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bokmal))
		})
		return model
	}
}

func bosnianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bosnian))
		})
		return model
	}
}

func bulgarianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Bulgarian))
		})
		return model
	}
}

func catalanFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Catalan))
		})
		return model
	}
}

func chineseFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Chinese))
		})
		return model
	}
}

func croatianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Croatian))
		})
		return model
	}
}

func czechFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Czech))
		})
		return model
	}
}

func danishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Danish))
		})
		return model
	}
}

func dutchFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Dutch))
		})
		return model
	}
}

func englishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(English))
		})
		return model
	}
}

func esperantoFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Esperanto))
		})
		return model
	}
}

func estonianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Estonian))
		})
		return model
	}
}

func finnishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Finnish))
		})
		return model
	}
}

func frenchFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(French))
		})
		return model
	}
}

func gandaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Ganda))
		})
		return model
	}
}

func georgianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Georgian))
		})
		return model
	}
}

func germanFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(German))
		})
		return model
	}
}

func greekFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Greek))
		})
		return model
	}
}

func gujaratiFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Gujarati))
		})
		return model
	}
}

func hebrewFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hebrew))
		})
		return model
	}
}

func hindiFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hindi))
		})
		return model
	}
}

func hungarianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Hungarian))
		})
		return model
	}
}

func icelandicFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Icelandic))
		})
		return model
	}
}

func indonesianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Indonesian))
		})
		return model
	}
}

func irishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Irish))
		})
		return model
	}
}

func italianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Italian))
		})
		return model
	}
}

func japaneseFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Japanese))
		})
		return model
	}
}

func kazakhFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Kazakh))
		})
		return model
	}
}

func koreanFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Korean))
		})
		return model
	}
}

func latinFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Latin))
		})
		return model
	}
}

func latvianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Latvian))
		})
		return model
	}
}

func lithuanianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Lithuanian))
		})
		return model
	}
}

func macedonianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Macedonian))
		})
		return model
	}
}

func malayFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Malay))
		})
		return model
	}
}

func maoriFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Maori))
		})
		return model
	}
}

func marathiFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Marathi))
		})
		return model
	}
}

func mongolianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Mongolian))
		})
		return model
	}
}

func nynorskFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Nynorsk))
		})
		return model
	}
}

func persianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Persian))
		})
		return model
	}
}

func polishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Polish))
		})
		return model
	}
}

func portugueseFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Portuguese))
		})
		return model
	}
}

func punjabiFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Punjabi))
		})
		return model
	}
}

func romanianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Romanian))
		})
		return model
	}
}

func russianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Russian))
		})
		return model
	}
}

func serbianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Serbian))
		})
		return model
	}
}

func shonaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Shona))
		})
		return model
	}
}

func sinhalaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Sinhala))
		})
		return model
	}
}

func slovakFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Slovak))
		})
		return model
	}
}

func sloveneFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Slovene))
		})
		return model
	}
}

func somaliFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Somali))
		})
		return model
	}
}

func sothoFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Sotho))
		})
		return model
	}
}

func spanishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Spanish))
		})
		return model
	}
}

func swahiliFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Swahili))
		})
		return model
	}
}

func swedishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Swedish))
		})
		return model
	}
}

func tagalogFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tagalog))
		})
		return model
	}
}

func tamilFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tamil))
		})
		return model
	}
}

func teluguFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Telugu))
		})
		return model
	}
}

func thaiFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Thai))
		})
		return model
	}
}

func tsongaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tsonga))
		})
		return model
	}
}

func tswanaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Tswana))
		})
		return model
	}
}

func turkishFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Turkish))
		})
		return model
	}
}

func ukrainianFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Ukrainian))
		})
		return model
	}
}

func urduFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Urdu))
		})
		return model
	}
}

func vietnameseFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Vietnamese))
		})
		return model
	}
}

func welshFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Welsh))
		})
		return model
	}
}

func xhosaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Xhosa))
		})
		return model
	}
}

func yorubaFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Yoruba))
		})
		return model
	}
}

func zuluFivegramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() languageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadFivegrams(Zulu))
		})
		return model
	}
}

func loadFivegrams(language Language) []byte {
	return loadJson(language, 5)
}
