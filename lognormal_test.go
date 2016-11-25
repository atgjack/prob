package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573225
func Test_LogNormal(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
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
      sample: sampleValues{
        mean:       12.18249396070347343807,
        variance:   255.0156343901585191873,
        epsilon:    0.015,
      },
    },
    distributionTest{
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
      sample: sampleValues{
        mean:       7.38905609893065022723,
        variance:   2926.359837008584035665,
        epsilon:    .333,
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
