package core

func If(c bool, a Element, b Element) Element {
	if c {
		return a
	}
	return b
}

type _case struct {
	Element
	condition func() bool
	state     *State
}

func Case(c func() bool, s *State, a Element) *_case {
	return &_case{
		Element:   a,
		state:     s,
		condition: c,
	}
}

func (c *_case) render() (res string) {

	println("render case ", c.condition())
	if c.condition() {
		res = "<case>" + c.Element.render() + "</case>"
	} else {
		res = "<case></case>"
	}
	c.state.AddElement(c)
	return
}
