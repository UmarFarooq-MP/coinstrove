package websocket

import (
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
)

type broadCastManager struct {
	handler *Handler
}

func NewBroadcastManager(handler *Handler) ports.BroadCastHandler {
	return &broadCastManager{
		handler: handler,
	}
}

func (broadcast *broadCastManager) BroadCast(data domain.Response) {
	broadcast.handler.BroadCast(data)
}
