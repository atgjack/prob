package distributions

import (
  "math"
  "math/rand"
)

//The Poisson Distribution is a discrete probability distribution
// with parameters μ > 0.
//
// See: https://en.wikipedia.org/wiki/Poisson_distribution
type Poisson struct {
  Mu  float64   `json:"mu"`
}

func NewPoisson(mu float64) (Poisson, error) {
  dist := Poisson{ mu }
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Poisson) Validate() error {
  if dist.Mu <= 0 {
    return InvalidParamsError{ "Mu must be greater than zero." }
  }
  return nil
}

func (dist Poisson) Mean() float64 {
  return dist.Mu
}

func (dist Poisson) Variance() float64 {
  return dist.Mu
}

func (dist Poisson) Skewness() float64 {
  result := math.Pow(dist.Mu, -0.5)
  return result
}

func (dist Poisson) Kurtosis() float64 {
  result := 1 / dist.Mu
  return result
}

func (dist Poisson) StdDev() float64 {
  result := math.Sqrt(dist.Mu)
  return result
}

func (dist Poisson) RelStdDev() float64 {
  result := math.Sqrt(dist.Mu) / dist.Mu
  return result
}

func (dist Poisson) Pdf(x float64) float64 {
  lg, _ := math.Lgamma(x + 1)
  result := math.Exp((math.Log(dist.Mu) * x) -lg - dist.Mu)
  return result
}

func (dist Poisson) Cdf(x float64) float64 {
  if (x < 0.0) {
    return 0.0
  }
  result := 1 - GammaIncLower(math.Floor(x + 1), dist.Mu)
  return result
}

func (dist Poisson) Random() float64 {
  mu := dist.Mu
  k := 0.0
  for mu > 10.0 {
    m := math.Floor((mu * (7.0/8.0)) + 0.5)
    x := Gamma{ Shape: m, Rate: 1.0 }.Random()
    if x >= mu {
      rand := Binomial{ Prob: mu / x, Trials: m - 1 }.Random()
      return k + rand
    }
    k += m
    mu -= x
  }
  prod := 1.0
  emu := math.Exp(-mu)
  for ok := true; ok; {
    prod *= rand.Float64()
    k++
    ok = prod > emu
  }
  return k - 1.0
}
