package distributions

import "math"

const gamma_epsilon = 1e-14
const beta_epsilon = 2.2204460492503131e-16
const beta_iterations = 1e9

// The  regularized lower incomplete gamma function.
// Code kanged from SAMTools: https://github.com/lh3/samtools/blob/master/bcftools/kfunc.c
func GammaIncLower(s float64, z float64) float64 {
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
func BinomialCoefficient(n, k float64) float64 {
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

// A variadic version of the Beta function.
// See: https://en.wikipedia.org/wiki/Beta_function
func BetaFn(a ...float64) float64 {
  product := 1.0
  sum := 0.0
  for _, ai := range a {
    product *= math.Gamma(ai)
    sum += ai
  }
  return product / math.Gamma(sum)
}

// The incomplete beta function.
// See: https://en.wikipedia.org/wiki/Beta_function#Incomplete_beta_function
func BetaInc(a, b, x float64) float64 {
  return RegBetaInc(a, b, x) * BetaFn(a, b)
}

// The regularized incomplete beta function.
// See: https://en.wikipedia.org/wiki/Beta_function#Incomplete_beta_function
func RegBetaInc(a, b, x float64) float64 {
  if x == 0.0 {
    return 0.0
  }
  if x == 1.0 {
    return 1.0
  }
  lab, _ := math.Lgamma(a + b)
  la, _ := math.Lgamma(a)
  lb, _ := math.Lgamma(b)
  lbeta := lab - la - lb + (a * math.Log(x)) + (b * math.Log(1-x))
  if x < (a + 1) / (a + b + 2) {
    return math.Exp(lbeta) * contFracBeta(a, b, x) / a
  }
  return 1 - math.Exp(lbeta) * contFracBeta(b, a, 1-x) / b
}

// Ref: https://malishoaib.wordpress.com/2014/04/15/the-beautiful-beta-functions-in-raw-python/
func contFracBeta(a, b, x float64) float64 {
  am, bm, az := 1.0, 1.0, 1.0
  qab := a + b
  qap := a + 1.0
  qam := a - 1.0
  bz := 1.0 - (qab * x / qap)
  for i := 0.0; i <= beta_iterations; i += 1.0 {
    em := i + 1.0
    tem := em + em
    d := em * (b - em) * x / ((qam + tem) * (a + tem))
    ap := az + (d * am)
    bp := bz + (d * bm)
    d = -(a + em) * (qab + em) * x / ((qap + tem) * (a + tem))
    app := ap + (d * az)
    bpp := bp + (d * bz)
    aold := az
    am = ap / bpp
    bm = bp / bpp
    az = app / bpp
    bz = 1.0
    if math.Abs(az - aold) < beta_epsilon * math.Abs(az) {
      return az
    }
  }
  return math.NaN()
}
