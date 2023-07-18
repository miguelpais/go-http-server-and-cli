package errors

import (
	"net"
)

type NotFoundHandler struct{}

func (u NotFoundHandler) Handle(_ string, connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 404 NOT FOUND\r\n" +
			"Content-Type: text/htlm\r\n\r\n"))
	connection.Close()
}
