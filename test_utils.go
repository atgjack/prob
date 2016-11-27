package distributions

import (
  "testing"
	"math"
  "fmt"
  "math/rand"
  "time"
)

const (
  numSamples = 1000000
  defaultEpsilon = 0.01
)

type inOut struct {
  in  float64
  out float64
}

// Epsilon is option. Will default to defaultEpsilon value.
type sampleValues struct {
  mean      float64
  variance  float64
  epsilon   float64
}

// Distribution test value struct.
type distributionTest struct {
  dist        Distribution
  mean        float64
  variance    float64
  stdDev      float64
  relStdDev   float64
  skewness    float64
  kurtosis    float64
  pdf         []inOut
  cdf         []inOut
}

// Run tests on distribtion examples.
func testValues(examples []distributionTest) error {
  rand.Seed(time.Now().UTC().UnixNano())
  for _, example := range examples {
    // Test mean.
    mean, err := example.dist.Mean()
    if err != nil || !floatsPicoEqual(mean, example.mean) {
      if !checkInf(mean, example.mean) && !checkNaN(mean, example.mean) {
        return fmt.Errorf("\nMean:\n  Expected: %f\n  Got: %f\n", example.mean, mean)
      }
    }
    // Test variance.
    variance, err := example.dist.Variance()
    if err != nil || !floatsPicoEqual(variance, example.variance) {
      if !checkInf(variance, example.variance) && !checkNaN(variance, example.variance) {
        return fmt.Errorf("\nVariance:\n  Expected: %f\n  Got: %f\n", example.variance, variance)
      }
    }
    // Test standard deviation.
    stdDev, err := example.dist.StdDev()
    if err != nil || !floatsPicoEqual(stdDev, example.stdDev) {
      if !checkInf(stdDev, example.stdDev) && !checkNaN(stdDev, example.stdDev) {
        return fmt.Errorf("\nStdDev:\n  Expected: %f\n  Got: %f\n", example.stdDev, stdDev)
      }
    }
    // Test relative standard deviation.
    relStdDev, err := example.dist.RelStdDev()
    if err != nil || !floatsPicoEqual(relStdDev, example.relStdDev) {
      if !checkInf(relStdDev, example.relStdDev) && !checkNaN(relStdDev, example.relStdDev) {
        return fmt.Errorf("\nRelStdDev:\n  Expected: %f\n  Got: %f\n", example.relStdDev, relStdDev)
      }
    }
    // Test skewness.
    skewness, err := example.dist.Skewness()
    if err != nil || !floatsPicoEqual(skewness, example.skewness) {
      if !checkInf(skewness, example.skewness) && !checkNaN(skewness, example.skewness) {
        return fmt.Errorf("\nSkewness:\n  Expected: %f\n  Got: %f\n", example.skewness, skewness)
      }
    }
    // Test excess kurtosis.
    kurtosis, err := example.dist.Kurtosis()
    if err != nil || !floatsPicoEqual(kurtosis, example.kurtosis) {
      if !checkInf(kurtosis, example.kurtosis) && !checkNaN(kurtosis, example.kurtosis) {
        return fmt.Errorf("\nKurtosis:\n  Expected: %f\n  Got: %f\n", example.kurtosis, kurtosis)
      }
    }
    // Test pdf values.
    for _, pdf := range example.pdf {
      out, err := example.dist.Pdf(pdf.in)
      if err != nil || !floatsPicoEqual(out, pdf.out) {
        return fmt.Errorf("\nPdf of %f:\n  Expected: %f\n  Got: %f\n", pdf.in, pdf.out, out)
      }
    }
    // Test cdf values.
    for _, cdf := range example.cdf {
      out, err := example.dist.Cdf(cdf.in)
      if err != nil || !floatsPicoEqual(out, cdf.out) {
        return fmt.Errorf("\nCdf of %f:\n  Expected: %f\n  Got: %f\n", cdf.in, cdf.out, out)
      }
    }
  }
  return nil
}

func testSamples(dist Distribution) error {
  // Generate samples.
  samples, err := Sample(dist, numSamples)
  if err != nil {
    return fmt.Errorf("\nCould not generate samples.")
  }
  // Test sample average against expected value if it exists.
  sampleMean := averageFloats(samples)
  actualMean, err := dist.Mean()
  if err != nil {
    return err
  }
  if !math.IsInf(actualMean,0) && !math.IsNaN(actualMean) {
    if !floatsEqual(actualMean, sampleMean, defaultEpsilon) {
      return fmt.Errorf("\nSample average:\n  Expected: %f\n  Got: %f\n", actualMean, sampleMean)
    }
  }
  // Test sample variance against expected variance if it exists.
  sampleVar := varianceFloats(samples, sampleMean)
  actualVar, err := dist.Mean()
  if err != nil {
    return err
  }
  if !math.IsInf(actualVar,0) && !math.IsNaN(actualVar) {
    if !floatsEqual(actualVar, sampleMean, defaultEpsilon) {
      return fmt.Errorf("\nSample variance:\n  Expected: %f\n  Got: %f\n", actualVar, sampleVar)
    }
  }
  return nil
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

func runBenchmark(b *testing.B, dist Distribution) {
  for n := 0; n <= b.N; n++ {
    dist.Random()
  }
}
