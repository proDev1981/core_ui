package ui

import . "app/core"
import "app/models"

func Prueba() Element {

	data := NewState("data", models.Persons)

	return List(Args{State: data},
		Box(Args{Class: "item-list"},
			Label(Args{Value: "{{.Age}}"}),
			Label(Args{Value: "{{.Name}}"}),
		),
	)
}
