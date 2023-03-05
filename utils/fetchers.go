package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type TickerData struct {
	Ticker string `json:"symbol"`
	Price  string `json:"price"`
}

type MapPrice map[string]string

func FetchData() ([]TickerData, error) {
	resp, err := http.Get("https://api.exchange.klever.io/v1/market/tickers")
	if err != nil {
		log.Fatalf("Error fetching data: %q", err)
	}
	defer resp.Body.Close()

	var tickersData []TickerData
	err = json.NewDecoder(resp.Body).Decode(&tickersData)
	if err != nil {
		log.Fatalf("Error decoding data: %q", err)
	}

	return tickersData, nil
}

func CleanUpData(t []TickerData) ([]MapPrice, error) {
	var newTickersData []MapPrice
	for _, data := range t {
		symbol := data.Ticker
		price := data.Price
		if strings.Contains(symbol, "USDT") {
			symbol = strings.TrimSuffix(symbol, "-USDT")
			newTickersData = append(newTickersData, MapPrice{symbol: price})
		}
	}
	return newTickersData, nil
}
