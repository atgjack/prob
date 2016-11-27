package distributions

import (
  "math"
  "math/rand"
)

//TheCauchy  Distribution is a continuous probability distribution
// with parameters x, γ >= 0.
//
// See: https://en.wikipedia.org/wiki/Cauchy_distribution
type Cauchy struct {
  Location  float64   `json:"location"`
  Scale     float64   `json:"scale"`
}

func (dist Cauchy) validate() error {
  if dist.Location <= 0 {
    return InvalidParamsError{ "Location must be greater than zero." }
  }
  return nil
}

func (dist Cauchy) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return math.NaN(), nil
}

func (dist Cauchy) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  diff := x - dist.Location
  denom := (diff * diff) + (dist.Scale * dist.Scale)
  result := dist.Scale / denom / math.Pi
  return result, nil
}

func (dist Cauchy) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (math.Atan((x - dist.Location) / dist.Scale) / math.Pi) + 0.5
  return result, nil
}

func (dist Cauchy) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  var u float64
  for u == 0.0 || u == 0.5 {
      u = rand.Float64()
  }
  result := dist.Location + (dist.Scale * math.Atan(math.Pi * u))
  return result, nil
}
