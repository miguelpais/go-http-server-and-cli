package routing

import "net"

type RouteHandler interface {
	Handle(request string, connection net.Conn)
}
