package http_server

import (
	"fmt"
	"http-server/internal/app/http_server/handler"
	"http-server/internal/app/http_server/handler/consumer"
	"http-server/internal/app/http_server/handler/responses"
	"http-server/internal/app/http_server/limiter"
	"http-server/internal/app/http_server/routing"
	"http-server/internal/app/routes/api"
	"http-server/internal/app/routes/ui"
	"net"
)

const (
	MAX_WORKERS int = 1
)

type HttpServer struct {
	rateLimiter *limiter.Limiter
}

func BuildHttpServer() HttpServer {
	routeDispatcher := routing.MakeRegisterRoute()
	routeDispatcher.RegisterRoute("/", ui.RouteUIIndex{})
	routeDispatcher.RegisterRoute("/api", api.RouteApi{})

	server := HttpServer{
		rateLimiter: limiter.MakeRateLimiter(),
	}

	for i := 0; i < MAX_WORKERS; i++ {
		go handler.SpawnHandler(server.rateLimiter.AcceptedConnectionsQueue, routeDispatcher)
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

		if proceed, err := h.rateLimiter.ProceedOrBufferConnection(clientConnection); err != nil {
			// connection could not be buffered
			consumer.Consumer{}.ConsumeAndRespond(clientConnection, responses.TooManyRequestsResponse{})
			continue
		} else if !proceed {
			continue
		}

		select {
		case h.rateLimiter.AcceptedConnectionsQueue <- clientConnection:
		default:
			fmt.Println("Our connections buffer is at max capacity, still answer with too many requests")
			consumer.Consumer{}.ConsumeAndRespond(clientConnection, responses.TooManyRequestsResponse{})
		}
	}
}
