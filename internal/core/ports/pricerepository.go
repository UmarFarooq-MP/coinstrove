package ports

import "coinscience/consts"

type PriceRepository interface {
	// Get When we want to query the external api to get the latest price
	Get(exchange consts.EXCHANGE) []map[string]interface{}
	// Publish When we want to publish the data to some que
	Publish()
}
