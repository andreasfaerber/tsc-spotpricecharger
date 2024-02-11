package main

import (
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type SpotPriceEntry struct {
	Id        int
	StartDate string
	EndDate   string
	Price     float64
}

//var SpotCount = 0

func GetSpotPrice() (float64, bool) {
	var SpotPrices []SpotPriceEntry
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetResult(&SpotPrices).
		Get(Cfg.Tscapi + "/api/ChargingCost/GetSpotPrices")
	if err != nil {
		log.Fatal(err)
		return 0, false
	}

	currentTime := time.Now()
	location, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	SpotPriceDateLayout := "2006-01-02T15:04:05"

	//SpotCount++

	// For Testing
	//if SpotCount > 10 {
	//	return 100, true
	//}
	//return -2, true

	for _, entry := range SpotPrices {
		StartDate, err := time.ParseInLocation(SpotPriceDateLayout, entry.StartDate, location)
		EndDate, err := time.ParseInLocation(SpotPriceDateLayout, entry.EndDate, location)
		if err != nil {
			panic(err)
		}
		if currentTime.After(StartDate) &&
			currentTime.Before(EndDate) {
			//fmt.Println("FOUND")
			return entry.Price, true
		}
	}

	return 0, false
}
