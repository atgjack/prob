package distributions

import (
  "math"
  "testing"
)

type studentstTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573203
func Test_StudentsT(t *testing.T) {
  examples := []studentstTest{
    studentstTest{
      dist:       StudentsT{10},
      mean:       0.0,
      variance:   1.25,
      stdDev:     1.118033988749895,
      relStdDev:  math.NaN(),
      skewness:   0.0,
      kurtosis:   4.0,
      pdf: []inOut{
        inOut{ in: 9.0,   out: 0.0000020670116801089978 },
        inOut{ in: 2.5,   out: 0.0269387276282444589776 },
        inOut{ in: 4.0,   out: 0.00203103391104121595875 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,   out: 0.9999979309754751589939 },
        inOut{ in: 2.5,   out: 0.9842765778816955978753 },
        inOut{ in: 4.0,   out: 0.9987408336876316538681 },
      },
    },
    // This is a Exponential distribution ;P
    studentstTest{
      dist:       StudentsT{2},
      mean:       0.0,
      variance:   math.Inf(1),
      stdDev:     math.Inf(1),
      relStdDev:  math.NaN(),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 9.0,   out: 0.00132246096373120901175 },
        inOut{ in: 2.5,   out: 0.0422006438680479607703 },
        inOut{ in: 4.0,   out: 0.0130945700219731023037 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,   out: 0.9939391699536065658886 },
        inOut{ in: 2.5,   out: 0.935194139889244595443 },
        inOut{ in: 4.0,   out: 0.9714045207910316829339 },
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
    if !math.IsInf(example.variance,0) {
      sampleVar := varianceFloats(samples, sampleMean)
      if !floatsEqual(example.variance, sampleVar, 1.5) {
        t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
      }
    }
  }
}
