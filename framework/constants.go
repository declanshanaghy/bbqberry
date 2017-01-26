package framework

import (
	"os"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"github.com/Polarishq/middleware/framework/log"
)

const vcc = 3.3
const analogMax = 1024
const r2 = 1000.0

// Absolute temperature limits
const tempLimitLowCelsius = -50
const tempLimitHighCelsius = 400

// Warn if temperature gets within this threshold of absolute limits
const tempWarnThreshold = 0.1

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

	Constants = constants{
		ServiceName: "bbqberry",
		Version:     "v1",
		Stub:        stub,
		Hardware: hardwareConfig{
			NumLEDPixels:           26,
			NumTemperatureProbes:   2,
			AmbientProbeNumber:     1,
			VCC:                    vcc,
			VDivR2:                 r2,
			AnalogVoltsPerUnit:     vcc / analogMax,
			MinTempWarnCelsius:     tempLimitLowCelsius - (tempLimitLowCelsius * tempWarnThreshold),
			MaxTempWarnCelsius:     tempLimitHighCelsius - (tempLimitHighCelsius * tempWarnThreshold),
		},
	}
}

// hardwareConfig represents the underlying physical hardware
type hardwareConfig struct {
	NumLEDPixels            int
	NumTemperatureProbes    int32
	AmbientProbeNumber      int32
	VCC, VDivR2, AnalogVoltsPerUnit             float32
	MinTempWarnCelsius, MaxTempWarnCelsius      float32
}

type constants struct {
	ServiceName string
	Version     string
	Stub        bool
	Hardware    hardwareConfig
}

// Constants contains static information about the running service
var Constants constants
