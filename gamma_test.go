package distributions

import (
  "testing"
)

type gammaTest struct {
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

// Test at http://keisan.casio.com/exec/system/1180573217
// Test site uses Scale, which is 1/Rate.
func Test_Gamma(t *testing.T) {
  examples := []gammaTest{
    gammaTest{
      dist:       Gamma{10, 2},
      mean:       5.0,
      variance:   2.5,
      stdDev:     1.5811388300841898,
      relStdDev:  0.31622776601683794,
      skewness:   0.6324555320336759,
      kurtosis:   0.6,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  0.0,  out: 0 },
        inOut{ in:  2.0,  out: 0.0264623833821006005231 },
        inOut{ in:  4.0,  out: 0.248153834578838979826 },
        inOut{ in:  6.0,  out: 0.1747287598060988675054 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  2.0,  out: 0.00813224279693386315568 },
        inOut{ in:  4.0,  out: 0.283375741272989098481 },
        inOut{ in:  6.0,  out: 0.7576078383294876513181 },
      },
    },
    gammaTest{
      dist:       Gamma{1, 4},
      mean:       .25,
      variance:   0.0625,
      stdDev:     .25,
      relStdDev:  1,
      skewness:   2,
      kurtosis:   6,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  0.0,  out: 4 },
        inOut{ in:  0.5,  out: 0.541341132946450767576 },
        inOut{ in:  1.0,  out: 0.07326255555493672117487 },
        inOut{ in:  2.0,  out: 0.001341850511610047355286 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  0.5,  out: 0.864664716763387308106 },
        inOut{ in:  1.0,  out: 0.9816843611112658197063 },
        inOut{ in:  2.0,  out: 0.9996645373720974881612 },
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
    if !floatsIntegerEqual(example.mean, sampleMean) {
      t.Fatalf("\nSample average:\n  Expected: %f\n  Got: %f\n", example.mean, sampleMean)
    }
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsIntegerEqual(example.variance, sampleVar) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance, sampleVar)
    }
  }
}
