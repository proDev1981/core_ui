package core

import "github.com/gorilla/websocket"

// struct event
type Event struct {
	Type   string `json:"type"`
	Call   func(*Event)
	Id     string `json:"id"`
	Value  string `json:"value"`
	Client *websocket.Conn
}

// get element and inyect comunication client
func (e *Event) Target() Element {
	ele := selector(dom, "#"+e.Id)
	ele.MotorRender().SetConn(e.Client)
	return ele
}
