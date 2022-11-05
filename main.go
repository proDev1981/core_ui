package main

import c "app/core"
import "fmt"

func main(){

  types:= c.NewState("types","on")
  web:= c.NewHtml()// creo motor renderizado modo html

  p:= c.Page(web,
        c.Header(),
        c.Styles("./styles.css"),
        c.Column(c.Args{ 
          Name:"body", 
        },
          c.Row(c.Args{ 
            Name:"title", 
          },
            c.Label(c.Args{ 
              Class:"{{.types}}",
              Value:"algo que decir",
              Type:"h1",
              State:types, 
            },
              c.Img(c.Args{ 
                Src:"http://logo.png",
                Alt:"logo.png", 
              }),
            ),//label
            c.Input(c.Args{ 
              Class:"inp",
              Type:"text",
            }),
          ),//row=>title
          c.Row(c.Args{ 
            Name:"controler",
          },
            c.Button(c.Args{ 
              Class:"btn-{{.types}}",
              Value:"press",
              State:types,
              Events:c.Listener{
                "click":ClickButton,
              }, 
            }),
          ),//row=>controler
        ),//column=>body
      )//page
  fmt.Println("26:main ",p.Render())
  types.Set("off")
}

func ClickButton(){}

