package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	Exchange  Exchange
	Symbol    Symbol
	Price     decimal.Decimal
	Timestamp time.Time
	TradeID   string
}
