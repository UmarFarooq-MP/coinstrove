package apirepository

import (
	"log"
)

func GetKrakenPriceBTC(resp map[string]interface{}) string {
	if results, found := resp["result"].(map[string]interface{}); found {
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
		//handle error - the map didn't contain this key
		err := "Failed to extract price from KrakenPriceBTC"
		log.Printf("%s\n", err)
		return ""
	}

	return ""
}

func GetKrakenPriceETH(resp map[string]interface{}) string {
	if results, found := resp["result"].(map[string]interface{}); found {
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
		//handle error - the map didn't contain this key
		err := "Failed to extract price from KrakenPriceETH"
		log.Printf("%s\n", err)
		return ""
	}

	return ""
}
