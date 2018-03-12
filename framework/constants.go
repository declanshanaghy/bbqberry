package framework

import (
	"os"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"github.com/declanshanaghy/bbqberry/models"
)

func init() {
	stub := false
	if os.Getenv("STUB") != "" {
		stub = true
		if os.Getenv("INFLUXDB") == "" {
			log.Warningf("Hardware is stubbed, resetting influx database")
			os.Setenv("INFLUXDB", "stub")
			influxdb.LoadConfig()
		}
	}
	/******************************** BEGIN PSEUDO CONSTANTS *********************************************/
	/**/
	// Electrical Constants
	vcc := float32(3.3)
	analogMax := int32(1024)

	// Accessories
	nPixels := int32(25)

	// Absolute temperature limits
	tempLimitAbsAmbientLowCelsius := int32(-50.0)
	tempLimitAbsAmbientHighCelsius := int32(400.0)
	tempLimitAbsCookingLowCelsius := -50.0
	tempLimitAbsCookingHighCelsius := 250.0

	// Warn if temperature gets within this threshold of absolute limits
	tempWarnThreshold := 0.2
	/**/
	/********************************* END PSEUDO CONSTANTS **********************************************/

	minTempWarnAmbCelsius := int32(float64(tempLimitAbsAmbientLowCelsius) -
		(float64(tempLimitAbsAmbientLowCelsius) * tempWarnThreshold))
	maxTempWarnAmbCelsius := int32(float64(tempLimitAbsAmbientHighCelsius) -
		(float64(tempLimitAbsAmbientHighCelsius) * tempWarnThreshold))

	minTempWarnCookingCelsius := int32(tempLimitAbsCookingLowCelsius - (tempLimitAbsCookingLowCelsius * tempWarnThreshold))
	maxTempWarnCookingCelsius := int32(tempLimitAbsCookingHighCelsius - (tempLimitAbsCookingHighCelsius * tempWarnThreshold))

	sAmb := "ambient"
	ambient := models.TemperatureLimits{
		ProbeType:      &sAmb,
		MinWarnCelsius: &minTempWarnAmbCelsius,
		MaxWarnCelsius: &maxTempWarnAmbCelsius,
		MinAbsCelsius:  &tempLimitAbsAmbientLowCelsius,
		MaxAbsCelsius:  &tempLimitAbsAmbientHighCelsius,
	}

	sCook := "cooking"
	tempLimitAbsCookingLowCelsiusI32 := int32(tempLimitAbsCookingLowCelsius)
	tempLimitAbsCookingHighCelsiusI32 := int32(tempLimitAbsCookingHighCelsius)
	cooking := models.TemperatureLimits{
		ProbeType:      &sCook,
		MinWarnCelsius: &minTempWarnCookingCelsius,
		MaxWarnCelsius: &maxTempWarnCookingCelsius,
		MinAbsCelsius:  &tempLimitAbsCookingLowCelsiusI32,
		MaxAbsCelsius:  &tempLimitAbsCookingHighCelsiusI32,
	}

	chamber := "Chamber"
	pa := "Probe A"
	pb := "Probe B"
	probes := models.TemperatureSettings{
		&models.TemperatureSetting{
			Label:      &chamber,
			TempLimits: &ambient,
		},
		&models.TemperatureSetting{
			Label:      &pa,
			TempLimits: &cooking,
		},
		&models.TemperatureSetting{
			Label:      &pb,
			TempLimits: &cooking,
		},
	}

	hwCfg := models.HardwareConfig{
		NumLedPixels: &nPixels,
		Vcc:          &vcc,
		AnalogMax:    &analogMax,
		Probes:       probes,
	}

	Constants = constants{
		ServiceName: "bbqberry",
		Version:     "v1",
		Stub:        stub,
		Hardware:    hwCfg,
	}
}

type constants struct {
	ServiceName string
	Version     string
	Stub        bool
	Hardware    models.HardwareConfig
}

// Constants contains static information about the running service
var Constants constants
