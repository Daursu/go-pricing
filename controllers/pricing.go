package controllers

import (
	"github.com/daursu/go-pricing/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Object wrapping all methods
type Price struct{}

func (Price) Calculate(c *gin.Context) {
	startDate, _ := c.GetPostForm("start_date")
	endDate, _ := c.GetPostForm("end_date")

	// Convert them to a time object
	startDateObject, _ := time.Parse(time.RFC3339, startDate)
	endDateObject, _ := time.Parse(time.RFC3339, endDate)

	pricingService := services.PricingService{
		PricePerHour: 1,
		PricePerDay:  5,
	}

	price := pricingService.Calculate(startDateObject, endDateObject)

	c.JSON(http.StatusOK, gin.H{
		"price": price,
	})
}
