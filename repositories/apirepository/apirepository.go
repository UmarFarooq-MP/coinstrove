package apirepository

import (
	"coinscience/consts"
	"coinscience/internal/core/ports"
	"coinscience/pkg/http"
)

type apirepository struct {
	client *http.Client
}

func NewAPIRepository() ports.PriceRepository {
	return &apirepository{
		client: http.NewHttpClientWithTimeout(2),
	}
}

func (repo *apirepository) Get(exchange consts.EXCHANGE) []map[string]interface{} {
	var responseMap []map[string]interface{}
	switch exchange {
	case consts.BINANCE:
		resp, err := repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT")
		if err == nil {
			responseMap = append(responseMap, resp)
		}
		resp, err = repo.client.Get("https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT")
		if err == nil {
			responseMap = append(responseMap, resp)
		}

	case consts.GATEIO:
		resp, err := repo.client.Get("https://data.gateapi.io/api2/1/ticker/btc_usdt")
		if err == nil {
			//since the response does not contain any information about coin type so appended that information
			//in response, so we can detect the coin name
			resp["currency"] = "btc"
			responseMap = append(responseMap, resp)
		}
		resp, err = repo.client.Get("https://data.gateapi.io/api2/1/ticker/eth_usdt")
		if err == nil {
			//since the response does not contain any information about coin type so appended that information
			//in response, so we can detect the coin name
			resp["currency"] = "eth"
			responseMap = append(responseMap, resp)
		}
	}
	return responseMap
}
