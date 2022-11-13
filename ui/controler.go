package ui

import . "app/core"
import . "app/model"

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
	Target(e).Selector(".title").GetState().Set(Data{Str: "pedro"})
	Target(e).Selector(".map").GetState().Set([]Person{{"julian", 20}, {"pedro", 18}})
}
