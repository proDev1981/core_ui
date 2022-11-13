package ui

import . "app/core"
import . "app/core/http"
import "log"

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

func ClickButton(e *Event) {
	log.Println(Selector("#" + e.Id).Parent())
}
