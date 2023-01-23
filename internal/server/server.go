package server

import (
	"net"
	"net/http"
)

type Server struct {
	Server   http.Server
	Listener net.Listener
}

func (s *Server) Run() error {
	return s.Server.Serve(s.Listener)
}
