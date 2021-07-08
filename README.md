![lingua](images/logo.png)

<br>

[![Build Status](https://github.com/pemistahl/lingua-go/workflows/build/badge.svg?branch=main)](https://github.com/pemistahl/lingua-go/actions?query=workflow%3A%22build%22+branch%3Amain)
[![codecov](https://codecov.io/gh/pemistahl/lingua-go/branch/main/graph/badge.svg)](https://codecov.io/gh/pemistahl/lingua-go)
[![supported languages](https://img.shields.io/badge/supported%20languages-76-green.svg)](#supported-languages)
[![Go Reference](https://pkg.go.dev/badge/github.com/pemistahl/lingua-go.svg)](https://pkg.go.dev/github.com/pemistahl/lingua-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/pemistahl/lingua-go)](https://goreportcard.com/report/github.com/pemistahl/lingua-go)
[![license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

## <a name="table-of-contents"></a> Table of Contents

1. [What does this library do?](#library-purpose)
2. [Why does this library exist?](#library-reason)
3. [Which languages are supported?](#supported-languages)
4. [How good is it?](#library-accuracy)
5. [Why is it better than other libraries?](#why-is-it-better)
6. [Test report generation](#report-generation)
7. [How to add it to your project?](#library-dependency)
8. [How to build?](#library-build)
9. [How to use?](#library-use)
10. [What's next for version 1.1.0?](#whats-next)
11. [Contributions](#contributions)

## 1. <a name="library-purpose"></a> What does this library do? <sup>[Top ▲](#table-of-contents)</sup>

Its task is simple: It tells you which language some provided textual data is written in.
This is very useful as a preprocessing step for linguistic data in natural language
processing applications such as text classification and spell checking.
Other use cases, for instance, might include routing e-mails to the right geographically
located customer service department, based on the e-mails' languages.

## 2. <a name="library-reason"></a> Why does this library exist? <sup>[Top ▲](#table-of-contents)</sup>

Language detection is often done as part of large machine learning frameworks or natural
language processing applications. In cases where you don't need the full-fledged
functionality of those systems or don't want to learn the ropes of those,
a small flexible library comes in handy.

So far, the only other comprehensive open source library in the Go ecosystem for
this task is [*Whatlanggo*](https://github.com/abadojack/whatlanggo).
Unfortunately, it has two major drawbacks:

1. Detection only works with quite lengthy text fragments. For very short text snippets
   such as Twitter messages, it does not provide adequate results.
2. The more languages take part in the decision process, the less accurate are the
   detection results.

*Lingua* aims at eliminating these problems. It nearly does not need any configuration and
yields pretty accurate results on both long and short text, even on single words and phrases.
It draws on both rule-based and statistical methods but does not use any dictionaries of words.
It does not need a connection to any external API or service either.
Once the library has been downloaded, it can be used completely offline.

## 3. <a name="supported-languages"></a> Which languages are supported? <sup>[Top ▲](#table-of-contents)</sup>

Compared to other language detection libraries, *Lingua's* focus is on *quality over quantity*, that is,
getting detection right for a small set of languages first before adding new ones.
Currently, the following 76 languages are supported:

- A
    - Afrikaans
    - Albanian
    - Arabic
    - Armenian
    - Azerbaijani
- B
    - Basque
    - Belarusian
    - Bengali
    - Norwegian Bokmal
    - Bosnian
    - Bulgarian
- C
    - Catalan
    - Chinese
    - Croatian
    - Czech
- D
    - Danish
    - Dutch
- E
    - English
    - Esperanto
    - Estonian
- F
    - Finnish
    - French
- G
    - Ganda
    - Georgian
    - German
    - Greek
    - Gujarati
- H
    - Hebrew
    - Hindi
    - Hungarian
- I
    - Icelandic
    - Indonesian
    - Irish
    - Italian
- J
    - Japanese
- K
    - Kazakh
    - Korean
- L
    - Latin
    - Latvian
    - Lithuanian
- M
    - Macedonian
    - Malay
    - Maori
    - Marathi
    - Mongolian
- N
    - Norwegian Nynorsk
- P
    - Persian
    - Polish
    - Portuguese
    - Punjabi
- R
    - Romanian
    - Russian
- S
    - Serbian
    - Shona
    - Sinhala
    - Slovak
    - Slovene
    - Somali
    - Sotho
    - Spanish
    - Swahili
    - Swedish
- T
    - Tagalog
    - Tamil
    - Telugu
    - Thai
    - Tsonga
    - Tswana
    - Turkish
- U
    - Ukrainian
    - Urdu
- V
    - Vietnamese
- W
    - Welsh
- X
    - Xhosa
- Y
    - Yoruba
- Z
    - Zulu

## 4. <a name="library-accuracy"></a> How good is it? <sup>[Top ▲](#table-of-contents)</sup>

*Lingua* is able to report accuracy statistics for some bundled test data available for each
supported language. The test data for each language is split into three parts:

1. a list of single words with a minimum length of 5 characters
2. a list of word pairs with a minimum length of 10 characters
3. a list of complete grammatical sentences of various lengths

Both the language models and the test data have been created from separate documents of the
[Wortschatz corpora](https://wortschatz.uni-leipzig.de) offered by Leipzig University, Germany.
Data crawled from various news websites have been used for training, each corpus comprising one
million sentences. For testing, corpora made of arbitrarily chosen websites have been used,
each comprising ten thousand sentences. From each test corpus, a random unsorted subset of
1000 single words, 1000 word pairs and 1000 sentences has been extracted, respectively.

Given the generated test data, I have compared the detection results of *Lingua* and *Whatlanggo*
running over the data of *Lingua's* supported 76 languages. Additionally, I have added Google's 
[CLD3](https://github.com/google/cld3/) to the comparison with the help of the 
[gocld3](https://github.com/jmhodges/gocld3) bindings. Languages that are not supported
by *CLD3* or *Whatlanggo* are simply ignored during the detection process.

The box plot below shows the distribution of the averaged accuracy values for all three performed tasks:
Single word detection, word pair detection and sentence detection. *Lingua* clearly outperforms its contenders.
Bar plots for each language and further box plots for the separate detection tasks can be found in the file
[ACCURACY_PLOTS.md](https://github.com/pemistahl/lingua-go/blob/main/ACCURACY_PLOTS.md).
Detailed statistics including mean, median and standard deviation values for each language and classifier are
available in the file
[ACCURACY_TABLE.md](https://github.com/pemistahl/lingua-go/blob/main/ACCURACY_TABLE.md).

<img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/images/plots/boxplot-average.png" alt="Average Detection Performance" />

## 5. <a name="why-is-it-better"></a> Why is it better than other libraries? <sup>[Top ▲](#table-of-contents)</sup>

Every language detector uses a probabilistic [n-gram](https://en.wikipedia.org/wiki/N-gram) model trained on the
character distribution in some training corpus. Most libraries only use n-grams of size 3 (trigrams) which is
satisfactory for detecting the language of longer text fragments consisting of multiple sentences. For short
phrases or single words, however, trigrams are not enough. The shorter the input text is, the less n-grams are
available. The probabilities estimated from such few n-grams are not reliable. This is why *Lingua* makes use
of n-grams of sizes 1 up to 5 which results in much more accurate prediction of the correct language.

A second important difference is that *Lingua* does not only use such a statistical model, but also a rule-based
engine. This engine first determines the alphabet of the input text and searches for characters which are unique
in one or more languages. If exactly one language can be reliably chosen this way, the statistical model is not
necessary anymore. In any case, the rule-based engine filters out languages that do not satisfy the conditions
of the input text. Only then, in a second step, the probabilistic n-gram model is taken into consideration.
This makes sense because loading less language models means less memory consumption and better runtime performance.

In general, it is always a good idea to restrict the set of languages to be considered in the classification process
using the respective [api methods](#library-use). If you know beforehand that certain languages are
never to occur in an input text, do not let those take part in the classifcation process. The filtering mechanism
of the rule-based engine is quite good, however, filtering based on your own knowledge of the input text is always preferable.

## 6. <a name="report-generation"></a> Test report generation <sup>[Top ▲](#table-of-contents)</sup>

If you want to reproduce the accuracy results above, you can generate the test reports yourself for both classifiers
and all languages by doing:

    go run cmd/accuracy_reporter.go

For each detector and language, a test report file is then written into
[`/accuracy-reports`](https://github.com/pemistahl/lingua-go/tree/main/accuracy-reports).
As an example, here is the current output of the *Lingua* German report:

```
##### German #####

>>> Accuracy on average: 89.13%

>> Detection of 1000 single words (average length: 9 chars)
Accuracy: 73.90%
Erroneously classified as Dutch: 2.30%, Danish: 2.10%, English: 2.00%, Latin: 1.90%, Bokmal: 1.60%, Basque: 1.20%, French: 1.20%, Italian: 1.20%, Esperanto: 1.10%, Swedish: 1.00%, Afrikaans: 0.80%, Tsonga: 0.70%, Nynorsk: 0.60%, Portuguese: 0.60%, Yoruba: 0.60%, Finnish: 0.50%, Sotho: 0.50%, Welsh: 0.50%, Estonian: 0.40%, Irish: 0.40%, Polish: 0.40%, Spanish: 0.40%, Swahili: 0.40%, Tswana: 0.40%, Bosnian: 0.30%, Icelandic: 0.30%, Tagalog: 0.30%, Albanian: 0.20%, Catalan: 0.20%, Croatian: 0.20%, Indonesian: 0.20%, Lithuanian: 0.20%, Maori: 0.20%, Romanian: 0.20%, Xhosa: 0.20%, Zulu: 0.20%, Latvian: 0.10%, Malay: 0.10%, Slovak: 0.10%, Slovene: 0.10%, Somali: 0.10%, Turkish: 0.10%

>> Detection of 1000 word pairs (average length: 18 chars)
Accuracy: 93.80%
Erroneously classified as Dutch: 0.90%, Latin: 0.80%, English: 0.70%, Swedish: 0.60%, Danish: 0.50%, French: 0.40%, Bokmal: 0.30%, Irish: 0.20%, Tagalog: 0.20%, Tsonga: 0.20%, Afrikaans: 0.10%, Esperanto: 0.10%, Estonian: 0.10%, Finnish: 0.10%, Italian: 0.10%, Maori: 0.10%, Nynorsk: 0.10%, Portuguese: 0.10%, Somali: 0.10%, Swahili: 0.10%, Turkish: 0.10%, Welsh: 0.10%, Xhosa: 0.10%, Zulu: 0.10%

>> Detection of 1000 sentences (average length: 111 chars)
Accuracy: 99.70%
Erroneously classified as Dutch: 0.20%, Latin: 0.10%
```

## 7. <a name="library-dependency"></a> How to add it to your project? <sup>[Top ▲](#table-of-contents)</sup>

    go get github.com/pemistahl/lingua-go@v1.0.1

## 8. <a name="library-build"></a> How to build? <sup>[Top ▲](#table-of-contents)</sup>

*Lingua* requires Go version 1.16.

```
git clone https://github.com/pemistahl/lingua-go.git
cd lingua-go
go build
```

The source code is accompanied by an extensive unit test suite. To run them, simply say:

    go test

## 9. <a name="library-use"></a> How to use? <sup>[Top ▲](#table-of-contents)</sup>

### 9.1 Basic usage

```go
package main

import (
    "fmt"
    "github.com/pemistahl/lingua-go"
)

func main() {
    languages := []lingua.Language{
        lingua.English,
        lingua.French,
        lingua.German,
        lingua.Spanish,
    }

    detector := lingua.NewLanguageDetectorBuilder().
        FromLanguages(languages...).
        Build()

    if language, exists := detector.DetectLanguageOf("languages are awesome"); exists {
        fmt.Println(language)
    }

    // Output: English
}
```

### 9.2 Minimum relative distance

By default, *Lingua* returns the most likely language for a given input text. However, there are
certain words that are spelled the same in more than one language. The word *prologue*, for
instance, is both a valid English and French word. *Lingua* would output either English or
French which might be wrong in the given context. For cases like that, it is possible to
specify a minimum relative distance that the logarithmized and summed up probabilities for
each possible language have to satisfy. It can be stated in the following way:

```go
package main

import (
    "fmt"
    "github.com/pemistahl/lingua-go"
)

func main() {
    languages := []lingua.Language{
        lingua.English,
        lingua.French,
        lingua.German,
        lingua.Spanish,
    }

    detector := lingua.NewLanguageDetectorBuilder().
        FromLanguages(languages...).
        WithMinimumRelativeDistance(0.25).
        Build()

    language, exists := detector.DetectLanguageOf("languages are awesome")

    fmt.Println(language)
    fmt.Println(exists)

    // Output:
    // Unknown
    // false
}
```

Be aware that the distance between the language probabilities is dependent on the length of the
input text. The longer the input text, the larger the distance between the languages. So if you
want to classify very short text phrases, do not set the minimum relative distance too high.
Otherwise [`Unknown`](https://github.com/pemistahl/lingua-go/blob/main/language.go#L106) will be
returned most of the time as in the example above. This is the return value for cases where
language detection is not reliably possible.

### 9.3 Confidence values

Knowing about the most likely language is nice but how reliable is the computed likelihood?
And how less likely are the other examined languages in comparison to the most likely one?
These questions can be answered as well:

```go
package main

import (
    "fmt"
    "github.com/pemistahl/lingua-go"
)

func main() {
    languages := []lingua.Language{
        lingua.English,
        lingua.French,
        lingua.German,
        lingua.Spanish,
    }

    detector := lingua.NewLanguageDetectorBuilder().
        FromLanguages(languages...).
        Build()

    confidenceValues := detector.ComputeLanguageConfidenceValues("languages are awesome")

    for _, elem := range confidenceValues {
        fmt.Printf("%s: %.2f\n", elem.Language(), elem.Value())
    }

    // Output:
    // English: 1.00
    // French: 0.79
    // German: 0.75
    // Spanish: 0.72
}
```

In the example above, a slice of 
[`ConfidenceValue`](https://github.com/pemistahl/lingua-go/blob/main/confidence.go#L21) 
is returned containing all possible languages sorted by their confidence value in descending 
order. The values that this method computes are part of a **relative** confidence metric, not of 
an absolute one. Each value is a number between 0.0 and 1.0. The most likely language is always 
returned with value 1.0. All other languages get values assigned which are lower than 1.0,
denoting how less likely those languages are in comparison to the most likely language.

The slice returned by this method does not necessarily contain all languages which the calling 
instance of `LanguageDetector` was built from. If the rule-based engine decides that a specific 
language is truly impossible, then it will not be part of the returned slice. Likewise, if no 
ngram probabilities can be found within the detector's languages for the given input text, the 
returned slice will be empty. The confidence value for each language not being part of the returned
slice is assumed to be 0.0.

### 9.4 Eager loading versus lazy loading

By default, *Lingua* uses lazy-loading to load only those language models on demand which are
considered relevant by the rule-based filter engine. For web services, for instance, it is
rather beneficial to preload all language models into memory to avoid unexpected latency while
waiting for the service response. If you want to enable the eager-loading mode, you can do it
like this:

```go
lingua.NewLanguageDetectorBuilder().
    FromAllLanguages().
    WithPreloadedLanguageModels().
    Build()
```

Multiple instances of `LanguageDetector` share the same language models in memory which are
accessed asynchronously by the instances.

### 9.5 Methods to build the LanguageDetector

There might be classification tasks where you know beforehand that your language data is
definitely not written in Latin, for instance (what a surprise :-). The detection accuracy can
become better in such cases if you exclude certain languages from the decision process or just
explicitly include relevant languages:

```go
// Including all languages available in the library
// consumes at least 2GB of memory and might
// lead to slow runtime performance.
lingua.NewLanguageDetectorBuilder().FromAllLanguages()

// Include only languages that are not yet extinct
// (= currently excludes Latin).
lingua.NewLanguageDetectorBuilder().FromAllSpokenLanguages()

// Include only languages written with Cyrillic script.
lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithCyrillicScript()

// Exclude only the Spanish language from the decision algorithm.
lingua.NewLanguageDetectorBuilder().FromAllLanguagesWithout(lingua.Spanish)

// Only decide between English and German.
lingua.NewLanguageDetectorBuilder().FromLanguages(lingua.English, lingua.German)

// Select languages by ISO 639-1 code.
lingua.NewLanguageDetectorBuilder().FromIsoCodes639_1(lingua.EN, lingua.DE)

// Select languages by ISO 639-3 code.
lingua.NewLanguageDetectorBuilder().FromIsoCodes639_3(lingua.ENG, lingua.DEU)
```

## 10. <a name="whats-next"></a> What's next for version 1.1.0? <sup>[Top ▲](#table-of-contents)</sup>

Take a look at the [planned issues](https://github.com/pemistahl/lingua-go/milestone/1).

## 11. <a name="contributions"></a> Contributions <sup>[Top ▲](#table-of-contents)</sup>

In case you want to contribute something to *Lingua*, please take a look at the file 
[CONTRIBUTING.md](https://github.com/pemistahl/lingua-go/blob/main/CONTRIBUTING.md).
