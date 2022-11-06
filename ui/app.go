package ui

import ."app/core"

func App(action *State)Element{

  return Column(Args{ Name:"body" },
          Title(action),
          List(),
          Controler(),
        )
}

