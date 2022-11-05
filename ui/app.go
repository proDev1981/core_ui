package ui

import c "app/core"

func App()c.Element{

  return c.Column(c.Args{ Name:"body" },Title(),List(),Controler())
}

