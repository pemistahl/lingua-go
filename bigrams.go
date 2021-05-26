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

var bigramModels = map[Language]lazyTrainingDataLanguageModel{
	Afrikaans:   afrikaansBigramModel(),
	Albanian:    albanianBigramModel(),
	Arabic:      arabicBigramModel(),
	Armenian:    armenianBigramModel(),
	Azerbaijani: azerbaijaniBigramModel(),
	Basque:      basqueBigramModel(),
	Belarusian:  belarusianBigramModel(),
	Bengali:     bengaliBigramModel(),
	Bokmal:      bokmalBigramModel(),
	Bosnian:     bosnianBigramModel(),
	Bulgarian:   bulgarianBigramModel(),
	Catalan:     catalanBigramModel(),
	Chinese:     chineseBigramModel(),
	Croatian:    croatianBigramModel(),
	Czech:       czechBigramModel(),
	Danish:      danishBigramModel(),
	Dutch:       dutchBigramModel(),
	English:     englishBigramModel(),
	Esperanto:   esperantoBigramModel(),
	Estonian:    estonianBigramModel(),
	Finnish:     finnishBigramModel(),
	French:      frenchBigramModel(),
	Ganda:       gandaBigramModel(),
	Georgian:    georgianBigramModel(),
	German:      germanBigramModel(),
	Greek:       greekBigramModel(),
	Gujarati:    gujaratiBigramModel(),
	Hebrew:      hebrewBigramModel(),
	Hindi:       hindiBigramModel(),
	Hungarian:   hungarianBigramModel(),
	Icelandic:   icelandicBigramModel(),
	Indonesian:  indonesianBigramModel(),
	Irish:       irishBigramModel(),
	Italian:     italianBigramModel(),
	Japanese:    japaneseBigramModel(),
	Kazakh:      kazakhBigramModel(),
	Korean:      koreanBigramModel(),
	Latin:       latinBigramModel(),
	Latvian:     latvianBigramModel(),
	Lithuanian:  lithuanianBigramModel(),
	Macedonian:  macedonianBigramModel(),
	Malay:       malayBigramModel(),
	Maori:       maoriBigramModel(),
	Marathi:     marathiBigramModel(),
	Mongolian:   mongolianBigramModel(),
	Nynorsk:     nynorskBigramModel(),
	Persian:     persianBigramModel(),
	Polish:      polishBigramModel(),
	Portuguese:  portugueseBigramModel(),
	Punjabi:     punjabiBigramModel(),
	Romanian:    romanianBigramModel(),
	Russian:     russianBigramModel(),
	Serbian:     serbianBigramModel(),
	Shona:       shonaBigramModel(),
	Slovak:      slovakBigramModel(),
	Slovene:     sloveneBigramModel(),
	Somali:      somaliBigramModel(),
	Sotho:       sothoBigramModel(),
	Spanish:     spanishBigramModel(),
	Swahili:     swahiliBigramModel(),
	Swedish:     swedishBigramModel(),
	Tagalog:     tagalogBigramModel(),
	Tamil:       tamilBigramModel(),
	Telugu:      teluguBigramModel(),
	Thai:        thaiBigramModel(),
	Tsonga:      tsongaBigramModel(),
	Tswana:      tswanaBigramModel(),
	Turkish:     turkishBigramModel(),
	Ukrainian:   ukrainianBigramModel(),
	Urdu:        urduBigramModel(),
	Vietnamese:  vietnameseBigramModel(),
	Welsh:       welshBigramModel(),
	Xhosa:       xhosaBigramModel(),
	Yoruba:      yorubaBigramModel(),
	Zulu:        zuluBigramModel(),
}

func afrikaansBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Afrikaans))
		})
		return &model
	}
}

func albanianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Albanian))
		})
		return &model
	}
}

func arabicBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Arabic))
		})
		return &model
	}
}

func armenianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Armenian))
		})
		return &model
	}
}

func azerbaijaniBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Azerbaijani))
		})
		return &model
	}
}

func basqueBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Basque))
		})
		return &model
	}
}

func belarusianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Belarusian))
		})
		return &model
	}
}

func bengaliBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bengali))
		})
		return &model
	}
}

func bokmalBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bokmal))
		})
		return &model
	}
}

func bosnianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bosnian))
		})
		return &model
	}
}

func bulgarianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Bulgarian))
		})
		return &model
	}
}

func catalanBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Catalan))
		})
		return &model
	}
}

func chineseBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Chinese))
		})
		return &model
	}
}

func croatianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Croatian))
		})
		return &model
	}
}

func czechBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Czech))
		})
		return &model
	}
}

func danishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Danish))
		})
		return &model
	}
}

func dutchBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Dutch))
		})
		return &model
	}
}

func englishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(English))
		})
		return &model
	}
}

func esperantoBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Esperanto))
		})
		return &model
	}
}

func estonianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Estonian))
		})
		return &model
	}
}

func finnishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Finnish))
		})
		return &model
	}
}

func frenchBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(French))
		})
		return &model
	}
}

func gandaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Ganda))
		})
		return &model
	}
}

func georgianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Georgian))
		})
		return &model
	}
}

func germanBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(German))
		})
		return &model
	}
}

func greekBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Greek))
		})
		return &model
	}
}

func gujaratiBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Gujarati))
		})
		return &model
	}
}

func hebrewBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hebrew))
		})
		return &model
	}
}

func hindiBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hindi))
		})
		return &model
	}
}

func hungarianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Hungarian))
		})
		return &model
	}
}

func icelandicBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Icelandic))
		})
		return &model
	}
}

func indonesianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Indonesian))
		})
		return &model
	}
}

func irishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Irish))
		})
		return &model
	}
}

func italianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Italian))
		})
		return &model
	}
}

func japaneseBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Japanese))
		})
		return &model
	}
}

func kazakhBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Kazakh))
		})
		return &model
	}
}

func koreanBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Korean))
		})
		return &model
	}
}

func latinBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Latin))
		})
		return &model
	}
}

func latvianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Latvian))
		})
		return &model
	}
}

func lithuanianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Lithuanian))
		})
		return &model
	}
}

func macedonianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Macedonian))
		})
		return &model
	}
}

func malayBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Malay))
		})
		return &model
	}
}

func maoriBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Maori))
		})
		return &model
	}
}

func marathiBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Marathi))
		})
		return &model
	}
}

func mongolianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Mongolian))
		})
		return &model
	}
}

func nynorskBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Nynorsk))
		})
		return &model
	}
}

func persianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Persian))
		})
		return &model
	}
}

func polishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Polish))
		})
		return &model
	}
}

func portugueseBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Portuguese))
		})
		return &model
	}
}

func punjabiBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Punjabi))
		})
		return &model
	}
}

func romanianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Romanian))
		})
		return &model
	}
}

func russianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Russian))
		})
		return &model
	}
}

func serbianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Serbian))
		})
		return &model
	}
}

func shonaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Shona))
		})
		return &model
	}
}

func slovakBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Slovak))
		})
		return &model
	}
}

func sloveneBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Slovene))
		})
		return &model
	}
}

func somaliBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Somali))
		})
		return &model
	}
}

func sothoBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Sotho))
		})
		return &model
	}
}

func spanishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Spanish))
		})
		return &model
	}
}

func swahiliBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Swahili))
		})
		return &model
	}
}

func swedishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Swedish))
		})
		return &model
	}
}

func tagalogBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tagalog))
		})
		return &model
	}
}

func tamilBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tamil))
		})
		return &model
	}
}

func teluguBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Telugu))
		})
		return &model
	}
}

func thaiBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Thai))
		})
		return &model
	}
}

func tsongaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tsonga))
		})
		return &model
	}
}

func tswanaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Tswana))
		})
		return &model
	}
}

func turkishBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Turkish))
		})
		return &model
	}
}

func ukrainianBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Ukrainian))
		})
		return &model
	}
}

func urduBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Urdu))
		})
		return &model
	}
}

func vietnameseBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Vietnamese))
		})
		return &model
	}
}

func welshBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Welsh))
		})
		return &model
	}
}

func xhosaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Xhosa))
		})
		return &model
	}
}

func yorubaBigramModel() lazyTrainingDataLanguageModel {
	var once sync.Once
	var model trainingDataLanguageModel
	return func() *trainingDataLanguageModel {
		once.Do(func() {
			model = newTrainingDataLanguageModelFromJson(loadBigrams(Yoruba))
		})
		return &model
	}
}

func zuluBigramModel() lazyTrainingDataLanguageModel {
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
