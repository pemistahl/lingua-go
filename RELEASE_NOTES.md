## Lingua 1.3.3 (released on 03 Feb 2023)

### Bug Fixes

- For long input texts, a panic occurred while computing the confidence
  values due to an accidental division by zero. This has been fixed. (#27)

## Lingua 1.3.2 (released on 30 Jan 2023)

### Improvements

- After applying some internal optimizations, language detection is now
  faster, at least between 20% and 30%, approximately. For long input texts,
  the speed improvement is greater than for short input texts.

## Lingua 1.3.1 (released on 08 Jan 2023)

### Bug Fixes

- For long input texts, an error occurred while computing the confidence 
  values due to numerical underflow when converting probabilities. 
  This has been fixed.

## Lingua 1.3.0 (released on 01 Jan 2023)

### Improvements

- The min-max normalization method for the confidence values has been
  replaced with applying the softmax function. This gives more realistic
  probabilities. (#25)

## Lingua 1.2.2 (released on 27 Dec 2022)

### Bug Fixes

- Under certain circumstances, calling the method
  `LanguageDetector.DetectMultipleLanguagesOf()` caused an index error.
  This has been fixed.

## Lingua 1.2.1 (released on 13 Dec 2022)

### Bug Fixes

- A misconfiguration in a `go.mod` file caused errors when trying to download
  the library via the `go get` command. This has been fixed. (#23)

## Lingua 1.2.0 (released on 12 Dec 2022)

### Features

- The new method `LanguageDetector.DetectMultipleLanguagesOf()` has been
  introduced. It allows to detect multiple languages in mixed-language text. (#9)

## Lingua 1.1.1 (released on 22 Nov 2022)

### Documentation

- Some documentation mistakes have been fixed and missing information has been added.

## Lingua 1.1.0 (released on 21 Nov 2022)

### Features

- The new method `LanguageDetectorBuilder.WithLowAccuracyMode()` has been
  introduced. By activating it, detection accuracy for short text is reduced 
  in favor of a smaller memory footprint and faster detection performance. (#17)

- The new method `LanguageDetector.ComputeLanguageConfidence()` has been
  introduced. It allows to retrieve the confidence value for one specific
  language only, given the input text. (#19)

### Improvements

- The computation of the confidence values has been revised and the min-max
  normalization algorithm is now applied to the values, making them better
  comparable by behaving more like real probabilities. (#16)

- The language models are now serialized as protocol buffers instead of json.
  Thanks to this change, they are now loaded into memory twice as fast as before. (#22)

### Bug Fixes

- The unigram counts in the statistics engine were not retrieved correctly.
  This has been fixed, producing more correct detection results. (#14)

### Compatibility

- The lowest supported Go version is 1.18 now. Older versions are no longer
  compatible with this library.

### Miscellaneous

- The library now has a fresh and colorful new logo. Why? Well, why not? (-:

## Lingua 1.0.5 (released on 25 Dec 2021)

### Bug Fixes

- The character *â* was erroneously not treated as a possible indicator
  for French.

### Improvements

- The dependencies to the other language detectors which are used for
  the accuracy comparisons were always downloaded together with the main
  library. They are only needed when you want to update the accuracy reports,
  therefore the `cmd/` subdirectory now contains its own Go module that defines
  those dependencies. They have now been removed from the main library.
  Thanks to @dim and @BoeingX for identifying this problem. (#8)

## Lingua 1.0.4 (released on 28 Nov 2021)

### Bug Fixes

- It was possible to include `lingua.Unknown` in the set of input languages
  for building the language detector. It is only meant as a return value,
  so it is now automatically removed from the set of input languages.
  Thanks to @marians for identifying this problem. (#7)

## Lingua 1.0.3 (released on 20 Oct 2021)

### Improvements

- By replacing [sync.Once](https://pkg.go.dev/sync#Once) with 
  [sync.Map](https://pkg.go.dev/sync#Map) for storing the language models
  at runtime, a large amount of code could be removed while preserving 
  the same functionality. This improves code maintenance significantly.

## Lingua 1.0.2 (released on 13 Oct 2021)

### Bug Fixes

- In very rare cases, the language returned by the detector was non-deterministic.
  This has been fixed. Big thanks to @FilipAlexander for identifying this problem. (#6)

## Lingua 1.0.1 (released on 27 Jun 2021)

### Bug Fixes

- The language models were not embedded into the compiled binary. 
  This resulted in problems when trying to use Lingua within a Docker container, 
  for instance. Big thanks to @dsxack for identifying this problem and providing a fix. (#2 #3)

## Lingua 1.0.0 (released on 21 Jun 2021)

This is the very first release of the Go implementation of Lingua. Enjoy! :-)
