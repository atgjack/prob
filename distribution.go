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
  Sample(int)   ([]float64, error)
}

// Signifies bad parameters for a distribution.
type InvalidParamsError struct{ S string }
func (e InvalidParamsError) Error() string { return e.S }
