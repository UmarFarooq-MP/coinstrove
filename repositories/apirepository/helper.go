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
		log.Println("GetKrakenPriceETH Error occurred while converting response into map")
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
