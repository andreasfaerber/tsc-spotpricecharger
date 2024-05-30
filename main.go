package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Cfg = Config{
		Tmapi:            "http://teslamateapi:8080",
		Tscapi:           "http://teslasolarcharger:7190",
		Checkinterval:    300,
		SpotChargePrice:  0.06,
		Carid:            1,
		ChargeSocLimit:   90,
		FallbackSocLimit: 90,
		Debug:            false,
		StartupDelay:     180,
		DryRun:           false,
		TimeZone:         "Europe/Berlin",
	}
	PreviousCarSettings = CarConfigurationEntry{}
	TSCSettings         TSCSettingsEntry

	ErrEnvVarEmpty = errors.New("getenv: environment variable empty")
	logger         = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

type Config struct {
	Tmapi            string
	Tscapi           string
	Gridbasecost     float64
	Checkinterval    int
	SpotChargePrice  float64
	Carid            int
	ChargeSocLimit   int
	FallbackSocLimit int
	Debug            bool
	StartupDelay     int
	DryRun           bool
	TimeZone         string
}

func checkPriceLoop(carid int, spotchargeprice float64) {
	var settingsCarid int
	for {
		SpotPrice, success := GetSpotPrice()
		if success == false {
			logger.Println("Unable to get find spot price!")
			continue
		}
		GetCarConfiguration(Cfg.Carid, false)
		for i, e := range TSCSettings.CarsToManage {
			if e.Id == carid { // found our car
				settingsCarid = i
			}
		}
		message := fmt.Sprintf("Car %d, PluggedIn: %t, SoC: %d, SocLimit (SpotCharge): %d, SoCLimit (Car): %d,"+
			" ChargerActualCurrent: %d", carid,
			TSCSettings.CarsToManage[settingsCarid].PluggedIn,
			TSCSettings.CarsToManage[settingsCarid].SoC,
			Cfg.ChargeSocLimit,
			TSCSettings.CarsToManage[settingsCarid].SocLimit,
			TSCSettings.CarsToManage[settingsCarid].ChargerActualCurrent)
		message += ", SpotPrice NOW: %f"
		if SpotPrice > spotchargeprice {
			message += " - above charge price of %f."
			if isCharging(carid) {
				if isSpotCharge(carid) {
					message += " Stopping charge."
					stopCharge(carid, Cfg.ChargeSocLimit)
				} else {
					message += " Charging (and/or climate running). Assuming non-spot charge, no action"
				}
			}
			logger.Printf(message, SpotPrice, spotchargeprice)
		} else {
			message += " - below charge price of %f."
			if isCharging(carid) {
				message += " Charging."
				scheduleCharge(carid, Cfg.ChargeSocLimit)
			} else {
				msg, ok := shouldCharge(carid)
				if ok {
					message += " Scheduling charge."
					scheduleCharge(carid, Cfg.ChargeSocLimit)
				} else {
					message += " " + msg
				}
			}
			message += "\n"
			logger.Printf(message, SpotPrice, spotchargeprice)
		}
		//logger.Printf("Next check in %d seconds\n", Cfg.Checkinterval)
		time.Sleep(time.Duration(Cfg.Checkinterval) * time.Second)
	}
}

func main() {
	getEnvironment()
	logger.Printf("Startup Delay: Sleeping %d seconds\n", Cfg.StartupDelay)
	time.Sleep(time.Duration(Cfg.StartupDelay) * time.Second)

	if _, success := GetCarConfiguration(Cfg.Carid, true); success == false {
		logger.Fatalf("Fatal: Unable to get TSC configuration for carid %d\n", Cfg.Carid)
	}

	checkPriceLoop(Cfg.Carid, Cfg.SpotChargePrice)
}
