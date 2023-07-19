package reader

import (
	"context"
	"errors"
	"io"
	"strings"
)

type RequestReader struct{}

func (r RequestReader) ReadHttpRequest(ctx context.Context, reader io.Reader) (string, error) {
	var request []byte
	var buffer = make([]byte, 1024)

readLoop:
	for {
		select {
		case <-ctx.Done():
			return "", errors.New("context cancelled, read timeout reached")
		default:
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

			request = append(request, buffer[:nRead]...)

			if detectEndOfHttpRequest(buffer[:nRead]) {
				break readLoop
			}
		}
	}

	return string(request), nil
}

func detectEndOfHttpRequest(buffer []byte) bool {
	return strings.IndexAny(string(buffer), "\r\n\r\n") != -1
}
