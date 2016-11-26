package distributions

import (
  "math"
  "testing"
)

// Test at http://keisan.casio.com/exec/system/1180573203
func Test_StudentsT(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       StudentsT{10.0},
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
    distributionTest{
      dist:       StudentsT{2.0},
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
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  // Using high degrees of freedom to keep variance low.
  // A custom test would be better, but there is no closed form MLE that I am aware of.
  sample := StudentsT{15.0}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
