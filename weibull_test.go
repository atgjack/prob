package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573175
// Test at http://www.wolframalpha.com/input/?i=weibull+distribution+scale%3D4+shape%3D5
// You must calculate PDF and CDF values on your own.
func Test_Weibull(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Weibull{10, 2.5},
      mean:       8.8726382,
      variance:   14.4146689,
      stdDev:     3.7966655,
      relStdDev:  0.4279072,
      skewness:   0.3586318,
      kurtosis:   2.8567831 - 3,
      pdf: []inOut{
        inOut{ in: 2.0,   out: 0.02196423624549381206548 },
        inOut{ in: 6.0,   out: 0.08791475759141008819744 },
        inOut{ in: 10.0,  out: 0.09196986029286058039888 },

      },
      cdf: []inOut{
        inOut{ in: 2.0,   out: 0.01772949362422154647455 },
        inOut{ in: 6.0,   out: 0.2433502399169036679423 },
        inOut{ in: 10.0,  out: 0.6321205588285576784045 },
      },
      sample: sampleValues{
        mean:       8.8726382,
        variance:   14.4146689,
      },
    },
    distributionTest{
      dist:       Weibull{1, 4},
      mean:       0.9064025,
      variance:   0.0646615,
      stdDev:     0.2542862,
      relStdDev:  0.2805445,
      skewness:   -0.0872370,
      kurtosis:   2.7478295 - 3,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 1.471517764685769286382 },
        inOut{ in: 0.5,   out: 0.4697065314067378930599 },
        inOut{ in: 1.5,   out: 0.08545115827105757878768 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.6321205588285576784045 },
        inOut{ in: 0.5,   out: 0.06058693718652421388029 },
        inOut{ in: 1.5,   out: 0.9936702845725142534231 },
      },
      sample: sampleValues{
        mean:       0.9064025,
        variance:   0.0646615,
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
