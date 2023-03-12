package core

type looper struct {
	Element
}

func For(args Args, children ...Element) *looper {
	m := &looper{NewElement("list", "div", args)}
	m.childs(children...)
	return m
}
