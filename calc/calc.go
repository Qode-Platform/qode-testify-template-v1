// Package calc is trivial code under test — replace it with the real package.
package calc

import "errors"

var ErrDivideByZero = errors.New("denominator must be non-zero")

func Divide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, ErrDivideByZero
	}
	return numerator / denominator, nil
}

func RunningTotal(values []float64) []float64 {
	out := make([]float64, 0, len(values))
	total := 0.0
	for _, v := range values {
		total += v
		out = append(out, total)
	}
	return out
}
