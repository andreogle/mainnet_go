package controllers

import (
	"fmt"

	"mainnet_go/utils"

	"github.com/gin-gonic/gin"
)

type params struct {
	ID string `json:"id"`
}

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

func Latest(c *gin.Context) {
	price, err :=

	c.JSON(200, gin.H{
		"message": price,
	})
}
