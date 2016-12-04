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

func NewChiSquared(degrees float64) (ChiSquared, error) {
  dist := ChiSquared{ degrees }
  if err := dist.validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist ChiSquared) validate() error {
  if dist.Degrees <= 0 {
    return InvalidParamsError{ "Degrees must be greater than zero." }
  }
  return nil
}

func (dist ChiSquared) Mean() float64 {
  return dist.Degrees
}

func (dist ChiSquared) Variance() float64 {
  result := 2 * dist.Degrees
  return result
}

func (dist ChiSquared) Skewness() float64 {
  result := math.Pow(2, 1.5) / math.Sqrt(dist.Degrees)
  return result
}

func (dist ChiSquared) Kurtosis() float64 {
  result := 3 + (12 / dist.Degrees)
  return result
}

func (dist ChiSquared) StdDev() float64 {
  result := math.Sqrt(2 * dist.Degrees)
  return result
}

func (dist ChiSquared) RelStdDev() float64 {
  result := math.Sqrt(2 / dist.Degrees)
  return result
}

func (dist ChiSquared) Pdf(x float64) float64 {
  if x < 0 {
    return 0.0
  }
  if dist.Degrees == 2 {
    result := math.Exp(-x / 2) / 2
    return result
  }
  lg, _ := math.Lgamma(dist.Degrees / 2)
  result := math.Exp((((dist.Degrees / 2) - 1) * math.Log(x / 2)) - (x / 2) - lg) / 2
  return result
}

func (dist ChiSquared) Cdf(x float64) float64 {
  if x < 0 {
    return 0.0
  }
  if dist.Degrees == 2 {
    result := 1 - math.Exp(-x / 2)
    return result
  }
  result := GammaIncLower(dist.Degrees / 2, x / 2);
  return result
}

func (dist ChiSquared) Random() float64 {
  random := Gamma{ Shape: dist.Degrees / 2, Rate: 1.0 }.Random()
  value := 2 * random
  return value
}
