package domain

// Exchange hold the information regarding a exchange that has different currencies involved
type Exchange struct {
	ExchangeName string     `json:"exchange_name" bson:"exchange_name"`
	Currencies   []Currency `json:"Currencies" bson:"Currencies"`
}

// Currency holds information regarding a particular currency
type Currency struct {
	Name  string `json:"name" bson:"name"`
	Price string `json:"price" bson:"price"`
}

// Response Websocket Response struct
type Response struct {
	Data         Exchange `json:"data" bson:"data"`
	ErrorMessage string   `bson:"errorMessage" json:"errorMessage"`
}
