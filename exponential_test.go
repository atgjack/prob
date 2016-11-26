package distributions

import "testing"

//Test at http://keisan.casio.com/exec/system/1180573224
func Test_Exponential(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
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
    distributionTest{
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
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := Exponential{4.0}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
