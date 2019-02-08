package main

import (
	"github.com/daursu/go-pricing/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/status", controllers.Status)
	priceGroup := router.Group("pricing")
	{
		priceController := controllers.Price{}
		priceGroup.POST("calculate", priceController.Calculate)
	}

	err := router.Run("127.0.0.1:3000")

	if err != nil {
		log.Panic(err)
	}
}
