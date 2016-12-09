package hardware

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/mocks/mock_embd"
	"github.com/kidoman/embd"
	// Enable RaspberryPi features by importing the embd host definitions
	// _ "github.com/kidoman/embd/host/rpi"
)

func init() {
	HardwareConfig = hardwareConfig{
		NumTemperatureProbes: 3,
	}
}

type hardwareConfig struct {
	NumTemperatureProbes int32
}

// HardwareConfig represents the underlying physical hardware
var HardwareConfig hardwareConfig

// Startup initializes the hardware, should be called before first access
func Startup() {
	//if err := embd.InitSPI(); err != nil {
	//	panic(err)
	//}
}

// Shutdown de-initializes the hardware, should be called when the service is shutting down
func Shutdown() {
	//if err := embd.CloseSPI(); err != nil {
	//	panic(err)
	//}
}

// NewStrandController provides an abstracted interface to the LED strands
func NewStrandController() TemperatureArray {
	bus := newSPIBus(0)
	return NewTemperatureArray(HardwareConfig.NumTemperatureProbes, bus)
}

// NewTemperatureReader provides an abstracted interface to the temperature probes
func NewTemperatureReader() TemperatureArray {
	bus := newSPIBus(1)
	return NewTemperatureArray(HardwareConfig.NumTemperatureProbes, bus)
}

func newSPIBus(channel byte) embd.SPIBus {
	log.Infof("action=NewSPIBus channel=%d", channel)
	return &mock_embd.MockSPIBus{}
	// return embd.NewSPIBus(embd.SPIMode0, channel, 1000000, 8, 0)
}
