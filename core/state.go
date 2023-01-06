package core

import "log"

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

type State struct {
	name     string
	value    any
	children []Element
}

// constructor State
func NewState(name string, value any) *State {
	return &State{value: value, name: name}
}

// append children element
func (s *State) AddElement(e Element) {
	s.children = append(s.children, e)
}

// getter value in state
func (s *State) Get() any {
	return s.value
}

// add in data state
func (s *State) Add(value any) {
	switch s.value.(type) {
	case int:
		s.Set(s.value.(int) + value.(int))
	case string:
		s.Set(s.value.(string) + value.(string))
	case float32, float64:
		s.Set(Float(s.value) + Float(value))
	default:
		//s.Set(Append(s.value, value)) <-- no funciona
	}
}

// setter value in state
func (s *State) Set(value any) {
	//if fmt.Sprint(s.value) != fmt.Sprint(value) {
	log.Println(value)
	s.value = value
	s.uploadElements()
}

// update and render subcriber element in state
func (s *State) uploadElements() {
	for _, item := range s.children {
		item.UpDate()
	}
}

// return last child observable
func (s *State) First() Element {
	return s.children[0]
}
