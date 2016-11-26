package distributions

import (
  "math"
  "math/rand"
)

//TheLogistic  Distribution is a continuous probability distribution
// with parameters 	μ, s >= 0.
//
// See: https://en.wikipedia.org/wiki/Logistic_distribution
type Logistic struct {
  Location  float64
  Scale     float64
}

func (dist Logistic) validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  return nil
}

func (dist Logistic) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Location, nil
}

func (dist Logistic) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Pi * math.Pi * dist.Scale * dist.Scale / 3
  return result, nil
}

func (dist Logistic) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 0.0, nil
}

func (dist Logistic) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 1.2, nil
}

func (dist Logistic) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Pi * dist.Scale / math.Sqrt(3)
  return result, nil
}

func (dist Logistic) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mean, _ := dist.Mean()
  stdDev, _ := dist.StdDev()
  result := stdDev / mean
  return result, nil
}

func (dist Logistic) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  exp := math.Exp(-(x - dist.Location) / dist.Scale)
  result := exp / (dist.Scale * (1 + exp) * (1 + exp))
  return result, nil
}

func (dist Logistic) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 1 / (1 + math.Exp(-(x - dist.Location) / dist.Scale))
  return result, nil
}

// Ref: http://www.stata.com/statalist/archive/2005-08/msg00131.html
func (dist Logistic) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  u := rand.Float64()
  value := dist.Location - (dist.Scale * math.Log((1 / u) - 1))
  return value, nil
}
