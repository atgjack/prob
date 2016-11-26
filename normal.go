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
  Mu      float64
  Sigma   float64
}

func (dist Normal) validate() error {
  if dist.Sigma < 0 {
    return InvalidParamsError{ "Sigma must be greater than zero." }
  }
  return nil
}

func (dist Normal) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Mu, nil
}

func (dist Normal) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Sigma * dist.Sigma
  return result, nil
}

func (dist Normal) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 0.0, nil
}

func (dist Normal) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 3.0, nil
}

func (dist Normal) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Sigma, nil
}

func (dist Normal) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Sigma / dist.Mu
  return result, nil
}

func (dist Normal) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, err := dist.Variance()
  if err != nil {
    return math.NaN(), err
  }
  diff := x - dist.Mu
  expo := -1 * diff * diff / (2 * variance)
  denom := math.Sqrt(2 * variance * math.Pi)
  result := math.Exp(expo) / denom
  return result, nil
}

func (dist Normal) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  inner := 1 + math.Erf( (x - dist.Mu) / (dist.Sigma * math.Sqrt(2)) )
  result := math.Abs(inner) / 2;
  return result, nil
}

func (dist Normal) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  var value float64
  if (skip) {
    value = dist.Mu + (next * dist.Sigma)
    skip = false
  } else {
    a := rand.Float64() * 2 * math.Pi
    b := math.Sqrt(-2.0 * math.Log(1.0 - rand.Float64()))
    z1 := math.Cos(a) * b
    next = math.Sin(a) * b
    value = dist.Mu + (z1 * dist.Sigma)
    skip = true
  }
  return value, nil
}
