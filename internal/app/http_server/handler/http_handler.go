package handler

import (
	"fmt"
	errorHandlers "http-server/internal/app/http_server/handler/errors"
	reader2 "http-server/internal/app/http_server/handler/reader"
	"http-server/internal/app/http_server/routing"
	"net"
)

func SpawnHandler(readChannel <-chan net.Conn, routeDispatcher *routing.RouteDispatcher) {
	for {
		select {
		case conn := <-readChannel:
			reader := reader2.RequestReader{}
			request, err := reader.ReadHttpRequest(conn)
			if err != nil {
				conn.Close()
				fmt.Sprintf("Could not read request, error was %s, closing connection...", err)
				
				continue
			}

			err = routeDispatcher.Route(request, conn)
			if err != nil {
				errorHandlers.BadRequestHandler{}.Handle("", conn)
				fmt.Println("Could not route request, error was ", err)
			}
		}
	}
}
