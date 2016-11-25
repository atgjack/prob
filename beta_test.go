package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573225
func Test_Beta(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Beta{1, 2},
      mean:       1.0 / 3.0,
      variance:   1.0 / 18.0,
      stdDev:     0.2357022603955158414669,
      relStdDev:  0.7071067811865475244008,
      skewness:   0.5656854249492380195207,
      kurtosis:   -3.0 / 5.0,
      pdf: []inOut{
        inOut{ in: 0.4,   out: 1.2 },
        inOut{ in: 0.6,   out: 0.8 },
        inOut{ in: 0.14,  out: 1.72 },

      },
      cdf: []inOut{
        inOut{ in: 0.4,   out: 0.64 },
        inOut{ in: 0.6,   out: 0.84 },
        inOut{ in: 0.14,  out: 0.2604 },
      },
      sample: sampleValues{
        mean:       1.0 / 3.0,
        variance:   1.0 / 18.0,
      },
    },
    distributionTest{
      dist:       Beta{5, 4},
      mean:       5.0 / 9.0,
      variance:   2.0 / 81.0,
      stdDev:     0.1571348402636772276446,
      relStdDev:  0.258198889747161125678,
      skewness:   -0.1285648693066450044365,
      kurtosis:   -21.0 / 44.0,
      pdf: []inOut{
        inOut{ in: 0.4,   out: 1.548288 },
        inOut{ in: 0.6,   out: 2.322432 },
        inOut{ in: 0.14,  out: 0.0684172364288 },
      },
      cdf: []inOut{
        inOut{ in: 0.4,   out: 0.1736704 },
        inOut{ in: 0.6,   out: 0.5940864 },
        inOut{ in: 0.14,  out: 0.002079010303104 },
      },
      sample: sampleValues{
        mean:       5.0 / 9.0,
        variance:   2.0 / 81.0,
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
