package main

import (
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
	"time"
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

	currentTime := time.Now()
	TimeToReachLayout := "2006-01-02T15:04:05"
	futureTime := currentTime.Add(10 * time.Minute)

	PostData = PreviousCarSettings
	PostData.MinimumSoC = chargelimit
	PostData.LatestTimeToReachSoC = futureTime.Format(TimeToReachLayout)

	//m, _ := json.MarshalIndent(PostData, "", "  ")
	//fmt.Printf("PostData: %s\n", m)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json-patch+json").
		SetBody(PostData).
		Put(Cfg.Tscapi + "/api/Config/UpdateCarConfiguration?carId=" + strconv.Itoa(carid))
	if err != nil {
		log.Fatal(err)
		successful = false
	}

	//fmt.Printf("Status Code: %d\n", resp.StatusCode())
	if resp.StatusCode() == 200 {
		successful = true
	}
	return successful
}

func stopCharge(carid int, chargelimit int) bool {
	var successful bool
	var PostData CarConfigurationEntry
	successful = false

	currentTime := time.Now()
	TimeToReachLayout := "2006-01-02T15:04:05"
	pastTime := currentTime.Add(-10 * time.Minute)

	//m, _ := json.MarshalIndent(PreviousCarSettings, "", "  ")
	//fmt.Printf("PreviousCarSettings: %s\n", m)

	PostData = PreviousCarSettings
	if PreviousCarSettings.MinimumSoC == chargelimit {
		PostData.MinimumSoC = Cfg.FallbackSocLimit
	} else {
		PostData.MinimumSoC = PreviousCarSettings.MinimumSoC
	}
	PostData.LatestTimeToReachSoC = pastTime.Format(TimeToReachLayout)

	//m, _ := json.MarshalIndent(PostData, "", "  ")
	//fmt.Printf("PostData: %s\n", m)
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json-patch+json").
		SetBody(PostData).
		Put(Cfg.Tscapi + "/api/Config/UpdateCarConfiguration?carId=" + strconv.Itoa(carid))
	if err != nil {
		log.Fatal(err)
		successful = false
	}

	//fmt.Printf("Status Code: %d\n", resp.StatusCode())
	if resp.StatusCode() == 200 {
		successful = true
	}
	return successful
}
