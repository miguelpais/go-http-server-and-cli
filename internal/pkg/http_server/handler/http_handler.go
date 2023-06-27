package handler

import (
	"fmt"
	"net"
)

type HttpHandler struct{}

func (h HttpHandler) Handle(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 2048)
	connection.Read(buffer)
	fmt.Sprintf("Received request: \n%s", buffer)
	content := "<HTML><body>LOL</body></html>"
	connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nLocation: http://0.0.0.0:8000/\r\nContent-Length: %d\r\nContent-Type: text/html; charset=UTF-8\r\n\n%s\r\n\r\n", len(content), content)))
}
