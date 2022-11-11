package ui

import . "app/core"

func Controler() Element {

	return Box(Args{Class: "controler"},
		Button(Args{Class: "btn",
			Events: Listener{"click": ClickButton},
			Value:  "Press"}),
		Button(Args{Class: "btn",
			Events: Listener{"click": ClickButton},
			Value:  "Exit"}),
	)
}

func ClickButton() {}
