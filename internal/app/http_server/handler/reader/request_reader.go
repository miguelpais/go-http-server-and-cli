package reader

import (
	"errors"
	"io"
	"strings"
)

type RequestReader struct{}

func (r RequestReader) ReadHttpRequest(reader io.Reader) (string, error) {
	var request []byte
	var buffer = make([]byte, 1024)

	for {
		nRead, err := reader.Read(buffer)
		if err == io.EOF {
			if len(request) > 0 {
				return string(request), nil
			} else {
				return "", errors.New("end of file got before content")
			}
		}
		if err != nil {
			return "", errors.New("Reading request failed")
		}

		if nRead == 0 {
			return "", errors.New("no bytes read")
		}

		request = append(request, buffer[:nRead]...)

		if detectEndOfHttpRequest(buffer[:nRead]) {
			break
		}
	}

	return string(request), nil
}

func detectEndOfHttpRequest(buffer []byte) bool {
	return strings.IndexAny(string(buffer), "\r\n\r\n") != -1
}
