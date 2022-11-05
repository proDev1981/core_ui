package core

//import "fmt"
import "log"
// init
func init(){
  log.SetFlags(log.Lshortfile)
}
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
func (s *State) Get()any{
  return s.value
}
func (s *State) Set(value any){
  //if fmt.Sprint(s.value) != fmt.Sprint(value) { 
    s.value= value 
    s.uploadElements()
  //}
}
func (s *State) uploadElements(){
  log.Println("name state=>",s.name)
  log.Println("len child=>",len(s.children))
  for _,item:= range s.children{
    item.UpDate()
  }
}
