package distributions

import (
  "math"
  "math/rand"
)

//The Uniform Distribution is a continuous probability distribution
// with parameters Min and Max, with Max < Min.
//
// See: https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)
type Uniform struct {
  Min  float64  `json:"min"`
  Max  float64  `json:"max"`
}

func NewUniform(min float64, max float64) (Uniform, error) {
  dist := Uniform{min, max}
  if err := dist.validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Uniform) validate() error {
  if dist.Max <= dist.Min {
    return InvalidParamsError{ "Max must be greater than Min." }
  }
  return nil
}

func (dist Uniform) Mean() float64 {
  result := (dist.Min + dist.Max) / 2
  return result
}

func (dist Uniform) Variance() float64 {
  diff := dist.Max - dist.Min
  result := diff * diff / 12
  return result
}

func (dist Uniform) Skewness() float64 {
  return 0.0
}

func (dist Uniform) Kurtosis() float64 {
  return -6.0 / 5.0
}

func (dist Uniform) StdDev() float64 {
  diff := dist.Max - dist.Min
  result := math.Sqrt(diff * diff / 12)
  return result
}

func (dist Uniform) RelStdDev() float64 {
  result := (dist.Max - dist.Min) / (math.Sqrt(3) * (dist.Max + dist.Min))
  return result
}

func (dist Uniform) Pdf(x float64) float64 {
  if x < dist.Min || x > dist.Max {
    return 0.0
  }
  result := 1 / (dist.Max - dist.Min)
  return result
}

func (dist Uniform) Cdf(x float64) float64 {
  if (x < dist.Min) {
    return 0.0
  }
  if (x >= dist.Max) {
    return 1.0
  }
  result := (x - dist.Min) / (dist.Max - dist.Min);
  return result
}

func (dist Uniform) Random() float64 {
  value := dist.Min + (rand.Float64() * (dist.Max - dist.Min))
  return value
}
