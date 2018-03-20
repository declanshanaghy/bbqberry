package hardware

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/util"
	"github.com/go-openapi/strfmt"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
)

// FakeTemps can be o.t to return o.ecific analog readings during tests
var FakeTemps = make(map[int32]int32, 0)

func init() {
	hwCfg := framework.Constants.Hardware

	if framework.Constants.Stub {
		nProbes := int32(len(hwCfg.Probes))
		for probe := int32(0); probe < nProbes; probe++ {
			limit := framework.Constants.Hardware.Probes[probe].Limits
			min := ConvertCelsiusToAnalog(-15)
			max := ConvertCelsiusToAnalog(*limit.MaxAbsCelsius)
			a := int32(rand.Intn(int(min)) + int(max-min))
			log.Infof("probe %d init to %d", probe, a)
			FakeTemps[probe] = a
		}
	}
}

// TemperatureReader provides an interface to read temperature values from the physical temperature probes
type TemperatureReader interface {
	// GetTemperatureReading reads the tempearature from the requested probe
	GetTemperatureReading(probe int32, reading *models.TemperatureReading) error
	// GetNumProbes returns the number of configured temperature probes
	GetNumProbes() int32
	// GetEnabledPobeIndices returns the indices of all enabled probes
	GetEnabledPobes() *[]int32
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

func (o *temperatureReader) Close() {
	log.Info("action=Close")
 o.bus.Close()
}

func (o *temperatureReader) GetNumProbes() int32 {
	return o.numProbes
}

func (o *temperatureReader) GetEnabledPobes() *[]int32 {
	enabled := make([]int32, 0)

	for probe := int32(0); probe < o.numProbes; probe++ {
		if *framework.Constants.Hardware.Probes[probe].Enabled {
			enabled = append(enabled, probe)
		}
	}
	
	return &enabled
}

func (o *temperatureReader) errorCheckProbeNumber(probe int32) error {
	if probe < 0 || probe > o.numProbes-1 {
		return fmt.Errorf("invalid probe: %d. Must be between 1 and %d", probe, o.numProbes)
	}
	return nil
}

func (o *temperatureReader) readProbe(probe int32) (a int32, err error) {
	if err := o.errorCheckProbeNumber(probe); err != nil {
		return 0, err
	}
	if framework.Constants.Stub {
		limit := framework.Constants.Hardware.Probes[probe].Limits
		min := ConvertCelsiusToAnalog(-15)
		max := ConvertCelsiusToAnalog(*limit.MaxAbsCelsius)
		a = FakeTemps[probe]

		if a >= max {
			a = min
		}
		FakeTemps[probe] = a + int32(rand.Intn(10))
	} else {
		iv, err := o.adc.AnalogValueAt(int(probe))
		a = int32(iv)
		if err != nil {
			return 0, err
		}
	}
	log.Debugf("action=readProbe probe=%d a=%d", probe, a)
	return int32(a), err
}

func (o *temperatureReader) GetTemperatureReading(probe int32, reading *models.TemperatureReading) error {
	analog, err := o.readProbe(probe)
	if err != nil {
		return err
	}

	hwCfg := framework.Constants.Hardware
	vOut := ConvertAnalogToVoltage(analog)

	physProbe := hwCfg.Probes[probe]
	probeLimits := physProbe.Limits

	tempK, tempC, tempF := adafruitAD8495ThermocoupleVtoKCF(vOut)
	log.Debugf("probe=%d A=%d V=%0.5f K=%d C=%d F=%d minC=%d maxC=%d",
		probe, analog, vOut, tempK, tempC, tempF, probeLimits.MinWarnCelsius, probeLimits.MaxWarnCelsius)

	if tempC < *probeLimits.MinWarnCelsius {
		_, f := ConvertCToKF(float32(*probeLimits.MinWarnCelsius))
		reading.Warning = fmt.Sprintf("%d °F exceeds low temperature limit of %d °F",
			int32(tempF), int32(f))
	}
	if tempC > *probeLimits.MaxWarnCelsius {
		_, f := ConvertCToKF(float32(*probeLimits.MaxWarnCelsius))
		reading.Warning = fmt.Sprintf("%d °F exceeds high temperature limit of %d °F",
			int32(tempF), int32(f))
	}

	t := strfmt.DateTime(time.Now())
	reading.Probe = &probe
	reading.DateTime = &t
	reading.Analog = &analog
	reading.Voltage = &vOut
	reading.Kelvin = &tempK
	reading.Celsius = &tempC
	reading.Fahrenheit = &tempF

	return nil
}

// adafruitAD8495ThermocoupleVtoKCF converts the voltage read from the Adafruit Thermocouple breakout board
// to temperatures in Kelvin, Celsius and Fahrenheit
func adafruitAD8495ThermocoupleVtoKCF(v float32) (tempK int32, tempC int32, tempF int32) {
	// https://www.adafruit.com/product/1778
	// Analog Output K-Type Thermocouple Amplifier - AD8495 Breakout
	// PRODUCT ID: 1778
	// Temperature = (Vout - 1.25) / 0.005 V
	// e.g:
	// v = 1.5VDC
	// The temperature is (1.5 - 1.25) / 0.005 = 50°C

	fTempC := (v - 1.25) / 0.005
	fTempK, fTempF := ConvertCToKF(fTempC)
	tempC = util.RoundFloat32ToInt32(fTempC)
	tempF = util.RoundFloat32ToInt32(fTempF)
	tempK = util.RoundFloat32ToInt32(fTempK)
	return
}

// ConvertVoltageToTemperature converts the given voltage value to its corresponding temperature values
//func ConvertVoltageToTemperature(v float32) (tempK int32, tempC int32, tempF int32) {
//	return adafruitAD8495ThermocoupleVtoKCF(v)
//}

// ConvertAnalogToVoltage converts an analog reading to its corresponding voltage value
func ConvertAnalogToVoltage(analog int32) float32 {
	hwCfg := framework.Constants.Hardware
	vcc := *hwCfg.Vcc
	// volts per analog unit = VCC / Analog max
	amax := float32(*hwCfg.AnalogMax)
	avpu := vcc / amax
	return float32(analog) * avpu
}

// ConvertVoltageToAnalog converts the given voltage to its corresponding analog value
func ConvertVoltageToAnalog(v float32) (a int32) {
	hwCfg := framework.Constants.Hardware
	vcc := *hwCfg.Vcc
	amax := float32(*hwCfg.AnalogMax)
	// volts per analog unit = VCC / Analog max
	avpu := vcc / amax
	// Therefore:
	// 	analog = volts / avpu
	a = util.RoundFloat32ToInt32(v / avpu)
	return
}

// ConvertCelsiusToVoltage converts a celsius temperature to its corresponding voltage
func ConvertCelsiusToVoltage(c int32) (v float32) {
	// According to adafruitAD8495ThermocoupleVtoKCF
	// 	c = (v - 1.25 ) / 0.005
	// Therefore:
	v = float32(c)*0.005 + 1.25
	return
}

// ConvertCelsiusToAnalog converts the given celsius to its corresponding analog value
func ConvertCelsiusToAnalog(c int32) (a int32) {
	v := ConvertCelsiusToVoltage(c)
	return ConvertVoltageToAnalog(v)
}

// ConvertAnalogToCF converts the given analog value to its corresponding celsius & fahrenheit value
func ConvertAnalogToCF(a int32) (int32, int32) {
	v := ConvertAnalogToVoltage(a)
	_, c, f := adafruitAD8495ThermocoupleVtoKCF(v)
	return c, f
}

// ConvertCToKF converts a celsius temperature to kelvin and fahrenheit
func ConvertCToKF(tempC float32) (tempK, tempF float32) {
	tempK = tempC + 273.15 // C to K
	tempF = tempC*1.8 + 32 // C to F
	return
}

// ConvertCToKFInt32 converts a celsius temperature to kelvin and fahrenheit
func ConvertCToKFInt32(tempC float32) (tempK, tempF int32) {
	k, f := ConvertCToKF(tempC)
	return int32(k), int32(f)
}
