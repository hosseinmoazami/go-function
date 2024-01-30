package simpleMath

import (
	"errors"
	"math"
)

func Sum(variables ...float64) float64 {
	total := 0.0
	for _, variable := range variables {
		total += variable
	}

	return total
}

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Divided(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("can not divided by Zero")
	}

	return p1 / p2, nil
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}
