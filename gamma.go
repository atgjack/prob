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
  Shape  float64
  Rate    float64
}

func (dist Gamma) validate() error {
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Shape must be greater than zero." }
  }
  if dist.Shape <= 0 {
    return InvalidParamsError{ "Rate must be greater than zero." }
  }
  return nil
}

func (dist Gamma) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := dist.Shape / dist.Rate
  return result, nil
}

func (dist Gamma) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := dist.Shape / math.Pow(dist.Rate, 2)
  return result, nil
}

func (dist Gamma) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := 2 / math.Sqrt(dist.Shape)
  return result, nil
}

func (dist Gamma) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := 6 / dist.Shape
  return result, nil
}

func (dist Gamma) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := math.Sqrt(dist.Shape / math.Pow(dist.Rate, 2))
  return result, nil
}

func (dist Gamma) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  result := 1 / math.Sqrt(dist.Shape)
  return result, nil
}

func (dist Gamma) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
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
    return 0.0, err
  }
  if (x <= 0) {
    return 0.0, nil
  }
  result := Lowerincgamma(dist.Shape, x * dist.Rate)
  return result, nil
}

func (dist Gamma) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return 0.0, err
  }
  if (dist.Shape < 1) {
    random := rand.Float64()
    newDist := Gamma{ Shape: dist.Shape + 1, Rate: dist.Rate }
    grandom, err := newDist.random()
    if err != nil {
      return 0.0, err
    }
    result := grandom * math.Pow(random, 1 / dist.Shape)
    if err != nil {
      return 0.0, err
    }
    return result, nil
  }
  var x, v float64
  d := dist.Shape - (1 / 3)
  c := 1 / math.Sqrt(9 * d)
  normal := Normal{ Mu: 0, Sigma: 1 }
  for {
    for ok := true; ok; ok = v <= 0 {
      random, _, err := normal.random()
      if err != nil {
        return 0.0, err
      }
      x = random
      v = 1 + (c * x)
    }
    v = math.Pow(v, 3)
    u := rand.Float64()
    if u < 1 - 0.331 * math.Pow(x, 4) {
      break
    }
    if math.Log(u) < (0.5 * math.Pow(x, 2)) + d * (1 - v + math.Log(v)) {
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
