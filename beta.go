package distributions

import (
  "math"
  "math/rand"
)

//TheBeta  Distribution is a continuous probability distribution
// with parameters α > 0, β >= 0.
//
// See: https://en.wikipedia.org/wiki/Beta_distribution
type Beta struct {
  Alpha   float64
  Beta    float64
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

func (dist Beta) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Alpha / (dist.Alpha + dist.Beta)
  return result, nil
}

func (dist Beta) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  comb := dist.Alpha + dist.Beta
  result := (dist.Alpha * dist.Beta) / (comb * comb * (comb + 1))
  return result, nil
}

func (dist Beta) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  numer := 2 * (dist.Beta - dist.Alpha) * math.Sqrt(dist.Alpha + dist.Beta + 1)
  result := numer / ((dist.Alpha + dist.Beta + 2) * math.Sqrt(dist.Alpha * dist.Beta))
  return result, nil
}

func (dist Beta) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  first := (dist.Alpha - dist.Beta) * (dist.Alpha - dist.Beta) * (dist.Alpha + dist.Beta + 1)
  second := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2)
  denom := dist.Alpha * dist.Beta * (dist.Alpha + dist.Beta + 2) * (dist.Alpha + dist.Beta + 3)
  result := 6 * (first - second) / denom
  return result, nil
}

func (dist Beta) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  variance, _ := dist.Variance()
  result := math.Sqrt(variance)
  return result, nil
}

func (dist Beta) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  mean, _ := dist.Mean()
  stdDev, _ := dist.StdDev()
  result := stdDev / mean
  return result, nil
}

func (dist Beta) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  beta := BetaFn(dist.Alpha, dist.Beta)
  result := math.Pow(x, dist.Alpha - 1) * math.Pow(1 - x, dist.Beta - 1) / beta
  return result, nil
}

func (dist Beta) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := RegBetaInc(dist.Alpha, dist.Beta, x)
  return result, nil
}

// Ref: https://compbio.soe.ucsc.edu/gen_sequence/gen_beta.c
// Ref: https://github.com/e-dard/godist/blob/master/beta.go
func (dist Beta) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  min := math.Min(dist.Alpha, dist.Beta)
  max := math.Max(dist.Alpha, dist.Beta)
  // Use Joehnk's algorithm.
  if max < 0.5 {
    u1 := rand.Float64()
    u2 := rand.Float64()
    result := math.Pow(u1, 1 / dist.Alpha) / (math.Pow(u1, 1 / dist.Alpha) + math.Pow(u2, 1 / dist.Beta))
    return result, nil
  // Use Cheng's BC Algorithm
  } else if min <= 1.0 {
    var u1, u2, v, w, y, z float64
    alpha := min + max
    beta := 1.0 / min
    delta := 1.0 + max - min
  	k1 := delta * (0.0138889 + (0.0416667 * min)) / ((max * beta) - 0.777778)
  	k2 := 0.25 + ((0.5 + 0.25) * min / delta)
    setvw := func() {
      v = beta * math.Log(u1 / (1.0 - u1))
      if v <= 709.78 {
        w = dist.Alpha * math.Exp(v)
        if math.IsInf(w,0) {
          w = math.MaxFloat64
        }
      } else {
        w = math.MaxFloat64
      }
    }
    for {
      u1 = rand.Float64()
      u2 = rand.Float64()
      if u1 < 0.5 {
        y = u1 * u2
        z = u1 * y
        if (0.25 * u2) + z - y >= k1 {
          continue
        }
      } else {
        z = u1 * u1 * u2
        if z <= 0.25 {
          setvw()
          break
        }
        if z >= k2 {
          continue
        }
      }
      setvw()
      if alpha * (math.Log(alpha / (min + w)) + v) - 1.3862944 >= math.Log(z) {
        break
      }
    }
    if dist.Alpha == min {
      return min / (min + w), nil
    }
    return w / (min + w), nil
  // Use Cheng's BB Algorithm
  } else {
    alpha := min + max
    beta := math.Sqrt((alpha - 2.0) / ((2.0 * min * max) - alpha))
    gamma := min + (1.0 / beta)
    var r, s, t, v, w, z float64
    for {
      u1 := rand.Float64()
      u2 := rand.Float64()
      v = beta * math.Log(u1 / (1.0 - u1))
      if v <= 709.78 {
        w = dist.Alpha * math.Exp(v)
        if math.IsInf(w,0) {
          w = math.MaxFloat64
        }
      } else {
        w = math.MaxFloat64
      }
      z = u1 * u1 * u2
      r = (gamma * v) - 1.3862944
      s = min + r - w
      if s + 2.609438 >= 5.0 * z {
        break
      }
      t = math.Log(z)
      if r + (alpha * math.Log(alpha / (max + w))) < t {
        break
      }
    }
    if dist.Alpha != min {
      return max / (max + w), nil
    }
    return w / (max + w), nil
  }
}

func (dist Beta) Sample(n int) ([]float64, error) {
  if err := dist.validate(); err != nil {
    return []float64{}, err
  }
  if n <= 0 {
    return []float64{}, nil
  }
  result := make([]float64, n)
  for i := 0; i < n; i++ {
    value, err := dist.random()
    if err != nil {
      return []float64{}, nil
    }
    result[i] = value
  }
  return result, nil
}
