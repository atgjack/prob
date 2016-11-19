package distributions

import (
  "math"
  "testing"
)

type weibullTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573175
// Test at http://www.wolframalpha.com/input/?i=weibull+distribution+scale%3D4+shape%3D5
// You must calculate PDF and CDF values on your own.
func Test_Weibull(t *testing.T) {
  examples := []weibullTest{
    weibullTest{
      dist:       Weibull{10, 2.5},
      mean:       8.8726382,
      variance:   14.4146689,
      stdDev:     3.7966655,
      relStdDev:  0.4279072,
      skewness:   0.3586318,
      kurtosis:   2.8567831 - 3,
      pdf: []inOut{
        inOut{ in: 2.0,   out: 0.02196423624549381206548 },
        inOut{ in: 6.0,   out: 0.08791475759141008819744 },
        inOut{ in: 10.0,  out: 0.09196986029286058039888 },

      },
      cdf: []inOut{
        inOut{ in: 2.0,   out: 0.01772949362422154647455 },
        inOut{ in: 6.0,   out: 0.2433502399169036679423 },
        inOut{ in: 10.0,  out: 0.6321205588285576784045 },
      },
    },
    weibullTest{
      dist:       Weibull{1, 4},
      mean:       0.9064025,
      variance:   0.0646615,
      stdDev:     0.2542862,
      relStdDev:  0.2805445,
      skewness:   -0.0872370,
      kurtosis:   2.7478295 - 3,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 1.471517764685769286382 },
        inOut{ in: 0.5,   out: 0.4697065314067378930599 },
        inOut{ in: 1.5,   out: 0.08545115827105757878768 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.6321205588285576784045 },
        inOut{ in: 0.5,   out: 0.06058693718652421388029 },
        inOut{ in: 1.5,   out: 0.9936702845725142534231 },
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
