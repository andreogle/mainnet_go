package exchanges

import (
	"fmt"
	"mainnet_go/utils"
)

type BinancePrice struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           uint   `json:"openTime"`
	CloseTime          uint   `json:"closeTime"`
	FirstID            uint   `json:"firstId"`
	LastID             uint   `json:"lastId"`
	Count              uint   `json:"count"`
}

type BinancePrices struct {
	Collection []BinancePrice
}

const baseURL string = "https://api.binance.com"

// FetchMarketTicker Fetches the latest price from Binance for a given symbol
func FetchMarketTicker(ticker string) (BinancePrice, error) {
	fmt.Printf("[Binance] Fetching latest prices for %s", ticker)

	url := fmt.Sprintf("%s/api/v1/ticker/24hr?symbol=%s", baseURL, ticker)

	price := BinancePrice{}
	utils.HttpGet(url, &price)

	return price, nil
}
