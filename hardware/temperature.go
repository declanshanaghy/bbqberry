package hardware

import (
	"fmt"
	"math"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
)

// TemperatureArray provides an interface to read temperature values from the physical temperature probes
type TemperatureArray interface {
	// GetTemperatureReading the tempearature from the requested probe and returns a TemperatureReading object
	GetTemperatureReading(probe int32) (*TemperatureReading, error)
	//GetAverageTemperatureReading(probe int32, samples n) (*TemperatureReading, error)
	// GetNumProbes returns the number of configured temperature probes
	GetNumProbes() int32
	// Close closes communication with the underlying hardware
	Close()
}

// TemperatureReading represents a single point temperature reading in various scales
type TemperatureReading struct {
	Probe, Analog               int32
	Time                        time.Time
	Kelvin, Celsius, Fahrenheit float32
}

type temperatureArray struct {
	numProbes int32
	bus       embd.SPIBus
	adc       *mcp3008.MCP3008
}

// NewTemperatureArray constructs a concrete implementation of
// TemperatureArray which can communicate with the underlying hardware
func NewTemperatureArray(numProbes int32, bus embd.SPIBus) TemperatureArray {
	return &temperatureArray{
		numProbes: numProbes,
		bus:       bus,
		adc:       mcp3008.New(mcp3008.SingleMode, bus),
	}
}

func (s *temperatureArray) Close() {
	log.Info("action=Close")
	s.bus.Close()
}

func (s *temperatureArray) GetNumProbes() int32 {
	return s.numProbes
}

func (s *temperatureArray) errorCheckProbeNumber(probe int32) error {
	if probe < 1 || probe > s.numProbes {
		return fmt.Errorf("Invalid probe: %d. Must be between 1 and %d", probe, s.numProbes)
	}
	return nil
}

func (s *temperatureArray) readProbe(probe int32) (int, error) {
	if err := s.errorCheckProbeNumber(probe); err != nil {
		return 0, err
	}
	v, err := s.adc.AnalogValueAt(int(probe-1))
	log.Infof("action=readProbe probe=%v v=%v", probe, v)
	return v, err
}

func (s *temperatureArray) GetAverageTemperatureReading(probe int32, samples int) (*TemperatureReading, error) {
	return nil, nil
}

func (s *temperatureArray) GetTemperatureReading(probe int32) (*TemperatureReading, error) {
	v, err := s.readProbe(probe)
	if err != nil {
		return nil, err
	}
	return newTemperatureReading(probe, v), nil
}

func newTemperatureReading(probe int32, v int) *TemperatureReading {
	k, c, f := convertVoltToTemp(v)
	return &TemperatureReading{
		Probe:      probe,
		Time:       time.Now(),
		Analog:	    int32(v),
		Kelvin:     float32(k),
		Celsius:    float32(c),
		Fahrenheit: float32(f),
	}
}

func convertVoltToTemp(volt int) (k, c, f float64) {
	// y = -400 x + 2.029000e+05
	// y := -400 * volt + 202900
	// y := -0.0025 * float64(volt) + 2029
	// log.Infof("action=convertVoltToTemp volt=%v y=%v", volt, y)

	// get the Kelvin temperature
	k = math.Log(10240000.0/float64(volt) - 10000)
	k = 1 / (0.001129148 + (0.000234125 * k) + (0.0000000876741 * k * k * k))

	// convert to Celsius and round to 1 decimal place
	c = k - 273.15

	// get the Fahrenheit temperature
	f = (c * 1.8) + 32

	// return all three temperature values
	return
}
