package main

import (
	"mainnet_go/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers.SetRoutes(r)

	r.Run(":4000")
}
