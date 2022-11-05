package core

import "fmt"

type Generic[T any] struct{
  Value T
}
func (g *Generic[T]) Print(){
  fmt.Println(g.Value)
}
func (g *Generic[T]) GetValue()T{
  return g.Value
}

 
  
