package controllers

import (
	"fmt"
	"mainnet/utils"

	"github.com/gin-gonic/gin"
)

type params struct {
	id string `json:"id"`
}

type marketPricesData struct {
	ticker string `json:"ticker"`
}

func Latest(c *gin.Context) {
	fmt.Printf("[Binance] Fetching latest prices for LINK_BTC")
	res, _ := utils.HttpGet("https://api.binance.com/api/v1/ticker/24hr?ticker=LINK_BTC")
	fmt.Printf("[Binance] Fetching latest prices for %v", res.Body)

	c.JSON(200, gin.H{
		"message": res.Body,
	})
}
