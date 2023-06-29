package ports

import "coinscience/internal/core/domain"

type Publisher interface {
	Publish(data domain.Response)
	Close()
	Init()
}
