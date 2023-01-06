package core

func Map[T any](data []T, fn func(index int, item T) Element) Element {
	var childs []Element
	for index, item := range data {
		childs = append(childs, fn(index, item))
	}
	return Box(Args{}, childs...)
}
