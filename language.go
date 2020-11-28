package lingua

//go:generate stringer -type=Language
type Language int

const (
	Afrikaans Language = iota
	Albanian
	Arabic
	Armenian
	Azerbaijani
	Basque
	Belarusian
	Bengali
	Bokmal
	Bosnian
	Bulgarian
	Catalan
	Chinese
	Croatian
	Czech
	Danish
	Dutch
	English
	Esperanto
	Estonian
	Finnish
	French
	Ganda
	Georgian
	German
	Greek
	Gujarati
	Hebrew
	Hindi
	Hungarian
	Icelandic
	Indonesian
	Irish
	Italian
	Japanese
	Kazakh
	Korean
	Latin
	Latvian
	Lithuanian
	Macedonian
	Malay
	Marathi
	Mongolian
	Nynorsk
	Persian
	Polish
	Portuguese
	Punjabi
	Romanian
	Russian
	Serbian
	Shona
	Slovak
	Slovene
	Somali
	Sotho
	Spanish
	Swahili
	Swedish
	Tagalog
	Tamil
	Telugu
	Thai
	Tsonga
	Tswana
	Turkish
	Ukrainian
	Urdu
	Vietnamese
	Welsh
	Xhosa
	Yoruba
	Zulu
)

func AllLanguages() []Language {
	languages := make([]Language, Zulu)
	for i := 0; i < int(Zulu); i++ {
		languages[i] = Language(i)
	}
	return languages
}

func AllSpokenLanguages() (languages []Language) {
	for _, language := range AllLanguages() {
		if language != Latin {
			languages = append(languages, language)
		}
	}
	return
}

func (language Language) uniqueCharacters() string {
	switch language {
	case Azerbaijani:
		return "Əə"
	case Catalan:
		return "Ïï"
	case Czech:
		return "ĚěŘřŮů"
	case Esperanto:
		return "ĈĉĜĝĤĥĴĵŜŝŬŭ"
	case German:
		return "ß"
	case Hungarian:
		return "ŐőŰű"
	case Kazakh:
		return "ӘәҒғҚқҢңҰұ"
	case Latvian:
		return "ĢģĶķĻļŅņ"
	case Lithuanian:
		return "ĖėĮįŲų"
	case Macedonian:
		return "ЃѓЅѕЌќЏџ"
	case Marathi:
		return "ळ"
	case Mongolian:
		return "ӨөҮү"
	case Polish:
		return "ŁłŃńŚśŹź"
	case Romanian:
		return "Țţ"
	case Serbian:
		return "ЂђЋћ"
	case Slovak:
		return "ĹĺĽľŔŕ"
	case Spanish:
		return "¿¡"
	case Ukrainian:
		return "ҐґЄєЇї"
	case Vietnamese:
		return "ẰằẦầẲẳẨẩẴẵẪẫẮắẤấẠạẶặẬậỀềẺẻỂểẼẽỄễẾếỆệỈỉĨĩỊịƠơỒồỜờỎỏỔổỞởỖỗỠỡỐốỚớỘộỢợƯưỪừỦủỬửŨũỮữỨứỤụỰựỲỳỶỷỸỹỴỵ"
	case Yoruba:
		return "ŌōṢṣ"
	default:
		return ""
	}
}
