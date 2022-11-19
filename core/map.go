package core

type mapper struct {
	Element
}

func Map(args Args, children ...Element) *mapper {
	m := &mapper{NewElement("map", "div", args)}
	m.childs(children...)
	return m
}
