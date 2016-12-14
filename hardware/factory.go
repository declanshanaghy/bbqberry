package hardware

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	"github.com/kidoman/embd"
	// Enable RaspberryPi features by importing the embd host definitions
	_ "github.com/kidoman/embd/host/rpi"
)

// StubBus can be set to a mock object for testing purposes
var StubBus *stubembd.StubSPIBus

func init() {
	HardwareConfig = hardwareConfig{
		NumLEDPixels:         18,
		NumTemperatureProbes: 3,
	}
}

type hardwareConfig struct {
	NumLEDPixels         int
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
func NewStrandController() ws2801.WS2801 {
	bus := newSPIBus(0)
	return ws2801.NewWS2801(HardwareConfig.NumLEDPixels, bus)
}

// NewTemperatureReader provides an abstracted interface to the temperature probes
func NewTemperatureReader() TemperatureArray {
	bus := newSPIBus(1)
	return NewTemperatureArray(HardwareConfig.NumTemperatureProbes, bus)
}

func newSPIBus(channel byte) embd.SPIBus {
	if framework.Constants.Stub {
		log.Warningf("action=NewSPIBus channel=%d STUBBED", channel)
		return StubBus
	}
	log.Infof("action=NewSPIBus channel=%d", channel)
	return embd.NewSPIBus(embd.SPIMode0, channel, 1000000, 8, 0)
}
