package main

import (
	"os"
	"strconv"
)

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetenvFloat(key string) (float64, error) {
	s, err := GetenvStr(key)
	f, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return 0, err
	}
	return f, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	b, err := strconv.ParseBool(s)
	if nil != err {
		return false, err
	}
	return b, nil
}

func getEnvironment() {
	var ident = "%-30s"
	// Read environment variables
	if s, err := GetenvStr("TSC_SPOT_TESLAMATEAPI_URL"); err != nil {
		logger.Printf(ident+" undefined, using default: %s\n", "TSC_SPOT_TESLAMATEAPI_URL", Cfg.Tmapi)
	} else {
		Cfg.Tmapi = s
	}
	if s, err := GetenvStr("TSC_SPOT_TSC_URL"); err != nil {
		logger.Printf(ident+" undefined, using default: %s\n", "TSC_SPOT_TSC_URL", Cfg.Tscapi)
	} else {
		Cfg.Tscapi = s
	}
	if f, err := GetenvFloat("TSC_SPOT_SPOTCHARGEPRICE"); err != nil {
		logger.Printf(ident+" undefined, using default: %f\n", "TSC_SPOT_SPOTCHARGEPRICE", Cfg.SpotChargePrice)
	} else {
		Cfg.SpotChargePrice = f
	}
	if i, err := GetenvInt("TSC_SPOT_CHECKINTERVAL"); err != nil {
		logger.Printf(ident+" undefined, using default: %d\n", "TSC_SPOT_CHECKINTERVAL", Cfg.Checkinterval)
	} else {
		Cfg.Checkinterval = i
	}
	if i, err := GetenvInt("TSC_SPOT_CARID"); err != nil {
		logger.Printf(ident+" undefined, using default: %d\n", "TSC_SPOT_CARID", Cfg.Carid)
	} else {
		Cfg.Carid = i
	}
	if i, err := GetenvInt("TSC_SPOT_CHARGESOCLIMIT"); err != nil {
		logger.Printf(ident+" undefined, using default: %d\n", "TSC_SPOT_CHARGESOCLIMIT", Cfg.ChargeSocLimit)
	} else {
		Cfg.ChargeSocLimit = i
	}
	if i, err := GetenvInt("TSC_SPOT_FALLBACKCHARGELIMIT"); err != nil {
		logger.Printf(ident+" undefined, using default: %d\n", "TSC_SPOT_FALLBACKCHARGELIMIT", Cfg.FallbackSocLimit)
	} else {
		Cfg.FallbackSocLimit = i
	}
	if i, err := GetenvInt("TSC_SPOT_STARTUPDELAY"); err != nil {
		logger.Printf(ident+" undefined, using default: %d\n", "TSC_SPOT_STARTUPDELAY", Cfg.ChargeSocLimit)
	} else {
		Cfg.StartupDelay = i
	}
	if b, err := GetenvBool("TSC_SPOT_DEBUG"); err != nil {
		logger.Printf(ident+" undefined, using default: %t\n", "TSC_SPOT_DEBUG", Cfg.Debug)
	} else {
		Cfg.Debug = b
	}
	if b, err := GetenvBool("TSC_SPOT_DRYRUN"); err != nil {
		logger.Printf(ident+" undefined, using default: %t\n", "TSC_SPOT_DRYRUN", Cfg.DryRun)
	} else {
		Cfg.DryRun = b
	}
	if s, err := GetenvStr("TZ"); err != nil {
		logger.Printf(ident+" undefined, using default: %t\n", "TZ", Cfg.DryRun)
	} else {
		Cfg.TimeZone = s
	}

	logger.Printf(ident+" %s\n", "TSC_SPOT_TESLAMATEAPI_URL", Cfg.Tmapi)
	logger.Printf(ident+" %s\n", "TSC_SPOT_TSC_URL", Cfg.Tscapi)
	logger.Printf(ident+" %f\n", "TSC_SPOT_SPOTCHARGEPRICE", Cfg.SpotChargePrice)
	logger.Printf(ident+" %d\n", "TSC_SPOT_CHECKINTERVAL", Cfg.Checkinterval)
	logger.Printf(ident+" %d\n", "TSC_SPOT_CARID", Cfg.Carid)
	logger.Printf(ident+" %d\n", "TSC_SPOT_CHARGESOCLIMIT", Cfg.ChargeSocLimit)
	logger.Printf(ident+" %d\n", "TSC_SPOT_FALLBACKSOCLIMIT", Cfg.FallbackSocLimit)
	logger.Printf(ident+" %d\n", "TSC_SPOT_STARTUPDELAY", Cfg.StartupDelay)
	logger.Printf(ident+" %t\n", "TSC_SPOT_DEBUG", Cfg.Debug)
	logger.Printf(ident+" %t\n", "TSC_SPOT_DRYRUN", Cfg.DryRun)
	logger.Printf(ident+" %s\n", "TZ", Cfg.TimeZone)

}
