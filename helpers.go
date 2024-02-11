package main

import (
	"fmt"
	"time"
)

func isCharging(carid int) bool {
	var sCarId, success = getSettingsCarId(carid)
	if success == false {
		if Cfg.Debug {
			logger.Printf("isCharging, returning false\n")
		}
		return false // better panic?
	}
	if TSCSettings.CarsToManage[sCarId].CarState.PluggedIn == false {
		if Cfg.Debug {
			logger.Printf("isCharging, returning false\n")
		}
		return false
	}
	if Cfg.Debug {
		logger.Printf("isCharging, returning %t\n",
			TSCSettings.CarsToManage[sCarId].CarState.ChargerActualCurrent > 0)
	}
	return TSCSettings.CarsToManage[sCarId].CarState.ChargerActualCurrent > 0
}

func shouldCharge(carid int) bool {
	var sCarId, success = getSettingsCarId(carid)
	if !success {
		return false
	}
	if TSCSettings.CarsToManage[sCarId].CarState.PluggedIn == false {
		return false
	}
	if TSCSettings.CarsToManage[sCarId].CarState.SoC >= Cfg.ChargeSocLimit {
		return false
	}
	return true
}

func isSpotCharge(carid int) bool {
	var sCarId, success = getSettingsCarId(carid)
	if !success {
		return false
	}
	timeStampLayout := "2006-01-02T15:04:05"
	timeLatestTimeToReachSoC := TSCSettings.CarsToManage[sCarId].CarConfiguration.LatestTimeToReachSoC
	latestTimeToReachSoC, err := time.Parse(timeStampLayout, timeLatestTimeToReachSoC)
	if err != nil {
		fmt.Printf("isSpotCharge, Error: %+v\n", err)
		return false
	}
	latestSecond := latestTimeToReachSoC.Second()
	if Cfg.Debug {
		logger.Printf("Debug, isSpotCharge, LatestTimeToReachSoC: %s\n", latestTimeToReachSoC.Format(timeStampLayout))
		logger.Printf("Debug, isSpotCharge, Second: %d\n", latestSecond)
	}
	if latestSecond != 59 {
		if Cfg.Debug {
			logger.Printf("Debug, isSpotCharge, returning false\n")
		}
		return false
	}
	if Cfg.Debug {
		logger.Printf("Debug, isSpotCharge, returning true\n")
	}
	return true
}
