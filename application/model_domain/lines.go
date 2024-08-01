package model_domain

import "fmt"

type Lines []*Pair[int16, int16]

type LinePair []*Pair[*Vector2[int16], *Vector2[int16]]

type Line[T Number] struct {
	Start T `json:"start"`
	End   T `json:"end"`
}

func (l Line[T]) String() string {
	return fmt.Sprintf("s: %v, e: %v", l.Start, l.End)
}
