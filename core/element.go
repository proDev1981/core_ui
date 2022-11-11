package core

import "fmt"
import "log"

// init
func init() {
	log.SetFlags(log.Lshortfile)
}

// struct element
type Ele struct {
	tag         string
	args        Args
	subtype     string
	motorRender Motor
	children    []Element
}

// contructor element
func NewElement(sub string, tag string, args Args, childs ...Element) *Ele {
	e := &Ele{subtype: sub, tag: tag, args: args}
	e.args.id = fmt.Sprintf("%p", e) // grabo la direccion de memoria como id
	if len(childs) > 0 {
		e.children = append(e.children, childs...)
	}
	return e
}

// getter subtype
func (e *Ele) GetSubType() string {
	return e.subtype
}

// setter children
func (e *Ele) childs(eles ...Element) {
	e.children = append(e.children, eles...)
}

// setter state
func (e *Ele) State(s *State) Element {
	s.children = append(s.children, e)
	return e
}

// setter eventen listener
func (e *Ele) AddEventListener(types string, call func()) {
	e.args.Events[types] = call
}

// render element
func (e *Ele) render() string {
	return e.motorRender.RenderElement(e)
}

// setter tag
func (e *Ele) SetTag(t string) {
	e.tag = t
}

// getter tag
func (e *Ele) Tag() string {
	return e.tag
}

// setter args
func (e *Ele) SetArgs(a Args) {
	e.args = a
}

// getter args
func (e *Ele) Args() Args {
	return e.args
}

// getter children
func (e *Ele) Children() []Element {
	return e.children
}

// update element render
func (e *Ele) UpDate() {
	log.Println("buscar elemento con id =>", e.args.id)
	log.Println("// nuevo elemento renderizado //")
	log.Println(e.render())
}

// setter motor render
func (e *Ele) SetMotorRender(m Motor) {
	e.motorRender = m
}

// getter motor render
func (e *Ele) MotorRender() Motor {
	return e.motorRender
}
