package distributions

import (
	"math"
)

type inOut struct {
  in  float64
  out float64
}

// floatsEqual determines if two values are within epsilon of each other.
func floatsEqual(f1, f2, epsilon float64) bool {
	return math.Abs(f1-f2) < epsilon
}

// floatsIntegerEqual determines if two values are within 10^0 of each other.
func floatsIntegerEqual(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1
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

func checkInf(f1, f2 float64) bool {
  if math.IsInf(f1,0) || math.IsInf(f2,0) {
    return math.IsInf(f1,0) && math.IsInf(f2,0)
  }
  return true
}

func checkNaN(f1, f2 float64) bool {
  if math.IsNaN(f1) || math.IsNaN(f2) {
   return math.IsNaN(f1) && math.IsNaN(f2)
  }
  return true
}

func averageFloats(values []float64) float64 {
  var total float64
  for _, value := range values {
    total += value
  }
  return total / float64(len(values))
}

func varianceFloats(values []float64, mean float64) float64 {
  var total, diff float64
  for _, value := range values {
    diff = value - mean
    total += diff * diff
  }
  return total / (float64(len(values)) - 1)
}
