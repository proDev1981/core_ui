package http

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Socket struct {
	clients         map[string]*websocket.Conn
	path            string
	clienteToServer []func(*Socket, string)
	upgrader        websocket.Upgrader
	stateclients    map[string]bool
	sms             string
}

func NewSocket() *Socket {
	return &Socket{
		clients:      make(map[string]*websocket.Conn),
		stateclients: make(map[string]bool),
	}
}

func (so *Socket) Run() *Socket {
	if so.path == "" {
		so.path = "/ws"
	}
	http.HandleFunc(so.path, so.reciver)
	return so
}

func (so *Socket) Path(path string) *Socket {
	so.path = path
	return so
}

func (so *Socket) AddHandle(check func(*Socket, string)) {
	so.clienteToServer = append(so.clienteToServer, check)
}

func (so *Socket) GetHandles() []func(*Socket, string) {
	return so.clienteToServer
}

func (so *Socket) GetReciver() string {
	return so.sms
}

func (so *Socket) Send(name string, sms string) {
	so.clients[name].WriteMessage(1, []byte(sms))
}

func (so *Socket) IsConnected(client string) bool {
	return so.stateclients[client]
}

func (so *Socket) reciver(w http.ResponseWriter, r *http.Request) {
	conn, _ := so.upgrader.Upgrade(w, r, nil)
	log.Println("new client conected")
	id := fmt.Sprintf("%p", conn)
	so.clients[id] = conn
	so.stateclients[id] = true
	defer conn.Close()

	for {
		_, tempSMS, _ := conn.ReadMessage()
		so.sms = string(tempSMS)
		for _, caller := range so.clienteToServer {
			caller(so, fmt.Sprintf("%p", conn))
		}
	}
}
