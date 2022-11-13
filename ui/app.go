package ui

import . "app/core"

var global *Provider

func App(gl *Provider) Element {
	global = gl

	return Box(Args{Name: "root"},
		Title(),
		List(),
		Controler(),
	)
}
