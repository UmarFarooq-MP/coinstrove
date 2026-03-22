package binance

import (
	"context"
	"fmt"

	"coinstrove/services/ingestion/internal/config"
	"coinstrove/services/ingestion/internal/domain"
)

type Connector struct {
	cfg    config.BinanceConfig
	client *wsClient
	trades chan domain.Trade
}

func NewConnector(cfg config.BinanceConfig) *Connector {
	return &Connector{
		cfg:    cfg,
		trades: make(chan domain.Trade),
	}
}

func (c *Connector) Connect(ctx context.Context) error {
	symbols := make([]domain.Symbol, len(c.cfg.Symbols))
	for i, s := range c.cfg.Symbols {
		// TODO: parse properly e.g. "btcusdt" → Symbol
		symbols[i] = domain.NewSymbol(s, "")
	}

	url := buildURL(symbols)
	client, err := newWSClient(url)
	if err != nil {
		return fmt.Errorf("binance connect: %w", err)
	}
	c.client = client

	go c.stream(ctx)

	return nil
}

func (c *Connector) Subscribe(symbols []domain.Symbol) error {
	// TODO: dynamic subscription after connect
	return nil
}

func (c *Connector) Trades() <-chan domain.Trade {
	return c.trades
}

func (c *Connector) Close() error {
	return c.client.close()
}

func (c *Connector) stream(ctx context.Context) {
	defer close(c.trades)

	msgs, errs := c.client.read(ctx)
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errs:
			if err != nil {
				// TODO: reconnect logic
				return
			}
		case msg, ok := <-msgs:
			if !ok {
				return
			}
			trade, err := toDomainTrade(msg)
			if err != nil {
				continue
			}
			c.trades <- trade
		}
	}
}
