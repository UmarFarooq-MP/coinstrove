package ports

import "coinstrove/internal/core/domain"

type Publisher interface {
	Publish(data domain.Response)
	Close()
	Init()
}
