package distributions

import (
  "math"
)

//The Log-Normal Distribution is a continuous probability distribution
// with parameters μ, σ >= 0.
//
// See: https://en.wikipedia.org/wiki/LogNormal_distribution
type LogNormal struct {
  Mu      float64   `json:"mu"`
  Sigma   float64   `json:"sigma"`
}

func NewLogNormal(mu float64, sigma float64) (LogNormal, error) {
  dist := LogNormal{mu, sigma}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist LogNormal) Validate() error {
  if dist.Sigma < 0 {
    return InvalidParamsError{ "Sigma must be greater than zero." }
  }
  return nil
}

func (dist LogNormal) Mean() float64 {
  result := math.Exp(dist.Mu + (dist.Sigma * dist.Sigma / 2))
  return result
}

func (dist LogNormal) Variance() float64 {
  first := math.Exp(dist.Sigma * dist.Sigma) - 1
  second := math.Exp((2 * dist.Mu) + (dist.Sigma * dist.Sigma))
  result := first * second
  return result
}

func (dist LogNormal) Skewness() float64 {
  first := math.Exp(dist.Sigma * dist.Sigma) + 2
  second := math.Exp(dist.Sigma * dist.Sigma) - 1
  result := first * math.Sqrt(second)
  return result
}

func (dist LogNormal) Kurtosis() float64 {
  sigsqr := dist.Sigma * dist.Sigma
  result := math.Exp(4 * sigsqr) + (2 * math.Exp(3 * sigsqr)) + (3 * math.Exp(2 * sigsqr)) - 6
  return result
}

func (dist LogNormal) StdDev() float64 {
  variance := dist.Variance()
  result := math.Sqrt(variance)
  return result
}

func (dist LogNormal) RelStdDev() float64 {
  mean := dist.Mean()
  stdDev := dist.StdDev()
  result := stdDev / mean
  return result
}

func (dist LogNormal) Pdf(x float64) float64 {
  lnDiff := math.Log(x) - dist.Mu
  numer := math.Exp(-(lnDiff * lnDiff) / (2 * dist.Sigma * dist.Sigma))
  denom := x * dist.Sigma * math.Sqrt(2 * math.Pi)
  result := numer / denom
  return result
}

func (dist LogNormal) Cdf(x float64) float64 {
  erf := math.Erf((math.Log(x) - dist.Mu) / (dist.Sigma * math.Sqrt(2)))
  result := 0.5 + (0.5 * erf);
  return result
}

// A lognormal random variate is e^Normal{mu, sigma}.
func (dist LogNormal) Random() float64 {
  random := Normal{ Mu: dist.Mu, Sigma: dist.Sigma }.Random()
  value := math.Exp(random)
  return value
}
