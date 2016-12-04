package distributions

import (
  "math"
  "math/rand"
)

var next float64
var skip bool

//The Normal(or Gaussian) Distribution is a continuous probability distribution
// with parameters μ, σ >= 0.
//
// See: https://en.wikipedia.org/wiki/Normal_distribution
type Normal struct {
  Mu      float64   `json:"mu"`
  Sigma   float64   `json:"sigma"`
}

func NewNormal(mu float64, sigma float64) (Normal, error) {
  dist := Normal{mu, sigma}
  if err := dist.validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Normal) validate() error {
  if dist.Sigma < 0 {
    return InvalidParamsError{ "Sigma must be greater than zero." }
  }
  return nil
}

func (dist Normal) Mean() float64 {
  return dist.Mu
}

func (dist Normal) Variance() float64 {
  result := dist.Sigma * dist.Sigma
  return result
}

func (dist Normal) Skewness() float64 {
  return 0.0
}

func (dist Normal) Kurtosis() float64 {
  return 3.0
}

func (dist Normal) StdDev() float64 {
  return dist.Sigma
}

func (dist Normal) RelStdDev() float64 {
  result := dist.Sigma / dist.Mu
  return result
}

func (dist Normal) Pdf(x float64) float64 {
  variance := dist.Variance()
  diff := x - dist.Mu
  expo := -1 * diff * diff / (2 * variance)
  denom := math.Sqrt(2 * variance * math.Pi)
  result := math.Exp(expo) / denom
  return result
}

func (dist Normal) Cdf(x float64) float64 {
  inner := 1 + math.Erf( (x - dist.Mu) / (dist.Sigma * math.Sqrt(2)) )
  result := math.Abs(inner) / 2;
  return result
}

func (dist Normal) Random() float64 {
  // var value float64
  // if (skip) {
  //   value = dist.Mu + (next * dist.Sigma)
  //   skip = false
  // } else {
  //   a := rand.Float64() * 2 * math.Pi
  //   b := math.Sqrt(-2.0 * math.Log(1.0 - rand.Float64()))
  //   z1 := math.Cos(a) * b
  //   next = math.Sin(a) * b
  //   value = dist.Mu + (z1 * dist.Sigma)
  //   skip = true
  // }
  value := rand.NormFloat64() * dist.Sigma + dist.Mu
  return value
}
