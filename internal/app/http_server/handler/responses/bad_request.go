package responses

import (
	"net"
)

type BadRequestResponse struct{}

func (u BadRequestResponse) Respond(connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 400 BAD REQUEST\r\n" +
			"Content-Type: text/htlm\r\n\r\n"))
	connection.Close()
}
