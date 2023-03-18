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
	NewObject(Element, string, any) string
	GetAttribute(Element, string) string
	SetAttribute(Element, string, string)
	Log(Element, ...string)
	Alert(Element, ...string)
	GetData(Element) map[string]string
	Reset(Element) Element
	Focus(Element) Element
	GetInner(Element) string
	SetInner(Element, string)
	GetValue(Element) string
	SetValue(Element, string)
	// themes
	SetBackgroundColor(string)
	SetBackgroundTitle(string)
	// apparence
	AddClass(Element, string)
	RemoveClass(Element, string)
}

// interface Element
type Element interface {
	render() string
	UpDate()
	childs(...Element)
	Children() []Element
	Parent() Element
	setParent(Element)
	State(*State) Element
	SetTag(string)
	Tag() string
	GetClass() string
	SetClass(string)
	AddClass(string)
	setId(string)
	SetArgs(Args)
	Args() Args
	SetReactive(bool)
	IsReactive() bool
	GetSubType() string
	AddEventListener(string, func(*Event))
	SetMotorRender(Motor)
	MotorRender() Motor
	Key(string) Element
	RootSelector(string) Element
	Selector(string) Element
	SelectorAll(string) []Element
	// js binding
	NewObject(string, any) string
	GetAttribute(string) string
	SetAttribute(string, string)
	Log(...string)
	Alert(...string)
	GetData() map[string]string
	Reset() Element
	Focus() Element
	GetInner() string
	SetInner(string)
	GetValue() string
	SetValue(string)
	// themes
	SetBackgroundColor(string)
	SetBackgroundTitle(string)
	// apparence
	Colapsed() Element
	Uncolapsed() Element
}
