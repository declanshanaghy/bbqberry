package framework

import (
	"os"
)

func init() {
	stub := false
	if os.Getenv("STUB") != "" {
		stub = true
	}

	Constants = constants{
		ServiceName: "bbqberry",
		Version:     "v1",
		Stub:        stub,
		Hardware: hardwareConfig{
			NumLEDPixels:         18,
			NumTemperatureProbes: 3,
		},
	}
}

func init() {
}

type hardwareConfig struct {
	NumLEDPixels         int
	NumTemperatureProbes int32
}

// HardwareConfig represents the underlying physical hardware

type constants struct {
	ServiceName string
	Version     string
	Stub        bool
	Hardware    hardwareConfig
}

// Constants contains static information about the running service
var Constants constants
