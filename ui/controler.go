package ui

import . "app/core"
import "app/models"

func Controler() Element {
	return Box(Args{},
		Button(Args{Value: "press", Events: Listener{"click": add}}),
	)
}
func add(e *Event) {
	this := e.Target()
	this.GetState("data").Set(append(models.Persons, models.Person{Name: "antonio", Age: 56}))
}
