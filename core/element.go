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
	parent      Element
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

// getter state
func (s *Ele) GetState() *State {
	return s.args.State
}

// setter eventen listener
func (e *Ele) AddEventListener(types string, call func(*Event)) {
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

// getter parent field
func (e *Ele) Parent() Element {
	return e.parent
}

// setter parent field
func (e *Ele) setParent(parent Element) {
	e.parent = parent
}

// update element render
func (e *Ele) UpDate() {
	e.MotorRender().Update(e)
}

// setter motor render
func (e *Ele) SetMotorRender(m Motor) {
	e.motorRender = m
}

// getter motor render
func (e *Ele) MotorRender() Motor {
	return e.motorRender
}

// search element by query
func (e *Ele) Selector(query string) Element {
	return e.MotorRender().Selector(query)
}

// create new object in js client
func (e *Ele) NewObject(name string, value any) *PROMISE {
	return e.MotorRender().NewObject(e, name, value)
}

// get attribute of element in dom html
func (e *Ele) GetAttribute(value string) *PROMISE {
	return e.MotorRender().GetAttribute(e, value)
}

// set attribute of element in dom html
func (e *Ele) SetAttribute(name string, value string) *PROMISE {
	return e.MotorRender().SetAttribute(e, name, value)
}

// log in js console
func (e *Ele) Log(value string) {
	e.MotorRender().Log(e, value)
}

// alert in js client
func (e *Ele) Alert(value string) {
	e.MotorRender().Alert(e, value)
}
