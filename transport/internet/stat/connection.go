package stat

import (
	"net"
)

type Connection interface {
	net.Conn
}

type CounterConnection struct {
	Connection
}

func (c *CounterConnection) Read(b []byte) (int, error) {
	nBytes, err := c.Connection.Read(b)

	return nBytes, err
}

func (c *CounterConnection) Write(b []byte) (int, error) {
	nBytes, err := c.Connection.Write(b)
	return nBytes, err
}
