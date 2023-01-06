package core

type container struct {
	Element
}

func Box(args Args, children ...Element) *container {
	r := &container{NewElement("box", "div", args)}
	r.childs(children...)
	return r
}
