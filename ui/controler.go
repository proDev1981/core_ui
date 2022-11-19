package ui

import . "app/core"
import . "app/model"
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
	this := e.Target()
	person := this.Selector(".map").GetState()
	person.Set([]Person{{Name: "julian", Age: 20}, {Name: "pedro", Age: 18}})

	classBtn := <-this.GetAttribute("class").Await()
	if classBtn == "btn" {
		log.Println("press a button")
		this.Alert("apretaste un button")
	}
}
