package distributions

import (
  "math"
  "testing"
)
// Test at http://keisan.casio.com/exec/system/1180573193
// Test at http://www.wolframalpha.com/input/?i=geometric+distribution+p%3D0.5
func Test_Geometric(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Geometric{0.5},
      mean:       1.0,
      variance:   2.0,
      stdDev:     math.Sqrt(2.0),
      relStdDev:  math.Sqrt(2.0),
      skewness:   2.12132,
      kurtosis:   6.5,
      pdf: []inOut{
        inOut{ in: 1.0,  out: 0.25 },
        inOut{ in: 3.0,  out: 0.0625 },
        inOut{ in: 5.0,  out: 0.015625 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,  out: 0.75 },
        inOut{ in: 3.0,  out: 0.9375 },
        inOut{ in: 5.0,  out: 0.984375 },
      },
    },
    distributionTest{
      dist:       Geometric{0.2},
      mean:       4.0,
      variance:   20.0,
      stdDev:     math.Sqrt(20),
      relStdDev:  math.Sqrt(20) / 4,
      skewness:   2.01246,
      kurtosis:   6.05,
      pdf: []inOut{
        inOut{ in: 1.0,  out: 0.16 },
        inOut{ in: 3.0,  out: 0.1024 },
        inOut{ in: 5.0,  out: 0.065536 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,  out: 0.36 },
        inOut{ in: 3.0,  out: 0.5904 },
        inOut{ in: 5.0,  out: 0.737856 },
      },
    },
  }

  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }

  sample := Geometric{0.4}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}

func Benchmark_Geometric(b *testing.B) {
  dist := Geometric{0.4}
  runBenchmark(b, dist)
}
