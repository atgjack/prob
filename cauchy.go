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

func NewCauchy(location float64, scale float64) (Cauchy, error) {
  dist := Cauchy{location, scale}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Cauchy) Validate() error {
  if dist.Location <= 0 {
    return InvalidParamsError{ "Location must be greater than zero." }
  }
  return nil
}

func (dist Cauchy) Mean() float64 {
  return math.NaN()
}

func (dist Cauchy) Variance() float64 {
  return math.NaN()
}

func (dist Cauchy) Skewness() float64 {
  return math.NaN()
}

func (dist Cauchy) Kurtosis() float64 {
  return math.NaN()
}

func (dist Cauchy) StdDev() float64 {
  return math.NaN()
}

func (dist Cauchy) RelStdDev() float64 {
  return math.NaN()
}

func (dist Cauchy) Pdf(x float64) float64 {
  diff := x - dist.Location
  denom := (diff * diff) + (dist.Scale * dist.Scale)
  result := dist.Scale / denom / math.Pi
  return result
}

func (dist Cauchy) Cdf(x float64) float64 {
  result := (math.Atan((x - dist.Location) / dist.Scale) / math.Pi) + 0.5
  return result
}

func (dist Cauchy) Random() float64 {
  var u float64
  for u == 0.0 || u == 0.5 {
      u = rand.Float64()
  }
  result := dist.Location + (dist.Scale * math.Atan(math.Pi * u))
  return result
}
