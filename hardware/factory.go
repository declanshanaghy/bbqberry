package hardware

import (
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/kidoman/embd"
)

//TODO: Abstract this entire file to work on Mac

func init() {
	HardwareConfig = HWconfig{
		NumTemperatureProbes: 3,
	}
}

type HWconfig struct {
	NumTemperatureProbes int32
}

var HardwareConfig HWconfig

func Startup() {
	//if err := embd.InitSPI(); err != nil {
	//	panic(err)
	//}
}

func Shutdown() {
	//if err := embd.CloseSPI(); err != nil {
	//	panic(err)
	//}
}

func NewStrandController() TemperatureArray {
	bus := newSPIBus(0)
	return NewBBQTempReader(HardwareConfig.NumTemperatureProbes, bus)
}

func NewTemperatureReader() TemperatureArray {
	bus := newSPIBus(1)
	return NewBBQTempReader(HardwareConfig.NumTemperatureProbes, bus)
}

func newSPIBus(channel byte) embd.SPIBus {
	log.Infof("action=NewSPIBus channel=%d", channel)
	return &MockSPIBus{}
}
