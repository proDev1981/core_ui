package core

import "github.com/gorilla/websocket"

// interface motor render
type Motor interface {
	RenderElement(Element) string
	RenderPage(*page) string
	AddEventListener(string, Listener)
	NewServer() *Server
	GetServer() *Server
	SetConn(*websocket.Conn)
	Conn() *websocket.Conn
	Key(string) Element
	RootSelector(string) Element
	Selector(Element, string) Element
	SelectorAll(Element, string) []Element
	Update(Element)
	// js binding
	NewObject(Element, string, any) *PROMISE
	GetAttribute(Element, string) *PROMISE
	SetAttribute(Element, string, string) *PROMISE
	Log(Element, ...string)
	Alert(Element, ...string)
	GetData(Element) map[string]string
	Reset(Element) Element
	Focus(Element) Element
	GetInner(Element) *PROMISE
	SetInner(Element, string)
	GetValue(Element) *PROMISE
	SetValue(Element, string)
	// themes
	SetBackgroundColor(string)
	SetBackgroundTitle(string)
}

// interface Element
type Element interface {
	render() string
	UpDate()
	childs(...Element)
	Children() []Element
	Parent() Element
	setParent(Element)
	State(s *State) Element
	SetTag(string)
	Tag() string
	SetArgs(Args)
	Args() Args
	GetSubType() string
	AddEventListener(string, func(*Event))
	SetMotorRender(Motor)
	MotorRender() Motor
	Key(string) Element
	RootSelector(string) Element
	Selector(string) Element
	SelectorAll(string) []Element
	// js binding
	NewObject(string, any) *PROMISE
	GetAttribute(string) *PROMISE
	SetAttribute(string, string) *PROMISE
	Log(...string)
	Alert(...string)
	GetData() map[string]string
	Reset() Element
	Focus() Element
	GetInner() *PROMISE
	SetInner(string)
	GetValue() *PROMISE
	SetValue(string)
	// themes
	SetBackgroundColor(string)
	SetBackgroundTitle(string)
}
