package main

import (
	"brawlstars-api-test/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/players/:player", controller.GetPlayer)
	}
	router.Run()
}
