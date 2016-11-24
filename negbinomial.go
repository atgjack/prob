package distributions

import (
  "math"
)

//The Negative Binomial Distribution is a discrete probability distribution
// with parameters r > 0, 1 > p > 0.
//
// See: https://en.wikipedia.org/wiki/Negative_binomial_distribution
type NegBinomial struct {
  Failures   float64
  Prob    float64
}

func (dist *NegBinomial) validate() error {
  dist.Failures = math.Floor(dist.Failures)
  if dist.Failures < 0.0 {
    return InvalidParamsError{ "Failures must be greater than zero." }
  }
  if dist.Prob < 0.0 || dist.Prob > 1.0 {
    return InvalidParamsError{ "Prob must be between zero and one." }
  }
  return nil
}

func (dist NegBinomial) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Failures * dist.Prob / (1.0 - dist.Prob)
  return result, nil
}

func (dist NegBinomial) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Failures * dist.Prob / ((1.0 - dist.Prob) * (1.0 - dist.Prob))
  return result, nil
}

func (dist NegBinomial) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (1.0 + dist.Prob) / math.Sqrt(dist.Failures * dist.Prob)
  return result, nil
}

func (dist NegBinomial) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (6.0 / dist.Failures) + ((1.0 - dist.Prob) * (1.0 - dist.Prob) / (dist.Failures * dist.Prob))
  return result, nil
}

func (dist NegBinomial) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(1.0 + dist.Prob) / (dist.Failures * dist.Prob)
  return result, nil
}

func (dist NegBinomial) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mean, _ := dist.Mean()
  stdDev, _ := dist.StdDev()
  result := mean / stdDev
  return result, nil
}

func (dist NegBinomial) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0.0 || x > dist.Failures {
    return 0.0, nil
  }
  x = math.Floor(x)
  cnk := Choose(x + dist.Failures - 1.0, x)
  pows := math.Pow(dist.Prob, x) * math.Pow(1.0 - dist.Prob, dist.Failures)
  result := cnk * pows
  return result, nil
}

func (dist NegBinomial) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 1.0 - RegBetaInc(x + 1.0, dist.Failures, dist.Prob)
  return result, nil
}

// Ref: https://github.com/ampl/gsl/blob/48fbd40c7c9c24913a68251d23bdbd0637bbda20/randist/nbinomial.c
func (dist NegBinomial) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  gamma := Gamma{ Shape: dist.Failures, Rate: 1.0 }
  g, err := gamma.random()
  if err != nil {
    return math.NaN(), err
  }
  poisson := Poisson{ Mu: g * (1 - dist.Prob) / dist.Prob }
  value, err := poisson.random()
  if err != nil {
    return math.NaN(), err
  }
  return value, nil
}

func (dist NegBinomial) Sample(n int) ([]float64, error) {
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
      return []float64{}, err
    }
    result[i] = value
  }
  return result, nil
}
