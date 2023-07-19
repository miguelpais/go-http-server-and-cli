package limiter

import (
	"errors"
	"http-server/internal/app/http_server/handler/consumer"
	"http-server/internal/app/http_server/handler/responses"
	"net"
	"sync"
	"time"
)

const (
	TOKENS_DEPTH_SIZE      int = 1
	MAX_QUEUED_CONNECTIONS int = 20
	RATE_LIMITED_BUFFER    int = 2000
)

type Limiter struct {
	tokensBucketDepth        int
	pendingConnectionsQueue  chan net.Conn
	AcceptedConnectionsQueue chan net.Conn
	tokensMutex              sync.Mutex
}

func MakeRateLimiter() *Limiter {
	limiter := Limiter{
		tokensBucketDepth:        TOKENS_DEPTH_SIZE,
		pendingConnectionsQueue:  make(chan net.Conn, RATE_LIMITED_BUFFER),
		AcceptedConnectionsQueue: make(chan net.Conn, MAX_QUEUED_CONNECTIONS),
		tokensMutex:              sync.Mutex{},
	}

	return &limiter
}

func (l *Limiter) ProceedOrBufferConnection(conn net.Conn) (bool, error) {
	l.tokensMutex.Lock()
	if l.tokensBucketDepth > 0 {
		l.tokensBucketDepth--
		l.tokensMutex.Unlock()

		go Refill(l)
		return true, nil
	}

	l.tokensMutex.Unlock()

	select {
	case l.pendingConnectionsQueue <- conn:
	default:
		return false, errors.New("buffer is full, message should be discarded")
	}
	return false, nil
}

func Refill(l *Limiter) {
	time.AfterFunc(5000*time.Microsecond, func() {
		l.tokensMutex.Lock()
		if l.tokensBucketDepth < TOKENS_DEPTH_SIZE {
			select {
			case conn := <-l.pendingConnectionsQueue:
				select {
				case l.AcceptedConnectionsQueue <- conn:
					go Refill(l)
				default:
					select {
					case l.pendingConnectionsQueue <- conn:
						l.tokensBucketDepth++
					default:
						consumer.Consumer{}.ConsumeAndRespond(conn, responses.TooManyRequestsResponse{})
					}
				}
			default:
				l.tokensBucketDepth++
			}
		}

		l.tokensMutex.Unlock()
	})
}
