package distributions

import (
  "testing"
)

type exponentialTest struct {
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
func Test_Exponential(t *testing.T) {
  examples := []exponentialTest{
    exponentialTest{
      dist:       Exponential{10},
      mean:       10.0,
      variance:   100.0,
      stdDev:     10.0,
      relStdDev:  1.0,
      skewness:   2.0,
      kurtosis:   9.0,
      pdf: []inOut{
        inOut{ in: 9.0,   out: 0.04065696597405991118835 },
        inOut{ in: 2.5,   out: 0.07788007830714048682452 },
        inOut{ in: 4.0,   out: 0.06703200460356393007444 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,   out: 0.5934303402594008881165 },
        inOut{ in: 2.5,   out: 0.2211992169285951317548 },
        inOut{ in: 4.0,   out: 0.3296799539643606992556 },
      },
    },
    exponentialTest{
      dist:       Exponential{2},
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
    if err != nil || mean != example.mean {
      t.Fatalf("\nMean:\n  Expected: %f\n  Got: %f\n", example.mean, mean)
    }
    variance, err := example.dist.Variance()
    if err != nil || variance != example.variance {
      t.Fatalf("\nVariance:\n  Expected: %f\n  Got: %f\n", example.variance, variance)
    }
    stdDev, err := example.dist.StdDev()
    if err != nil || stdDev != example.stdDev {
      t.Fatalf("\nStdDev:\n  Expected: %f\n  Got: %f\n", example.stdDev, stdDev)
    }
    relStdDev, err := example.dist.RelStdDev()
    if err != nil || relStdDev != example.relStdDev {
      t.Fatalf("\nRelStdDev:\n  Expected: %f\n  Got: %f\n", example.relStdDev, relStdDev)
    }
    skewness, err := example.dist.Skewness()
    if err != nil || skewness != example.skewness {
      t.Fatalf("\nSkewness:\n  Expected: %f\n  Got: %f\n", example.skewness, skewness)
    }
    kurtosis, err := example.dist.Kurtosis()
    if err != nil || kurtosis != example.kurtosis {
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
    if !floatsIntegerEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
