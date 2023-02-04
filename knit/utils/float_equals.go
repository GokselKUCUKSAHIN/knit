package utils

import (
	"knit/knit/constraints"
	"math"
)

func FloatEquals[T constraints.Number](first, second T) bool {
	return math.Abs(float64(first-second)) < 1e-7
}
