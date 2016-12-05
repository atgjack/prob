package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573199
func Test_Binomial(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       &Binomial{10.0, 0.5},
      mean:       5.0,
      variance:   2.5,
      stdDev:     1.581138830084189665999,
      relStdDev:  0.3162277660168379331999,
      skewness:   0.0,
      kurtosis:   2.8,
      pdf: []inOut{
        inOut{ in: 0.0,  out: 0.0009765625 },
        inOut{ in: 1.0,  out: 0.009765625 },
        inOut{ in: 5.0,  out: 0.24609375 },
      },
      cdf: []inOut{
        inOut{ in: 0.0,  out: 0.0009765625 },
        inOut{ in: 1.0,  out: 0.0107421875 },
        inOut{ in: 5.0,  out: 0.623046875 },
      },
    },
  }

  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }

  sample := &Binomial{10.0, 0.5}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}

func Benchmark_Binomial(b *testing.B) {
  dist := &Binomial{10.0, 0.5}
  runBenchmark(b, dist)
}
