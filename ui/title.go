package ui

import ."app/core"

func Title(action *State)Element{


  return Row(Args{ Class:"container-title" },
            Label(Args{ Class:"title-{{.Class}}",State:action,
              Value:"algo que decir {{.Str}} " }),
          )
}
