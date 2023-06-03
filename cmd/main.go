package main

import (
	"coinscience/internal/core/ports"
	"coinscience/internal/core/services/realtimeprice/binance"
	gate_io "coinscience/internal/core/services/realtimeprice/gate.io"
	"coinscience/repositories/apirepository"
	"fmt"
	"log"
	"time"
)

func writeToQueue() {

	fmt.Println("Writing to kafka")
}

func getPriceAfterFiveSeconds(priceService []ports.PriceService) {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		log.Println("Getting Latest Prices")
		for _, value := range priceService {
			value.GetThePrice()
		}
		writeToQueue()
	}
}

func main() {

	apiRepo := apirepository.NewAPIRepository()
	priceService := []ports.PriceService{
		binance.NewBinanceService(apiRepo),
		gate_io.NewGateIOService(apiRepo),
	}

	go getPriceAfterFiveSeconds(priceService)
	select {}
}
