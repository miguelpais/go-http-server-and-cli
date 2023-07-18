package routing

import (
	"errors"
	"fmt"
	errorHandlers "http-server/internal/app/http_server/handler/errors"
	"net"
	"regexp"
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
		return error
	}
	handler, ok := r.routes[path]
	if !ok {
		errorHandlers.NotFoundHandler{}.Handle("", connection)
		return errors.New(fmt.Sprintf("RouteHandler has no route registered for path %s, returning NOT FOUND", path))
	}

	handler.Handle(request, connection)
	return nil
}

func GetRequestPath(request string) (string, error) {
	lines := strings.Split(request, "\r\n")
	if len(lines) < 1 {
		return "", fmt.Errorf("Request is a single line, which is incorrect")
	}

	firstLineParts := strings.Split(lines[0], " ")
	if len(firstLineParts) != 3 {
		return "", fmt.Errorf("First line is not composed of three space separated tokens")
	}

	pattern, err := regexp.Compile("^/[a-zA-Z_0-9!$&'+()*,;=:-@.~/]*$")
	if err != nil {
		return "", fmt.Errorf("Could not compile regexp pattern")
	}
	if !pattern.Match([]byte(firstLineParts[1])) {
		return "", fmt.Errorf("path: %s does not match an HTTP path", firstLineParts[1])
	}

	return firstLineParts[1], nil
}
