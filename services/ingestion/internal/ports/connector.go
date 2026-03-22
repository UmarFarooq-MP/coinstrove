package ports

import (
	"context"

	"coinstrove/services/ingestion/internal/domain"
)

type ExchangeConnector interface {
	Connect(ctx context.Context) error
	Subscribe(symbols []domain.Symbol) error
	Trades() <-chan domain.Trade
	Close() error
}
