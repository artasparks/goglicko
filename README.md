goglicko
========

A Glicko2 Implementation in Go (golang)

To install, all you should need is

    go get github.com/Kashomon/goglicko

Usage
-----

To use:

    import (
      github.com/Kashomon/goglicko
    )

    func Rate() {
      player := DefaultRating()

      opponents := []*Rating{
        NewRating(1400, 30, DefaultVol),
        NewRating(1550, 100, DefaultVol),
        NewRating(1700, 300, DefaultVol),
      }

      results := []Result{1, 0, 0}

      newRating := CalculateRating(p, o, res)
    }

Speed
-----

For the simple example above, the algorithm is quite quick:

    Benchmarks:
      BenchmarkSimpleExample	 1000000	      1089 ns/op

So roughly 1μs/call. Thus, we can expect that 1,000,000 such ratings should take
about a second.
