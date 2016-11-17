package distributions

import (
  "math"
  "testing"
)

type uniformTest struct {
  epsilon     float64
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

//Test at http://keisan.casio.com/exec/system/1180573224
func Test_Uniform(t *testing.T) {
  examples := []uniformTest{
    uniformTest{
      epsilon:    .0001,
      dist:       Uniform{0.0, 1.0},
      mean:       0.5,
      variance:   1.0/12.0,
      stdDev:     math.Sqrt(1.0/12.0),
      relStdDev:  math.Sqrt(1.0/12.0) * 2.0,
      skewness:   0.0,
      kurtosis:   -6.0/5.0,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in: 2.0,   out: 0 },
        inOut{ in: 0.5,   out: 1 },
        inOut{ in: 0.25,  out: 1 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in: 2.0,   out: 1 },
        inOut{ in: 0.5,   out: 0.5 },
        inOut{ in: 0.25,  out: 0.25 },
      },
    },
    uniformTest{
      epsilon:    10.0,
      dist:       Uniform{420.0, 666.0},
      mean:       543.0,
      variance:   5043.0,
      stdDev:     math.Sqrt(5043.0),
      relStdDev:  math.Sqrt(5043.0) / 543.0,
      skewness:   0.0,
      kurtosis:   -6.0/5.0,
      pdf: []inOut{
        inOut{ in: 350.0,   out: 0.0 },
        inOut{ in: 700.0,   out: 0.0 },
        inOut{ in: 444.0,   out: 1.0 / 246.0 },
        inOut{ in: 555.0,   out: 1.0 / 246.0 },
      },
      cdf: []inOut{
        inOut{ in: 350.0,   out: 0.0 },
        inOut{ in: 700.0,   out: 1.0 },
        inOut{ in: 444.0,   out: 0.0975609756097561 },
        inOut{ in: 555.0,   out: 0.5487804878048781 },
      },
    },
  }

  for _, example := range examples {
    mean, err := example.dist.Mean()
    if err != nil || !floatsPicoEqual(mean, example.mean) {
      t.Fatalf("\nMean:\n  Expected: %f\n  Got: %f\n", example.mean, mean)
    }
    variance, err := example.dist.Variance()
    if err != nil || !floatsPicoEqual(variance, example.variance) {
      t.Fatalf("\nVariance:\n  Expected: %f\n  Got: %f\n", example.variance, variance)
    }
    stdDev, err := example.dist.StdDev()
    if err != nil || !floatsPicoEqual(stdDev, example.stdDev) {
      t.Fatalf("\nStdDev:\n  Expected: %f\n  Got: %f\n", example.stdDev, stdDev)
    }
    relStdDev, err := example.dist.RelStdDev()
    if err != nil || !floatsPicoEqual(relStdDev, example.relStdDev) {
      t.Fatalf("\nRelStdDev:\n  Expected: %f\n  Got: %f\n", example.relStdDev, relStdDev)
    }
    skewness, err := example.dist.Skewness()
    if err != nil || !floatsPicoEqual(skewness, example.skewness) {
      t.Fatalf("\nSkewness:\n  Expected: %f\n  Got: %f\n", example.skewness, skewness)
    }
    kurtosis, err := example.dist.Kurtosis()
    if err != nil || !floatsPicoEqual(kurtosis, example.kurtosis) {
      t.Fatalf("\nKurtosis:\n  Expected: %f\n  Got: %f\n", example.kurtosis, kurtosis)
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
    // Set example epsilon to adjust for magnitude difference.
    if !floatsEqual(example.variance, sampleVar, example.epsilon) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
