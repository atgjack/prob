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
  Min  float64
  Max  float64
}

func (dist Uniform) validate() error {
  if dist.Max <= dist.Min {
    return InvalidParamsError{ "Max must be greater than Min." }
  }
  return nil
}

func (dist Uniform) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (dist.Min + dist.Max) / 2
  return result, nil
}

func (dist Uniform) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  diff := dist.Max - dist.Min
  result := diff * diff / 12
  return result, nil
}

func (dist Uniform) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return 0.0, nil
}

func (dist Uniform) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  return -6.0 / 5.0, nil
}

func (dist Uniform) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  diff := dist.Max - dist.Min
  result := math.Sqrt(diff * diff / 12)
  return result, nil
}

func (dist Uniform) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (dist.Max - dist.Min) / (math.Sqrt(3) * (dist.Max + dist.Min))
  return result, nil
}

func (dist Uniform) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < dist.Min || x > dist.Max {
    return 0.0, nil
  }
  result := 1 / (dist.Max - dist.Min)
  return result, nil
}

func (dist Uniform) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x < dist.Min) {
    return 0.0, nil
  }
  if (x >= dist.Max) {
    return 1.0, nil
  }
  result := (x - dist.Min) / (dist.Max - dist.Min);
  return result, nil
}

func (dist Uniform) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  value := dist.Min + (rand.Float64() * (dist.Max - dist.Min))
  return value, nil
}

func (dist Uniform) Sample(n int) ([]float64, error) {
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
