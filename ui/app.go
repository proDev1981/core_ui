package ui

import . "app/core"

func App() Element {

	return Box(Args{Name: "root"},
		Title("Bien venido a CoreUI"),
	)
}
