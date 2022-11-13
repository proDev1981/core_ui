package core

import "github.com/gorilla/websocket"

// interface motor render
type Motor interface {
	RenderElement(Element) string
	RenderPage(*page) string
	AddEventListener(string, Listener)
	NewServer() *Server
	SetConn(*websocket.Conn)
	Conn() *websocket.Conn
	Selector(string) Element
	Update(Element)
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
	GetState() *State
	SetTag(string)
	Tag() string
	SetArgs(Args)
	Args() Args
	GetSubType() string
	AddEventListener(string, func(*Event))
	SetMotorRender(Motor)
	MotorRender() Motor
	Selector(string) Element
}
