package core

import "log"
import "reflect"

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

// append children element
func (s *State) AddElement(e Element) {
	s.children = append(s.children, e)
}

// getter value in state
func (s *State) Get() any {
	return s.value
}

//getter value with referency
func (s *State) Fill(value any) {
	reflect.ValueOf(value).Elem().Set(reflect.ValueOf(s.value))
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
		s.Set(reflect.Append(reflect.ValueOf(s.value), reflect.ValueOf(value)).Interface())
	}
}

// setter value in state
func (s *State) Set(value any) {
	//if fmt.Sprint(s.value) != fmt.Sprint(value) {
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
