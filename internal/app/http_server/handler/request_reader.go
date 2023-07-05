package handler

import (
	"errors"
	"io"
	"strings"
)

type RequestReader struct{}

func (r RequestReader) ReadHttpRequest(reader io.Reader) (string, error) {
	var response []byte
	var buffer = make([]byte, 1024)
	for {
		nRead, error := reader.Read(buffer)

		if error == io.EOF {
			return string(response), nil
		}
		if error != nil {
			return "", errors.New("Reading request failed")
		}

		response = append(response, buffer[:nRead]...)

		if detectEndOfHttpRequest(buffer[:nRead]) {
			break
		}
	}

	return string(response), nil
}

func detectEndOfHttpRequest(buffer []byte) bool {
	return strings.IndexAny(string(buffer), "\r\n\r\n") != -1
}
