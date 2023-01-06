package core

type mapper struct {
	Element
}

func List(args Args, children ...Element) *mapper {
	m := &mapper{NewElement("list", "div", args)}
	m.childs(children...)
	return m
}
