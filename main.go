package main

import (
	"mainnet/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers.SetRoutes(r)

	r.Run(":4000")
}
