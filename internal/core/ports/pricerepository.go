package ports

import (
	"coinstrove/consts"
	"coinstrove/internal/core/domain"
)

type PriceRepository interface {
	// Get When we want to query the external api to get the latest price
	Get(exchange consts.EXCHANGE) domain.Response
}
