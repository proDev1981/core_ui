package main

import . "app/core"
import "app/ui"

func main() {

	pwa := HtmlBuild(
		Page("es",
			Header(

				Meta(Args{Name: "theme-color", Content: "white"}),
			),
			Script(Args{Src: "./index.js"}),
			ui.App(),
		))
	pwa.NewServer().AndSocket().Listen()

}
