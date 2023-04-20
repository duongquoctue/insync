package corews

type subscriber struct {
	writePump chan ClientMessage
	ReadPump  chan ClientMessage
}

func (c *Client) SendMessage(msg ClientMessage) {
	c.writePump <- msg
}
