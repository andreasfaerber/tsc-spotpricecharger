package main

func isCharging(carid int) bool {
	var sCarId, success = getSettingsCarId(carid)
	if TSCSettings.CarsToManage[sCarId].CarState.PluggedIn == false {
		return false
	}
	if success == false {
		return false // better panic?
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
