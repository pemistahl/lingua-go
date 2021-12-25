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
