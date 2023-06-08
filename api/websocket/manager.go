package websocket

import "coinscience/internal/core/ports"

type broadCastManager struct {
	handler *Handler
}

func NewBroadcastManager(handler *Handler) ports.BroadCastHandler {
	return &broadCastManager{
		handler: handler,
	}
}

func (broadcast *broadCastManager) BroadCast(data []map[string]interface{}) {
	broadcast.handler.BroadCast(data)
}
