package services

import (
	"testing"
	"time"
)

func TestCalculatesHourlyPrice(t *testing.T) {
	pricingService := PricingService{
		PricePerHour: 1,
		PricePerDay:  10,
	}

	var eps = 0.00000001
	price := pricingService.Calculate(time.Now(), time.Now().Add(time.Hour*time.Duration(1)))

	if price-1 > eps {
		t.Errorf("Expected price to be %f, instead got %f", float64(1), price)
	}
}

func TestItCalculatesPricePerDay(t *testing.T) {
	pricingService := PricingService{
		PricePerHour: 1,
		PricePerDay:  5,
	}

	var eps = 0.00000001
	var expected = 5.0
	price := pricingService.Calculate(time.Now(), time.Now().Add(time.Hour*time.Duration(24)))

	if price-expected > eps {
		t.Errorf("Expected price to be %f, instead got %f", expected, price)
	}
}

func TestPricePerDaySuperceedesPricePerHour(t *testing.T) {
	pricingService := PricingService{
		PricePerHour: 1,
		PricePerDay:  5,
	}

	var eps = 0.00000001
	var expected = 5.0
	price := pricingService.Calculate(time.Now(), time.Now().Add(time.Hour*time.Duration(6)))

	if price-expected > eps {
		t.Errorf("Expected price to be %f, instead got %f", expected, price)
	}
}

func TestItDoesntReturnANegativePrice(t *testing.T) {
	pricingService := PricingService{
		PricePerHour: 1,
		PricePerDay:  5,
	}

	var expected = 0.0
	price := pricingService.Calculate(time.Now(), time.Now().Add(time.Hour*time.Duration(6)*-1))

	if price < 0 {
		t.Errorf("Expected price to be %f, instead got %f", expected, price)
	}
}
