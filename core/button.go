package core


type button struct {
	Element
}

func Button(args Args, childs ...Element) *button {
	return &button{NewElement("button", "button", args, childs...)}
}
