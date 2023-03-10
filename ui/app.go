package ui

import "app/core"
import "app/ui/controler"

func App() core.Element {

	return core.Box(
		core.Args{
			Name: "root",
		},
		controler.New(),
	)
}
