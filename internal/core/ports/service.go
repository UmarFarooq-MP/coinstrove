package ports

type PriceService interface {
	// Publish When we want to publish the data to some que
	Publish()
	// GetThePrice When we want to query the external api to get the latest price
	GetThePrice()
}
