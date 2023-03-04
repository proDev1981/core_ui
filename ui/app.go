package ui

import "app/core"
import "app/components/post"
import "app/ui/controler"
import "app/store"

func App() core.Element {

	return core.Box(
		core.Args{
			Name: "root",
		},
		core.List(
			core.Args{
				State: store.States["data"],
			},
			post.New(
				post.Args{
					Title:   "{{.Email}}",
					Image:   "man.png",
					Content: "{{.Content}}",
				}),
		),
		controler.New(),
	)
}
