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

func (dist Exponential) validate() error {
  if dist.Lambda <= 0 {
    return InvalidParamsError{ "Lambda must be greater than zero." }
  }
  return nil
}

func (dist Exponential) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Lambda, nil
}

func (dist Exponential) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Lambda * dist.Lambda
  return result, nil
}

func (dist Exponential) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 2.0, nil
}

func (dist Exponential) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 9.0, nil
}

func (dist Exponential) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Lambda, nil
}

func (dist Exponential) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 1.0, nil
}

func (dist Exponential) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0 {
    return 0.0, nil
  }
  result := math.Exp(-1 * x / dist.Lambda) / dist.Lambda
  return result, nil
}

func (dist Exponential) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x <= 0) {
    return 0.0, nil
  }
  result := 1 - math.Exp(-1 * x / dist.Lambda);
  return result, nil
}

func (dist Exponential) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  // value := -1 * dist.Lambda * math.Log1p(-1 * rand.Float64())
  value := rand.ExpFloat64() * dist.Lambda
  return value, nil
}
