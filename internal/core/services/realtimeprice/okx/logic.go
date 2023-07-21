package okx

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
	"coinstrove/internal/core/ports"
)

type newOkxService struct {
	priceRepo        ports.PriceRepository
	broadcastHandler ports.BroadCastHandler
	data             domain.Response
	publisher        ports.Publisher
}

func NewOkxService(priceRepo ports.PriceRepository, broadcaster ports.BroadCastHandler, publisher ports.Publisher) ports.PriceService {
	return &newOkxService{
		priceRepo:        priceRepo,
		broadcastHandler: broadcaster,
		publisher:        publisher,
	}
}

func (okx *newOkxService) GetThePrice() {
	okx.data = okx.priceRepo.Get(consts.OKX)
	okx.BroadCast()
	okx.WriteToQue()
}

func (okx *newOkxService) BroadCast() {
	okx.broadcastHandler.BroadCast(okx.data)
}

func (okx *newOkxService) WriteToQue() {
	okx.publisher.Publish(okx.data)
}
