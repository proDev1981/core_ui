package core

type container struct {
	Element
}

func Box(args Args, children ...Element) *container {
	r := &container{NewElement("box", "div", args)}
	r.childs(children...)
	return r
}

func Row(args Args, children ...Element) *container {
	children = append(children, Styles(Sheet{"": &Rule{Display: "flex", FlexDirection: "row"}}.Parse()))
	return Box(args, children...)
}
func Column(args Args, children ...Element) *container {
	children = append(children, Styles(Sheet{"": &Rule{Display: "flex", FlexDirection: "column"}}.Parse()))
	return Box(args, children...)
}
