package controllers

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine) {
	r.GET("/api/market-prices", Latest)
}
