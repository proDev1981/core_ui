package core

import "fmt"

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
  Direction string
}
func (a Args) GetString()string{
  return fmt.Sprintf("%+v\n",a)
  // sacar con reflexion la lista de argumentos con valor menos la id
}
