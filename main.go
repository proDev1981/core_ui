package main

import . "app/core"
import "app/ui"

func main() {
	var global = NewProvider()

	pwa := HtmlBuild(
		Page(
			Header(
				Link(Args{Href: "./styles.css"}),
			),
			ui.App(global),
		))
	pwa.NewServer().AndSocket().Listen()
}
