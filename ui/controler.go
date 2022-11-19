package ui

import . "app/core"
import . "app/model"
import "fmt"

func Controler() Element {

	return Box(Args{Class: "controler"},
		Button(Args{Class: "btn",
			Events: Listener{"click": handlePress},
			Value:  "Press"}),
		Button(Args{Class: "btn",
			Events: Listener{"click": handleExit},
			Value:  "Exit"}),
		Input(Args{Class: "name"}),
		Input(Args{Class: "age"}),
	)
}

func handlePress(e *Event) {
	this := e.Target()
	person := this.RootSelector(".map").GetState()
	person.Set([]Person{{Name: "julian", Age: 20}, {Name: "pedro", Age: 18}})
	////
	data := this.Parent().GetData()
	this.Alert(fmt.Sprintf("tu nombre es %s y tienes %s años",data["name"],data["age"]))

}
func handleExit(e *Event) {

}
