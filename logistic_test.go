package distributions

import "testing"

// Test at http://keisan.casio.com/exec/system/1180573209
func Test_Logistic(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Logistic{1, 2},
      mean:       1.0,
      variance:   13.15947253478581149178,
      stdDev:     3.627598728468435701188,
      relStdDev:  0.7071067811865475244008,
      skewness:   0.0,
      kurtosis:   1.2,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.125 },
        inOut{ in: 3.0,   out: 0.09830596662074092626871 },
        inOut{ in: 9.0,   out: 0.008831353106645558210781 },

      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.5 },
        inOut{ in: 3.0,   out: 0.7310585786300048792512 },
        inOut{ in: 9.0,   out: 0.9820137900379084419732 },
      },
      sample: sampleValues{
        mean:       1.0,
        variance:   13.15947253478581149178,
      },
    },
    distributionTest{
      dist:       Logistic{5, 4},
      mean:       5.0,
      variance:   52.63789013914324596712,
      stdDev:     7.255197456936871402376,
      relStdDev:  1.451039491387374280475,
      skewness:   0.0,
      kurtosis:   1.2,
      pdf: []inOut{
        inOut{ in: 1.0,   out: 0.04915298331037046313436 },
        inOut{ in: 3.0,   out: 0.05875092805039862226733 },
        inOut{ in: 9.0,   out: 0.04915298331037046313436 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,   out: 0.2689414213699951207488 },
        inOut{ in: 3.0,   out: 0.3775406687981454353611 },
        inOut{ in: 9.0,   out: 0.7310585786300048792512 },
      },
      sample: sampleValues{
        mean:       5.0,
        variance:   52.63789013914324596712,
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
