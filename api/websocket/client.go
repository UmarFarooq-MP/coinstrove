package websocket

import (
	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	handler    *Handler
}

func NewClient(conn *websocket.Conn, handler *Handler) *Client {
	return &Client{
		connection: conn,
		handler:    handler,
	}
}
