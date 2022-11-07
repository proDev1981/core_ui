package ui

import . "app/core"
import . "app/model"

func List() Element {

	data := NewState([]Person{{Name: "alberto", Age: 42}, {Name: "paco", Age: 39}})

	return Map(Args{Class: "map", State: data},
		Label(Args{
			Value: "tu nombre es: {{.Name}}",
		}),
		Label(Args{
			Value: "tu edad es:{{.Age}}",
		}),
	)
}
