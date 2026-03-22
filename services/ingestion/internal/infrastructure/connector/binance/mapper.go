package binance

import (
	"fmt"
	"time"

	"coinstrove/services/ingestion/internal/domain"
	"github.com/shopspring/decimal"
)

func toDomainTrade(msg wsMessage) (domain.Trade, error) {
	price, err := decimal.NewFromString(msg.Data.Price)
	if err != nil {
		return domain.Trade{}, fmt.Errorf("invalid price: %w", err)
	}

	return domain.Trade{
		Exchange:  domain.Binance,
		Symbol:    parseSymbol(msg.Data.Symbol),
		Price:     price,
		Timestamp: time.UnixMilli(msg.Data.Timestamp),
		TradeID:   fmt.Sprintf("%d", msg.Data.TradeID),
	}, nil
}

// "BTCUSDT" → Symbol{Base: "BTC", Quote: "USDT"}
func parseSymbol(raw string) domain.Symbol {
	// TODO: handle more quote currencies
	quotes := []string{"USDT", "BTC", "ETH", "BNB"}
	for _, quote := range quotes {
		if len(raw) > len(quote) && raw[len(raw)-len(quote):] == quote {
			return domain.NewSymbol(raw[:len(raw)-len(quote)], quote)
		}
	}
	return domain.NewSymbol(raw, "")
}
