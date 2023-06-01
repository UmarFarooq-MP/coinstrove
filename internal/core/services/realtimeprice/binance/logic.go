package binance

import (
	"coinscience/internal/core/ports"
)

type newBinanceService struct {
	priceRepo ports.PriceRepository
}

func NewBinanceService(priceRepo ports.PriceRepository) ports.PriceService {
	return &newBinanceService{
		priceRepo: priceRepo,
	}
}

func (repo *newBinanceService) GetThePrice() {

}

func (repo *newBinanceService) Publish() {

}
