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

var maxIncomingConnections = 50

type HttpServer struct {
	connectionsQueue chan net.Conn
}

func BuildHttpServer() HttpServer {
	routeDispatcher := routing.MakeRegisterRoute()
	routeDispatcher.RegisterRoute("/", ui.RouteUIIndex{})
	routeDispatcher.RegisterRoute("/api", api.RouteApi{})

	server := HttpServer{
		connectionsQueue: make(chan net.Conn, maxIncomingConnections),
	}

	for i := 0; i < 10; i++ {
		go handler.SpawnHandler(server.connectionsQueue, routeDispatcher)
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
		fmt.Println("Number of requests in queue: %d", len(h.connectionsQueue))

		clientConnection, err := connection.Accept()
		if err != nil {
			panic("Could not accept connection")
		}

		select {
		case h.connectionsQueue <- clientConnection:
		default:
			errors.TooManyRequestsHandler{}.Handle("", clientConnection)
		}
	}
}
