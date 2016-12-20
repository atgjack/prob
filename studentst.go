package prob

import (
  "math"
)

//The Student's t-Distribution is a continuous probability distribution
// with parameters df > 0.
//
// See: https://en.wikipedia.org/wiki/Student's_t-distribution
type StudentsT struct {
  Degrees  float64  `json:"degrees"`
}

func NewStudentsT(degrees float64) (StudentsT, error) {
  dist := StudentsT{ degrees }
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist StudentsT) Validate() error {
  if dist.Degrees <= 0 {
    return InvalidParamsError{ "Degrees must be greater than zero." }
  }
  return nil
}

func (dist StudentsT) Mean() float64 {
  if (dist.Degrees <= 1) {
    return math.NaN()
  }
  return 0.0
}

func (dist StudentsT) Variance() float64 {
  if (dist.Degrees < 1) {
    return math.NaN()
  }
  if (dist.Degrees <= 2) {
    return math.Inf(1)
  }
  result := dist.Degrees / (dist.Degrees - 2)
  return result
}

func (dist StudentsT) Skewness() float64 {
  if (dist.Degrees <= 3) {
    return math.NaN()
  }
  return 0.0
}

func (dist StudentsT) Kurtosis() float64 {
  if (dist.Degrees <= 4) {
    return math.NaN()
  }
  result := 3 * (dist.Degrees - 2) / (dist.Degrees - 4)
  return result
}

func (dist StudentsT) StdDev() float64 {
  if (dist.Degrees < 1) {
    return math.NaN()
  }
  if (dist.Degrees <= 2) {
    return math.Inf(1)
  }
  result := math.Sqrt(dist.Degrees / (dist.Degrees - 2))
  return result
}

func (dist StudentsT) RelStdDev() float64 {
  return math.NaN()
}

func (dist StudentsT) Pdf(x float64) float64 {
  lg1, _ := math.Lgamma(dist.Degrees / 2)
  lg2, _ := math.Lgamma((dist.Degrees + 1) / 2)
  result := math.Exp(lg2 - lg1) * math.Pow(1 + (x * x / dist.Degrees), -(dist.Degrees + 1) / 2) / math.Sqrt(math.Pi * dist.Degrees)
  return result
}

// Ref: https://github.com/chbrown/nlp/blob/master/src/main/java/cc/mallet/util/StatFunctions.java
func (dist StudentsT) Cdf(x float64) float64 {
  g1 := 1.0 / math.Pi
  idf := dist.Degrees
  a := x / math.Sqrt(idf)
  b := idf / (idf + x * x)
  im2 := dist.Degrees - 2.0
  ioe := math.Mod(idf, 2.0)
  s := 1.0
  c := 1.0
  idf = 1.0
  ks := 2.0 + ioe
  fk := ks
  if im2 >= 2.0 {
    for k := ks; k <= im2; k += 2.0 {
      c = c * b * (fk - 1.0) / fk
      s += c
      if s != idf {
        idf = s
        fk += 2.0
      }
    }
  }
  if ioe != 1 {
    result := 0.5 + (0.5  * a * math.Sqrt(b) * s)
    return result
  }
  if dist.Degrees == 1 {
    s = 0
  }
  result := 0.5 + (((a * b * s) + math.Atan(a)) * g1)
  return result
}

// Ref: https://github.com/ampl/gsl/blob/master/randist/tdist.c
func (dist StudentsT) Random() float64 {
  if (dist.Degrees <= 2) {
    y1 := Normal{ Mu: 0, Sigma: 1 }.Random()
    y2 := ChiSquared{ Degrees: dist.Degrees }.Random()
    result := y1 / math.Sqrt(y2 / dist.Degrees)
    return result
  } else {
    var y1, y2, z float64
    ok := true
    for ok {
      y1 = Normal{ Mu: 0, Sigma: 1 }.Random()
      y2 = Exponential{ Lambda: 1 / ((dist.Degrees / 2) - 1) }.Random()
      z = y1 * y2 / (dist.Degrees - 2)
      ok = 1 - z < 0 || math.Exp(-y2 - z) > 1 - z
    }
    result := y1 / math.Sqrt((1 - (2 / dist.Degrees)) * (1 - z))
    return result
  }
}
