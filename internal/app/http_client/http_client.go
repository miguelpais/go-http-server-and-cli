package http_client

import (
	"fmt"
	"net"
)

type HttpClient struct{}

func BuildHttpClient() HttpClient {
	return HttpClient{}
}

func (h HttpClient) Get(protocol, host, path string) string {
	connection, err := net.Dial("tcp", host)
	if err != nil {
		return fmt.Sprintf("Could not connect to host %s on path %s, error is %s", host, path, err)
	}
	http_request := []byte(fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\nUser-Agent: Mozilla/5.0\r\nAccept: */*\r\n\r\n", path, host))
	//fmt.Printf("Sending following http GET request\n%s", http_request)
	_, err = connection.Write(http_request)
	if err != nil {
		return "Could not perform request to " + path
	}
	buffer := make([]byte, 2048)
	readLength, err := connection.Read(buffer)
	if err != nil {
		return fmt.Sprintf("Could not read response, error is %s", err)
	}

	return string(buffer[:readLength])
}
