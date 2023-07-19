package handler

import (
	"context"
	"fmt"
	reader2 "http-server/internal/app/http_server/handler/reader"
	errorHandlers "http-server/internal/app/http_server/handler/responses"
	"http-server/internal/app/http_server/routing"
	"net"
	"time"
)

func SpawnHandler(readChannel <-chan net.Conn, routeDispatcher *routing.RouteDispatcher) {
	for {
		select {
		case conn := <-readChannel:
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

			reader := reader2.RequestReader{}
			request, err := reader.ReadHttpRequest(ctx, conn)
			if err != nil {
				conn.Close()
				continue
			}

			err = routeDispatcher.Route(request, conn)
			if err != nil {
				errorHandlers.BadRequestResponse{}.Respond(conn)
				fmt.Println("Could not route request, error was ", err)
			}
			cancel()
		}
	}
}
