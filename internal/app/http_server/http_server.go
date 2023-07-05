package http_server

import (
	"fmt"
	"http-server/internal/app/http_server/handler"
	"http-server/internal/app/http_server/handler/errors"
	"http-server/internal/app/http_server/routing"
	"http-server/internal/app/routes/api"
	"http-server/internal/app/routes/ui"
	"net"
)

type HttpServer struct {
	connections_queue chan net.Conn
}

func BuildHttpServer() HttpServer {
	routeDispatcher := routing.MakeRegisterRoute()
	routeDispatcher.RegisterRoute("/", ui.RouteUIIndex{})
	routeDispatcher.RegisterRoute("/api", api.RouteApi{})

	server := HttpServer{
		connections_queue: make(chan net.Conn, 5000),
	}

	for i := 0; i < 500; i++ {
		go handler.SpawnHandler(server.connections_queue, routeDispatcher)
	}

	return server
}

func (h HttpServer) Serve(host, path string) {
	connection, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		panic("Could not listen at address")
	}
	fmt.Println("Accepting connections..")
	for true {
		clientConnection, err := connection.Accept()
		if err != nil {
			panic("Could not accept connection")
		}
		fmt.Println("Incoming request received, handling...")
		select {
		case h.connections_queue <- clientConnection:
		default:
			errors.TooManyRequestsHandler{}.Handle("", clientConnection)
		}
	}
}
