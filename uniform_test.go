package distributions

import (
  "math"
  "testing"
)

//Test at http://keisan.casio.com/exec/system/1180573224
func Test_Uniform(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Uniform{0.0, 1.0},
      mean:       0.5,
      variance:   1.0/12.0,
      stdDev:     math.Sqrt(1.0/12.0),
      relStdDev:  math.Sqrt(1.0/12.0) * 2.0,
      skewness:   0.0,
      kurtosis:   -6.0/5.0,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0.0 },
        inOut{ in: 2.0,   out: 0.0 },
        inOut{ in: 0.5,   out: 1.0 },
        inOut{ in: 0.25,  out: 1.0 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0.0 },
        inOut{ in: 2.0,   out: 1.0 },
        inOut{ in: 0.5,   out: 0.5 },
        inOut{ in: 0.25,  out: 0.25 },
      },
    },
    distributionTest{
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
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := Uniform{0.0, 10.0}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
