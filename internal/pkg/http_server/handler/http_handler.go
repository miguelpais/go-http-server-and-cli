package handler

import (
	"fmt"
	"http-server/internal/pkg/http_server/routes"
	"net"
)

type HttpHandler struct{}

func (h HttpHandler) Handle(connection net.Conn) {
	defer connection.Close()
	reader := RequestReader{}
	request, error := reader.ReadHttpRequest(connection)
	if error != nil {
		fmt.Sprintf("Could not read request, error was %s, disregarding...", error)
	}

	fmt.Printf("Received request: \n%s", request)
	routes.RouteDispatcherSingleton().Route(request, connection)
}
