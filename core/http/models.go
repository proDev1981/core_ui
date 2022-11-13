package http

import "github.com/gorilla/websocket"

// struct event
type Event struct {
	Type   string `json:"type"`
	Call   func(*Event)
	Id     string `json:"id"`
	Value  string `json:"value"`
	Client *websocket.Conn
}

// struct reciver sms to  client
type reciverSms struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Event Event  `json:"event"`
}
