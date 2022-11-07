package ui

import . "app/core"
import . "app/model"

func Title() Element {

	action := NewState(DataTitle{Class: "on", Str: "Alberto"}).
		Provider("action", global)

	return Box(Args{Class: "container-title"},
		Label(Args{Class: "title-{{.Class}}", State: action,
			Value: "algo que decir {{.Str}} "}),
	)
}
