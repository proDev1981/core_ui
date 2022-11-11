package main

import . "app/core"
import "app/ui"

func main() {
	var global = NewProvider()
	html := HtmlBuild(Page(ui.App(global)))
	html.NewServer().Socket().Listen()
}
