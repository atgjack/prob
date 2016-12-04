package distributions

import (
  "math"
)

// Distirbution is an interface for impementing continuous probability distributions.
//
// See: https://en.wikipedia.org/wiki/Probability_distribution
type Distribution interface {
  Validate()    error
  Mean()        float64
  Variance()    float64
  Kurtosis()    float64
  Skewness()    float64
  StdDev()      float64
  RelStdDev()   float64
  Pdf(float64)  float64
  Cdf(float64)  float64
  Random()      float64
}

// Signifies bad parameters for a distribution.
type InvalidParamsError struct{ S string }
func (e InvalidParamsError) Error() string { return e.S }

// Takes n samples from a distribution.
func Sample(dist Distribution, n int) []float64 {
  if n <= 0 {
    return []float64{}
  }
  result := make([]float64, n)
  for i := 0; i < n; i++ {
    value := dist.Random()
    if math.IsNaN(value) {
      return []float64{}
    }
    result[i] = value
  }
  return result
}
