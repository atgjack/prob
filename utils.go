package distributions

import "math"

const gamma_epsilon = 1e-14

// The  regularized lower incomplete gamma function.
// Code kanged from SAMTools: https://github.com/lh3/samtools/blob/master/bcftools/kfunc.c
func Lowerincgamma(s float64, z float64) float64 {
  var k, x, sum float64
  sum = 1
  x = 1
  for k = 1; k < 100; k++ {
    x *= z / (s + k)
    sum += x
    if (x / sum < gamma_epsilon) {
      break
    }
  }
  lgamma, _ := math.Lgamma(s + 1)
  result := math.Exp((s * math.Log(z)) - z - lgamma + math.Log(sum))
  return result
}

// Choose k elements from a set of n elements.
// See: https://en.wikipedia.org/wiki/Binomial_coefficient
func Choose(n, k float64) float64 {
  if k > n {
    return math.NaN()
  }
  r := 1.0
  for d := 1.0; d <= k; d++ {
    r *= n
    r /= d
    n -= 1
  }
  return r
}
