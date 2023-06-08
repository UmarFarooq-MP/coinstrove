package ports

type PriceService interface {
	// GetThePrice When we want to query the external api to get the latest price
	GetThePrice()
	// BroadCast will broadcast the latest price
	BroadCast()
}
