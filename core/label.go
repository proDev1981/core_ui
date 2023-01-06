package core

type label struct {
	Element
}

func Label(args Args, childs ...Element) *label {
	var name string
	if args.Type != "" {
		name = args.Type
	} else {
		name = "h1"
	}
	l := &label{NewElement("label", name, args, childs...)}
	return l
}
