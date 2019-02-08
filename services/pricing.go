package services

import (
	"math"
	"time"
)

type PricingService struct {
	PricePerHour float64
	PricePerDay  float64
}

func (p PricingService) Calculate(startDate time.Time, endDate time.Time) float64 {
	duration := endDate.Sub(startDate)

	days := math.Floor(duration.Hours() / 24)
	hours := math.Mod(duration.Hours(), 24)
	pricePerHour := hours * p.PricePerHour

	if pricePerHour > p.PricePerDay {
		pricePerHour = p.PricePerDay
	}

	return math.Max(0, days*p.PricePerDay+pricePerHour)
}
