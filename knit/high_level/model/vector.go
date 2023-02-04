package model

import (
	"fmt"
	"knit/knit/constraints"
)

type Vector struct {
	X int16
	Y int16
}

func (base *Vector) String() string {
	return fmt.Sprintf("x: %d, y: %d", base.X, base.Y)
}

func CreateVector[T constraints.Number](x T, y T) *Vector {
	return &Vector{X: int16(x), Y: int16(y)}
}

func (base *Vector) Slope(end *Vector) float64 {
	return float64(end.Y-base.Y) / float64(end.X-base.X)
}
