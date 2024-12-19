package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/bnozir/todoes/internal/routes"
	"github.com/bnozir/todoes/internal/server"
)

func main() {
	r := &routes.Routes{ServeMux: &http.ServeMux{}}
	r.Init()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	s := &server.Server{
		Listener: listener,
		Server: http.Server{
			Handler: r.ServeMux,
		}}

	fmt.Println("Server is running on port 1234")
	err = s.Run()
	if err != nil {
		panic(err)
	}
}
