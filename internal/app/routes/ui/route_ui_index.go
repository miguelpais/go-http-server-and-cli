package ui

import (
	"fmt"
	"net"
)

type RouteUIIndex struct{}

func (u RouteUIIndex) Handle(_ string, connection net.Conn) {
	content := "<HTML><body>LOL</body></html>"
	connection.Write([]byte(fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"Location: http://0.0.0.0:8000/\r\n"+
			"Content-Length: %d\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
			"%s",
		len(content), content)))
	connection.Close()
}
