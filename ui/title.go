package ui

import . "app/core"

func Title(str string) Element {

	title := NewState("title", str)

	return Box(Args{Class: "container-title"},
		Label(Args{Class: "title", State: title,
			Value: "{{.state}}"}),
	)
}
