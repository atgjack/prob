package distributions

import (
  "math"
  "testing"
)

type cauchyTest struct {
  magnitude   float64
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

// Test at http://keisan.casio.com/exec/system/1180573169
func Test_Cauchy(t *testing.T) {
  examples := []cauchyTest{
    cauchyTest{
      dist:       Cauchy{10.0, 2.0},
      mean:       math.NaN(),
      variance:   math.NaN(),
      stdDev:     math.NaN(),
      relStdDev:  math.NaN(),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 4.0,   out: 0.01591549430918953357689 },
        inOut{ in: 6.0,   out: 0.03183098861837906715378 },
        inOut{ in: 12.0,  out: 0.07957747154594766788444 },
      },
      cdf: []inOut{
        inOut{ in: 4.0,   out: 0.1024163823495667258246 },
        inOut{ in: 6.0,   out: 0.1475836176504332741754 },
        inOut{ in: 12.0,  out: 0.75 },
      },
    },
    cauchyTest{
      dist:       Cauchy{1.0, 4.0},
      mean:       math.NaN(),
      variance:   math.NaN(),
      stdDev:     math.NaN(),
      relStdDev:  math.NaN(),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 2.0,   out: 0.07489644380795074624418 },
        inOut{ in: 0.5,   out: 0.07835320275293308837853 },
        inOut{ in: 8.0,   out: 0.01958830068823327209463 },
      },
      cdf: []inOut{
        inOut{ in: 2.0,   out: 0.5779791303773693254605 },
        inOut{ in: 0.5,   out: 0.4604165758394344579891 },
        inOut{ in: 8.0,   out: 0.8347506594614320903617 },
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
    // The variance is undefined becuase of the heavy tails, so it's not here.
    // Currently checks expected value very arbitrarily.Open to better ideas.
    sampleMean := averageFloats(samples)
    var expectedValue float64
    if dist, ok := example.dist.(Cauchy); ok {
      expectedValue = dist.Location
    } else {
      t.Fatalf("\nCould not generate 1,000,000 samples.")
    }
    if !floatsEqual(expectedValue, sampleMean, 10) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", expectedValue, sampleMean)
    }
  }
}
