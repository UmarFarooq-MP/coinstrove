package gate_io

import (
	"coinscience/consts"
	"coinscience/internal/core/ports"
	"log"
)

type newGateIOService struct {
	priceRepo ports.PriceRepository
	data      []map[string]interface{}
}

func NewGateIOService(priceRepo ports.PriceRepository) ports.PriceService {
	return &newGateIOService{
		priceRepo: priceRepo,
	}
}

func (repo *newGateIOService) GetThePrice() {
	repo.data = repo.priceRepo.Get(consts.GATEIO)
	log.Println(repo.data)
}

func (repo *newGateIOService) Publish() {

}
