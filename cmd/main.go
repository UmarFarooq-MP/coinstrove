package main

import (
	"coinscience/internal/core/ports"
	"coinscience/internal/core/services/realtimeprice/binance"
	gate_io "coinscience/internal/core/services/realtimeprice/gate.io"
	"coinscience/repositories/apirepository"
)

func main() {
	apiRepo := apirepository.NewAPIRepository()
	var priceService []ports.PriceService
	priceService = append(priceService, binance.NewBinanceService(apiRepo))
	priceService = append(priceService, gate_io.NewGateIOService(apiRepo))

	for _, value := range priceService {
		value.GetThePrice()
	}
}
