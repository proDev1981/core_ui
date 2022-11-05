package core

import "fmt"
import "reflect"
import "strings"

type Args struct{
  Class string
  Name string
  id string
  Style Style
  Type string
  Value string
  Src string
  Alt string
  State *State
  Events Listener
}
func (a Args) GetString()string{
  return fmt.Sprintf("%+v\n",a)
  // sacar con reflexion la lista de argumentos con valor menos la id
}
func (s Args) ToHtml()string{
  sType:= reflect.TypeOf(s) 
  sValue:= reflect.ValueOf(s)
  sLen:= sType.NumField()
  var res string

  for i:= 0 ; i<sLen ; i++{
    value:= fmt.Sprint(sValue.Field(i))
    name:= strings.ToLower(sType.Field(i).Name)
    types:= fmt.Sprint(sType.Field(i).Type)
    if s.State != nil {
      value= strings.ReplaceAll(value,"{{."+s.State.name+"}}",fmt.Sprint(s.State.value))
    }
    if types == "string" && name != "id" && name != "value" && len(value) > 0{
      res += fmt.Sprint(" ",name,"='",value,"'")
    }
  }
  return res
}
