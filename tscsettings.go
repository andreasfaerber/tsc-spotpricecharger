package main

import (
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type TSCSettingsEntry struct {
	IsNewVersionAvailable bool        `json:"isNewVersionAvailable"`
	InverterPower         interface{} `json:"inverterPower"`
	Overage               interface{} `json:"overage"`
	PowerBuffer           interface{} `json:"powerBuffer"`
	CarsToManage          []struct {
		Id               int    `json:"id"`
		Vin              string `json:"vin"`
		CarConfiguration struct {
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
		} `json:"carConfiguration"`
		CarState struct {
			Name                            string        `json:"name"`
			ShouldStartChargingSince        interface{}   `json:"shouldStartChargingSince"`
			EarliestSwitchOn                interface{}   `json:"earliestSwitchOn"`
			ShouldStopChargingSince         time.Time     `json:"shouldStopChargingSince"`
			EarliestSwitchOff               time.Time     `json:"earliestSwitchOff"`
			ScheduledChargingStartTime      interface{}   `json:"scheduledChargingStartTime"`
			SoC                             int           `json:"soC"`
			SocLimit                        int           `json:"socLimit"`
			IsHomeGeofence                  bool          `json:"isHomeGeofence"`
			TimeUntilFullCharge             string        `json:"timeUntilFullCharge"`
			ReachingMinSocAtFullSpeedCharge time.Time     `json:"reachingMinSocAtFullSpeedCharge"`
			AutoFullSpeedCharge             bool          `json:"autoFullSpeedCharge"`
			LastSetAmp                      int           `json:"lastSetAmp"`
			ChargerPhases                   int           `json:"chargerPhases"`
			ActualPhases                    int           `json:"actualPhases"`
			ChargerVoltage                  int           `json:"chargerVoltage"`
			ChargerActualCurrent            int           `json:"chargerActualCurrent"`
			ChargerPilotCurrent             int           `json:"chargerPilotCurrent"`
			ChargerRequestedCurrent         int           `json:"chargerRequestedCurrent"`
			PluggedIn                       bool          `json:"pluggedIn"`
			ClimateOn                       bool          `json:"climateOn"`
			Latitude                        float64       `json:"latitude"`
			Longitude                       float64       `json:"longitude"`
			DistanceToHomeGeofence          int           `json:"distanceToHomeGeofence"`
			ChargingPowerAtHome             int           `json:"chargingPowerAtHome"`
			State                           int           `json:"state"`
			Healthy                         bool          `json:"healthy"`
			ReducedChargeSpeedWarning       bool          `json:"reducedChargeSpeedWarning"`
			PlannedChargingSlots            []interface{} `json:"plannedChargingSlots"`
		} `json:"carState"`
	} `json:"carsToManage"`
	HomeBatterySoc                  interface{}   `json:"homeBatterySoc"`
	HomeBatteryPower                interface{}   `json:"homeBatteryPower"`
	ActiveIssues                    []interface{} `json:"activeIssues"`
	ControlledACarAtLastCycle       bool          `json:"controlledACarAtLastCycle"`
	LastPvValueUpdate               time.Time     `json:"lastPvValueUpdate"`
	AverageHomeGridVoltage          int           `json:"averageHomeGridVoltage"`
	TeslaApiRequestCounter          int           `json:"teslaApiRequestCounter"`
	CrashedOnStartup                bool          `json:"crashedOnStartup"`
	StartupCrashMessage             interface{}   `json:"startupCrashMessage"`
	FleetApiProxyNeeded             bool          `json:"fleetApiProxyNeeded"`
	AllowUnlimitedFleetApiRequests  bool          `json:"allowUnlimitedFleetApiRequests"`
	LastFleetApiRequestAllowedCheck time.Time     `json:"lastFleetApiRequestAllowedCheck"`
	Cars                            []struct {
		Id               int    `json:"id"`
		Vin              string `json:"vin"`
		CarConfiguration struct {
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
		} `json:"carConfiguration"`
		CarState struct {
			Name                            string        `json:"name"`
			ShouldStartChargingSince        interface{}   `json:"shouldStartChargingSince"`
			EarliestSwitchOn                interface{}   `json:"earliestSwitchOn"`
			ShouldStopChargingSince         time.Time     `json:"shouldStopChargingSince"`
			EarliestSwitchOff               time.Time     `json:"earliestSwitchOff"`
			ScheduledChargingStartTime      interface{}   `json:"scheduledChargingStartTime"`
			SoC                             int           `json:"soC"`
			SocLimit                        int           `json:"socLimit"`
			IsHomeGeofence                  bool          `json:"isHomeGeofence"`
			TimeUntilFullCharge             string        `json:"timeUntilFullCharge"`
			ReachingMinSocAtFullSpeedCharge time.Time     `json:"reachingMinSocAtFullSpeedCharge"`
			AutoFullSpeedCharge             bool          `json:"autoFullSpeedCharge"`
			LastSetAmp                      int           `json:"lastSetAmp"`
			ChargerPhases                   int           `json:"chargerPhases"`
			ActualPhases                    int           `json:"actualPhases"`
			ChargerVoltage                  int           `json:"chargerVoltage"`
			ChargerActualCurrent            int           `json:"chargerActualCurrent"`
			ChargerPilotCurrent             int           `json:"chargerPilotCurrent"`
			ChargerRequestedCurrent         int           `json:"chargerRequestedCurrent"`
			PluggedIn                       bool          `json:"pluggedIn"`
			ClimateOn                       bool          `json:"climateOn"`
			Latitude                        float64       `json:"latitude"`
			Longitude                       float64       `json:"longitude"`
			DistanceToHomeGeofence          int           `json:"distanceToHomeGeofence"`
			ChargingPowerAtHome             int           `json:"chargingPowerAtHome"`
			State                           int           `json:"state"`
			Healthy                         bool          `json:"healthy"`
			ReducedChargeSpeedWarning       bool          `json:"reducedChargeSpeedWarning"`
			PlannedChargingSlots            []interface{} `json:"plannedChargingSlots"`
		} `json:"carState"`
	} `json:"cars"`
}

func GetCarConfiguration(carid int, init bool) (TSCSettingsEntry, bool) {
	var successful bool
	client := resty.New()

	successful = false
	_, err := client.R().
		EnableTrace().
		SetResult(&TSCSettings).
		Get(Cfg.Tscapi + "/api/Config/GetSettings")
	if err != nil {
		log.Fatal(err)
		return TSCSettingsEntry{}, successful
	}
	successful = true
	if init == false {
		return TSCSettings, successful
	}
	for _, element := range TSCSettings.CarsToManage {
		//fmt.Printf("idx: %d\n", idx)
		//fmt.Printf("element.id: %+v\n", element.Id)
		//fmt.Printf("element.vin: %s\n", element.Vin)
		if element.Id == carid {
			if init == true {
				PreviousCarSettings = element.CarConfiguration
			}
		}
	}

	return TSCSettings, successful
}

func getSettingsCarId(carid int) (int, bool) {
	var settingsCarid int
	settingsCarid = -1
	GetCarConfiguration(Cfg.Carid, false)
	for i, e := range TSCSettings.CarsToManage {
		if e.Id == carid { // found our car
			settingsCarid = i
		}
	}
	if Cfg.Debug {
		logger.Printf("getSettingsCardId, CardID: %d, settingsCarid != -1: %t\n", carid, settingsCarid != -1)
	}
	return settingsCarid, settingsCarid != -1
}
