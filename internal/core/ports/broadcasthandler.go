package ports

import "coinstrove/internal/core/domain"

type BroadCastHandler interface {
	BroadCast(response domain.Response)
}
