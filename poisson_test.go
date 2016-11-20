package distributions

import (
  "math"
  "testing"
)

type poissonTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573179
func Test_Poisson(t *testing.T) {
  examples := []poissonTest{
    poissonTest{
      dist:       Poisson{10.0},
      mean:       10.0,
      variance:   10.0,
      stdDev:     math.Sqrt(10),
      relStdDev:  math.Sqrt(10) / 10,
      skewness:   1 / math.Sqrt(10),
      kurtosis:   0.1,
      pdf: []inOut{
        inOut{ in: 9.0,  out: 0.125110035721133298985 },
        inOut{ in: 2.0,  out: 0.00226999648812424257678 },
        inOut{ in: 4.0,  out: 0.0189166374010353548065 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,  out: 0.45792971447185220831 },
        inOut{ in: 2.0,  out: 0.00276939571551157594367 },
        inOut{ in: 4.0,  out: 0.0292526880769610726728 },
      },
    },
    poissonTest{
      dist:       Poisson{2.0},
      mean:       2.0,
      variance:   2.0,
      stdDev:     math.Sqrt(2),
      relStdDev:  math.Sqrt(2) / 2,
      skewness:   1 / math.Sqrt(2),
      kurtosis:   0.5,
      pdf: []inOut{
        inOut{ in: 1.0,  out: 0.270670566473225383788 },
        inOut{ in: 3.0,  out: 0.180447044315483589192 },
        inOut{ in: 5.0,  out: 0.0360894088630967178384 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,  out: 0.406005849709838075682 },
        inOut{ in: 3.0,  out: 0.857123460498547048662 },
        inOut{ in: 5.0,  out: 0.9834363915193855610964 },
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
    if !floatsIntegerEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
