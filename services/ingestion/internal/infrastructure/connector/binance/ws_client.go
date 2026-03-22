package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type wsClient struct {
	conn *websocket.Conn
}

func newWSClient(rawURL string) (*wsClient, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid ws url: %w", err)
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("ws dial failed: %w", err)
	}

	return &wsClient{conn: conn}, nil
}

func (c *wsClient) read(ctx context.Context) (<-chan wsMessage, <-chan error) {
	msgs := make(chan wsMessage)
	errs := make(chan error, 1)

	go func() {
		defer close(msgs)
		defer close(errs)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, raw, err := c.conn.ReadMessage()
				if err != nil {
					errs <- err
					return
				}

				var msg wsMessage
				if err := json.Unmarshal(raw, &msg); err != nil {
					// skip malformed messages
					continue
				}

				msgs <- msg
			}
		}
	}()

	return msgs, errs
}

func (c *wsClient) close() error {
	return c.conn.Close()
}
