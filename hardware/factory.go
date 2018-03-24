package hardware

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	"github.com/kidoman/embd"
	// Enable RaspberryPi features by importing the embd host definitions
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/declanshanaghy/bbqberry/hardware/ads1x15"
	"github.com/kidoman/embd/convertors/mcp3008"
)

// StubSPIBus can be set to a mock object for testing purposes
var StubSPIBus *stubembd.StubSPIBus

// StubSPIBus can be set to a mock object for testing purposes
var StubI2CBus *stubembd.StubI2CBus

// ADC provdes an interface for communicating with an ADS1x15 analog to digital converter chip
type ADC interface {
	AnalogValueAt(chanNum int) (int, error)
}

// NewADS1115 creates an abstracted ADC based on the ADS1115 I2C chip
func NewADS1115(bus embd.I2CBus) ADC {
	return ads1x15.NewADS1115(ads1x15.ADS1x15_DEFAULT_ADDRESS, bus)
}

// NewADS1115 creates an abstracted ADC based on the MCP3008 SPI chip
func NewMCP3008(bus embd.SPIBus) ADC {
	return mcp3008.New(mcp3008.SingleMode, bus)
}

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

	//bus := newSPIBus(1)
	bus := newI2CBus(1)
	return newTemperatureReader(int32(len(config.Probes)), bus)
}

func newSPIBus(channel byte) embd.SPIBus {
	if framework.Constants.Stub {
		//log.Warningf("action=NewSPIBus channel=%d STUBBED", channel)
		if StubSPIBus == nil {
			StubSPIBus = stubembd.NewStubSPIBus(channel)
		}
		return StubSPIBus
	}
	log.Debugf("action=NewSPIBus channel=%d", channel)
	return embd.NewSPIBus(embd.SPIMode0, channel, 1000000, 8, 0)
}

func newI2CBus(address byte) embd.I2CBus {
	if framework.Constants.Stub {
		//log.Warningf("action=NewSPIBus channel=%d STUBBED", channel)
		if StubI2CBus == nil {
			StubI2CBus = stubembd.NewStubI2CBus()
		}
		return StubI2CBus
	}
	log.Debugf("action=newI2CBus address=%d", address)
	return embd.NewI2CBus(address)
}
