package distributions

import (
  "math"
  "math/rand"
)

//The Pareto Distribution is a continuous probability distribution
// with parameters α > 0, k > 0.
//
// See: https://en.wikipedia.org/wiki/Pareto_distribution
type Pareto struct {
  Scale  float64  `json:"scale"`
  Shape  float64  `json:"shape"`
}

func NewPareto(scale float64, shape float64) (Pareto, error) {
  dist := Pareto{scale, shape}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Pareto) Validate() error {
  if dist.Scale <= 0 {
    return InvalidParamsError{ "Scale must be greater than zero." }
  }
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  return nil
}

func (dist Pareto) Mean() float64 {
  if (dist.Shape <= 1.0) {
    return math.Inf(1)
  }
  result := (dist.Shape * dist.Scale) / (dist.Shape - 1)
  return result
}

func (dist Pareto) Variance() float64 {
  if (dist.Shape <= 2.0) {
    return math.Inf(1)
  }
  result := (dist.Shape * dist.Scale * dist.Scale) / ((dist.Shape - 1) * (dist.Shape - 1) * (dist.Shape - 2))
  return result
}

func (dist Pareto) Skewness() float64 {
  if (dist.Shape < 3.0) {
    return math.NaN()
  }
  result := 2 * (1 + dist.Shape) / (dist.Shape - 3) * math.Sqrt((dist.Shape - 2) / dist.Shape)
  return result
}

func (dist Pareto) Kurtosis() float64 {
  if (dist.Shape < 3.0) {
    return math.NaN()
  }
  result := 3 * (dist.Shape - 2) * ((3 * dist.Shape * dist.Shape) + dist.Shape + 2) / ((dist.Shape - 4) * (dist.Shape - 3) * dist.Shape)
  return result
}

func (dist Pareto) StdDev() float64 {
  variance := dist.Variance()
  if math.IsInf(variance, 0) {
    return math.Inf(1)
  }
  result := math.Sqrt(variance)
  return result
}

func (dist Pareto) RelStdDev() float64 {
  variance := dist.Variance()
  if math.IsInf(variance, 0) {
    return math.Inf(1)
  }
  mean := dist.Mean()
  if math.IsInf(mean, 0) {
    return math.Inf(1)
  }
  result := math.Sqrt(variance) / mean
  return result
}

func (dist Pareto) Pdf(x float64) float64 {
  if x < dist.Scale {
    return 0.0
  }
  result := dist.Shape * math.Pow(dist.Scale, dist.Shape) / math.Pow(x, dist.Shape + 1)
  return result
}

func (dist Pareto) Cdf(x float64) float64 {
  if (x < dist.Scale) {
    return 0.0
  }
  result := 1 - math.Pow(dist.Scale / x, dist.Shape)
  return result
}

func (dist Pareto) Random() float64 {
  value := dist.Scale / math.Pow(rand.Float64(), 1 / dist.Shape)
  return value
}
