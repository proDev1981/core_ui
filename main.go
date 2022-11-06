package main

import ."app/core"
import "app/ui"
import ."log"

func main(){

  type dataTitle struct{
    Class string
    Str   string
  }

  action:= NewState(dataTitle{"on","Alberto"})

  context:= HtmlRender(Page(ui.App(action)))
  Print(context,"\n")

  action.Set(dataTitle{"off","Paco"})
}


