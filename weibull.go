package distributions

import (
  "math"
  "math/rand"
)

//The Weibull Distribution is a continuous probability distribution
// with parameters α > 0, k > 0.
//
// See: https://en.wikipedia.org/wiki/Weibull_distribution
type Weibull struct {
  Scale  float64
  Shape  float64
}

func (dist Weibull) validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  return nil
}

func (dist Weibull) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Scale * math.Gamma(1 + (1 / dist.Shape))
  return result, nil
}

func (dist Weibull) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  second := math.Gamma(1 + (1 / dist.Shape))
  result := (dist.Scale * dist.Scale * math.Gamma(1 + (2 / dist.Shape))) - (second * second)
  return result, nil
}

func (dist Weibull) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  stdDev, _ := dist.StdDev()
  mean, _ := dist.Mean()
  first := (math.Gamma(1 + (3 / dist.Shape)) * dist.Scale * dist.Scale * dist.Scale)
  middle := 3 * mean * stdDev * stdDev
  last := mean * mean * mean / stdDev / stdDev / stdDev
  result := first - middle - last
  return result, nil
}

func (dist Weibull) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  stdDev, _ := dist.StdDev()
  mean, _ := dist.Mean()
  skewness, _ := dist.Skewness()
  first := math.Gamma(1 - (4 / dist.Shape)) * dist.Scale * dist.Scale * dist.Scale
  middle := 4 * skewness * stdDev * stdDev * stdDev * mean
  last := (6 * mean * mean * stdDev * stdDev) - (mean * mean * mean * mean)
  result := ((first - middle - last) / (stdDev * stdDev * stdDev * stdDev)) - 3
  return result, nil
}

func (dist Weibull) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  result := math.Sqrt(variance)
  return result, nil
}

func (dist Weibull) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  mean, _ := dist.Mean()
  result := math.Sqrt(variance) / mean
  return result, nil
}

func (dist Weibull) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0.0 {
    return 0.0, nil
  }
  result := (dist.Shape / dist.Scale) * math.Pow(x / dist.Scale, dist.Shape - 1) * math.Exp(-math.Pow(x / dist.Scale, dist.Shape))
  return result, nil
}

func (dist Weibull) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x < 0.0) {
    return 0.0, nil
  }
  result := 1 - math.Exp(-math.Pow(x / dist.Scale, dist.Shape))
  return result, nil
}

func (dist Weibull) Random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  value := dist.Scale * math.Pow(-math.Log(rand.Float64()), 1 / dist.Shape)
  return value, nil
}
