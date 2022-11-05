package main

import c "app/core"
import "app/ui"
import "log"

func main(){
  log.SetFlags(log.Lshortfile)
  log.Println(c.HtmlRender(c.Page(ui.App())))
}


