package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573196
// Test at http://www.wolframalpha.com/input/?i=chi-squared+distribution+df%3D10
func Test_ChiSquared(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       ChiSquared{10.0},
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
    distributionTest{
      dist:       ChiSquared{2.0},
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
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := ChiSquared{2.0}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
