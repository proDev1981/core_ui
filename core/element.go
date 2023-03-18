package core

import "fmt"
import "log"
import "strings"

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
	if strings.Contains(args.Value, "{{") && strings.Contains(args.Value, "}}") {
		args.reactive = true
	}
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
func (e *Ele) AddEventListener(types string, call func(*Event)) {
	e.args.Events[types] = call
}

// render element
func (e *Ele) render() string {
	// return DebugAndRes(
	// 	fmt.Sprint(
	// 		"\n",
	// 		"render duration in :",
	// 		"\n-tag=>", e.tag,
	// 		"\n-class=>", e.Args().Class,
	// 		"\n-id=>", e.Args().id,
	// 		"\n",
	// 	),
	// 	func() string { return e.motorRender.RenderElement(e) },
	// )
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

// getter class
func (e *Ele) GetClass() string {
	return e.args.Class
}

// setter class
func (e *Ele) SetClass(c string) {
	e.args.Class = c
}

// add class
func (e *Ele) AddClass(c string) {
	e.args.Class += " " + c
}

// setter id
func (e *Ele) setId(id string) {
	e.args.id = id
}

// getter id
func (e *Ele) GetId() string {
	return e.args.id
}

// setter args
func (e *Ele) SetArgs(a Args) {
	e.args = a
}

// getter args
func (e *Ele) Args() Args {
	return e.args
}

// getter reactive
func (e *Ele) IsReactive() bool {
	return e.args.reactive
}

// setter reactive
func (e *Ele) SetReactive(v bool) {
	e.args.reactive = v
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
func (e *Ele) RootSelector(query string) Element {
	return e.MotorRender().RootSelector(query)
}

// search element in map by key
func (e *Ele) Key(key string) Element {
	return e.MotorRender().Key(key)
}

// search element by query
func (e *Ele) Selector(query string) Element {
	return e.MotorRender().Selector(e, query)
}

// search element by query
func (e *Ele) SelectorAll(query string) []Element {
	return e.MotorRender().SelectorAll(e, query)
}

// create new object in js client
func (e *Ele) NewObject(name string, value any) string {
	return e.MotorRender().NewObject(e, name, value)
}

// get attribute of element in dom html
func (e *Ele) GetAttribute(value string) string {
	return e.MotorRender().GetAttribute(e, value)
}

// set attribute of element in dom html
func (e *Ele) SetAttribute(name string, value string) {
	e.MotorRender().SetAttribute(e, name, value)
}

// log in js console
func (e *Ele) Log(value ...string) {
	e.MotorRender().Log(e, value...)
}

// alert in js client
func (e *Ele) Alert(value ...string) {
	e.MotorRender().Alert(e, value...)
}

// return data childs input in element
func (e *Ele) GetData() map[string]string {
	return e.MotorRender().GetData(e)
}

// reset value elemnet
func (e *Ele) Reset() Element {
	e.MotorRender().Reset(e)
	return e
}

// focus in element
func (e *Ele) Focus() Element {
	e.MotorRender().Focus(e)
	return e
}

// get innerHtml element
func (e *Ele) GetInner() string {
	return e.MotorRender().GetInner(e)
}

// set innerHTML element
func (e *Ele) SetInner(value string) {
	e.MotorRender().SetInner(e, value)
}

// get value element
func (e *Ele) GetValue() string {
	return e.MotorRender().GetValue(e)
}

// set value element
func (e *Ele) SetValue(value string) {
	e.MotorRender().SetValue(e, value)
}

// themes
func (e *Ele) SetBackgroundColor(color string) {
	e.MotorRender().SetBackgroundColor(color)
}
func (e *Ele) SetBackgroundTitle(color string) {
	e.MotorRender().SetBackgroundTitle(color)
}

// apparence
func (e *Ele) Colapsed() Element {
	if e.MotorRender() != nil {
		e.MotorRender().AddClass(e, "colapsed")
	} else {
		e.AddClass("colapsed")
	}
	return e
}
func (e *Ele) Uncolapsed() Element {
	e.MotorRender().RemoveClass(e, "colapsed")
	return e
}
