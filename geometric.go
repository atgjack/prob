package distributions

import (
  "math"
  "math/rand"
)

//The Geometric Distribution is a discrete probability distribution
// with parameters p > 0 >= 1. Note: this is models the number of failures
// until the first success is achieved.
//
// See: https://en.wikipedia.org/wiki/Geometric_distribution
type Geometric struct {
  Prob  float64   `json:"prob"`
}

func NewGeometric(prob float64) (Geometric, error) {
  dist := Geometric{ prob }
  if err := dist.validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist *Geometric) validate() error {
  if dist.Prob <= 0 || dist.Prob > 1 {
    return InvalidParamsError{ "Mu must be between zero and one or one." }
  }
  return nil
}

func (dist Geometric) Mean() float64 {
  result := (1 - dist.Prob) / dist.Prob
  return result
}

func (dist Geometric) Variance() float64 {
  result := (1 - dist.Prob) / (dist.Prob * dist.Prob)
  return result
}

func (dist Geometric) Skewness() float64 {
  result := (2 - dist.Prob) / math.Sqrt(1 - dist.Prob)
  return result
}

func (dist Geometric) Kurtosis() float64 {
  result := 6 + ((dist.Prob * dist.Prob) / (1 - dist.Prob))
  return result
}

func (dist Geometric) StdDev() float64 {
  result := math.Sqrt(1 - dist.Prob) / dist.Prob
  return result
}

func (dist Geometric) RelStdDev() float64 {
  result := math.Sqrt(1 - dist.Prob) * dist.Prob / (dist.Prob * (1 - dist.Prob))
  return result
}

func (dist Geometric) Pdf(x float64) float64 {
  if x < 0 {
    return 0.0
  }
  result := dist.Prob * math.Pow(1 - dist.Prob, x)
  return result
}

func (dist Geometric) Cdf(x float64) float64 {
  if (x < 0.0) {
    return 0.0
  }
  result := 1 - math.Pow(1 - dist.Prob, math.Floor(x) + 1)
  return result
}

// Ref: http://math.stackexchange.com/questions/485448/prove-the-way-to-generate-geometrically-distributed-random-numbers
func (dist Geometric) Random() float64 {
  value := math.Floor(math.Log(rand.Float64()) / math.Log(1 - dist.Prob))
  return value
}
