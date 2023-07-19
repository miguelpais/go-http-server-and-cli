package responses

import (
	"net"
)

type NotFoundResponse struct{}

func (u NotFoundResponse) Respond(connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 404 NOT FOUND\r\n" +
			"Content-Type: text/htlm\r\n\r\n"))
	connection.Close()
}
