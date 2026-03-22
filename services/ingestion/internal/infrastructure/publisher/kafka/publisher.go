package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"coinstrove/services/ingestion/internal/config"
	"coinstrove/services/ingestion/internal/domain"
)

type Publisher struct {
	cfg config.KafkaConfig
	// TODO: add kafka producer (e.g. github.com/segmentio/kafka-go)
}

func NewPublisher(cfg config.KafkaConfig) *Publisher {
	return &Publisher{cfg: cfg}
}

func (p *Publisher) Publish(ctx context.Context, trade domain.Trade) error {
	msg := Message{
		Exchange:  string(trade.Exchange),
		Symbol:    trade.Symbol.String(),
		Price:     trade.Price.String(),
		Timestamp: trade.Timestamp.UnixMilli(),
		TradeID:   trade.TradeID,
	}

	_, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal trade: %w", err)
	}

	// TODO: produce to kafka topic
	return nil
}

func (p *Publisher) Close() error {
	// TODO: flush + close producer
	return nil
}
