package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573217
// Test site uses Scale, which is 1/Rate.
func Test_Gamma(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Gamma{10, 2},
      mean:       5.0,
      variance:   2.5,
      stdDev:     1.5811388300841898,
      relStdDev:  0.31622776601683794,
      skewness:   0.6324555320336759,
      kurtosis:   0.6,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  0.0,  out: 0 },
        inOut{ in:  2.0,  out: 0.0264623833821006005231 },
        inOut{ in:  4.0,  out: 0.248153834578838979826 },
        inOut{ in:  6.0,  out: 0.1747287598060988675054 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0 },
        inOut{ in:  2.0,  out: 0.00813224279693386315568 },
        inOut{ in:  4.0,  out: 0.283375741272989098481 },
        inOut{ in:  6.0,  out: 0.7576078383294876513181 },
      },
    },
    distributionTest{
      dist:       Gamma{1, 4},
      mean:       0.25,
      variance:   0.0625,
      stdDev:     0.25,
      relStdDev:  1.0,
      skewness:   2.0,
      kurtosis:   6.0,
      pdf: []inOut{
        inOut{ in: -1.0,  out: 0.0 },
        inOut{ in:  0.0,  out: 4.0 },
        inOut{ in:  0.5,  out: 0.541341132946450767576 },
        inOut{ in:  1.0,  out: 0.07326255555493672117487 },
        inOut{ in:  2.0,  out: 0.001341850511610047355286 },
      },
      cdf: []inOut{
        inOut{ in: -1.0,  out: 0.0 },
        inOut{ in:  0.5,  out: 0.864664716763387308106 },
        inOut{ in:  1.0,  out: 0.9816843611112658197063 },
        inOut{ in:  2.0,  out: 0.9996645373720974881612 },
      },
    },
  }
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := Gamma{10, 4}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
