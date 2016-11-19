package distributions

import (
  "testing"
)

type chiSquaredTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573196
// Test at http://www.wolframalpha.com/input/?i=chi-squared+distribution+df%3D10
func Test_ChiSquared(t *testing.T) {
  examples := []chiSquaredTest{
    chiSquaredTest{
      dist:       ChiSquared{10},
      mean:       10.0,
      variance:   20.0,
      stdDev:     4.472135954999579392818,
      relStdDev:  0.4472135954999579392818,
      skewness:   0.8944271909999158785637,
      kurtosis:   4.2,
      pdf: []inOut{
        inOut{ in: 9.0,   out: 0.094903810270062204324 },
        inOut{ in: 2.5,   out: 0.0145723875356135101484 },
        inOut{ in: 4.0,   out: 0.045111761078870897298 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,   out: 0.467896423625284522439 },
        inOut{ in: 2.5,   out: 0.009124279218395273144 },
        inOut{ in: 4.0,   out: 0.052653017343711156742 },
      },
    },
    // This is a Exponential distribution ;P
    chiSquaredTest{
      dist:       ChiSquared{2},
      mean:       2.0,
      variance:   4.0,
      stdDev:     2.0,
      relStdDev:  1.0,
      skewness:   2.0,
      kurtosis:   9.0,
      pdf: []inOut{
        inOut{ in: 9.0,   out: 0.005554498269121153248072 },
        inOut{ in: 2.5,   out: 0.1432523984300950501624 },
        inOut{ in: 4.0,   out: 0.067667641618306345947 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,   out: 0.9888910034617576935039 },
        inOut{ in: 2.5,   out: 0.7134952031398098996751 },
        inOut{ in: 4.0,   out: 0.864664716763387308106 },
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
    if !floatsIntegerEqual(example.mean, sampleMean) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", example.mean, sampleMean)
    }
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsEqual(example.variance, sampleVar, 10) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
