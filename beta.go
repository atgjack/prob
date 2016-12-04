package distributions

import (
  "math"
)

//TheBeta  Distribution is a continuous probability distribution
// with parameters α > 0, β >= 0.
//
// See: https://en.wikipedia.org/wiki/Beta_distribution
type Beta struct {
  Alpha   float64   `json:"alpha"`
  Beta    float64   `json:"beta"`
}

func NewBeta(alpha float64, beta float64) (Beta, error) {
  dist := Beta{alpha, beta}
  if err := dist.validate(); err != nil {
    return dist, err
  }
  return dist, nil
}

func (dist Beta) validate() error {
  if dist.Alpha <= 0 {
    return InvalidParamsError{ "Alpha must be greater than zero." }
  }
  if dist.Beta <= 0 {
    return InvalidParamsError{ "Beta must be greater than zero." }
  }
  return nil
}

func (dist Beta) Mean() float64 {
  result := dist.Alpha / (dist.Alpha + dist.Beta)
  return result
}

func (dist Beta) Variance() float64 {
  comb := dist.Alpha + dist.Beta
  result := (dist.Alpha * dist.Beta) / (comb * comb * (comb + 1))
  return result
}

func (dist Beta) Skewness() float64 {
  numer := 2 * (dist.Beta - dist.Alpha) * math.Sqrt(dist.Alpha + dist.Beta + 1)
  result := numer / ((dist.Alpha + dist.Beta + 2) * math.Sqrt(dist.Alpha * dist.Beta))
  return result
}

func (dist Beta) Kurtosis() float64 {
  first := (dist.Alpha - dist.Beta) * (dist.Alpha - dist.Beta) * (dist.Alpha + dist.Beta + 1)
  second := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2)
  denom := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2) * (dist.Alpha + dist.Beta + 3)
  result := 6 * (first - second) / denom
  return result
}

func (dist Beta) StdDev() float64 {
  variance := dist.Variance()
  result := math.Sqrt(variance)
  return result
}

func (dist Beta) RelStdDev() float64 {
  mean := dist.Mean()
  stdDev := dist.StdDev()
  result := stdDev / mean
  return result
}

func (dist Beta) Pdf(x float64) float64 {
  beta := BetaFn(dist.Alpha, dist.Beta)
  result := math.Pow(x, dist.Alpha - 1) * math.Pow(1 - x, dist.Beta - 1) / beta
  return result
}

func (dist Beta) Cdf(x float64) float64 {
  result := RegBetaInc(dist.Alpha, dist.Beta, x)
  return result
}

// Ref: https://github.com/ampl/gsl/blob/master/randist/beta.c
func (dist Beta) Random() float64 {
  u1 := Gamma{ Shape: dist.Alpha, Rate: 1.0 }.Random()
  u2 := Gamma{ Shape: dist.Beta, Rate: 1.0 }.Random()
  result := u1 / (u1 + u2)
  return result
}
