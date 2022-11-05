package ui

import c "app/core"

func Title()c.Element{

  types:= c.NewState("types","on")

  return c.Row(c.Args{ Class:"container-title" },
            c.Label(c.Args{ Class:"title",State:types,
              Value:"algo que decir" }),
          )
}
