package distributions

import (
  "math"
  "testing"
)

// Test at http://keisan.casio.com/exec/system/1180573179
func Test_Poisson(t *testing.T) {
  examples := []distributionTest{
    distributionTest{
      dist:       Poisson{10.0},
      mean:       10.0,
      variance:   10.0,
      stdDev:     math.Sqrt(10.0),
      relStdDev:  math.Sqrt(10.0) / 10.0,
      skewness:   1.0 / math.Sqrt(10.0),
      kurtosis:   0.1,
      pdf: []inOut{
        inOut{ in: 9.0,  out: 0.125110035721133298985 },
        inOut{ in: 2.0,  out: 0.00226999648812424257678 },
        inOut{ in: 4.0,  out: 0.0189166374010353548065 },
      },
      cdf: []inOut{
        inOut{ in: 9.0,  out: 0.45792971447185220831 },
        inOut{ in: 2.0,  out: 0.00276939571551157594367 },
        inOut{ in: 4.0,  out: 0.0292526880769610726728 },
      },
    },
    distributionTest{
      dist:       Poisson{2.0},
      mean:       2.0,
      variance:   2.0,
      stdDev:     math.Sqrt(2.0),
      relStdDev:  math.Sqrt(2.0) / 2.0,
      skewness:   1.0 / math.Sqrt(2.0),
      kurtosis:   0.5,
      pdf: []inOut{
        inOut{ in: 1.0,  out: 0.270670566473225383788 },
        inOut{ in: 3.0,  out: 0.180447044315483589192 },
        inOut{ in: 5.0,  out: 0.0360894088630967178384 },
      },
      cdf: []inOut{
        inOut{ in: 1.0,  out: 0.406005849709838075682 },
        inOut{ in: 3.0,  out: 0.857123460498547048662 },
        inOut{ in: 5.0,  out: 0.9834363915193855610964 },
      },
    },
  }
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  if err := testValues(examples); err != nil {
    t.Fatal(err)
  }
  
  sample := Poisson{10.0}
  if err := testSamples(sample); err != nil {
    t.Fatal(err)
  }
}
