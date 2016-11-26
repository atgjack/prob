package distributions

import (
  "math"
)

//TheBeta  Distribution is a continuous probability distribution
// with parameters α > 0, β >= 0.
//
// See: https://en.wikipedia.org/wiki/Beta_distribution
type Beta struct {
  Alpha   float64
  Beta    float64
}

func (dist Beta) validate() error {
  if dist.Alpha <= 0 {
    return InvalidParamsError{ "Alpha must be greater than zero." }
  }
  if dist.Beta <= 0 {
    return InvalidParamsError{ "Beta must be greater than zero." }
  }
  return nil
}

func (dist Beta) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Alpha / (dist.Alpha + dist.Beta)
  return result, nil
}

func (dist Beta) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  comb := dist.Alpha + dist.Beta
  result := (dist.Alpha * dist.Beta) / (comb * comb * (comb + 1))
  return result, nil
}

func (dist Beta) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  numer := 2 * (dist.Beta - dist.Alpha) * math.Sqrt(dist.Alpha + dist.Beta + 1)
  result := numer / ((dist.Alpha + dist.Beta + 2) * math.Sqrt(dist.Alpha * dist.Beta))
  return result, nil
}

func (dist Beta) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  first := (dist.Alpha - dist.Beta) * (dist.Alpha - dist.Beta) * (dist.Alpha + dist.Beta + 1)
  second := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2)
  denom := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2) * (dist.Alpha + dist.Beta + 3)
  result := 6 * (first - second) / denom
  return result, nil
}

func (dist Beta) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  result := math.Sqrt(variance)
  return result, nil
}

func (dist Beta) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mean, _ := dist.Mean()
  stdDev, _ := dist.StdDev()
  result := stdDev / mean
  return result, nil
}

func (dist Beta) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  beta := BetaFn(dist.Alpha, dist.Beta)
  result := math.Pow(x, dist.Alpha - 1) * math.Pow(1 - x, dist.Beta - 1) / beta
  return result, nil
}

func (dist Beta) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := RegBetaInc(dist.Alpha, dist.Beta, x)
  return result, nil
}

// Ref: https://github.com/ampl/gsl/blob/master/randist/beta.c
func (dist Beta) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  u1, err := Gamma{ Shape: dist.Alpha, Rate: 1.0 }.Random()
  if err != nil {
    return math.NaN(), err
  }
  u2, err := Gamma{ Shape: dist.Beta, Rate: 1.0 }.Random()
  if err != nil {
    return math.NaN(), err
  }
  result := u1 / (u1 + u2)
  return result, nil
}
