package distributions

import (
  "math"
  "math/rand"
)

//The Normal(or Gaussian) Distribution is a continuous probability distribution
// with parameters μ, σ >= 0.
//
// See: https://en.wikipedia.org/wiki/Normal_distribution
type Normal struct {
  Mu      float64
  Sigma   float64
}

func (dist Normal) validate() error {
  if dist.Sigma < 0 {
    return InvalidParamsError{ "Sigma must be greater than zero." }
  }
  return nil
}

func (dist Normal) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return dist.Mu, nil
}

func (dist Normal) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := dist.Sigma * dist.Sigma
  return result, nil
}

func (dist Normal) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 0.0, nil
}

func (dist Normal) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return 3.0, nil
}

func (dist Normal) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  return dist.Sigma, nil
}

func (dist Normal) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := dist.Sigma / dist.Mu
  return result, nil
}

func (dist Normal) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  variance, err := dist.Variance()
  if err != nil {
    return 0.0, err
  }
  diff := x - dist.Mu
  expo := -1 * diff * diff / (2 * variance)
  denom := math.Sqrt(2 * variance * math.Pi)
  result := math.Exp(expo) / denom
  return result, nil
}

func (dist Normal) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  inner := 1 + math.Erf( (x - dist.Mu) / (dist.Sigma * math.Sqrt(2)) )
  result := math.Abs(inner) / 2;
  return result, nil
}

func (dist Normal) random() (float64, float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, 0.0, err
  }
  a := rand.Float64() * 2 * math.Pi
  b := math.Sqrt(-2.0 * math.Log(1.0 - rand.Float64()))
  z1 := math.Cos(a) * b
  z2 := math.Sin(a) * b
  r1 := dist.Mu + (z1 * dist.Sigma)
  r2 := dist.Mu + (z2 * dist.Sigma)
  return r1, r2, nil
}

// Generates n samples by using a Box–Muller transform.
func (dist Normal) Sample(n int) ([]float64, error) {
  if err := dist.validate(); err != nil {
    return []float64{}, err
  }
  if n <= 0 {
    return []float64{}, nil
  }
  var next, last float64
  var skipGen bool
  result := make([]float64, n)
  for i := 0; i < n; i++ {
    if (skipGen) {
      next = last
      skipGen = false
    } else {
      var err error
      next, last, err = dist.random()
      if err != nil {
        return []float64{}, err
      }
      skipGen = true
    }
    result[i] = next
  }
  return result, nil
}
