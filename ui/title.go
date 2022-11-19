package ui

import . "app/core"
import . "app/model"

func Title() Element {

	action := NewState(Data{Str: "Alberto"})

	return Box(Args{Class: "container-title"},
		Label(Args{Class: "title", State: action,
			Value: "algo que decir {{.Str}} "}),
	)
}
