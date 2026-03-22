package ports

import (
	"context"

	"coinstrove/services/ingestion/internal/domain"
)

type EventPublisher interface {
	Publish(ctx context.Context, trade domain.Trade) error
	Close() error
}
