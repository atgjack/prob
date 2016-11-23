package distributions

import (
  "testing"
)

type logNormalTest struct {
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
func Test_LogNormal(t *testing.T) {
  examples := []logNormalTest{
    logNormalTest{
      dist:       LogNormal{2.0, 1.0},
      mean:       12.18249396070347343807,
      variance:   255.0156343901585191873,
      stdDev:     15.96920894691275930978,
      relStdDev:  14.68685915835065706638,
      skewness:   6.184877138632554794835,
      kurtosis:   110.9363921763115252417,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.05399096651318805195056 },
        inOut{ in: 3.0,   out: 0.08858429229609990301838 },
        inOut{ in: 5.0,   out: 0.0739293170121196250158 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.02275013194817920720028 },
        inOut{ in: 3.0,   out: 0.1836911064379448915778 },
        inOut{ in: 5.0,   out: 0.3480604769177561100325 },
      },
    },
    logNormalTest{
      dist:       LogNormal{0.0, 2.0},
      mean:       7.38905609893065022723,
      variance:   2926.359837008584035665,
      stdDev:     54.09583936874058742019,
      relStdDev:  0.258198889747161125678,
      skewness:   414.3593433001470351088,
      kurtosis:   9220556.977307005663203,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.19947114020071633897 },
        inOut{ in: 3.0,   out: 0.05717911197597461990717 },
        inOut{ in: 5.0,   out: 0.02885967677529817685605 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.5 },
        inOut{ in: 3.0,   out: 0.7086023142840820900523 },
        inOut{ in: 5.0,   out: 0.789509060951236854941 },
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
    // Variance is off by a bit. Prob has to do with ln. Need to fix.
    sampleVar := varianceFloats(samples, sampleMean)
    if !floatsIntegerEqual(example.variance/sampleVar, 1) {
      t.Fatalf("\nSample variance:\n  Expected: %f\n  Got: %f\n", example.variance/sampleVar, sampleVar)
    }
  }
}
