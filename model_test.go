/*
 * Copyright © 2021-today Peter M. Stahl pemistahl@gmail.com
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
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const text = `These sentences are intended for testing purposes.
⚠ Do not use them in production
By the way, they consist of 23 words in total.`

func linesOfText(text string) []string {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		lines = append(lines, strings.ToLower(scanner.Text()))
	}
	return lines
}

var expectedUnigrams = mapStringsToNgrams(
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "l",
	"m", "n", "o", "p", "r", "s", "t", "u", "w", "y")

var expectedUnigramAbsoluteFrequencies = mapKeysToNgrams(map[string]uint32{
	"a": 3, "b": 1, "c": 3, "d": 5, "e": 14, "f": 2, "g": 1, "h": 4, "i": 6,
	"l": 1, "m": 1, "n": 10, "o": 10, "p": 3, "r": 5, "s": 10, "t": 13,
	"u": 3, "w": 2, "y": 3,
})

var expectedUnigramRelativeFrequencies = mapKeysToNgrams(map[string]float64{
	"a": 0.03, "b": 0.01, "c": 0.03, "d": 0.05, "e": 0.14,
	"f": 0.02, "g": 0.01, "h": 0.04, "i": 0.06, "l": 0.01,
	"m": 0.01, "n": 0.1, "o": 0.1, "p": 0.03, "r": 0.05,
	"s": 0.1, "t": 0.13, "u": 0.03, "w": 0.02, "y": 0.03,
})

var expectedBigrams = mapStringsToNgrams(
	"de", "pr", "pu", "do", "uc", "ds", "du", "ur", "us", "ed", "in",
	"io", "em", "en", "is", "al", "es", "ar", "rd", "re", "ey", "nc", "nd",
	"ay", "ng", "ro", "rp", "no", "ns", "nt", "fo", "wa", "se", "od", "si",
	"by", "of", "wo", "on", "st", "ce", "or", "os", "ot", "co", "ta", "te",
	"ct", "th", "ti", "to", "he", "po")

var expectedBigramAbsoluteFrequencies = mapKeysToNgrams(map[string]uint32{
	"de": 1, "pr": 1, "pu": 1, "do": 1, "uc": 1, "ds": 1, "du": 1,
	"ur": 1, "us": 1, "ed": 1, "in": 4, "io": 1, "em": 1, "en": 3,
	"is": 1, "al": 1, "es": 4, "ar": 1, "rd": 1, "re": 1, "ey": 1,
	"nc": 1, "nd": 1, "ay": 1, "ng": 1, "ro": 1, "rp": 1, "no": 1,
	"ns": 1, "nt": 2, "fo": 1, "wa": 1, "se": 4, "od": 1, "si": 1,
	"of": 1, "by": 1, "wo": 1, "on": 2, "st": 2, "ce": 1, "or": 2,
	"os": 1, "ot": 2, "co": 1, "ta": 1, "ct": 1, "te": 3, "th": 4,
	"ti": 2, "to": 1, "he": 4, "po": 1,
})

var expectedBigramRelativeFrequencies = mapKeysToNgrams(map[string]float64{
	"de": 0.2, "pr": 1.0 / 3, "pu": 1.0 / 3, "do": 0.2, "uc": 1.0 / 3,
	"ds": 0.2, "du": 0.2, "ur": 1.0 / 3, "us": 1.0 / 3, "ed": 1.0 / 14,
	"in": 2.0 / 3, "io": 1.0 / 6, "em": 1.0 / 14, "en": 3.0 / 14,
	"is": 1.0 / 6, "al": 1.0 / 3, "es": 2.0 / 7, "ar": 1.0 / 3, "rd": 0.2,
	"re": 0.2, "ey": 1.0 / 14, "nc": 0.1, "nd": 0.1, "ay": 1.0 / 3, "ng": 0.1,
	"ro": 0.2, "rp": 0.2, "no": 0.1, "ns": 0.1, "nt": 0.2, "fo": 0.5, "wa": 0.5,
	"se": 2.0 / 5, "od": 0.1, "si": 0.1, "of": 0.1, "by": 1, "wo": 0.5,
	"on": 0.2, "st": 0.2, "ce": 1.0 / 3, "or": 0.2, "os": 0.1, "ot": 0.2,
	"co": 1.0 / 3, "ta": 1.0 / 13, "ct": 1.0 / 3, "te": 3.0 / 13, "th": 4.0 / 13,
	"ti": 2.0 / 13, "to": 1.0 / 13, "he": 1, "po": 1.0 / 3,
})

var expectedTrigrams = mapStringsToNgrams(
	"rds", "ose", "ded", "con", "use", "est", "ion", "ist", "pur",
	"hem", "hes", "tin", "cti", "tio", "wor", "ten", "hey", "ota", "tal",
	"tes", "uct", "sti", "pro", "odu", "nsi", "rod", "for", "ces", "nce",
	"not", "are", "pos", "tot", "end", "enc", "sis", "sen", "nte", "ses",
	"ord", "ing", "ent", "int", "nde", "way", "the", "rpo", "urp", "duc",
	"ons", "ese")

var expectedTrigramAbsoluteFrequencies = mapKeysToNgrams(map[string]uint32{
	"rds": 1, "ose": 1, "ded": 1, "con": 1, "use": 1, "est": 1, "ion": 1,
	"ist": 1, "pur": 1, "hem": 1, "hes": 1, "tin": 1, "cti": 1, "wor": 1,
	"tio": 1, "ten": 2, "ota": 1, "hey": 1, "tal": 1, "tes": 1, "uct": 1,
	"sti": 1, "pro": 1, "odu": 1, "nsi": 1, "rod": 1, "for": 1, "ces": 1,
	"nce": 1, "not": 1, "pos": 1, "are": 1, "tot": 1, "end": 1, "enc": 1,
	"sis": 1, "sen": 1, "nte": 2, "ord": 1, "ses": 1, "ing": 1, "ent": 1,
	"way": 1, "nde": 1, "int": 1, "rpo": 1, "the": 4, "urp": 1, "duc": 1,
	"ons": 1, "ese": 1,
})

var expectedTrigramRelativeFrequencies = mapKeysToNgrams(map[string]float64{
	"rds": 1, "ose": 1, "ded": 1, "con": 1, "use": 1, "est": 0.25, "ion": 1,
	"ist": 1, "pur": 1, "hem": 0.25, "hes": 0.25, "tin": 0.5, "cti": 1,
	"wor": 1, "tio": 0.5, "ten": 2.0 / 3, "ota": 0.5, "hey": 0.25, "tal": 1,
	"tes": 1.0 / 3, "uct": 1, "sti": 0.5, "pro": 1, "odu": 1, "nsi": 1,
	"rod": 1, "for": 1, "ces": 1, "nce": 1, "not": 1, "pos": 1, "are": 1,
	"tot": 1, "end": 1.0 / 3, "enc": 1.0 / 3, "sis": 1, "sen": 0.25, "nte": 1,
	"ord": 0.5, "ses": 0.25, "ing": 0.25, "ent": 1.0 / 3, "way": 1, "nde": 1,
	"int": 0.25, "rpo": 1, "the": 1, "urp": 1, "duc": 1, "ons": 0.5, "ese": 0.25,
})

var expectedQuadrigrams = mapStringsToNgrams(
	"onsi", "sist", "ende", "ords", "esti", "tenc", "nces", "oduc",
	"tend", "thes", "rpos", "ting", "nten", "nsis", "they", "tota", "cons",
	"tion", "prod", "ence", "test", "otal", "pose", "nded", "oses", "inte",
	"urpo", "them", "sent", "duct", "stin", "ente", "ucti", "purp", "ctio",
	"rodu", "word", "hese")

var expectedQuadrigramAbsoluteFrequencies = mapKeysToNgrams(map[string]uint32{
	"onsi": 1, "sist": 1, "ende": 1, "ords": 1, "esti": 1, "oduc": 1,
	"nces": 1, "tenc": 1, "tend": 1, "thes": 1, "rpos": 1, "ting": 1,
	"nsis": 1, "nten": 2, "tota": 1, "they": 1, "cons": 1, "tion": 1,
	"prod": 1, "otal": 1, "test": 1, "ence": 1, "pose": 1, "oses": 1,
	"nded": 1, "inte": 1, "them": 1, "urpo": 1, "duct": 1, "sent": 1,
	"stin": 1, "ucti": 1, "ente": 1, "purp": 1, "ctio": 1, "rodu": 1,
	"word": 1, "hese": 1,
})

var expectedQuadrigramRelativeFrequencies = mapKeysToNgrams(map[string]float64{
	"onsi": 1, "sist": 1, "ende": 1, "ords": 1, "esti": 1, "oduc": 1, "nces": 1,
	"tenc": 0.5, "tend": 0.5, "thes": 0.25, "rpos": 1, "ting": 1, "nsis": 1,
	"nten": 1, "tota": 1, "they": 0.25, "cons": 1, "tion": 1, "prod": 1,
	"otal": 1, "test": 1, "ence": 1, "pose": 1, "oses": 1, "nded": 1, "inte": 1,
	"them": 0.25, "urpo": 1, "duct": 1, "sent": 1, "stin": 1, "ucti": 1,
	"ente": 1, "purp": 1, "ctio": 1, "rodu": 1, "word": 1, "hese": 1,
})

var expectedFivegrams = mapStringsToNgrams(
	"testi", "sente", "ences", "tende", "these", "ntenc", "ducti",
	"ntend", "onsis", "total", "uctio", "enten", "poses", "ction", "produ",
	"inten", "nsist", "words", "sting", "tence", "purpo", "estin", "roduc",
	"urpos", "ended", "rpose", "oduct", "consi")

var expectedFivegramAbsoluteFrequencies = mapKeysToNgrams(map[string]uint32{
	"testi": 1, "sente": 1, "ences": 1, "tende": 1, "ducti": 1,
	"ntenc": 1, "these": 1, "onsis": 1, "ntend": 1, "total": 1,
	"uctio": 1, "enten": 1, "poses": 1, "ction": 1, "produ": 1,
	"inten": 1, "nsist": 1, "words": 1, "sting": 1, "purpo": 1,
	"tence": 1, "estin": 1, "roduc": 1, "urpos": 1, "rpose": 1,
	"ended": 1, "oduct": 1, "consi": 1,
})

var expectedFivegramRelativeFrequencies = mapKeysToNgrams(map[string]float64{
	"testi": 1, "sente": 1, "ences": 1, "tende": 1, "ducti": 1, "ntenc": 0.5,
	"these": 1, "onsis": 1, "ntend": 0.5, "total": 1, "uctio": 1, "enten": 1,
	"poses": 1, "ction": 1, "produ": 1, "inten": 1, "nsist": 1, "words": 1,
	"sting": 1, "purpo": 1, "tence": 1, "estin": 1, "roduc": 1, "urpos": 1,
	"rpose": 1, "ended": 1, "oduct": 1, "consi": 1,
})

func TestNewTrainingDataLanguageModel(t *testing.T) {
	params := []struct {
		ngramLength                   int
		expectedAbsoluteFrequencies   map[ngram]uint32
		expectedRelativeFrequencies   map[ngram]float64
		lowerNgramAbsoluteFrequencies map[ngram]uint32
	}{
		{1, expectedUnigramAbsoluteFrequencies, expectedUnigramRelativeFrequencies, map[ngram]uint32{}},
		{2, expectedBigramAbsoluteFrequencies, expectedBigramRelativeFrequencies, expectedUnigramAbsoluteFrequencies},
		{3, expectedTrigramAbsoluteFrequencies, expectedTrigramRelativeFrequencies, expectedBigramAbsoluteFrequencies},
		{4, expectedQuadrigramAbsoluteFrequencies, expectedQuadrigramRelativeFrequencies, expectedTrigramAbsoluteFrequencies},
		{5, expectedFivegramAbsoluteFrequencies, expectedFivegramRelativeFrequencies, expectedQuadrigramAbsoluteFrequencies},
	}
	for i := range params {
		model := newTrainingDataLanguageModel(
			linesOfText(text),
			English,
			params[i].ngramLength,
			"\\p{L}&&\\p{Latin}",
			params[i].lowerNgramAbsoluteFrequencies,
		)
		assert.Equal(t, English, model.language)
		assert.Equal(t, params[i].expectedAbsoluteFrequencies, model.absoluteFrequencies)
		assert.Equal(t, params[i].expectedRelativeFrequencies, model.relativeFrequencies)

	}
}

func TestNewTestDataLanguageModel(t *testing.T) {
	params := []struct {
		ngramLength    int
		expectedNgrams map[ngram]struct{}
	}{
		{1, expectedUnigrams},
		{2, expectedBigrams},
		{3, expectedTrigrams},
		{4, expectedQuadrigrams},
		{5, expectedFivegrams},
	}
	for i := range params {
		model := newTestDataLanguageModel(strings.ToLower(text), params[i].ngramLength)
		assert.Equal(t, params[i].expectedNgrams, model.ngrams)
	}
}

func mapStringsToNgrams(strings ...string) map[ngram]struct{} {
	ngrams := make(map[ngram]struct{})
	for _, s := range strings {
		ngrams[newNgram(s)] = struct{}{}
	}
	return ngrams
}

func mapKeysToNgrams[V any](m map[string]V) map[ngram]V {
	ngrams := make(map[ngram]V)
	for key, value := range m {
		ngrams[newNgram(key)] = value
	}
	return ngrams
}
