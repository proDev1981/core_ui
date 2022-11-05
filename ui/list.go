package ui

import c "app/core"
import "app/model"

func List()c.Element{

  data:= c.NewState("persons",[]model.Person{{"alberto",42},{"paco",39}})

  return c.Map(c.Args{ Class:"map",State:data },
            c.Label(c.Args{ Class:"name",Type:"h1",
                Value:"tu nombre es: {{.Name}}" }),
            c.Label(c.Args{ Class:"age",Type:"h1",
                Value:"tu edad es:{{.Age}}" }),
          )
}
