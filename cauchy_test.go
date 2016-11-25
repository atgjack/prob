package distributions

import (
  "math"
  "testing"
)

// Test at http://keisan.casio.com/exec/system/1180573169
func Test_Cauchy(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Cauchy{10.0, 2.0},
      mean:       math.NaN(),
      variance:   math.NaN(),
      stdDev:     math.NaN(),
      relStdDev:  math.NaN(),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 4.0,   out: 0.01591549430918953357689 },
        inOut{ in: 6.0,   out: 0.03183098861837906715378 },
        inOut{ in: 12.0,  out: 0.07957747154594766788444 },
      },
      cdf: []inOut{
        inOut{ in: 4.0,   out: 0.1024163823495667258246 },
        inOut{ in: 6.0,   out: 0.1475836176504332741754 },
        inOut{ in: 12.0,  out: 0.75 },
      },
      sample: sampleValues{
        mean:       10.0,
        variance:   math.NaN(),
        epsilon:    0.16,
      },
    },
    distributionTest{
      dist:       Cauchy{1.0, 4.0},
      mean:       math.NaN(),
      variance:   math.NaN(),
      stdDev:     math.NaN(),
      relStdDev:  math.NaN(),
      skewness:   math.NaN(),
      kurtosis:   math.NaN(),
      pdf: []inOut{
        inOut{ in: 2.0,   out: 0.07489644380795074624418 },
        inOut{ in: 0.5,   out: 0.07835320275293308837853 },
        inOut{ in: 8.0,   out: 0.01958830068823327209463 },
      },
      cdf: []inOut{
        inOut{ in: 2.0,   out: 0.5779791303773693254605 },
        inOut{ in: 0.5,   out: 0.4604165758394344579891 },
        inOut{ in: 8.0,   out: 0.8347506594614320903617 },
      },
      sample: sampleValues{
        mean:       math.NaN(),
        variance:   math.NaN(),
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
}
