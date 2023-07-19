package responses

import "net"

type Responder interface {
	Respond(conn net.Conn)
}
