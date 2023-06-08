package gate_io

import (
	"coinscience/consts"
	"coinscience/internal/core/ports"
)

type newGateIOService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             []map[string]interface{}
}

func NewGateIOService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler) ports.PriceService {
	return &newGateIOService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
	}
}

func (gateio *newGateIOService) GetThePrice() {
	gateio.data = gateio.priceRepo.Get(consts.GATEIO)
}

func (gateio *newGateIOService) BroadCast() {
	gateio.broadcastHandler.BroadCast(gateio.data)
}
