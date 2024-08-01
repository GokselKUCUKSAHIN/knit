package model_domain

type Pair[T, U any] struct {
	first  T
	second U
}

func (pair Pair[T, U]) First() T {
	return pair.first
}

func (pair Pair[T, U]) Second() U {
	return pair.second
}

func NewPair[T, U any](first T, second U) *Pair[T, U] {
	return &Pair[T, U]{
		first:  first,
		second: second,
	}
}
