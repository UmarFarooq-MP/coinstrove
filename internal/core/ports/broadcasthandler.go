package ports

import "coinscience/internal/core/domain"

type BroadCastHandler interface {
	BroadCast(response domain.Response)
}
