package main

import (
	"coinscience/api/websocket"
	"coinscience/internal/core/ports"
	"coinscience/internal/core/publisher"
	"coinscience/internal/core/services/realtimeprice/binance"
	gate_io "coinscience/internal/core/services/realtimeprice/gate.io"
	"coinscience/repositories/apirepository"
	"log"
	"net/http"
	"time"
)

func getPriceAfterFiveSeconds(priceService []ports.PriceService) {
	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		for _, value := range priceService {
			go value.GetThePrice()
		}
	}
}

func startServer() {
	log.Println("Starting Server on port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func main() {

	apiRepo := apirepository.NewAPIRepository()
	handler := websocket.NewHandler()
	broadCastManager := websocket.NewBroadcastManager(handler)

	quePublisher, err := publisher.NewRabbitMQPublisher("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error while Initiating Rabbit MQ Connection with message %v", err)
	}
	quePublisher.Init()
	defer quePublisher.Close()

	websocket.NewRouter(handler)
	go startServer()

	priceService := []ports.PriceService{
		binance.NewBinanceService(apiRepo, broadCastManager, quePublisher),
		gate_io.NewGateIOService(apiRepo, broadCastManager, quePublisher),
	}

	// WebSocket Endpoints
	go getPriceAfterFiveSeconds(priceService)

	select {}
}
