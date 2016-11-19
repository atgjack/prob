package distributions

import (
  "math"  
  "testing"
)

type lowerIncGamma struct { s, x, out float64 }

// Test at http://keisan.casio.com/exec/system/1180573447
// Have to regularize it here.
func Test_Lowerincgamma(t *testing.T) {
  examples := []lowerIncGamma{
    lowerIncGamma{ 1,  2, 0.864664716763387308106 },
    lowerIncGamma{ 1,  3, 0.9502129316321360570207 },
    lowerIncGamma{ 4,  2, 0.857259237008717708028 },
    lowerIncGamma{ 4,  3, 2.116608667306612447611 },
    lowerIncGamma{ 10, 2, 16.87322146226469073825 },
    lowerIncGamma{ 10, 3, 400.0708926563052888277 },
  }
  for _, example := range examples {
    result := Lowerincgamma(example.s, example.x)
    out := example.out / math.Gamma(example.s)
    if !floatsPicoEqual(result, out) {
      t.Fatalf("\n  Expected: %f\n  Got: %f\n", out, result)
    }
  }
}
