package core

import "log"
import "reflect"
import "fmt"

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

// return string value
func (s *State) ToString() string {
	return fmt.Sprint(s.Get())
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

// sub in data state
func (s *State) Sub(value int) {
	switch s.value.(type) {
	case int:
		s.Set(s.value.(int) - value)
	case float32, float64:
		s.Set(Float(s.value) - Float(value))
	default:
		data := reflect.ValueOf(s.value)
		switch {
		case data.Len() <= value:
			log.Println("len is <= value insert")
		default:
			s.Set(reflect.AppendSlice(data.Slice(0, value), data.Slice(value+1, data.Len())).Interface())
		}

	}
}

// return len of data in state
func (s *State) Len() int {
	return reflect.ValueOf(s.value).Len()
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
