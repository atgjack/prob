package distributions

import (
  "math"
)

//The ChiSquared Distribution is a continuous probability distribution
// with parameters df > 0.
//
// See: https://en.wikipedia.org/wiki/Chi-squared_distribution
type ChiSquared struct {
  Degrees  float64  `json:"degrees"`
}

func (dist ChiSquared) validate() error {
  if dist.Degrees <= 0 {
    return InvalidParamsError{ "Degrees must be greater than zero." }
  }
  return nil
}

func (dist ChiSquared) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Degrees, nil
}

func (dist ChiSquared) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 2 * dist.Degrees
  return result, nil
}

func (dist ChiSquared) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Pow(2, 1.5) / math.Sqrt(dist.Degrees)
  return result, nil
}

func (dist ChiSquared) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 3 + (12 / dist.Degrees)
  return result, nil
}

func (dist ChiSquared) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(2 * dist.Degrees)
  return result, nil
}

func (dist ChiSquared) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(2 / dist.Degrees)
  return result, nil
}

func (dist ChiSquared) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0 {
    return 0.0, nil
  }
  if dist.Degrees == 2 {
    result := math.Exp(-x / 2) / 2
    return result, nil
  }
  lg, _ := math.Lgamma(dist.Degrees / 2)
  result := math.Exp((((dist.Degrees / 2) - 1) * math.Log(x / 2)) - (x / 2) - lg) / 2
  return result, nil
}

func (dist ChiSquared) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0 {
    return 0.0, nil
  }
  if dist.Degrees == 2 {
    result := 1 - math.Exp(-x / 2)
    return result, nil
  }
  result := GammaIncLower(dist.Degrees / 2, x / 2);
  return result, nil
}

func (dist ChiSquared) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  random, err := Gamma{ Shape: dist.Degrees / 2, Rate: 1.0 }.Random()
  if err != nil {
    return math.NaN(), err
  }
  value := 2 * random
  return value, nil
}
