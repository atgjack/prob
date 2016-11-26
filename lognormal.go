package distributions

import (
  "math"
)

//The Log-Normal Distribution is a continuous probability distribution
// with parameters μ, σ >= 0.
//
// See: https://en.wikipedia.org/wiki/LogNormal_distribution
type LogNormal struct {
  Mu      float64
  Sigma   float64
}

func (dist LogNormal) validate() error {
  if dist.Sigma < 0 {
    return InvalidParamsError{ "Sigma must be greater than zero." }
  }
  return nil
}

func (dist LogNormal) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Exp(dist.Mu + (dist.Sigma * dist.Sigma / 2))
  return result, nil
}

func (dist LogNormal) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  first := math.Exp(dist.Sigma * dist.Sigma) - 1
  second := math.Exp((2 * dist.Mu) + (dist.Sigma * dist.Sigma))
  result := first * second
  return result, nil
}

func (dist LogNormal) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  first := math.Exp(dist.Sigma * dist.Sigma) + 2
  second := math.Exp(dist.Sigma * dist.Sigma) - 1
  result := first * math.Sqrt(second)
  return result, nil
}

func (dist LogNormal) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  sigsqr := dist.Sigma * dist.Sigma
  result := math.Exp(4 * sigsqr) + (2 * math.Exp(3 * sigsqr)) + (3 * math.Exp(2 * sigsqr)) - 6
  return result, nil
}

func (dist LogNormal) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  result := math.Sqrt(variance)
  return result, nil
}

func (dist LogNormal) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mean, _ := dist.Mean()
  stdDev, _ := dist.StdDev()
  result := stdDev / mean
  return result, nil
}

func (dist LogNormal) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  lnDiff := math.Log(x) - dist.Mu
  numer := math.Exp(-(lnDiff * lnDiff) / (2 * dist.Sigma * dist.Sigma))
  denom := x * dist.Sigma * math.Sqrt(2 * math.Pi)
  result := numer / denom
  return result, nil
}

func (dist LogNormal) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  erf := math.Erf((math.Log(x) - dist.Mu) / (dist.Sigma * math.Sqrt(2)))
  result := 0.5 + (0.5 * erf);
  return result, nil
}

// A lognormal random variate is e^Normal{mu, sigma}.
func (dist LogNormal) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  random, err := Normal{ Mu: dist.Mu, Sigma: dist.Sigma }.Random()
  if err != nil {
    return math.NaN(), err
  }
  value := math.Exp(random)
  return value, nil
}
