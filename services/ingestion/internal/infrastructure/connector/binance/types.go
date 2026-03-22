package binance

// raw Binance WebSocket message shape
type wsMessage struct {
	Stream string    `json:"stream"`
	Data   tradeData `json:"data"`
}

type tradeData struct {
	Symbol    string `json:"s"` // BTCUSDT
	Price     string `json:"p"` // "94250.10"
	TradeID   int64  `json:"t"`
	Timestamp int64  `json:"T"`
}
