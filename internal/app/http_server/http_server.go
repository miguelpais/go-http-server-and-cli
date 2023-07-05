package http_server

import (
	"fmt"
	"http-server/internal/app/http_server/handler"
	"http-server/internal/app/http_server/routing"
	"http-server/internal/app/routes/api"
	"http-server/internal/app/routes/ui"
	"net"
)

type HttpServer struct{}

var routeDispatcher *routes.RouteDispatcher

func init() {
	routeDispatcher = routes.MakeRegisterRoute()
	routeDispatcher.RegisterRoute("/", ui.RouteUIIndex{})
	routeDispatcher.RegisterRoute("/api", api.RouteApi{})
}

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
		clientConnection, err := connection.Accept()
		if err != nil {
			panic("Could not accept connection")
		}
		fmt.Println("Incoming request received, handling...")
		go handler.HttpHandler{}.Handle(clientConnection, routeDispatcher)
	}
}
