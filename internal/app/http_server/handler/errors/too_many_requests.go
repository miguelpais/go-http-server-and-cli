package errors

import (
	"net"
)

type TooManyRequestsHandler struct{}

func (u TooManyRequestsHandler) Handle(_ string, connection net.Conn) {
	connection.Write([]byte(
		"HTTP/1.1 429 TOO MANY REQUESTS\r\n" +
			"Content-Type: text/htlm\r\n" +
			"Retry-After: 3600\r\n\r\n"))
}
