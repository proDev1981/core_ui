package ui

import ."app/core"

func Controler()Element{

  return Row(Args{Class:"controler"},
          Button(Args{Class:"btn",
            Events:Listener{ "click":ClickButton },
            Value:"Press" }),
          )
}

func ClickButton(){}
