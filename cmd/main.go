package main

import (
	"coinscience/pkg/http"
	"fmt"
)

func main() {
	client := http.NewHttpClientWithTimeout(2)
	fmt.Println(client.Get("https://data.gateapi.io/api2/1/ticker/eth_usdt"))
}
