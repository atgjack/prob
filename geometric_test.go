package distributions

import (
  "math"
  "testing"
)

type geometricTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573193
// Test at http://www.wolframalpha.com/input/?i=geometric+distribution+p%3D0.5
func Test_Geometric(t *testing.T) {
  examples := []geometricTest{
    geometricTest{
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
    geometricTest{
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
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsDeciEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
