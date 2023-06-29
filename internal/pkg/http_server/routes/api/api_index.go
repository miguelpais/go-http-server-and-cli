package api

import (
	"fmt"
	"net"
)

type ApiIndex struct{}

func (u ApiIndex) Handle(_ string, connection net.Conn) {
	content := "{ \"response\": \"ok\" }"
	connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nLocation: http://0.0.0.0:8000/\r\nContent-Length: %d\r\nContent-Type: application/json; charset=UTF-8\r\n\n%s\r\n\r\n", len(content), content)))
}
