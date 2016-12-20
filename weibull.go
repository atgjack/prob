package prob

import (
  "math"
  "math/rand"
)

//The Weibull Distribution is a continuous probability distribution
// with parameters α > 0, k > 0.
//
// See: https://en.wikipedia.org/wiki/Weibull_distribution
type Weibull struct {
  Scale  float64  `json:"scale"`
  Shape  float64  `json:"shape"`
}

func NewWeibull(scale float64, shape float64) (Weibull, error) {
  dist := Weibull{scale, shape}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Weibull) Validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  return nil
}

func (dist Weibull) Mean() float64 {
  result := dist.Scale * math.Gamma(1 + (1 / dist.Shape))
  return result
}

func (dist Weibull) Variance() float64 {
  second := math.Gamma(1 + (1 / dist.Shape))
  result := dist.Scale * dist.Scale * (math.Gamma(1 + (2 / dist.Shape)) - (second * second))
  return result
}

func (dist Weibull) Skewness() float64 {
  stdDev := dist.StdDev()
  mean := dist.Mean()
  first := math.Gamma(1 + (3 / dist.Shape)) * dist.Scale * dist.Scale * dist.Scale
  middle := 3 * mean * stdDev * stdDev
  last := mean * mean * mean
  result := (first - middle - last) / (stdDev * stdDev * stdDev)
  return result
}

func (dist Weibull) Kurtosis() float64 {
  gamma1 := math.Gamma(1 + (1 / dist.Shape))
  gamma2 := math.Gamma(1 + (2 / dist.Shape))
  gamma3 := math.Gamma(1 + (3 / dist.Shape))
  gamma4 := math.Gamma(1 + (4 / dist.Shape))
  numer := (-6.0 * gamma1 * gamma1 * gamma1 * gamma1) + (12 * gamma1 * gamma1 * gamma2) - (3 * gamma2 * gamma2) - (4 * gamma1 * gamma3) + gamma4
  denom := gamma2 - (gamma1 * gamma1)
  result := numer / (denom * denom)
  return result
}

func (dist Weibull) StdDev() float64 {
  variance := dist.Variance()
  result := math.Sqrt(variance)
  return result
}

func (dist Weibull) RelStdDev() float64 {
  variance := dist.Variance()
  mean := dist.Mean()
  result := math.Sqrt(variance) / mean
  return result
}

func (dist Weibull) Pdf(x float64) float64 {
  if x < 0.0 {
    return 0.0
  }
  result := (dist.Shape / dist.Scale) * math.Pow(x / dist.Scale, dist.Shape - 1) * math.Exp(-math.Pow(x / dist.Scale, dist.Shape))
  return result
}

func (dist Weibull) Cdf(x float64) float64 {
  if (x < 0.0) {
    return 0.0
  }
  result := 1 - math.Exp(-math.Pow(x / dist.Scale, dist.Shape))
  return result
}

func (dist Weibull) Random() float64 {
  value := dist.Scale * math.Pow(-math.Log(rand.Float64()), 1 / dist.Shape)
  return value
}
