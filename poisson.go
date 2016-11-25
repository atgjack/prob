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
  Mu  float64
}

func (dist *Poisson) validate() error {
  dist.Mu = math.Floor(dist.Mu)
  if dist.Mu <= 0 {
    return InvalidParamsError{ "Mu must be greater than zero." }
  }
  return nil
}

func (dist Poisson) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Mu, nil
}

func (dist Poisson) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return dist.Mu, nil
}

func (dist Poisson) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Pow(dist.Mu, -0.5)
  return result, nil
}

func (dist Poisson) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 1 / dist.Mu
  return result, nil
}

func (dist Poisson) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(dist.Mu)
  return result, nil
}

func (dist Poisson) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(dist.Mu) / dist.Mu
  return result, nil
}

func (dist Poisson) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  lg, _ := math.Lgamma(x + 1)
  result := math.Exp((math.Log(dist.Mu) * x) -lg - dist.Mu)
  return result, nil
}

func (dist Poisson) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x < 0.0) {
    return 0.0, nil
  }
  result := 1 - Lowerincgamma(math.Floor(x + 1), dist.Mu)
  return result, nil
}

func (dist Poisson) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mu := dist.Mu
  k := 0.0
  for mu > 10.0 {
    m := math.Floor((mu * (7.0/8.0)) + 0.5)
    x, err := Gamma{ Shape: m, Rate: 1.0 }.random()
    if err != nil {
      return math.NaN(), err
    }
    if x >= mu {
      rand, err := Binomial{ Prob: mu / x, Trials: m - 1 }.random()
      if err != nil {
        return math.NaN(), err
      }
      return k + rand, nil
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
  return k - 1.0, nil
}

func (dist Poisson) Sample(n int) ([]float64, error) {
  if err := dist.validate(); err != nil {
    return []float64{}, err
  }
  if n <= 0 {
    return []float64{}, nil
  }
  result := make([]float64, n)
  for i := 0; i < n; i++ {
    value, err := dist.random()
    if err != nil {
      return []float64{}, nil
    }
    result[i] = value
  }
  return result, nil
}
