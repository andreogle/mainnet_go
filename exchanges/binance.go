package exchanges

type BinancePrice struct {
	symbol             string `json:"symbol"`
	priceChange        string `json:"priceChange"`
	priceChangePercent string `json:"priceChangePercent"`
	weightedAvgPrice   string `json:"weightedAvgPrice"`
	prevClosePrice     string `json:"prevClosePrice"`
	lastPrice          string `json:"lastPrice"`
	lastQty            string `json:"lastQty"`
	bidPrice           string `json:"bidPrice"`
	askPrice           string `json:"askPrice"`
	openPrice          string `json:"openPrice"`
	highPrice          string `json:"highPrice"`
	lowPrice           string `json:"lowPrice"`
	volume             string `json:"volume"`
	quoteVolume        string `json:"quoteVolume"`
	openTime           uint   `json:"openTime"`
	closeTime          uint   `json:"openTime"`
	openTime           uint   `json:"openTime"`
	openTime           uint   `json:"openTime"`
}
