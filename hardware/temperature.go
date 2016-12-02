package hardware

import (
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
	"math"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"time"
	"fmt"
	"errors"
)

type TemperatureArray interface {
	/*
	Reads the tempearature from the requested probe and returns the value in Kelvin, Celcius & Fahrenheit
	 */
	GetTemperatureReading(probe int32) (*TemperatureReading, error)
	GetNumProbes() int32
	Close()
}

type TemperatureReading struct {
	Probe int32
	Time time.Time
	Kelvin, Celcius, Fahrenheit float32
}

type BBQTemp struct {
	numProbes int32
	bus embd.SPIBus
	adc *mcp3008.MCP3008
}

func NewBBQTempReader(numProbes int32, bus embd.SPIBus) TemperatureArray {
	return &BBQTemp{
		numProbes: numProbes,
		bus: bus,
		adc: mcp3008.New(mcp3008.SingleMode, bus),
	}
}

func (s *BBQTemp) Close() {
	log.Info("action=Close")
	s.bus.Close()
}

func (s *BBQTemp) GetNumProbes() int32 {
	return s.numProbes
}

func (s *BBQTemp) GetTemperatureReading(probe int32) (*TemperatureReading, error) {
	if probe < 1 || probe > s.numProbes {
		return nil, errors.New(fmt.Sprintf("Invalid probe: %d. Must be between 1 and %d", probe, s.numProbes))
	}
	
	log.Debugf("action=GetTemp probe=%d", probe)
	v, err := s.adc.AnalogValueAt(int(probe))
	if err != nil {
		panic(err)
	}
	k, c, f := convertVoltToTemp(v)
	return &TemperatureReading{
		Probe: probe,
		Time: time.Now(),
		Kelvin: float32(k),
		Celcius: float32(c),
		Fahrenheit: float32(f),
	}, nil
}

func convertVoltToTemp(volt int) (k, c, f float64) {
	// get the Kelvin temperature
	k = math.Log(10240000.0 / float64(volt) - 10000);
	k = 1 / (0.001129148 + (0.000234125 * k) + (0.0000000876741 * k * k * k));

	// convert to Celsius and round to 1 decimal place
	c = k - 273.15;

	// get the Fahrenheit temperature
	f = (c * 1.8) + 32;

	// return all three temperature values
	return
}
