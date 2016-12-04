package distributions

import (
  "math"
  "math/rand"
)

//The Exponential Distribution is a continuous probability distribution
// with parameters λ > 0.
//
// See: https://en.wikipedia.org/wiki/Exponential_distribution
type Exponential struct {
  Lambda  float64   `json:"lambda"`
}

func NewExponential(lambda float64) (Exponential, error) {
  dist := Exponential{ lambda }
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Exponential) Validate() error {
  if dist.Lambda <= 0 {
    return InvalidParamsError{ "Lambda must be greater than zero." }
  }
  return nil
}

func (dist Exponential) Mean() float64 {
  return dist.Lambda
}

func (dist Exponential) Variance() float64 {
  result := dist.Lambda * dist.Lambda
  return result
}

func (dist Exponential) Skewness() float64 {
  return 2.0
}

func (dist Exponential) Kurtosis() float64 {
  return 9.0
}

func (dist Exponential) StdDev() float64 {
  return dist.Lambda
}

func (dist Exponential) RelStdDev() float64 {
  return 1.0
}

func (dist Exponential) Pdf(x float64) float64 {
  if x < 0 {
    return 0.0
  }
  result := math.Exp(-1 * x / dist.Lambda) / dist.Lambda
  return result
}

func (dist Exponential) Cdf(x float64) float64 {
  if (x <= 0) {
    return 0.0
  }
  result := 1 - math.Exp(-1 * x / dist.Lambda);
  return result
}

func (dist Exponential) Random() float64 {
  // value := -1 * dist.Lambda * math.Log1p(-1 * rand.Float64())
  value := rand.ExpFloat64() * dist.Lambda
  return value
}
