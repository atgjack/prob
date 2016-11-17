package distributions

import (
	"math"
)

// floatsEqual determines if two values are within epsilon of each other.
func floatsEqual(f1, f2, epsilon float64) bool {
	return math.Abs(f1-f2) < epsilon
}

// floatsDeciEqual determines if two values are within 10^-1 of each other.
func floatsDeciEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.1
}

// floatsCentiEqual determines if two values are within 10^-2 of each other.
func floatsCentiEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.01
}

// floatsMilliEqual determines if two values are within 10^-3 of each other.
func floatsMilliEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.001
}

// floatsNanoEqual determines if two values are within 10^-9 of each other.
func floatsNanoEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.000000001
}

// floatsPicoEqual determines if two values are within 10^-12 of each other.
func floatsPicoEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 0.000000000001
}

func averageFloats(values []float64) float64 {
  var total float64
  for _, value := range values {
    total += value
  }
  return total / float64(len(values))
}

func varianceFloats(values []float64, mean float64) float64 {
  var total float64
  for _, value := range values {
    total += math.Pow(value - mean, 2)
  }
  return total / (float64(len(values)) - 1)
}
