package api

import (
	"fmt"
	"net"
)

type RouteApi struct{}

func (u RouteApi) Handle(_ string, connection net.Conn) {
	content := "{ \"response\": \"ok\" }"
	connection.Write([]byte(fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"Location: http://0.0.0.0:8000/\r\n"+
			"Content-Length: %d\r\nContent-Type: application/json; charset=UTF-8\r\n\r\n"+
			"%s",
		len(content), content)))
	connection.Close()
}
