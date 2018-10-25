package controllers

import "github.com/gin-gonic/gin"

type params struct {
	id string `json:"id"`
}

type marketPricesData struct {
	ticker string `json:"ticker"`
}

func Latest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
