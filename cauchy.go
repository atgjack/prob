package distributions

import (
  "math"
  "math/rand"
)

//TheCauchy  Distribution is a continuous probability distribution
// with parameters x, γ >= 0.
//
// See: https://en.wikipedia.org/wiki/Cauchy_distribution
type Cauchy struct {
  Location  float64
  Scale     float64
}

func (dist Cauchy) validate() error {
  if dist.Location <= 0 {
    return InvalidParamsError{ "Location must be greater than zero." }
  }
  return nil
}

func (dist Cauchy) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, IndeterminateError
}

func (dist Cauchy) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  denom := math.Pow(x - dist.Location, 2) + math.Pow(dist.Scale, 2)
  result := dist.Scale / denom / math.Pi
  return result, nil
}

func (dist Cauchy) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := (math.Atan((x - dist.Location) / dist.Scale) / math.Pi) + 0.5
  return result, nil
}

func (dist Cauchy) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  var u float64
  for u == 0.0 || u == 0.5 {
      u = rand.Float64()
  }
  result := dist.Location + (dist.Scale * math.Atan(math.Pi * u))
  return result, nil
}

func (dist Cauchy) Sample(n int) ([]float64, error) {
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
