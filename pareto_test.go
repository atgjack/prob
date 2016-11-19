package distributions

import (
  "math"
  "testing"
)

type paretoTest struct {
  dist        Distribution
  mean        float64
  variance    float64
  stdDev      float64
  relStdDev   float64
  skewness    float64
  kurtosis    float64
  pdf         []inOut
  cdf         []inOut
}

// Test at http://www.wolframalpha.com/input/?i=pareto+distribution+k%3D4+alpha%3D5
// Calc at http://keisan.casio.com/calculator
// You must calculate PDF and CDF values on your own.
func Test_Pareto(t *testing.T) {
  examples := []paretoTest{
    paretoTest{
      dist:       Pareto{1, 2},
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
    paretoTest{
      dist:       Pareto{4, 5},
      mean:       5.0,
      variance:   1.666666666666666666667,
      stdDev:     1.290994448735805628393,
      relStdDev:  0.258198889747161125678,
      skewness:   4.64758,
      kurtosis:   70.8,
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

  for _, example := range examples {
    mean, err := example.dist.Mean()
    if err != nil || !floatsPicoEqual(mean, example.mean) {
      if !checkInf(mean, example.mean) && !checkNaN(mean, example.mean) {
        t.Fatalf("\nMean:\n  Expected: %f\n  Got: %f\n", example.mean, mean)
      }
    }
    variance, err := example.dist.Variance()
    if err != nil || !floatsPicoEqual(variance, example.variance) {
      if !checkInf(variance, example.variance) && !checkNaN(variance, example.variance) {
        t.Fatalf("\nVariance:\n  Expected: %f\n  Got: %f\n", example.variance, variance)
      }
    }
    stdDev, err := example.dist.StdDev()
    if err != nil || !floatsPicoEqual(stdDev, example.stdDev) {
      if !checkInf(stdDev, example.stdDev) && !checkNaN(stdDev, example.stdDev) {
        t.Fatalf("\nStdDev:\n  Expected: %f\n  Got: %f\n", example.stdDev, stdDev)
      }
    }
    relStdDev, err := example.dist.RelStdDev()
    if err != nil || !floatsPicoEqual(relStdDev, example.relStdDev) {
      if !checkInf(relStdDev, example.relStdDev) && !checkNaN(relStdDev, example.relStdDev) {
        t.Fatalf("\nRelStdDev:\n  Expected: %f\n  Got: %f\n", example.relStdDev, relStdDev)
      }
    }
    skewness, err := example.dist.Skewness()
    if err != nil || !floatsPicoEqual(skewness, example.skewness) {
      if !checkInf(skewness, example.skewness) && !checkNaN(skewness, example.skewness) {
        t.Fatalf("\nSkewness:\n  Expected: %f\n  Got: %f\n", example.skewness, skewness)
      }
    }
    kurtosis, err := example.dist.Kurtosis()
    if err != nil || !floatsPicoEqual(kurtosis, example.kurtosis) {
      if !checkInf(kurtosis, example.kurtosis) && !checkNaN(kurtosis, example.kurtosis) {
        t.Fatalf("\nKurtosis:\n  Expected: %f\n  Got: %f\n", example.kurtosis, kurtosis)
      }
    }
    for _, pdf := range example.pdf {
      out, err := example.dist.Pdf(pdf.in)
      if err != nil || !floatsPicoEqual(out, pdf.out) {
        t.Fatalf("\nPdf of %f:\n  Expected: %f\n  Got: %f\n", pdf.in, pdf.out, out)
      }
    }
    for _, cdf := range example.cdf {
      out, err := example.dist.Cdf(cdf.in)
      if err != nil || !floatsPicoEqual(out, cdf.out) {
        t.Fatalf("\nCdf of %f:\n  Expected: %f\n  Got: %f\n", cdf.in, cdf.out, out)
      }
    }
    samples, err := example.dist.Sample(1000000)
    if err != nil {
      t.Fatalf("\nCould not generate 1,000,000 samples.")
    }
    sampleMean := averageFloats(samples)
    if !floatsDeciEqual(example.mean, sampleMean) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", example.mean, sampleMean)
    }
    if !math.IsInf(example.variance,0) {
      sampleVar := varianceFloats(samples, sampleMean)
      if !floatsIntegerEqual(example.variance, sampleVar) {
        t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
      }
    }
  }
}
