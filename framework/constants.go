package framework

import (
	"os"

	"github.com/declanshanaghy/bbqberry/models"
)

// DefaultDB is the default database name that should be used if an override is not provided
const DefaultDB = "bbqberry"

var Enabled = true
var Disabled = false

const HUE_KEY = "5EPXHOGHzm7TGha3IFumdF2bTLdcwuae-21iQguC"
const HUE_ALERT_GROUP = "Patio"
const AWS_DEFAULT_REGION = "us-east-1"

func init() {
	stub := false
	if os.Getenv("STUB") != "" {
		stub = true
	}
	/******************************** BEGIN PSEUDO CONSTANTS *********************************************/
	// Electrical Constants
	vcc := float32(3.3)
	analogMax := int32(26453)

	// Accessories
	nPixels := int32(25)

	// Absolute temperature limits
	tempLimitAbsAmbientLowCelsius := int32(-50.0)
	tempLimitAbsAmbientHighCelsius := int32(400.0)
	tempLimitAbsCookingLowCelsius := -50.0
	tempLimitAbsCookingHighCelsius := 250.0

	// Warn if temperature gets within this threshold of absolute limits
	tempWarnThreshold := 0.2
	/********************************* END PSEUDO CONSTANTS **********************************************/

	minTempWarnAmbCelsius := int32(95)
	maxTempWarnAmbCelsius := int32(205)

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

	chamber 	:= "Chamber"
	chamberAlt 	:= "Chamber Alt"
	food 		:= "Food"
	foodAlt 	:= "Food Alt"
	probes 		:= []*models.TemperatureProbe{
		{
			Label:  	&chamberAlt,
			Enabled:	&Disabled,
			Limits: 	&ambient,
		},
		{
			Label:  	&food,
			Enabled:	&Disabled,
			Limits: 	&cooking,
		},
		{
			Label:  	&foodAlt,
			Enabled:	&Disabled,
			Limits: 	&cooking,
		},
		{
			Label:  	&chamber,
			Enabled:	&Enabled,
			Limits: 	&ambient,
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