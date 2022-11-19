package core

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// struct socket
type Socket struct {
	clients         map[string]*websocket.Conn
	path            string
	clienteToServer []func(*Socket, string)
	initialEvents   map[string]Event
	upgrader        websocket.Upgrader
	conection       bool
	loaded          map[string]bool
	sms             string
	functions       map[string]func(*Event)
	response        map[string]func(string)
	promises        map[string]*PROMISE
}

// constructor socket
func NewSocket() *Socket {
	return &Socket{
		clients:       make(map[string]*websocket.Conn),
		loaded:        make(map[string]bool),
		initialEvents: make(map[string]Event),
		functions:     make(map[string]func(*Event)),
		response:      make(map[string]func(string)),
		promises:      make(map[string]*PROMISE),
	}
}

// run struct socket witch configutarion this
func (so *Socket) Run() *Socket {
	if so.path == "" {
		so.path = "/ws"
	}
	http.HandleFunc(so.path, so.reciver)
	return so
}

// setter field path
func (so *Socket) Path(path string) *Socket {
	so.path = path
	return so
}

// add handle in socket
func (so *Socket) AddHandle(check func(*Socket, string)) {
	so.clienteToServer = append(so.clienteToServer, check)
}

// add handles in socket
func (so *Socket) GetHandles() []func(*Socket, string) {
	return so.clienteToServer
}

// add handle response
func (so *Socket) AddPromise(promise *PROMISE) {
	so.promises[promise.id] = promise
}

// getter sms recived
func (so *Socket) GetReciver() string {
	return so.sms
}

// method for send sms to client
func (so *Socket) Send(name string, sms string) {
	so.clients[name].WriteMessage(1, []byte(sms))
}

// method for send sms all clients
func (so *Socket) SendAll(sms string) {
	for _, item := range so.clients {
		item.WriteMessage(1, []byte(sms))
	}
}

// response if socket is connected
func (so *Socket) IsConnected() bool {
	return so.conection
}

// response if socket is loaded
func (so *Socket) IsLoaded(idConn string) bool {
	return so.loaded[idConn]
}

// assets for run action initials
func (so *Socket) initialActons(conn *websocket.Conn) {
	log.Printf("new client conected => %p\n", conn)
	id := fmt.Sprintf("%p", conn)
	so.clients[id] = conn
	so.loaded[id] = true
	so.conection = true
	so.sendInitialEvents(id, conn)
}

// asset  for sender initials event to client
func (so *Socket) sendInitialEvents(idconn string, conn *websocket.Conn) {
	for id, ev := range so.initialEvents {
		name := ev.Type + id
		registerFunctions := ComposeBind(name)
		registerEventsinElements := ComposeEventListener(id, ev.Type, name)
		so.Send(idconn, registerFunctions)
		so.Send(idconn, registerEventsinElements)
	}
}

// asset search match namo to functions in registre
func (so *Socket) searchFunctions(sms reciverSms) {
	if sms.Type == "event" {
		log.Printf("llamando a %s\n", sms.Name)
		go so.functions[sms.Name](&sms.Event)
	}
}

// convert json to struct
func (so *Socket) ConvertSms(data []byte) reciverSms {
	sms := reciverSms{}
	if err := json.Unmarshal(data, &sms); err != nil {
		log.Println("sms no valid error sytax")
	} else {
		//log.Println(sms)
	}
	return sms
}

// inyection conn cliente in sms reciver
func (so *Socket) AddClient(sms reciverSms, client *websocket.Conn) reciverSms {
	sms.Event.Client = client
	return sms
}

// manager promise in socket
func (so *Socket) PromiseManager(sms reciverSms) {
	if sms.Type == "response" {
		log.Println("promise reciver")
		if so.promises[sms.Name] != nil {
			promise := so.promises[sms.Name]
			promise.ChangeState("resolved")
			promise.data = sms.Event.Value
			promise.resolved(promise)
		}
	}
}

// handle reciver socket
func (so *Socket) reciver(w http.ResponseWriter, r *http.Request) {
	conn, _ := so.upgrader.Upgrade(w, r, nil)
	so.initialActons(conn)
	defer conn.Close()

	for {
		_, sms, _ := conn.ReadMessage()
		so.sms = string(sms)
		dataSms := so.AddClient(so.ConvertSms(sms), conn)
		log.Printf("el cliente %p manda =>%s", conn, so.sms) // debugger print
		so.searchFunctions(dataSms)
		so.PromiseManager(dataSms)
		for _, caller := range so.clienteToServer {
			caller(so, fmt.Sprintf("%p", conn))
		}
	}
}
