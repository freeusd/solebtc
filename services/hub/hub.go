package hub

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Hub maintains active connections and broadcast messages
type Hub interface {
	PutConn(Conn)
	Broadcast([]byte)
	Len() int
}

// Conn describes a connection's behaviour
type Conn interface {
	Write([]byte) error
	Close() error
}

type connWrapper struct {
	write func([]byte) error
	close func() error
	mutex sync.Mutex
}

func (c *connWrapper) Write(raw []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.write(raw)
}

func (c *connWrapper) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.close()
}

func newConnWrapper(write func([]byte) error, close func() error) Conn {
	return &connWrapper{write: write, close: close}
}

// WrapPutWebsocketConn wraps func(Conn) to func(*websocket.Conn)
func WrapPutWebsocketConn(put func(Conn)) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		put(newConnWrapper(
			func(raw []byte) error {
				return c.WriteMessage(websocket.TextMessage, raw)
			},
			c.Close,
		))
	}
}
