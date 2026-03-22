package binance

import (
	"fmt"
	"strings"

	"coinstrove/services/ingestion/internal/domain"
)

const baseURL = "wss://stream.binance.com:9443/stream"

// buildURL builds the multiplexed stream URL
// e.g. wss://stream.binance.com:9443/stream?streams=btcusdt@trade/ethusdt@trade
func buildURL(symbols []domain.Symbol) string {
	streams := make([]string, len(symbols))
	for i, s := range symbols {
		streams[i] = fmt.Sprintf("%s@trade", strings.ToLower(s.String()))
	}
	return fmt.Sprintf("%s?streams=%s", baseURL, strings.Join(streams, "/"))
}
