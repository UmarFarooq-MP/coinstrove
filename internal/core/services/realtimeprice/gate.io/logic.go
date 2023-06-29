package gate_io

import (
	"coinscience/consts"
	"coinscience/internal/core/domain"
	"coinscience/internal/core/ports"
)

type newGateIOService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewGateIOService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newGateIOService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (gateio *newGateIOService) GetThePrice() {
	gateio.data = gateio.priceRepo.Get(consts.GATEIO)
	gateio.BroadCast()
	gateio.WriteToQue()
}

func (gateio *newGateIOService) BroadCast() {
	gateio.broadcastHandler.BroadCast(gateio.data)
}

func (gateio *newGateIOService) WriteToQue() {
	gateio.publisher.Publish(gateio.data)
}
