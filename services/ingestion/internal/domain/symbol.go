package domain

import "strings"

type Symbol struct {
	Base  string // BTC
	Quote string // USDT
}

func NewSymbol(base, quote string) Symbol {
	return Symbol{
		Base:  strings.ToUpper(base),
		Quote: strings.ToUpper(quote),
	}
}

func (s Symbol) String() string {
	return s.Base + s.Quote
}
