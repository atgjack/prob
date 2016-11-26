package distributions

// Distirbution is an interface for impementing continuous probability distributions.
//
// See: https://en.wikipedia.org/wiki/Probability_distribution
type Distribution interface {
  Mean()        (float64, error)
  Variance()    (float64, error)
  Kurtosis()    (float64, error)
  Skewness()    (float64, error)
  StdDev()      (float64, error)
  RelStdDev()   (float64, error)
  Pdf(float64)  (float64, error)
  Cdf(float64)  (float64, error)
  Random()      (float64, error)
}

// Signifies bad parameters for a distribution.
type InvalidParamsError struct{ S string }
func (e InvalidParamsError) Error() string { return e.S }

// Takes n samples from a distribution.
func Sample(dist Distribution, n int) ([]float64, error) {
  if n <= 0 {
    return []float64{}, nil
  }
  result := make([]float64, n)
  for i := 0; i < n; i++ {
    value, err := dist.Random()
    if err != nil {
      return []float64{}, err
    }
    result[i] = value
  }
  return result, nil
}
