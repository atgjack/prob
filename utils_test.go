package distributions

import (
  "math"
  "testing"
)

type lowerIncGamma struct { s, x, out float64 }
type nChoosek struct { n, k, out float64 }
type betaFn struct { a, b, out float64 }
type betaIncFn struct { x, a, b, out float64 }

// Test at http://keisan.casio.com/exec/system/1180573447
// Have to regularize it here.
func Test_Utils_GammaIncLower(t *testing.T) {
  examples := []lowerIncGamma{
    lowerIncGamma{ 1,  2, 0.864664716763387308106  },
    lowerIncGamma{ 1,  3, 0.9502129316321360570207 },
    lowerIncGamma{ 4,  2, 0.857259237008717708028  },
    lowerIncGamma{ 4,  3, 2.116608667306612447611  },
    lowerIncGamma{ 10, 2, 16.87322146226469073825  },
    lowerIncGamma{ 10, 3, 400.0708926563052888277  },
  }
  for _, example := range examples {
    result := GammaIncLower(example.s, example.x)
    out := example.out / math.Gamma(example.s)
    if !floatsPicoEqual(result, out) {
      t.Fatalf("\n  Expected: %f\n  Got: %f\n", out, result)
    }
  }
}

func Test_Utils_BinomialCoefficient(t *testing.T) {
  examples := []nChoosek {
    nChoosek{ 10, 2,  45    },
    nChoosek{ 14, 2,  91    },
    nChoosek{ 18, 13, 8568  },
    nChoosek{ 23, 22, 23    },
    nChoosek{  9,  5, 126   },
    nChoosek{ 20, 14, 38760 },
  }
  for _, example := range examples {
    result := BinomialCoefficient(example.n, example.k)
    if result != example.out {
      t.Fatalf("\n  Expected: %f\n  Got: %f\n", example.out, result)
    }
  }
}

func Test_Utils_BetaFn(t *testing.T) {
  examples := []betaFn {
    betaFn{ 10, 2,  0.00909090909090909090909 },
    betaFn{ 14, 2,  0.00476190476190476190476 },
    betaFn{  8,  3, 0.0027777777777777777778  },
    betaFn{  4,  2, 0.05 },
    betaFn{  5,  5, 0.00158730158730158730159 },
    betaFn{  5,  1, 0.2 },
  }
  for _, example := range examples {
    result := BetaFn(example.a, example.b)
    if result != example.out {
      t.Fatalf("\n  Expected: %f\n  Got: %f\n", example.out, result)
    }
  }
}

func Test_Utils_BetaInc(t *testing.T) {
  examples := []betaIncFn {
    betaIncFn{ 0.8, 10,  2, 0.00292838679272727272727 },
    betaIncFn{ 0.9, 14,  2, 0.00261449056628125714286 },
    betaIncFn{ 0.6,  2,  4, 0.045648 },
    betaIncFn{ 0.6,  1,  4, 0.2436 },
    betaIncFn{ 0.5,  3,  5, 0.00736607142857142857143 },
    betaIncFn{ 0.2,  1, 17, 0.057498941286067501176 },
  }
  for _, example := range examples {
    result := BetaInc(example.a, example.b, example.x)
    if !floatsPicoEqual(result, example.out) {
      t.Fatalf("\n  Expected: %b\n  Got: %b\n", example.out, result)
    }
  }
}

func Test_Utils_RegBetaInc(t *testing.T) {
  examples := []betaIncFn {
    betaIncFn{ 0.8, 10,  2, 0.3221225472 },
    betaIncFn{ 0.9, 14,  2, 0.549043018919064 },
    betaIncFn{ 0.6,  2,  4, 0.91296 },
    betaIncFn{ 0.6,  1,  4, 0.9744 },
    betaIncFn{ 0.5,  3,  5, 0.7734375 },
    betaIncFn{ 0.2,  1, 17, 0.97748200186314752 },
  }
  for _, example := range examples {
    result := RegBetaInc(example.a, example.b, example.x)
    if !floatsPicoEqual(result, example.out) {
      t.Fatalf("\n  Expected: %b\n  Got: %b\n", example.out, result)
    }
  }
}
