package handlers

import (
	"http-server/internal/pkg/http_client"
	"strings"
)

type CurlHandler struct{}

const HTTP = "http://"

func (c CurlHandler) Identifier() string {
	return "curl"
}
func (c CurlHandler) Help() string {
	return "curl <url>: makes an HTTP GET request against the specified url and prints the HTTP response"
}

func (c CurlHandler) Process(params []string) string {
	if len(params) != 1 {
		return "curl: Url not provided"
	}

	httpClient := http_client.BuildHttpClient()
	baseUrl := params[0]
	hostIdx := strings.Index(baseUrl, HTTP)
	if hostIdx != -1 {
		hostIdx += len(HTTP)
		baseUrl = baseUrl[hostIdx:]
	}
	slashIdx := strings.Index(baseUrl, "/")

	var host, path string
	if slashIdx == -1 {
		path = "/"
		host = baseUrl
	} else {
		host = baseUrl[:slashIdx]
		path = baseUrl[slashIdx:]
	}

	return httpClient.Get("http", host, path)
}
