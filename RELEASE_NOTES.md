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
