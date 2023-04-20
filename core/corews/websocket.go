package corews

import (
	"net/http"

	"nhooyr.io/websocket"
)

const (
	InternalError = websocket.StatusInternalError
	NormalClosure = websocket.StatusNormalClosure
)

func NewConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}
