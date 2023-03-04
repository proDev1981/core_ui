package title

import . "app/core"

func New(str string) Element {

	title := NewState("title", str)

	return Box(Args{
		Class: "container-title",
	},
		Label(Args{
			Class: "title",
			State: title,
			Value: "{{.state}}"}),
	)
}
