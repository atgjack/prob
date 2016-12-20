package prob

import (
  "math"
)

//The Negative Binomial Distribution is a discrete probability distribution
// with parameters r > 0, 1 > p > 0.
//
// See: https://en.wikipedia.org/wiki/Negative_binomial_distribution
type NegBinomial struct {
  Failures  float64   `json:"failures"`
  Prob      float64   `json:"prob"`
}

func NewNegBinomial(failures float64, prob float64) (NegBinomial, error) {
  dist := NegBinomial{failures, prob}
  if err := dist.Validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist *NegBinomial) Validate() error {
  dist.Failures = math.Floor(dist.Failures)
  if dist.Failures < 0.0 {
    return InvalidParamsError{ "Failures must be greater than zero." }
  }
  if dist.Prob < 0.0 || dist.Prob > 1.0 {
    return InvalidParamsError{ "Prob must be between zero and one." }
  }
  return nil
}

func (dist NegBinomial) Mean() float64 {
  result := dist.Failures * dist.Prob / (1.0 - dist.Prob)
  return result
}

func (dist NegBinomial) Variance() float64 {
  result := dist.Failures * dist.Prob / ((1.0 - dist.Prob) * (1.0 - dist.Prob))
  return result
}

func (dist NegBinomial) Skewness() float64 {
  result := (1.0 + dist.Prob) / math.Sqrt(dist.Failures * dist.Prob)
  return result
}

func (dist NegBinomial) Kurtosis() float64 {
  result := (6.0 / dist.Failures) + ((1.0 - dist.Prob) * (1.0 - dist.Prob) / (dist.Failures * dist.Prob))
  return result
}

func (dist NegBinomial) StdDev() float64 {
  result := math.Sqrt(dist.Failures * (1.0 - dist.Prob)) / dist.Prob
  return result
}

func (dist NegBinomial) RelStdDev() float64 {
  mean := dist.Mean()
  stdDev := dist.StdDev()
  result := stdDev / mean
  return result
}

func (dist NegBinomial) Pdf(x float64) float64 {
  if x < 0.0 {
    return 0.0
  }
  x = math.Floor(x)
  cnk := BinomialCoefficient(x + dist.Failures - 1.0, x)
  pows := math.Pow(dist.Prob, x) * math.Pow(1.0 - dist.Prob, dist.Failures)
  result := cnk * pows
  return result
}

func (dist NegBinomial) Cdf(x float64) float64 {
  result := 1.0 - RegBetaInc(x + 1.0, dist.Failures, dist.Prob)
  return result
}

// Ref: https://github.com/ampl/gsl/blob/48fbd40c7c9c24913a68251d23bdbd0637bbda20/randist/nbinomial.c
func (dist NegBinomial) Random() float64 {
  rate := (1.0 - dist.Prob) / dist.Prob
  g := Gamma{ Shape: dist.Failures, Rate: rate }.Random()
  p := Poisson{ Mu: g }.Random()
  value := math.Floor(p + 0.5)
  return value
}
