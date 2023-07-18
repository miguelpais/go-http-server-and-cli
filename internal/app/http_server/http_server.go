package http_server

import (
	"fmt"
	"http-server/internal/app/http_server/handler"
	"http-server/internal/app/http_server/handler/errors"
	requestReader "http-server/internal/app/http_server/handler/reader"
	"http-server/internal/app/http_server/routing"
	"http-server/internal/app/routes/api"
	"http-server/internal/app/routes/ui"
	"net"
)

var maxWorkers = 20
var maxQueuedConnections = 20

type HttpServer struct {
	connectionsQueue chan net.Conn
}

func BuildHttpServer() HttpServer {
	routeDispatcher := routing.MakeRegisterRoute()
	routeDispatcher.RegisterRoute("/", ui.RouteUIIndex{})
	routeDispatcher.RegisterRoute("/api", api.RouteApi{})

	server := HttpServer{
		connectionsQueue: make(chan net.Conn, maxQueuedConnections),
	}

	for i := 0; i < maxWorkers; i++ {
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
		clientConnection, err := connection.Accept()
		if err != nil {
			panic("Could not accept connection")
		}

		select {
		case h.connectionsQueue <- clientConnection:
		default:
			reader := requestReader.RequestReader{}
			_, err := reader.ReadHttpRequest(clientConnection)
			if err != nil {
				fmt.Println("Could not read request, error was:")
				fmt.Println(err)
				clientConnection.Close()
			} else {
				errors.TooManyRequestsHandler{}.Handle("", clientConnection)
			}
		}
	}
}
