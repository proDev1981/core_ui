package core

import (
	"fmt"
	"log"
	"net/http"
)

// struct server
type Server struct {
	port          string
	root          string
	routes        []string
	contain       string
	static        string
	socket        *Socket
	initialEvents map[string]Event
	exit          bool
}

// constructor server
func NewServer() *Server {
	return &Server{initialEvents: make(map[string]Event)}
}

// function init server default
func InitDefaultServer() {
	NewServer().Static("./src").Listen()
}

// setter field port
func (s *Server) Port(port string) *Server {
	s.port = port
	return s
}

// setter field root server
func (s *Server) Root(root string) *Server {
	s.root = root
	return s
}

// setter field contain html
func (s *Server) Contain(contain string) *Server {
	s.contain = contain
	return s
}

// setter path folder static server
func (s *Server) Static(static string) *Server {
	s.static = static
	return s
}

// setter option constructor socket in server
func (s *Server) AndSocket() *Server {
	s.socket = NewSocket().Run()
	return s
}

// getter socket in server
func (s *Server) GetSocket() *Socket {
	return s.socket
}

// setter path socket in server
func (s *Server) SocketWitchPath(path string) *Server {
	s.socket = NewSocket().Path(path).Run()
	return s
}

// setter handles socket in server
func (s *Server) AddHandles(calls ...func(*Socket, string)) *Server {
	for _, item := range calls {
		s.socket.AddHandle(item)
	}
	return s
}

// setter initial events of page html
func (s *Server) SetInitialEvents(id string, types string, call func(*Event)) {
	s.initialEvents[id] = Event{Type: types, Call: call}
}

// asset sender initial event page html to socket
func (s *Server) sendInitialEvents() {
	for na, ev := range s.initialEvents {
		name := ev.Type + na
		s.socket.initialEvents[na] = ev
		s.socket.functions[name] = ev.Call
	}
}

// await exit to ui
func (s *Server) Await() {
	s.sendInitialEvents()
	for s.socket.conection {
	}
}

// configuration default server
func (s *Server) Default() bool {
	if s.port == "" {
		s.port = ":3000"
	}
	if s.root == "" {
		s.root = "/"
	}
	if s.static == "" {
		s.static = "./src"
	}
	if s.contain != "" {
		return true
	}
	return false
}

// run and listen server
func (s *Server) Listen() {
	if s.Default() {
		http.HandleFunc(s.root, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, s.contain)
		})
	} else {
		http.Handle(s.root, http.FileServer(http.Dir(s.static)))
	}
	log.Println("Server listen in =>", s.port)
	go http.ListenAndServe(s.port, nil)
	s.Await()
}
