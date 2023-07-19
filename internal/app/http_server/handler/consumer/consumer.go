package consumer

import (
	"context"
	"fmt"
	requestReader "http-server/internal/app/http_server/handler/reader"
	"http-server/internal/app/http_server/handler/responses"
	"net"
	"time"
)

type Consumer struct{}

func (c Consumer) ConsumeAndRespond(connection net.Conn, responder responses.Responder) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	reader := requestReader.RequestReader{}
	_, err := reader.ReadHttpRequest(ctx, connection)
	cancel()
	if err != nil {
		fmt.Printf("Could not read request, error was: %s", err)
		connection.Close()
	}
	responder.Respond(connection)
}
