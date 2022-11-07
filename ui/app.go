package ui

import . "app/core"

var global *Provider

func App(gl *Provider) Element {
	global = gl

	return Box(Args{Name: "body"},
		Title(),
		List(),
		Controler(),
	)
}
