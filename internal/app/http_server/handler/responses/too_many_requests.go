package responses

import (
	"net"
)

type TooManyRequestsResponse struct{}

func (u TooManyRequestsResponse) Respond(connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 429 TOO MANY REQUESTS\r\n" +
			"Content-Type: text/htlm\r\n" +
			"Retry-After: 3600\r\n\r\n"))
	connection.Close()
}
