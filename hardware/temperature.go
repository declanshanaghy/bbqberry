package hardware

import (
	"fmt"
	"math"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/go-openapi/strfmt"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
)

// TemperatureArray provides an interface to read temperature values from the physical
// temperature probes
type TemperatureArray interface {
	// GetTemperatureReading reads the tempearature from the requested probe
	GetTemperatureReading(probe int32, reading *models.TemperatureReading) error
	// GetNumProbes returns the number of configured temperature probes
	GetNumProbes() int32
	// Close closes communication with the underlying hardware
	Close()
}

type temperatureArray struct {
	numProbes int32
	bus       embd.SPIBus
	adc       *mcp3008.MCP3008
}

var fakeTemps = make(map[int32]int, 0)

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

func (s *temperatureArray) readProbe(probe int32) (int32, error) {
	var v int
	var err error

	if err := s.errorCheckProbeNumber(probe); err != nil {
		return 0, err
	}
	if framework.Constants.Stub {
		v = fakeTemps[probe] + 1
		if v == 1024 {
			v = 0
		}
		fakeTemps[probe] = v
	} else {
		v, err = s.adc.AnalogValueAt(int(probe - 1))
		if err != nil {
			return 0, err
		}
	}
	log.Infof("action=readProbe probe=%v v=%v", probe, v)
	return int32(v), err
}

func (s *temperatureArray) GetTemperatureReading(probe int32, reading *models.TemperatureReading) error {
	a, err := s.readProbe(probe)
	if err != nil {
		return err
	}
	k, c, f, v, o := SteinhartHart(a)

	time := strfmt.DateTime(time.Now())
	reading.Probe = &probe
	reading.Time = &time
	reading.Analog = &a
	reading.Voltage = &v
	reading.Resistance = &o
	reading.Kelvin = &k
	reading.Celsius = &c
	reading.Fahrenheit = &f

	return nil
}

// SteinhartHart calculates temperature from the given analog value using the Steinhart Hart formula
func SteinhartHart(analog int32) (tempK float32, tempC float32, tempF float32, voltage float32, resistance int32) {
	// iBBQ probe is 100.8K at 25c

	volts := (float64(analog) * 3.3) / 1024 // calculate the voltage
	voltage = float32(volts)
	ohms := ((1 / volts) * 3300) - 1000 // calculate the resistance of the thermististor
	resistance = int32(ohms)

	lnohm := math.Log1p(ohms) // take ln(ohms)

	a := framework.Constants.SteinhartHart.A
	b := framework.Constants.SteinhartHart.B
	c := framework.Constants.SteinhartHart.C

	// Steinhart Hart Equation
	// T = 1/(a + b[ln(ohm)] + c[ln(ohm)]^3)
	t1 := (b * lnohm)     // b[ln(ohm)]
	c2 := c * lnohm       // c[ln(ohm)]
	t2 := math.Pow(c2, 3) // c[ln(ohm)]^3

	tempK = float32(1 / (a + t1 + t2)) // Calculate temperature in Kelvin
	tempC = tempK - 273.15 - 4         // K to C (the -4 is error correction for bad python math)
	tempF = tempC*9/5 + 32             // Fahrenheit

	return
}
