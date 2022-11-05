package core

import "fmt"

type Element interface{
  render()string
  UpDate()
  childs(...Element)
  State(s *State)Element
  SetTag(string)
  Tag()string
  Children()[]Element
  SetArgs(Args)
  Args()Args
  AddEventListener( string, func())
  SetMotorRender(Motor)
}

type Ele struct{
  motorRender Motor
  tag string
  args Args
  children []Element
}
// contructor element
func NewElement(tag string, args Args,childs ...Element)*Ele{
  e := &Ele{ tag:tag,args:args}
  e.args.id = fmt.Sprintf("%p",e)// grabo la direccion de memoria como id
  if len(childs) > 0 { e.children = append(e.children,childs...) }
  return e
}
// setter children
func (e *Ele) childs(eles ...Element){
  e.children = append(e.children,eles...)
}
// setter state
func (e *Ele) State(s *State)Element{
  s.children = append(s.children,e)
  return e
}
// setter eventen listener
func (e *Ele) AddEventListener(types string, call func()){
  e.args.Events[types] = call
}
// render element
func (e *Ele) render()string{
  return e.motorRender.RenderElement(e)
}
// setter tag
func (e *Ele) SetTag(t string){
  e.tag = t
}
// getter tag
func (e *Ele) Tag()string{
  return e.tag
}
// setter args
func (e *Ele) SetArgs(a Args){
  e.args = a
}
// getter args
func (e *Ele) Args()Args{
  return e.args
}
// getter children
func (e *Ele) Children()[]Element{
  return e.children
}
// update element render
func (e *Ele) UpDate(){
  fmt.Println("70:element buscar elemento con id =>",e.args.id)
  fmt.Println("71:element // nuevo elemento renderizado //")
  fmt.Println("72:element ",e.render(),"\n")
}
// setter motor render
func (e *Ele) SetMotorRender(m Motor){
  e.motorRender= m
}
