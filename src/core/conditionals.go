package core

func If(c bool, a Element, b Element) Element {
	if c {
		return a
	}
	return b
}
func Then(c bool, a Element) Element {
	if c {
		return a
	}
	return Box(Args{})
}
