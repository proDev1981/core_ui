package core

//import "fmt"
import "log"

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

type State struct {
	value    any
	children []Element
}

// constructor State
func NewState(value any) *State {
	return &State{value: value}
}

// append state in props maganer
func (s *State) Provider(name string, p *Provider) *State {
	p.dataBase[name] = s
	return s
}

// append children element
func (s *State) Add(e Element) {
	s.children = append(s.children, e)
}

// getter value in state
func (s *State) Get() any {
	return s.value
}

// setter value in state
func (s *State) Set(value any) {
	//if fmt.Sprint(s.value) != fmt.Sprint(value) {
	s.value = value
	s.uploadElements()
	//}
}

// update and render subcriber element in state
func (s *State) uploadElements() {
	log.Println("len child=>", len(s.children))
	for _, item := range s.children {
		item.UpDate()
	}
}
