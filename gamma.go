package distributions

import (
  "math"
  "math/rand"
)

//The Gamma Distribution is a continuous probability distribution
// with parameters α > 0, β > 0.
//
// See: https://en.wikipedia.org/wiki/Gamma_distribution
type Gamma struct {
  Shape   float64
  Rate    float64
}

func (dist Gamma) validate() error {
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  if dist.Rate <= 0 {
    return InvalidParamsError{ "Rate must be greater than zero." }
  }
  return nil
}

func (dist Gamma) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Shape / dist.Rate
  return result, nil
}

func (dist Gamma) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Shape / (dist.Rate * dist.Rate)
  return result, nil
}

func (dist Gamma) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 2 / math.Sqrt(dist.Shape)
  return result, nil
}

func (dist Gamma) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 6 / dist.Shape
  return result, nil
}

func (dist Gamma) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(dist.Shape / (dist.Rate * dist.Rate))
  return result, nil
}

func (dist Gamma) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 1 / math.Sqrt(dist.Shape)
  return result, nil
}

func (dist Gamma) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0 {
    return 0.0, nil
  }
  if (x == 0) {
    if (dist.Shape == 1) {
      return dist.Rate, nil
    }
    return 0.0, nil
  }
  if (dist.Shape == 1) {
    return math.Exp((-1 * x) * dist.Rate) * dist.Rate, nil
  }
  first := (dist.Shape - 1) * math.Log(x * dist.Rate) - (x * dist.Rate)
  lgamma, _ := math.Lgamma(dist.Shape)
  result := math.Exp(first - lgamma) * dist.Rate
  return result, nil
}

func (dist Gamma) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x <= 0) {
    return 0.0, nil
  }
  result := Lowerincgamma(dist.Shape, x * dist.Rate)
  return result, nil
}

// Ref: https://github.com/ampl/gsl/blob/master/randist/gamma.c
func (dist Gamma) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (dist.Shape < 1.0) {
    random := rand.Float64()
    grandom, err := Gamma{ Shape: dist.Shape + 1.0, Rate: dist.Rate }.random()
    if err != nil {
      return math.NaN(), err
    }
    result := grandom * math.Pow(random, 1.0 / dist.Shape)
    if err != nil {
      return math.NaN(), err
    }
    return result, nil
  }
  var x, v float64
  d := dist.Shape - (1.0 / 3.0)
  c := 1.0 / math.Sqrt(9.0 * d)
  for {
    for {
      random, _, err := Normal{ Mu: 0.0, Sigma: 1.0 }.random()
      if err != nil {
        return math.NaN(), err
      }
      x = random
      v = 1.0 + (c * x)
      if v > 0.0 {
        break
      }
    }
    v = v * v * v
    u := rand.Float64()
    if u < 1.0 - 0.0331 * x * x * x * x {
      break
    }
    if math.Log(u) < (0.5 * x * x) + d * (1.0 - v + math.Log(v)) {
      break
    }
  }
  result := d * v / dist.Rate
  return result, nil
}

func (dist Gamma) Sample(n int) ([]float64, error) {
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
