package kafka

// Message is the shape published to Kafka
type Message struct {
	Exchange  string `json:"exchange"`
	Symbol    string `json:"symbol"`
	Price     string `json:"price"`
	Timestamp int64  `json:"timestamp"`
	TradeID   string `json:"trade_id"`
}
