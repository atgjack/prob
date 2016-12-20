package prob

import (
  "math"
  "math/rand"
)

//TheLogistic  Distribution is a continuous probability distribution
// with parameters 	μ, s >= 0.
//
// See: https://en.wikipedia.org/wiki/Logistic_distribution
type Logistic struct {
  Location  float64   `json:"location"`
  Scale     float64   `json:"scale"`
}

func NewLogistic(location float64, scale float64) (Logistic, error) {
  dist := Logistic{location, scale}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Logistic) Validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  return nil
}

func (dist Logistic) Mean() float64 {
  return dist.Location
}

func (dist Logistic) Variance() float64 {
  result := math.Pi * math.Pi * dist.Scale * dist.Scale / 3
  return result
}

func (dist Logistic) Skewness() float64 {
  return 0.0
}

func (dist Logistic) Kurtosis() float64 {
  return 1.2
}

func (dist Logistic) StdDev() float64 {
  result := math.Pi * dist.Scale / math.Sqrt(3)
  return result
}

func (dist Logistic) RelStdDev() float64 {
  mean := dist.Mean()
  stdDev := dist.StdDev()
  result := stdDev / mean
  return result
}

func (dist Logistic) Pdf(x float64) float64 {
  exp := math.Exp(-(x - dist.Location) / dist.Scale)
  result := exp / (dist.Scale * (1 + exp) * (1 + exp))
  return result
}

func (dist Logistic) Cdf(x float64) float64 {
  result := 1 / (1 + math.Exp(-(x - dist.Location) / dist.Scale))
  return result
}

// Ref: http://www.stata.com/statalist/archive/2005-08/msg00131.html
func (dist Logistic) Random() float64 {
  u := rand.Float64()
  value := dist.Location - (dist.Scale * math.Log((1 / u) - 1))
  return value
}
