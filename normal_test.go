package distributions

import "testing"

//Test at http://keisan.casio.com/exec/system/1180573188
func Test_Normal(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Normal{1.0, 4.0},
      mean:       1.0,
      variance:   16.0,
      stdDev:     4.0,
      relStdDev:  4.0,
      skewness:   0.0,
      kurtosis:   3.0,
      pdf: []inOut{
        inOut{ in: -4.0,  out: 0.04566227134725547624776 },
        inOut{ in: 0.5,   out: 0.09895942173618737103265 },
        inOut{ in: 12.0,  out: 0.002273390625397763192514 },
      },
      cdf: []inOut{
        inOut{ in: -4.0,  out: 0.1056497736668552576888 },
        inOut{ in: 0.5,   out: 0.4502617751698871070207 },
        inOut{ in: 12.0,  out: 0.9970202367649454432457 },
      },
      sample: sampleValues{
        mean:       1.0,
        variance:   16.0,
      },
    },
    distributionTest{
      dist:       Normal{10.0, 2.0},
      mean:       10.0,
      variance:   4.0,
      stdDev:     2.0,
      relStdDev:  0.2,
      skewness:   0.0,
      kurtosis:   3.0,
      pdf: []inOut{
        inOut{ in: 4.0,   out: 0.002215924205969003587801 },
        inOut{ in: 6.0,   out: 0.02699548325659402597528 },
        inOut{ in: 16.0,  out: 0.002215924205969003587801 },
      },
      cdf: []inOut{
        inOut{ in: 4.0,   out: 0.001349898031630094526652 },
        inOut{ in: 6.0,   out: 0.02275013194817920720028 },
        inOut{ in: 16.0,  out: 0.9986501019683699054733 },
      },
      sample: sampleValues{
        mean:       10.0,
        variance:   4.0,
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
