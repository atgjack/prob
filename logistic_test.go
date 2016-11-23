package distributions

import (
  "testing"
)

type logisticTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573209
func Test_Logistic(t *testing.T) {
  examples := []logisticTest{
    logisticTest{
      dist:       Logistic{1, 2},
      mean:       1.0,
      variance:   13.15947253478581149178,
      stdDev:     3.627598728468435701188,
      relStdDev:  0.7071067811865475244008,
      skewness:   0.0,
      kurtosis:   1.2,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.125 },
        inOut{ in: 3.0,   out: 0.09830596662074092626871 },
        inOut{ in: 9.0,   out: 0.008831353106645558210781 },

      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.5 },
        inOut{ in: 3.0,   out: 0.7310585786300048792512 },
        inOut{ in: 9.0,   out: 0.9820137900379084419732 },
      },
    },
    logisticTest{
      dist:       Logistic{5, 4},
      mean:       5.0,
      variance:   52.63789013914324596712,
      stdDev:     7.255197456936871402376,
      relStdDev:  1.451039491387374280475,
      skewness:   0.0,
      kurtosis:   1.2,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.04915298331037046313436 },
        inOut{ in: 3.0,   out: 0.05875092805039862226733 },
        inOut{ in: 9.0,   out: 0.04915298331037046313436 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.2689414213699951207488 },
        inOut{ in: 3.0,   out: 0.3775406687981454353611 },
        inOut{ in: 9.0,   out: 0.7310585786300048792512 },
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
