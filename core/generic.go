package core

import "fmt"

type Generic[T any] struct {
	Value T
}

func New[T any](value T) *Generic[T] {
	return &Generic[T]{value}
}
func (g *Generic[T]) Print() {
	fmt.Println(g.Value)
}
func (g *Generic[T]) GetValue() T {
	return g.Value
}
