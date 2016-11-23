package distributions

import (
  "testing"
)

type betaTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573225
func Test_Beta(t *testing.T) {
  examples := []betaTest{
    betaTest{
      dist:       Beta{1, 2},
      mean:       1.0 / 3.0,
      variance:   1.0 / 18.0,
      stdDev:     0.2357022603955158414669,
      relStdDev:  0.7071067811865475244008,
      skewness:   0.5656854249492380195207,
      kurtosis:   -3.0 / 5.0,
      pdf: []inOut{
        inOut{ in: 0.4,   out: 1.2 },
        inOut{ in: 0.6,   out: 0.8 },
        inOut{ in: 0.14,  out: 1.72 },

      },
      cdf: []inOut{
        inOut{ in: 0.4,   out: 0.64 },
        inOut{ in: 0.6,   out: 0.84 },
        inOut{ in: 0.14,  out: 0.2604 },
      },
    },
    betaTest{
      dist:       Beta{5, 4},
      mean:       5.0 / 9.0,
      variance:   2.0 / 81.0,
      stdDev:     0.1571348402636772276446,
      relStdDev:  0.258198889747161125678,
      skewness:   -0.1285648693066450044365,
      kurtosis:   -21.0 / 44.0,
      pdf: []inOut{
        inOut{ in: 0.4,   out: 1.548288 },
        inOut{ in: 0.6,   out: 2.322432 },
        inOut{ in: 0.14,  out: 0.0684172364288 },

      },
      cdf: []inOut{
        inOut{ in: 0.4,   out: 0.1736704 },
        inOut{ in: 0.6,   out: 0.5940864 },
        inOut{ in: 0.14,  out: 0.002079010303104 },
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
    if !floatsIntegerEqual(example.mean, sampleMean) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", example.mean, sampleMean)
    }
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsIntegerEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
