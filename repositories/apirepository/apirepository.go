package apirepository

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
	"coinstrove/pkg/http"
)

type apirepository struct {
	client *http.Client
}

func NewAPIRepository() ports.PriceRepository {
	return &apirepository{
		client: http.NewHttpClientWithTimeout(2),
	}
}

func (repo *apirepository) Get(exchange consts.EXCHANGE) domain.Response {
	var responseMap domain.Response
	switch exchange {
	case consts.OKX:
		resp, err := repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=BTC-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=ETH-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=ADA-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=DOGE-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=DOT-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=LTC-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=BCH-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=XRP-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=SOL-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.okx.com/api/v5/market/ticker?instId=LINK-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetOkxPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		responseMap.Data.ExchangeName = "Okx"
	case consts.KUCOIN:
		resp, err := repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=BTC-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=ETH-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=ADA-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=DOGE-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=DOT-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=BCH-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=XRP-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=SOL-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kucoin.com/api/v1/market/orderbook/level1?symbol=LINK-USDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetKucoinPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Kucoin"
	case consts.BITSTAMP:

		resp, err := repo.client.Get("https://www.bitstamp.net/api/v2/ticker/btcusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/ethusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/adausd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/dogeusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/dotusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/ltcusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/bchusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/xrpusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/solusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://www.bitstamp.net/api/v2/ticker/linkusd/")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetBitstampPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		responseMap.Data.ExchangeName = "Bitstamp"

	case consts.HUOBI:
		resp, err := repo.client.Get("https://api.huobi.pro/market/trade?symbol=btcusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=ethusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=adausdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=dogeusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=dotusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=ltcusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=bchusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=xrpusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=solusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.huobi.pro/market/trade?symbol=linkusdt")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetHuobiPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Huobi"
	case consts.BITFINEX:
		resp, err := repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tBTCUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tETHUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tADAUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tDOGE:USD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tDOTUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tLTCUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tBCHN:USD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tXRPUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tSOLUSD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api-pub.bitfinex.com/v2/ticker/tLINK:USD")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetBitfinexPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		responseMap.Data.ExchangeName = "Bitfinex"
	case consts.BITPAY:
		resp, err := repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=BTCUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=ETHUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=ADAUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=DOGEUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=DOTUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=LTCUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=BCHUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=XRPUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=SOLUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.bybit.com/public/linear/recent-trading-records?symbol=LINKUSDT&limit=1")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetBitPayPrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		responseMap.Data.ExchangeName = "BitPay"

	case consts.COINBASE:
		resp, err := repo.client.Get("https://api.coinbase.com/v2/prices/BTC-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/ETH-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/ADA-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/DOGE-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/DOT-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/LTC-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/BCH-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/XRP-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/SOL-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.coinbase.com/v2/prices/LINK-USD/buy")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetCoinBasePrice(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Coinbase"

	case consts.KRAKEN:
		resp, err := repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=BTCUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetKrakenPriceBTC(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=ETHUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetKrakenPriceETH(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=ADAUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetKrakenPriceADA(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=DOGEUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetKrakenPriceDOGE(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=DOTUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetKrakenPriceDOT(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=LTCUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetKrakenPriceLTC(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=BCHUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetKrakenPriceBCH(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=XRPUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetKrakenPriceXRP(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=SOLUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetKrakenPriceSOL(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.kraken.com/0/public/Ticker?pair=LINKUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetKrakenPriceLINK(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Kraken"

	case consts.BINANCE:
		resp, err := repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=ADAUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=DOGEUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=DOTUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=LTCUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=BCHUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=XRPUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=SOLUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=LINKUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetPriceForBinance(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Binance"

	case consts.GATEIO:
		resp, err := repo.client.Get("https://data.gateapi.io/api2/1/ticker/btc_usdt")
		if err == nil {
			//since the response does not contain any information about coin type so appended that information
			//in response, so we can detect the coin name
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/eth_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/doge_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOGE",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/ada_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ADA",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/dot_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "DOT",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/ltc_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LTC",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/BCH_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BCH",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/XRP_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "XRP",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/sol_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "SOL",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/link_usdt")
		if err == nil {

			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "LINK",
				Price: GetPriceForGateIO(resp),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Gate.io"
	}
	return responseMap
}
