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
		responseMap.Data.ExchangeName = "Kraken"

	case consts.BINANCE:
		resp, err := repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "BTC",
				Price: resp["price"].(string),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT")
		if err == nil {
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: resp["price"].(string),
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
				Price: resp["last"].(string),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}

		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/eth_usdt")
		if err == nil {
			//since the response does not contain any information about coin type so appended that information
			//in response, so we can detect the coin name
			responseMap.Data.Currencies = append(responseMap.Data.Currencies, domain.Currency{
				Name:  "ETH",
				Price: resp["last"].(string),
			})
		} else {
			responseMap.ErrorMessage = err.Error()
		}
		responseMap.Data.ExchangeName = "Gate.io"
	}
	return responseMap
}
