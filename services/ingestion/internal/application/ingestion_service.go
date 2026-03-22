package application

import (
	"context"

	"coinstrove/services/ingestion/internal/ports"
)

type IngestionService struct {
	connectors []ports.ExchangeConnector
	publisher  ports.EventPublisher
}

func NewIngestionService(publisher ports.EventPublisher, connectors ...ports.ExchangeConnector) *IngestionService {
	return &IngestionService{
		connectors: connectors,
		publisher:  publisher,
	}
}

func (s *IngestionService) Run(ctx context.Context) error {
	// TODO: fan-in all connectors → publish to kafka
	return nil
}
