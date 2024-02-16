package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"time"
	_ "time/tzdata"
)

type CarConfigurationEntry struct {
	ChargeMode                     int    `json:"chargeMode"`
	MinimumSoC                     int    `json:"minimumSoC"`
	LatestTimeToReachSoC           string `json:"latestTimeToReachSoC"`
	IgnoreLatestTimeToReachSocDate bool   `json:"ignoreLatestTimeToReachSocDate"`
	MaximumAmpere                  int    `json:"maximumAmpere"`
	MinimumAmpere                  int    `json:"minimumAmpere"`
	UsableEnergy                   int    `json:"usableEnergy"`
	ShouldBeManaged                bool   `json:"shouldBeManaged"`
	ShouldSetChargeStartTimes      bool   `json:"shouldSetChargeStartTimes"`
	ChargingPriority               int    `json:"chargingPriority"`
}

func scheduleCharge(carid int, chargelimit int) bool {
	var successful bool
	var PostData CarConfigurationEntry
	successful = false

	location, err := time.LoadLocation(Cfg.TimeZone)
	if err != nil {
		panic(err)
	}
	currentTime := time.Now().UTC().In(location)
	TimeToReachLayout := "2006-01-02T15:04:05"

	if Cfg.Debug {
		logger.Printf("futureTime: %s\n", currentTime.Format(TimeToReachLayout))
	}

	fSeconds := currentTime.Second()
	fAdd := 59 - fSeconds
	futureTime := currentTime.Add(10 * time.Minute).Add(time.Duration(fAdd) * time.Second)

	PostData = PreviousCarSettings
	PostData.MinimumSoC = chargelimit
	PostData.LatestTimeToReachSoC = futureTime.Format(TimeToReachLayout)

	if Cfg.Debug {
		m, _ := json.MarshalIndent(PostData, "", "  ")
		fmt.Printf("scheduleCharge, PostData: %s\n", m)
	}

	if Cfg.DryRun == false {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json-patch+json").
			SetBody(PostData).
			Put(Cfg.Tscapi + "/api/Config/UpdateCarConfiguration?carId=" + strconv.Itoa(carid))
		if err != nil {
			log.Printf("Error during scheduleCharge: %+v\n", err)
			successful = false
		}

		//fmt.Printf("Status Code: %d\n", resp.StatusCode())
		if resp.StatusCode() == 200 {
			successful = true
		}
	} else {
		logger.Printf("Dry run!\n")
	}
	return successful
}

func stopCharge(carid int, chargelimit int) bool {
	var successful bool
	var PostData CarConfigurationEntry
	successful = false

	currentTime := time.Now()
	//TimeToReachLayout := "2006-01-02T15:04:05"
	TimeToReachLayout := time.RFC3339

	if !isSpotCharge(carid) {
		return false
	}

	if Cfg.Debug {
		logger.Printf("stopCharge, currentTime: %s\n", currentTime.Format(TimeToReachLayout))
	}
	pSeconds := currentTime.Second()
	pAdd := 59 - pSeconds
	pastTime := currentTime.Add(-10 * time.Minute).Add(time.Duration(pAdd) * time.Second)

	PostData = PreviousCarSettings
	if PreviousCarSettings.MinimumSoC == chargelimit {
		PostData.MinimumSoC = Cfg.FallbackSocLimit
	} else {
		PostData.MinimumSoC = PreviousCarSettings.MinimumSoC
	}
	PostData.LatestTimeToReachSoC = pastTime.Format(TimeToReachLayout)

	if Cfg.Debug {
		m, _ := json.MarshalIndent(PostData, "", "  ")
		fmt.Printf("Debug, stopCharge, PostData: %s\n", m)
	}
	if Cfg.DryRun == false {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Content-Type", "application/json-patch+json").
			SetBody(PostData).
			Put(Cfg.Tscapi + "/api/Config/UpdateCarConfiguration?carId=" + strconv.Itoa(carid))
		if err != nil {
			log.Printf("Error during scheduleCharge: %+v\n", err)
			successful = false
		}

		//fmt.Printf("Status Code: %d\n", resp.StatusCode())
		if resp.StatusCode() == 200 {
			successful = true
		} else {
			logger.Printf("Dry run!\n")
		}
	}
	return successful
}
