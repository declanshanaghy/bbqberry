package hardware

import (
	"fmt"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/go-openapi/strfmt"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
)

// FakeTemps can be set to return specific analog readings during tests
var FakeTemps = make(map[int32]int32, 0)

func init() {
	hardware := framework.Constants.Hardware

	if framework.Constants.Stub {
		for i := int32(1); i <= hardware.NumTemperatureProbes; i++ {
			FakeTemps[i] = i
		}
	}
}

// TemperatureReader provides an interface to read temperature values from the physical temperature probes
type TemperatureReader interface {
	// GetTemperatureReading reads the tempearature from the requested probe
	GetTemperatureReading(probe int32, reading *models.TemperatureReading) error
	// GetNumProbes returns the number of configured temperature probes
	GetNumProbes() int32
	// Close closes communication with the underlying hardware
	Close()
}

type temperatureReader struct {
	numProbes int32
	bus       embd.SPIBus
	adc       *mcp3008.MCP3008
}

// newTemperatureReader constructs a concrete implementation of
// TemperatureReader which can communicate with the underlying hardware
func newTemperatureReader(numProbes int32, bus embd.SPIBus) TemperatureReader {
	return &temperatureReader{
		numProbes: numProbes,
		bus:       bus,
		adc:       mcp3008.New(mcp3008.SingleMode, bus),
	}
}

func (s *temperatureReader) Close() {
	log.Info("action=Close")
	s.bus.Close()
}

func (s *temperatureReader) GetNumProbes() int32 {
	return s.numProbes
}

func (s *temperatureReader) errorCheckProbeNumber(probe int32) error {
	if probe < 1 || probe > s.numProbes {
		return fmt.Errorf("Invalid probe: %d. Must be between 1 and %d", probe, s.numProbes)
	}
	return nil
}

func (s *temperatureReader) readProbe(probe int32) (v int32, err error) {
	if err := s.errorCheckProbeNumber(probe); err != nil {
		return 0, err
	}
	if framework.Constants.Stub {
		v = FakeTemps[probe] + 1
		if v == 1024 {
			v = 0
		}
		FakeTemps[probe] = v
	} else {
		iv, err := s.adc.AnalogValueAt(int(probe - 1))
		v = int32(iv)
		if err != nil {
			return 0, err
		}
	}
	log.Debugf("action=readProbe probe=%v v=%v", probe, v)
	return int32(v), err
}

func (s *temperatureReader) GetTemperatureReading(probe int32, reading *models.TemperatureReading) error {
	analog, err := s.readProbe(probe)
	if err != nil {
		return err
	}

	/*
		Voltage divider configuration
			vcc	(3.3v)
		    /\
			|
			r1	(Temp Sensor)
			|
			|------> vOut
			|
			r2	(1k)
			|
		   ---
			gnd	(0v)
	*/
	hwCfg := framework.Constants.Hardware
	vOut := float32(analog) * hwCfg.AnalogVoltsPerUnit
	r1 := int32(((hwCfg.VCC * hwCfg.VDivR2) / vOut) - hwCfg.VDivR2)

	tempK, tempC, tempF := adafruitAD8495ThermocoupleVtoKCF(vOut)
	log.Debugf("probe=%d, A=%d, R=%d, V=%0.5f, K=%0.5f, C=%0.5f, F=%0.5f", probe, analog, r1, vOut, tempK, tempC, tempF)
	
	if tempC < hwCfg.MinTempWarnCelsius {
		reading.Warning = fmt.Sprintf("Low temperature limit exceeded: actual=%0.2f °C < threshold=%0.2f °C",
			tempC, hwCfg.MinTempWarnCelsius)
	}
	if tempC > hwCfg.MaxTempWarnCelsius {
		reading.Warning = fmt.Sprintf("High temperature limit exceeded: actual=%0.2f °C > threshold=%0.2f °C",
			tempC, hwCfg.MaxTempWarnCelsius)
	}
	
	t := strfmt.DateTime(time.Now())
	reading.Probe = &probe
	reading.DateTime = &t
	reading.Analog = &analog
	reading.Voltage = &vOut
	reading.Resistance = &r1
	reading.Kelvin = &tempK
	reading.Celsius = &tempC
	reading.Fahrenheit = &tempF

	return nil
}

// adafruitAD8495ThermocoupleVtoKCF converts the voltage read from the Adafruit Thermocouple breakout board
// to temperatures in Kelvin, Celsius and Fahrenheit
func adafruitAD8495ThermocoupleVtoKCF(v float32) (tempK float32, tempC float32, tempF float32) {
	// https://www.adafruit.com/product/1778
	// Analog Output K-Type Thermocouple Amplifier - AD8495 Breakout
	// PRODUCT ID: 1778
	// Temperature = (Vout - 1.25) / 0.005 V
	// e.g:
	// v = 1.5VDC
	// The temperature is (1.5 - 1.25) / 0.005 = 50°C

	tempC = (v - 1.25) / 0.005
	tempK, tempF = convertCToKF(tempC)
	return
}

// convertCToKF converts a celsius temperature to kelvin and fahrenheit
func convertCToKF(tempC float32) (tempK float32, tempF float32) {
	tempK = tempC + 273.15 // C to K
	tempF = tempC*1.8 + 32 // C to F
	return
}
