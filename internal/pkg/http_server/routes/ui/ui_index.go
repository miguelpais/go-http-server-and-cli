package ui

import (
	"fmt"
	"net"
)

type UiIndex struct{}

func (u UiIndex) Handle(_ string, connection net.Conn) {
	content := "<HTML><body>LOL</body></html>"
	connection.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nLocation: http://0.0.0.0:8000/\r\nContent-Length: %d\r\nContent-Type: text/html; charset=UTF-8\r\n\n%s\r\n\r\n", len(content), content)))
}
