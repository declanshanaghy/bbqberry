package sensors

import (
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
	"math"
	"github.com/Polarishq/middleware/framework/log"
)

type TemperatureArray interface {
	/*
	Reads the tempearature from the requested probe and returns the value in Kelvin, Celcius & Fahrenheit
	 */
	GetTemp(probe int) TemperatureReading
}

type TemperatureReading struct {
	K, C, F float64
}

type BBQTemp struct {
	channel byte
	bus embd.SPIBus
	adc *mcp3008.MCP3008
}

func NewTemperature(bus embd.SPIBus) TemperatureArray {
	return &BBQTemp{
		bus: bus,
		adc: mcp3008.New(mcp3008.SingleMode, bus),
	}
}

func (s *BBQTemp) Close() {
	log.Info("action=Close")
	s.bus.Close()
}

func (s *BBQTemp) GetTemp(probe int) TemperatureReading {
	log.Debugf("action=GetTemp probe=%d", probe)
	v, err := s.adc.AnalogValueAt(probe)
	if err != nil {
		panic(err)
	}
	return convertVoltToTemp(v)
}

func convertVoltToTemp(volt int) TemperatureReading {
	// get the Kelvin temperature
	k := math.Log(10240000.0 / float64(volt) - 10000);
	k = 1 / (0.001129148 + (0.000234125 * k) + (0.0000000876741 * k * k * k));

	// convert to Celsius and round to 1 decimal place
	c := k - 273.15;

	// get the Fahrenheit temperature
	f := (c * 1.8) + 32;

	// return all three temperature values
	return TemperatureReading{K: k, C: c, F: f}
}
