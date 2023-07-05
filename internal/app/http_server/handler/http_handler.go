package handler

import (
	"fmt"
	"http-server/internal/app/http_server/routing"
	"net"
)

func SpawnHandler(readChannel <-chan net.Conn, routeDispatcher *routing.RouteDispatcher) {
	for {
		select {
		case conn := <-readChannel:
			defer conn.Close()
			reader := RequestReader{}
			request, error := reader.ReadHttpRequest(conn)
			if error != nil {
				fmt.Sprintf("Could not read request, error was %s, disregarding...", error)
			}

			fmt.Printf("Received request: \n%s", request)
			err := routeDispatcher.Route(request, conn)
			if err != nil {
				fmt.Println("Could not route request, error was ", err)
			}
		}
	}
}
