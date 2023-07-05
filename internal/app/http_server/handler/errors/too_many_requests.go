package errors

import (
	"fmt"
	"http-server/internal/app/http_server/handler"
	"net"
)

type TooManyRequestsHandler struct{}

func (u TooManyRequestsHandler) Handle(_ string, connection net.Conn) {
	reader := handler.RequestReader{}
	_, error := reader.ReadHttpRequest(connection)
	if error != nil {
		fmt.Println("Could not read request")
	}
	connection.Write([]byte(
		"HTTP/1.1 429 TOO MANY REQUESTS\r\n" +
			"Content-Type: text/htlm\r\n" +
			"Retry-After: 3600\r\n\r\n"))
	connection.Close()
}
