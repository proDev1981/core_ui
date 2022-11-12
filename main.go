package main

import . "app/core"
import "app/ui"

func main() {
	var global = NewProvider()
	pwa := HtmlBuild(Page(ui.App(global)))
	pwa.NewServer().AndSocket().Listen()
}
