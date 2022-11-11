package http

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port    string
	root    string
	routes  []string
	contain string
	static  string
	socket  *Socket
}

func NewServer() *Server {
	return &Server{}
}
func InitDefaultServer() {
	NewServer().Static("./src").Listen()
}
func (s *Server) Port(port string) *Server {
	s.port = port
	return s
}
func (s *Server) Root(root string) *Server {
	s.root = root
	return s
}
func (s *Server) Contain(contain string) *Server {
	s.contain = contain
	return s
}
func (s *Server) Static(static string) *Server {
	s.static = static
	return s
}
func (s *Server) Socket() *Server {
	s.socket = NewSocket().Run()
	return s
}
func (s *Server) SocketWitchPath(path string) *Server {
	s.socket = NewSocket().Path(path).Run()
	return s
}

func (s *Server) AddHandles(calls ...func(*Socket, string)) *Server {
	for _, item := range calls {
		s.socket.AddHandle(item)
	}
	return s
}

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
func (s *Server) Listen() {
	if s.Default() {
		http.HandleFunc(s.root, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, s.contain)
		})
	} else {
		http.Handle(s.root, http.FileServer(http.Dir(s.static)))
	}
	log.Println("Server listen in =>", s.port)
	http.ListenAndServe(s.port, nil)
}
