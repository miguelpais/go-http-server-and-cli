package routing

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type RouteDispatcher struct {
	routes map[string]RouteHandler
}

func MakeRegisterRoute() *RouteDispatcher {
	return &RouteDispatcher{
		routes: make(map[string]RouteHandler, 0),
	}
}

func (r *RouteDispatcher) RegisterRoute(url string, handler RouteHandler) {
	r.routes[url] = handler
}

func (r *RouteDispatcher) Route(request string, connection net.Conn) error {
	path, error := GetRequestPath(request)
	if error != nil {
		return errors.New("Request path not found for request, disregarding...")
	}
	handler, ok := r.routes[path]
	if !ok {
		return errors.New(fmt.Sprintf("RouteHandler has no route registered for path %s, disregarding", path))
	}

	handler.Handle(request, connection)
	return nil
}

func GetRequestPath(request string) (string, error) {
	lines := strings.Split(request, "\n")
	if len(lines) < 1 {
		return "", errors.New("Request is a single line, which is incorrect")
	}

	firstLineParts := strings.Split(lines[0], " ")
	if len(firstLineParts) < 2 {
		return "", errors.New("Could not obtain request path for request")
	}

	return firstLineParts[1], nil
}
