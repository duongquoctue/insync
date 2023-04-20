package corews

import (
	"sync"

	"nhooyr.io/websocket"
)

type Channel struct {
	UID string

	conn *websocket.Conn

	subscribersMu sync.Mutex
	subscribers   map[*subscriber]struct{}
}

func (c *Channel) AddClient(client *Client) {
	c.Clients[client] = true
}
