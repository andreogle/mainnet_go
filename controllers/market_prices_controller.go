package controllers

import (
	"fmt"

	"mainnet_go/utils"

	"github.com/gin-gonic/gin"
)

type params struct {
	ID string `json:"id"`
}

type MarketPrice struct {
	Ticker string `json:"ticker"`
}

func Latest(c *gin.Context) {
	fmt.Printf("[Binance] Fetching latest prices for LINK_BTC")
	res, _ := utils.HttpGet("https://api.binance.com/api/v1/ticker/24hr?ticker=LINK_BTC")
	fmt.Printf("[Binance] Fetching latest prices for %v", res.Body)

	c.JSON(200, gin.H{
		"message": res.Body,
	})
}
