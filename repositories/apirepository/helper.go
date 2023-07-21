package apirepository

import (
	"fmt"
	"log"
)

// GetKrakenPriceBTC  is a function which is parsing the response
func GetKrakenPriceBTC(resp interface{}) string {
	result, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetKrakenPriceBTC Error occurred while converting response into map")
		return ""
	}
	if results, found := result["result"].(map[string]interface{}); found {
		if xbtusdt, ok := results["XBTUSDT"].(map[string]interface{}); ok {
			if a, ok := xbtusdt["a"].([]interface{}); ok {
				if len(a) > 0 {
					if price, ok := a[0].(string); ok {
						return price
					}
				}
			}
		}
	} else {
		log.Print("Failed to extract price from KrakenPriceBTC")
		return ""
	}
	return ""
}

func GetOkxPrice(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetOkxPrice Error occurred while converting response into map")
		return ""
	}

	if data, ok := results["data"].([]interface{}); ok {

		if len(data) > 0 {
			lastData, ok := data[0].(map[string]interface{})
			if !ok {
				log.Println("Error occurred while extracting latest data")
				return ""
			}
			price, ok := lastData["last"].(string)
			if !ok {
				log.Println("Error occurred while extracting price")
				return ""
			}
			return price
		}
	}

	return ""
}

// GetKrakenPriceETH  is a function which is parsing the response
func GetKrakenPriceETH(resp interface{}) string {
	result, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetKrakenPriceETH Error occurred while converting response into map")
		return ""
	}

	if results, found := result["result"].(map[string]interface{}); found {
		if ethusdt, ok := results["ETHUSDT"].(map[string]interface{}); ok {
			if a, ok := ethusdt["a"].([]interface{}); ok {
				if len(a) > 0 {
					if price, ok := a[0].(string); ok {
						return price
					}
				}
			}
		}
	} else {
		log.Print("Failed to extract price from KrakenPriceETH")
		return ""
	}
	return ""
}

// GetBitPayPrice  is a function which is parsing the response to fetch the price
func GetBitPayPrice(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetBitPayPriceETH Error occurred while converting response into map")
		return ""
	}

	result, ok := results["result"].([]interface{})
	if !ok || len(result) == 0 {
		return ""
	}

	firstTrade, ok := result[0].(map[string]interface{})
	if !ok {
		return ""
	}

	price, ok := firstTrade["price"].(float64)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v", price)
}

// GetBitfinexPrice  is a function which is parsing the response to fetch the price
func GetBitfinexPrice(resp interface{}) string {

	dataSlice, ok := resp.([]interface{})
	if !ok || len(dataSlice) < 3 {
		return ""
	}

	thirdValue, ok := dataSlice[2].(float64)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%g", thirdValue)
}

// GetCoinBasePrice  is a function which is parsing the response
func GetCoinBasePrice(resp interface{}) string {
	response := resp.(map[string]interface{})
	data := response["data"].(map[string]interface{})
	return data["amount"].(string)
}

// GetPriceForGateIO  is a function which is parsing the response
func GetPriceForGateIO(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetPriceForGateIO Error occurred while converting response into map")
		return ""
	}
	return results["last"].(string)
}

// GetPriceForBinance  is a function which is parsing the response
func GetPriceForBinance(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetPriceForBinance Error occurred while converting response into map")
		return ""
	}
	return results["price"].(string)
}

/*
	func GetHuobiPrice(resp interface{}) string {
		results, ok := resp.(map[string]interface{})
		if !ok {
			log.Println("GetHuobiPrice Error occurred while converting response into map")
			return ""
		}

		data, ok := results["result"].([]interface{})
		if !ok || len(data) == 0 {
			return ""
		}

		price, ok := data[0].(map[string]interface{})
		if !ok {
			return ""
		}

		price, ok := ["price"].(float64)
		if !ok {
			return ""
		}
		return fmt.Sprintf("%v", price)
	}
*/
func GetHuobiPrice(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetHuobiPrice Error occurred while converting response into map")
		return ""
	}
	if tick, found := results["tick"].(map[string]interface{}); found {

		if data, ok := tick["data"].([]interface{}); ok {
			if len(data) > 0 {
				firstTrade, ok := data[0].(map[string]interface{})
				if !ok {
					log.Println("Error in GetHuobiPrice")
					return ""
				}
				price, ok := firstTrade["price"].(float64)
				if !ok {
					log.Println("Error occurred while extracting price")
					return ""
				}
				return fmt.Sprintf("%f", price)
			}
		}
	}
	return ""
}

func GetBitstampPrice(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetPriceForBitstamp Error occurred while converting response into map")
		return ""
	}
	return results["ask"].(string)
}

// GetKucoinPrice  is a function which is parsing the response
func GetKucoinPrice(resp interface{}) string {
	results, ok := resp.(map[string]interface{})
	if !ok {
		log.Println("GetKucoinPrice Error occurred while converting response into map")
		return ""
	}
	if data, found := results["data"].(map[string]interface{}); found {
		price, ok := data["price"].(string)
		if !ok {
			log.Println("GetKucoinPrice Error occurred")
		}
		return fmt.Sprintf("%s", price)

	}
	return ""
}
