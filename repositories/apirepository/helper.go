package apirepository

func GetKrakenPrice(resp map[string]interface{}) string {
	if results, found := resp["result"].(map[string]interface{}); found {
		if xbtusdt, ok := results["XBTUSDT"].(map[string]interface{}); ok {
			return xbtusdt["o"].(string)
		}
	} else {
		//handle error - the map didn't contain this key
	}
	return ""
}
