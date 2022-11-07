package main

import . "app/core"
import . "app/model"
import "app/ui"
import . "log"

func main() {

	var global = NewProvider()

	Println(HtmlRender(Page(ui.App(global))))
	global.GetState("action").
		Set(DataTitle{Class: "off", Str: "Paco"})
}
