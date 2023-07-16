package consts

type CURRENCY string

const (
	ETH  CURRENCY = "ETC"
	BTC           = "BTC"
	USDT          = "USDT"
	BUSD          = "BUSD"
)

type EXCHANGE string

const (
	BINANCE  EXCHANGE = "BINANCE"
	GATEIO            = "GATE.IO"
	KRAKEN            = "KRAKEN"
	COINBASE          = "COINBASE"
	BITPAY            = "BITPAY"
	BITFINEX          = "BITFINEX"
)
