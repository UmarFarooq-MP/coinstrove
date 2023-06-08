package binance

import (
	"coinscience/consts"
	"coinscience/internal/core/ports"
	"log"
)

type newBinanceService struct {
	priceRepo        ports.PriceRepository
	data             []map[string]interface{}
	broadcastHandler ports.BroadCastHandler
}

func NewBinanceService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler) ports.PriceService {
	return &newBinanceService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
	}
}

func (binance *newBinanceService) GetThePrice() {
	binance.data = binance.priceRepo.Get(consts.BINANCE)
	binance.BroadCast()
}

func (binance *newBinanceService) BroadCast() {
	log.Println("BroadCasting binance data")
	binance.broadcastHandler.BroadCast(binance.data)

}
