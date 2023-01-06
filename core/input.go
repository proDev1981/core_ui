package core

type input struct {
	Element
}

func Input(args Args) *input {
	return &input{NewElement("input", "input", args)}
}
