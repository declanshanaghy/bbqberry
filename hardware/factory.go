package hardware

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	"github.com/kidoman/embd"
	// Enable RaspberryPi features by importing the embd host definitions
	_ "github.com/kidoman/embd/host/rpi"
)

// StubBus can be set to a mock object for testing purposes
var StubBus *stubembd.StubSPIBus

// Startup initializes the hardware, should be called before first access
func Startup() {
	// no-op when stubbed
	if framework.Constants.Stub {
		return
	}

	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
}

// Shutdown de-initializes the hardware, should be called when the service is shutting down
func Shutdown() {
	// no-op when stubbed
	if framework.Constants.Stub {
		return
	}

	if err := embd.CloseSPI(); err != nil {
		panic(err)
	}
}

// NewStrandController provides an abstracted interface to the LED strands
func NewStrandController() WS2801 {
	config := framework.Constants.Hardware

	bus := newSPIBus(0)
	return newWS2801(*config.NumLedPixels, bus)
}

// NewTemperatureReader provides an abstracted interface to the temperature probes
func NewTemperatureReader() TemperatureReader {
	config := framework.Constants.Hardware

	bus := newSPIBus(1)
	return newTemperatureReader(int32(len(config.Probes)), bus)
}

func newSPIBus(channel byte) embd.SPIBus {
	if framework.Constants.Stub {
		//log.Warningf("action=NewSPIBus channel=%d STUBBED", channel)
		if StubBus == nil {
			StubBus = stubembd.NewStubSPIBus()
		}
		return StubBus
	}
	log.Infof("action=NewSPIBus channel=%d", channel)
	return embd.NewSPIBus(embd.SPIMode0, channel, 1000000, 8, 0)
}
