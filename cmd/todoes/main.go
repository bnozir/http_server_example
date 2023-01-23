package main

import (
	"net"
	"net/http"

	"github.com/bnozir/todoes/internal/routes"
	"github.com/bnozir/todoes/internal/server"
)

func main() {
	routes := &routes.Routes{ServeMux: &http.ServeMux{}}
	routes.Init()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &server.Server{
		Listener: listener,
		Server: http.Server{
			Handler: routes.ServeMux,
		}}

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
