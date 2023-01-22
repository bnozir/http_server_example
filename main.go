package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", greating)

	server := &http.Server{}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server.Serve(listener)
}

func greating(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(response, http.StatusText(http.StatusMethodNotAllowed)+"\n")

		return
	}

	fmt.Fprint(response, "Hi, Bala!\n")
}
