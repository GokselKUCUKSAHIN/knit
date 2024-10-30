package model_domain

import "fmt"

type Vector2[T Number] struct {
	X T `json:"x"`
	Y T `json:"y"`
}

func (v Vector2[T]) String() string {
	return fmt.Sprintf("x: %v, y: %v", v.X, v.Y)
}
