package errors

import (
	"net"
)

type BadRequestHandler struct{}

func (u BadRequestHandler) Handle(_ string, connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 400 BAD REQUEST\r\n" +
			"Content-Type: text/htlm\r\n\r\n"))
	connection.Close()
}
