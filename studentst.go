package distributions

import (
  "math"
)

//The Student's t-Distribution is a continuous probability distribution
// with parameters df > 0.
//
// See: https://en.wikipedia.org/wiki/Student's_t-distribution
type StudentsT struct {
  Degrees  float64
}

func (dist StudentsT) validate() error {
  if dist.Degrees <= 0 {
    return InvalidParamsError{ "Degrees must be greater than zero." }
  }
  return nil
}

func (dist StudentsT) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees <= 1) {
    return math.NaN(), nil
  }
  return 0.0, nil
}

func (dist StudentsT) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees < 1) {
    return math.NaN(), nil
  }
  if (dist.Degrees <= 2) {
    return math.Inf(1), nil
  }
  result := dist.Degrees / (dist.Degrees - 2)
  return result, nil
}

func (dist StudentsT) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees <= 3) {
    return math.NaN(), nil
  }
  return 0.0, nil
}

func (dist StudentsT) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees <= 4) {
    return math.NaN(), nil
  }
  result := 3 * (dist.Degrees - 2) / (dist.Degrees - 4)
  return result, nil
}

func (dist StudentsT) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees < 1) {
    return math.NaN(), nil
  }
  if (dist.Degrees <= 2) {
    return math.Inf(1), nil
  }
  result := math.Sqrt(dist.Degrees / (dist.Degrees - 2))
  return result, nil
}

func (dist StudentsT) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist StudentsT) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  lg1, _ := math.Lgamma(dist.Degrees / 2)
  lg2, _ := math.Lgamma((dist.Degrees + 1) / 2)
  result := math.Exp(lg2 - lg1) * math.Pow(1 + (x * x / dist.Degrees), -(dist.Degrees + 1) / 2) / math.Sqrt(math.Pi * dist.Degrees)
  return result, nil
}

// Ref: https://github.com/chbrown/nlp/blob/master/src/main/java/cc/mallet/util/StatFunctions.java
func (dist StudentsT) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
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
    return result, nil
  }
  if dist.Degrees == 1 {
    s = 0
  }
  result := 0.5 + (((a * b * s) + math.Atan(a)) * g1)
  return result, nil
}

// Ref: https://github.com/ampl/gsl/blob/master/randist/tdist.c
func (dist StudentsT) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Degrees <= 2) {
    y1, err := Normal{ Mu: 0, Sigma: 1 }.Random()
    if err != nil {
      return math.NaN(), err
    }
    y2, err := ChiSquared{ Degrees: dist.Degrees }.Random()
    if err != nil {
      return math.NaN(), err
    }
    result := y1 / math.Sqrt(y2 / dist.Degrees)
    return result, nil
  } else {
    var y1, y2, z float64
    var err error
    ok := true
    for ok {
      y1, err = Normal{ Mu: 0, Sigma: 1 }.Random()
      if err != nil {
        return math.NaN(), err
      }
      y2, err = Exponential{ Lambda: 1 / ((dist.Degrees / 2) - 1) }.Random()
      if err != nil {
        return math.NaN(), err
      }
      z = y1 * y2 / (dist.Degrees - 2)
      ok = 1 - z < 0 || math.Exp(-y2 - z) > 1 - z
    }
    result := y1 / math.Sqrt((1 - (2 / dist.Degrees)) * (1 - z))
    return result, nil
  }
}
