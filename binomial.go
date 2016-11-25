package distributions

import (
  "math"
  "math/rand"
)

const (
  binv_cutoff = 110
  small_mean = 14
  far_from_mean = 20
)

//The Binomial Distribution is a discrete probability distribution
// with parameters n > 0, 1 > p > 0.
//
// See: https://en.wikipedia.org/wiki/Binomial_distribution
type Binomial struct {
  Trials  float64
  Prob    float64
}

func (dist *Binomial) validate() error {
  dist.Trials = math.Floor(dist.Trials)
  if dist.Trials < 0 {
    return InvalidParamsError{ "Trials must be greater than zero." }
  }
  if dist.Prob < 0 || dist.Prob > 1 {
    return InvalidParamsError{ "Prob must be between zero and one." }
  }
  return nil
}

func (dist Binomial) Mean() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Trials * dist.Prob
  return result, nil
}

func (dist Binomial) Variance() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := dist.Trials * dist.Prob * (1 - dist.Prob)
  return result, nil
}

func (dist Binomial) Skewness() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := (1 - (2 * dist.Prob)) / math.Sqrt(dist.Trials * dist.Prob * (1 - dist.Prob))
  return result, nil
}

func (dist Binomial) Kurtosis() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := 3 - (6 / dist.Trials) + (1 / (dist.Trials * dist.Prob * (1 - dist.Prob)))
  return result, nil
}

func (dist Binomial) StdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt(dist.Trials * dist.Prob * (1 - dist.Prob))
  return result, nil
}

func (dist Binomial) RelStdDev() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  result := math.Sqrt((1 - dist.Prob) / (dist.Trials * dist.Prob))
  return result, nil
}

func (dist Binomial) Pdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if x < 0.0 || x > dist.Trials {
    return 0.0, nil
  }
  x = math.Floor(x)
  if dist.Prob == 0 {
    if x == 0 {
      return 1.0, nil
    }
    return 0.0, nil
  }
  if dist.Prob == 1 {
    if x == dist.Trials {
      return 1.0, nil
    }
    return 0.0, nil
  }
  cnk := BinomialCoefficient(dist.Trials, x)
  pows := math.Pow(dist.Prob, x) * math.Pow(1 - dist.Prob, dist.Trials - x)
  if math.IsInf(cnk, 0) {
    return 0.0, nil
  }
  result := cnk * pows
  return result, nil
}

func (dist Binomial) Cdf(x float64) (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if (x < 0.0) {
    return 0.0, nil
  }
  if (x > dist.Trials) {
    return 1.0, nil
  }
  result := 0.0
  end := math.Floor(x) + 1
  for i := 0.0; i < end; i++ {
    current := BinomialCoefficient(dist.Trials, i)
    pows := math.Pow(dist.Prob, i) * math.Pow(1 - dist.Prob, dist.Trials - i)
    result += current * pows
  }
  return result, nil
}

// Ref: https://github.com/ampl/gsl/blob/48fbd40c7c9c24913a68251d23bdbd0637bbda20/randist/binomial_tpe.c
func (dist Binomial) random() (float64, error) {
  if err := dist.validate(); err != nil {
    return math.NaN(), err
  }
  if dist.Trials == 0 {
    return 0.0, nil
  }
  flipped := false
  ix := 0.0
  prob := dist.Prob
  if prob > 0.5 {
    flipped = true
    prob = 1 - prob
  }
  q := 1 - prob
  s := prob / q
  np := dist.Trials * prob
  if np < small_mean {
    f0 := math.Pow(q, dist.Trials)
    for {
      f := f0
      u := rand.Float64()
      for ix = 0; ix <= binv_cutoff; ix++ {
        if u < f {
          goto Finish
        }
        u -= f
        f *= s * (dist.Trials - ix) / (ix + 1)
      }
    }
  } else {
    ffm := np + prob
    fm := math.Floor(ffm)
    xm := fm + 0.5
    npq := np * q
    p1 := math.Floor((2.195 * math.Sqrt(npq)) - (4.6 * q)) + 0.5
    xl := xm - p1
    xr := xm + p1
    c := 0.134 + (20.5 / (15.3 + fm))
    p2 := p1 * (1.0 + c + c)
    al := (ffm - xl) / (ffm - (xl * prob))
    lambda_l := al * (1.5 * al)
    ar := (xr - ffm) / (xr * q)
    lambda_r := ar * (1.5 * ar)
    p3 := p2 + (c / lambda_l)
    p4 := p3 + (c / lambda_r)
    var varr, accept, u, v float64

    TryAgain:
      u = rand.Float64() * p4
      v = rand.Float64()
      if u <= p1 {
        ix = math.Floor(xm - (p1 * v) + u)
        goto Finish
      } else if u <= p2 {
        x := xl + ((u - p1) / c)
        v = (v * c) + 1.0 - (math.Abs(x - xm) / p1)
        if v > 1.0 || v <= 0.0 {
          goto TryAgain
        }
        ix = x
      } else if u <= p3 {
        ix = math.Floor(xl + (math.Log(v) / lambda_l))
        if ix < 0 {
          goto TryAgain
        }
        v *= (u - p2) * lambda_r
      } else {
        ix = math.Floor(xr - (math.Log(v) / lambda_r))
        if ix > dist.Trials {
          goto TryAgain
        }
        v *= (u - p3) * lambda_r
      }
      varr = math.Log(v)
      // Skipping Squeeze methods - See Ref
      lg1, _ := math.Lgamma(fm)
      lg2, _ := math.Lgamma(dist.Trials - fm)
      lg3, _ := math.Lgamma(ix)
      lg4, _ := math.Lgamma(dist.Trials - ix)
      accept = lg1 + lg2 - lg3 - lg4 + ((ix - fm) * math.Log(q / prob))
      if varr <= accept {
        goto Finish
      } else {
        goto TryAgain
      }
  }

  Finish:
    value := math.Floor(ix)
    if flipped {
      value = dist.Trials - ix
    }
    return value, nil
}

func (dist Binomial) Sample(n int) ([]float64, error) {
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
