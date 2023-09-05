<div align="center">

  ![lingua](https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/logo.png)
  
  [![Build Status](https://github.com/pemistahl/lingua-go/workflows/build/badge.svg?branch=main)](https://github.com/pemistahl/lingua-go/actions?query=workflow%3A%22build%22+branch%3Amain)
  [![codecov](https://codecov.io/gh/pemistahl/lingua-go/branch/main/graph/badge.svg)](https://codecov.io/gh/pemistahl/lingua-go)
  [![supported languages](https://img.shields.io/badge/supported%20languages-75-green.svg)](#3-which-languages-are-supported)
  [![Go Reference](https://pkg.go.dev/badge/github.com/pemistahl/lingua-go.svg)](https://pkg.go.dev/github.com/pemistahl/lingua-go)
  [![Go Report Card](https://goreportcard.com/badge/github.com/pemistahl/lingua-go)](https://goreportcard.com/report/github.com/pemistahl/lingua-go)
  [![license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
</div>

<br>

## 1. What does this library do?

Its task is simple: It tells you which language some text is written in.
This is very useful as a preprocessing step for linguistic data in natural language
processing applications such as text classification and spell checking.
Other use cases, for instance, might include routing e-mails to the right geographically
located customer service department, based on the e-mails' languages.

## 2. Why does this library exist?

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

*Lingua* aims at eliminating these problems. She nearly does not need any configuration and
yields pretty accurate results on both long and short text, even on single words and phrases.
She draws on both rule-based and statistical methods but does not use any dictionaries of words.
She does not need a connection to any external API or service either.
Once the library has been downloaded, it can be used completely offline.

## 3. Which languages are supported?

Compared to other language detection libraries, *Lingua's* focus is on *quality over quantity*, that is,
getting detection right for a small set of languages first before adding new ones.
Currently, the following 75 languages are supported:

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

## 4. How good is it?

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
running over the data of *Lingua's* supported 75 languages. Additionally, I have added Google's 
[CLD3](https://github.com/google/cld3/) to the comparison with the help of the 
[gocld3](https://github.com/jmhodges/gocld3) bindings. Languages that are not supported
by *CLD3* or *Whatlanggo* are simply ignored during the detection process.

Each of the following sections contains two plots. The bar plot shows the detailed accuracy
results for each supported language. The box plot illustrates the distributions of the
accuracy values for each classifier. The boxes themselves represent the areas which the
middle 50 % of data lie within. Within the colored boxes, the horizontal lines mark the
median of the distributions.

### 4.1 Single word detection

<br/>

<img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/boxplot-single-words.png" alt="Single Word Detection Performance" />

<br/>

<details>
    <summary>Bar plot</summary>
    <img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/barplot-single-words.png" alt="Single Word Detection Performance" />
</details>

<br/><br/>

### 4.2 Word pair detection

<br/>

<img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/boxplot-word-pairs.png" alt="Word Pair Detection Performance" />

<br/>

<details>
    <summary>Bar plot</summary>
    <img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/barplot-word-pairs.png" alt="Word Pair Detection Performance" />
</details>

<br/><br/>

### 4.3 Sentence detection

<br/>

<img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/boxplot-sentences.png" alt="Sentence Detection Performance" />

<br/>

<details>
    <summary>Bar plot</summary>
    <img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/barplot-sentences.png" alt="Sentence Detection Performance" />
</details>

<br/><br/>

### 4.4 Average detection

<br/>

<img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/boxplot-average.png" alt="Average Detection Performance" />

<br/>

<details>
    <summary>Bar plot</summary>
    <img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/plots/barplot-average.png" alt="Average Detection Performance" />
</details>

<br/><br/>

### 4.5 Mean, median and standard deviation

The table below shows detailed statistics for each language and classifier
including mean, median and standard deviation.

<details>
  <summary>Open table</summary>
  <table>
    <tr>
        <th>Language</th>
        <th colspan="4">Average</th>
        <th colspan="4">Single Words</th>
        <th colspan="4">Word Pairs</th>
        <th colspan="4">Sentences</th>
    </tr>
    <tr>
        <th></th>
        <th>Lingua<br>(high accuracy mode)</th>
        <th>Lingua<br>(low accuracy mode)</th>
        <th>&nbsp;&nbsp;CLD3&nbsp;&nbsp;</th>
        <th>Whatlang</th>
        <th>Lingua<br>(high accuracy mode)</th>
        <th>Lingua<br>(low accuracy mode)</th>
        <th>&nbsp;&nbsp;CLD3&nbsp;&nbsp;</th>
        <th>Whatlang</th>
        <th>Lingua<br>(high accuracy mode)</th>
        <th>Lingua<br>(low accuracy mode)</th>
        <th>&nbsp;&nbsp;CLD3&nbsp;&nbsp;</th>
        <th>Whatlang</th>
        <th>Lingua<br>(high accuracy mode)</th>
        <th>Lingua<br>(low accuracy mode)</th>
        <th>&nbsp;&nbsp;CLD3&nbsp;&nbsp;</th>
        <th>Whatlang</th>
    </tr>
    <tr>
		<td>Afrikaans</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 21</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
	</tr>
	<tr>
		<td>Albanian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 18</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Arabic</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
	</tr>
	<tr>
		<td>Armenian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Azerbaijani</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
	</tr>
	<tr>
		<td>Basque</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Belarusian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
	</tr>
	<tr>
		<td>Bengali</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Bokmal</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 50</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 27</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 15</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 28</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
	</tr>
	<tr>
		<td>Bosnian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 35</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 23</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 19</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 35</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 28</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 41</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Bulgarian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 37</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
	</tr>
	<tr>
		<td>Catalan</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 19</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Chinese</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Croatian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 28</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
	</tr>
	<tr>
		<td>Czech</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 50</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 31</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
	</tr>
	<tr>
		<td>Danish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 24</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
	</tr>
	<tr>
		<td>Dutch</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
	</tr>
	<tr>
		<td>English</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 17</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 35</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
	</tr>
	<tr>
		<td>Esperanto</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
	</tr>
	<tr>
		<td>Estonian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 41</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
	</tr>
	<tr>
		<td>Finnish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
	</tr>
	<tr>
		<td>French</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 37</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
	</tr>
	<tr>
		<td>Ganda</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Georgian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>German</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
	</tr>
	<tr>
		<td>Greek</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Gujarati</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Hebrew</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
	</tr>
	<tr>
		<td>Hindi</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 11</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 27</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 20</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
	</tr>
	<tr>
		<td>Hungarian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 37</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
	</tr>
	<tr>
		<td>Icelandic</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Indonesian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
	</tr>
	<tr>
		<td>Irish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Italian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 31</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 47</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
	</tr>
	<tr>
		<td>Japanese</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
	</tr>
	<tr>
		<td>Kazakh</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Korean</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Latin</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 58</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Latvian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
	</tr>
	<tr>
		<td>Lithuanian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
	</tr>
	<tr>
		<td>Macedonian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 30</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 55</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
	</tr>
	<tr>
		<td>Malay</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 31</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 31</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 11</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 28</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 35</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Maori</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 43</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Marathi</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 16</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 30</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
	</tr>
	<tr>
		<td>Mongolian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Nynorsk</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 41</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 10</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 24</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
	</tr>
	<tr>
		<td>Persian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
	</tr>
	<tr>
		<td>Polish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
	</tr>
	<tr>
		<td>Portuguese</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 42</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 21</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
	</tr>
	<tr>
		<td>Punjabi</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Romanian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 24</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
	</tr>
	<tr>
		<td>Russian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 59</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
	</tr>
	<tr>
		<td>Serbian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 34</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
	</tr>
	<tr>
		<td>Shona</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 51</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
	</tr>
	<tr>
		<td>Slovak</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 32</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Slovene</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 29</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
	</tr>
	<tr>
		<td>Somali</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 38</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
	</tr>
	<tr>
		<td>Sotho</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 43</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 15</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Spanish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 48</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 16</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 19</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 32</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
	</tr>
	<tr>
		<td>Swahili</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 60</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 43</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 25</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Swedish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 49</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 30</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 24</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 56</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
	</tr>
	<tr>
		<td>Tagalog</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 52</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 23</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 43</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 90</td>
	</tr>
	<tr>
		<td>Tamil</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Telugu</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
	</tr>
	<tr>
		<td>Thai</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
	</tr>
	<tr>
		<td>Tsonga</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 89</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Tswana</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 88</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Turkish</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 54</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 41</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 100</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
	</tr>
	<tr>
		<td>Ukrainian</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 86</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 84</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 75</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 71</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 95</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 93</td>
	</tr>
	<tr>
		<td>Urdu</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 57</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 80</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 39</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 31</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 53</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 46</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
	</tr>
	<tr>
		<td>Vietnamese</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 73</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 79</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 76</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 26</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 36</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
	</tr>
	<tr>
		<td>Welsh</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 91</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 78</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 43</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 87</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 99</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Xhosa</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 82</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 69</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 66</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 64</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 40</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 85</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 67</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 65</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/grey.png"> -</td>
	</tr>
	<tr>
		<td>Yoruba</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 74</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 15</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 22</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 50</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 33</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 5</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 11</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 77</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 61</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 11</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/red.png"> 14</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 96</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 28</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 41</td>
	</tr>
	<tr>
		<td>Zulu</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 81</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 70</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 62</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 45</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/orange.png"> 35</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> 44</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 83</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 72</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 63</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> 68</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 97</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 94</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 92</td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> 98</td>
	</tr>
	<tr>
		<td colspan="16"></td>
	</tr>
	<tr>
		<td><strong>Mean</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>86</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>77</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>69</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>67</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>74</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>61</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> <strong>48</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/yellow.png"> <strong>48</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>89</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>78</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>67</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/lightgreen.png"> <strong>63</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>96</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>93</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>93</strong></td>
		<td><img src="https://raw.githubusercontent.com/pemistahl/lingua-go/main/cmd/images/green.png"> <strong>91</strong></td>
	</tr>
	<tr>
		<td colspan="16"></td>
	</tr>
	<tr>
		<td>Median</td>
		<td>89.0</td>
		<td>80.0</td>
		<td>68.0</td>
		<td>62.0</td>
		<td>74.0</td>
		<td>57.0</td>
		<td>41.0</td>
		<td>38.0</td>
		<td>94.0</td>
		<td>81.0</td>
		<td>66.0</td>
		<td>57.0</td>
		<td>99.0</td>
		<td>97.0</td>
		<td>98.0</td>
		<td>94.0</td>
	</tr>
	<tr>
		<td>Standard Deviation</td>
		<td>13.08</td>
		<td>17.29</td>
		<td>19.04</td>
		<td>20.22</td>
		<td>18.41</td>
		<td>24.9</td>
		<td>27.86</td>
		<td>29.28</td>
		<td>13.12</td>
		<td>18.93</td>
		<td>21.83</td>
		<td>24.22</td>
		<td>11.05</td>
		<td>11.91</td>
		<td>13.95</td>
		<td>11.24</td>
	</tr>
  </table>
</details>

## 5. Why is it better than other libraries?

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
using the respective api methods. If you know beforehand that certain languages are
never to occur in an input text, do not let those take part in the classifcation process. The filtering mechanism
of the rule-based engine is quite good, however, filtering based on your own knowledge of the input text is always preferable.

## 6. Test report generation

If you want to reproduce the accuracy results above, you can generate the test reports yourself for both classifiers
and all languages by doing:

    cd cmd
    go run accuracy_reporter.go

For *gocld3* to run successfully, you need to install the exact 
[version 3.17.3](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3) of Google's protocol buffers which is a bit
unfortunate. For each detector and language, a test report file is then written into
[`/accuracy-reports`](https://github.com/pemistahl/lingua-go/tree/main/cmd/accuracy-reports).
As an example, here is the current output of the *Lingua* German report:

```
##### German #####

>>> Accuracy on average: 89.23%

>> Detection of 1000 single words (average length: 9 chars)
Accuracy: 73.90%
Erroneously classified as Dutch: 2.30%, Danish: 2.10%, English: 2.00%, Latin: 1.90%, Bokmal: 1.60%, Basque: 1.20%, French: 1.20%, Italian: 1.20%, Esperanto: 1.10%, Swedish: 1.00%, Afrikaans: 0.80%, Tsonga: 0.70%, Nynorsk: 0.60%, Portuguese: 0.60%, Yoruba: 0.60%, Finnish: 0.50%, Sotho: 0.50%, Welsh: 0.50%, Estonian: 0.40%, Irish: 0.40%, Polish: 0.40%, Spanish: 0.40%, Swahili: 0.40%, Tswana: 0.40%, Bosnian: 0.30%, Icelandic: 0.30%, Tagalog: 0.30%, Albanian: 0.20%, Catalan: 0.20%, Croatian: 0.20%, Indonesian: 0.20%, Lithuanian: 0.20%, Maori: 0.20%, Romanian: 0.20%, Xhosa: 0.20%, Zulu: 0.20%, Latvian: 0.10%, Malay: 0.10%, Slovak: 0.10%, Slovene: 0.10%, Somali: 0.10%, Turkish: 0.10%

>> Detection of 1000 word pairs (average length: 18 chars)
Accuracy: 94.10%
Erroneously classified as Dutch: 0.90%, Latin: 0.80%, English: 0.70%, Swedish: 0.60%, Danish: 0.50%, French: 0.40%, Bokmal: 0.30%, Irish: 0.20%, Tagalog: 0.20%, Afrikaans: 0.10%, Esperanto: 0.10%, Estonian: 0.10%, Finnish: 0.10%, Italian: 0.10%, Maori: 0.10%, Nynorsk: 0.10%, Somali: 0.10%, Swahili: 0.10%, Tsonga: 0.10%, Turkish: 0.10%, Welsh: 0.10%, Zulu: 0.10%

>> Detection of 1000 sentences (average length: 111 chars)
Accuracy: 99.70%
Erroneously classified as Dutch: 0.20%, Latin: 0.10%
```

## 7. How to add it to your project?

    go get github.com/pemistahl/lingua-go

## 8. How to build?

*Lingua* requires at least Go version 1.18.

```
git clone https://github.com/pemistahl/lingua-go.git
cd lingua-go
go build
```

The source code is accompanied by an extensive unit test suite. To run the tests, simply say:

    go test

## 9. How to use?

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
        WithMinimumRelativeDistance(0.9).
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
Otherwise [`Unknown`](https://github.com/pemistahl/lingua-go/blob/main/language.go#L107) will be
returned most of the time as in the example above. This is the return value for cases where
language detection is not reliably possible. This value is not meant to be included in the set
of input languages when building the language detector. If you include it, it will be
automatically removed from the set of input languages.

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
    // English: 0.93
    // French: 0.04
    // German: 0.02
    // Spanish: 0.01
}
```

In the example above, a slice of 
[`ConfidenceValue`](https://github.com/pemistahl/lingua-go/blob/main/confidence.go#L21) 
is returned containing all possible languages sorted by their confidence value in descending 
order. Each value is a probability between 0.0 and 1.0. The probabilities of all languages 
will sum to 1.0. If the language is unambiguously identified by the rule engine, the value 1.0
will always be returned for this language. The other languages will receive a value of 0.0. 

There is also a method for returning the confidence value for one specific language only:

```go
confidence := detector.ComputeLanguageConfidence("languages are awesome", lingua.French)
fmt.Printf("%.2f", confidence)

// Output:
// 0.04
```

The value that this method computes is a number between 0.0 and 1.0.
If the language is unambiguously identified by the rule engine, the value
1.0 will always be returned. If the given language is not supported by
this detector instance, the value 0.0 will always be returned.

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

### 9.5 Low accuracy mode versus high accuracy mode

*Lingua's* high detection accuracy comes at the cost of being noticeably slower
than other language detectors. The large language models also consume significant
amounts of memory. These requirements might not be feasible for systems running low
on resources. If you want to classify mostly long texts or need to save resources,
you can enable a *low accuracy mode* that loads only a small subset of the language
models into memory:

```go
lingua.NewLanguageDetectorBuilder().
    FromAllLanguages().
    WithLowAccuracyMode().
    Build()
```

The downside of this approach is that detection accuracy for short texts consisting
of less than 120 characters will drop significantly. However, detection accuracy for
texts which are longer than 120 characters will remain mostly unaffected.

In high accuracy mode (the default), the language detector consumes approximately
1,800 MB of memory if all language models are loaded. In low accuracy mode, memory
consumption is reduced to approximately 110 MB. The goal is to further reduce memory 
consumption in later releases.

An alternative for a smaller memory footprint and faster performance is to reduce the set
of languages when building the language detector. In most cases, it is not advisable to
build the detector from all supported languages. When you have knowledge about
the texts you want to classify you can almost always rule out certain languages as impossible
or unlikely to occur.

### 9.6 Detection of multiple languages in mixed-language texts

In contrast to most other language detectors, *Lingua* is able to detect multiple languages 
in mixed-language texts. This feature can yield quite reasonable results but it is still
in an experimental state and therefore the detection result is highly dependent on the input
text. It works best in high-accuracy mode with multiple long words for each language.
The shorter the phrases and their words are, the less accurate are the results. Reducing the
set of languages when building the language detector can also improve accuracy for this task
if the languages occurring in the text are equal to the languages supported by the respective
language detector instance.

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
    }

    detector := lingua.NewLanguageDetectorBuilder().
        FromLanguages(languages...).
        Build()

    sentence := "Parlez-vous français? " + 
        "Ich spreche Französisch nur ein bisschen. " +
        "A little bit is better than nothing."

    for _, result := range detector.DetectMultipleLanguagesOf(sentence) {
        fmt.Printf("%s: '%s'\n", result.Language(), sentence[result.StartIndex():result.EndIndex()])
    }

    // Output:
    // French: 'Parlez-vous français? '
    // German: 'Ich spreche Französisch nur ein bisschen. '
    // English: 'A little bit is better than nothing.'
}
```

In the example above, a slice of [`DetectionResult`](https://github.com/pemistahl/lingua-go/blob/main/result.go#L22)
is returned. Each entry in the slice describes a contiguous single-language text section,
providing start and end indices of the respective substring.

### 9.7 Methods to build the LanguageDetector

There might be classification tasks where you know beforehand that your language data is
definitely not written in Latin, for instance. The detection accuracy can become better 
in such cases if you exclude certain languages from the decision process or just
explicitly include relevant languages:

```go
// Include all languages available in the library.
lingua.NewLanguageDetectorBuilder().FromAllLanguages()

// Include only languages that are not yet extinct (= currently excludes Latin).
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

## 10. What's next for version 1.5.0?

Take a look at the [planned issues](https://github.com/pemistahl/lingua-go/milestone/6).

## 11. Contributions

Any contributions to *Lingua* are very much appreciated. Please read the instructions
in [`CONTRIBUTING.md`](https://github.com/pemistahl/lingua-go/blob/main/CONTRIBUTING.md)
for how to add new languages to the library.
