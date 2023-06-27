package http_server

import (
	"fmt"
	"http-server/internal/pkg/http_server/handler"
	"net"
)

type HttpServer struct{}

func BuildHttpServer() HttpServer {
	return HttpServer{}
}

func (h HttpServer) Serve(host, path string) {
	connection, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		panic("Could not listen at address")
	}
	fmt.Println("Accepting connections..")
	for true {
		client_connection, err := connection.Accept()
		if err != nil {
			panic("Could not accept connection")
		}
		fmt.Println("Incoming request received, handling...")
		go handler.HttpHandler{}.Handle(client_connection)
	}
}
