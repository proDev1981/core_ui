package ui

import c "app/core"

func Controler()c.Element{

  return c.Row(c.Args{Class:"controler"},
            c.Button(c.Args{Class:"btn",Events:c.Listener{"click":ClickButton},
                Value:"Press" }),
            )
}

func ClickButton(){}
