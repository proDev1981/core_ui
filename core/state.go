package core

type State struct{
  name string
  value any
  children []Element
}

func NewState(name string, value any)*State{
  return &State{ name:name,value:value }
}
func (s *State) Add(e Element){
  s.children = append(s.children,e)
}
func (s *State) Set(value any){
  if s.value != value { 
    s.value= value 
    s.uploadElements()
  }
}
func (s *State) uploadElements(){
  for _,item:= range s.children{
    item.UpDate()
  }
}
