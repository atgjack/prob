package distributions

import (
  "math"
  "testing"
  "fmt"
)

// Test at http://www.wolframalpha.com/input/?i=pareto+distribution+k%3D4+alpha%3D5
// Calc at http://keisan.casio.com/calculator
// You must calculate PDF and CDF values on your own.
func Test_Pareto(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Pareto{1.0, 2.0},
      mean:       2.0,
      variance:   math.Inf(1),
      stdDev:     math.Inf(1),
      relStdDev:  math.Inf(1),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 4.0,   out: 0.03125 },
        inOut{ in: 6.0,   out: 0.009259259259259259259259 },
        inOut{ in: 14.0,  out: 0.000728862973760932944606 },
      },
      cdf: []inOut{
        inOut{ in: 4.0,   out: 0.9375 },
        inOut{ in: 6.0,   out: 0.9722222222222222222222 },
        inOut{ in: 14.0,  out: 0.9948979591836734693878 },
      },
    },
    distributionTest{
      dist:       Pareto{4.0, 5.0},
      mean:       5.0,
      variance:   1.666666666666666666667,
      stdDev:     1.290994448735805628393,
      relStdDev:  0.258198889747161125678,
      skewness:   4.647580015448900262215,
      kurtosis:   73.8,
      pdf: []inOut{
        inOut{ in: 5.0,   out: 0.32768 },
        inOut{ in: 10.0,  out: 0.00512 },
        inOut{ in: 13.0,  out: 0.00106074220048897729328 },
      },
      cdf: []inOut{
        inOut{ in: 5.0,   out: 0.67232 },
        inOut{ in: 10.0,  out: 0.98976 },
        inOut{ in: 13.0,  out: 0.9972420702787286590375 },
      },
    },
  }

  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }

  sample := Pareto{10.0, 5.0}
  if err := estimatePareto(sample); err != nil {
    t.Fatal(err)
  }
}


// This computes and compares parameters to MLE results.
func estimatePareto(dist Pareto) error {
  samples := Sample(dist, numSamples)
  if len(samples) != numSamples {
    return fmt.Errorf("\nCould not generate samples.")
  }
  n := float64(numSamples)
  min := math.Inf(1)
  sum := 0.0
  for _, num := range samples {
    if num < min {
      min = num
    }
    sum += math.Log(num)
  }
  avg := sum / n
  lnmin := math.Log(min)
  ahat := 1 / (avg - lnmin)
  if !floatsCentiEqual(ahat, dist.Shape) {
    return fmt.Errorf("\nAhat: %f\nShape: %f\n", ahat, dist.Shape)
  }
  if !floatsCentiEqual(min, dist.Scale) {
    return fmt.Errorf("\nMin: %f\nScale: %f\n", min, dist.Scale)
  }
  return nil
}

func Benchmark_Pareto(b *testing.B) {
  dist := Pareto{10.0, 5.0}
  runBenchmark(b, dist)
}
