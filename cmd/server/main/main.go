package main

import (
	"http-server/internal/app/http_server"
)

func main() {
	server := http_server.BuildHttpServer()
	server.Serve("0.0.0.0:8000", "/")
}
