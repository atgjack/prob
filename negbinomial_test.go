package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573210
func Test_NegBinomial(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       NegBinomial{10.0, 0.5},
      mean:       10.0,
      variance:   20.0,
      stdDev:     4.472135954999579392818,
      relStdDev:  0.4472135954999579392818,
      skewness:   0.6708203932499369089228,
      kurtosis:   0.65,
      pdf: []inOut{
        inOut{ in: 1.0,  out: 0.0048828125 },
        inOut{ in: 3.0,  out: 0.02685546875 },
        inOut{ in: 5.0,  out: 0.06109619140625 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,  out: 0.005859375 },
        inOut{ in: 3.0,  out: 0.046142578125 },
        inOut{ in: 5.0,  out: 0.15087890625 },
      },
    },
  }
  
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := NegBinomial{10.0, 0.5}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
