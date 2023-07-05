package handler

import (
	"fmt"
	"http-server/internal/app/http_server/routing"
	"net"
)

type HttpHandler struct{}

func (h HttpHandler) Handle(connection net.Conn, routeDispatcher *routes.RouteDispatcher) {
	defer connection.Close()
	reader := RequestReader{}
	request, error := reader.ReadHttpRequest(connection)
	if error != nil {
		fmt.Sprintf("Could not read request, error was %s, disregarding...", error)
	}

	fmt.Printf("Received request: \n%s", request)
	routeDispatcher.Route(request, connection)
}
