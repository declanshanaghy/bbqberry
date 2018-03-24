package hardware

import (
	"fmt"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/go-openapi/strfmt"
	"github.com/kidoman/embd"
)

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
	//bus       embd.SPIBus
	bus       embd.I2CBus
	adc       ADC
}

// newTemperatureReader constructs a concrete implementation of
// TemperatureReader which can communicate with the underlying hardware
func newTemperatureReader(numProbes int32, bus embd.I2CBus) TemperatureReader {
	return &temperatureReader{
		numProbes: numProbes,
		bus:       bus,
		//adc:       NewMCP3008(bus),
		adc:       NewADS1115(bus),
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

	iv, err := o.adc.AnalogValueAt(int(probe))
	a = int32(iv)
	if err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{
		"probe": probe,
		"a": a,
	}).Infof("Read probe")

	return int32(a), err
}

func (o *temperatureReader) GetTemperatureReading(probe int32, reading *models.TemperatureReading) error {
	analog, err := o.readProbe(probe)
	if err != nil {
		return err
	}

	hwCfg := framework.Constants.Hardware
	vOut := framework.ConvertAnalogToVoltage(analog)

	physProbe := hwCfg.Probes[probe]
	probeLimits := physProbe.Limits

	tempK, tempC, tempF := framework.AdafruitAD8495ThermocoupleVtoKCF(vOut)
	log.Infof("probe=%d A=%d V=%0.5f K=%d C=%d F=%d minC=%d maxC=%d",
		probe, analog, vOut, tempK, tempC, tempF, probeLimits.MinWarnCelsius, probeLimits.MaxWarnCelsius)

	if tempC < *probeLimits.MinWarnCelsius {
		_, f := framework.ConvertCToKF(float32(*probeLimits.MinWarnCelsius))
		reading.Warning = fmt.Sprintf("%d °F exceeds low temperature limit of %d °F",
			int32(tempF), int32(f))
	}
	if tempC > *probeLimits.MaxWarnCelsius {
		_, f := framework.ConvertCToKF(float32(*probeLimits.MaxWarnCelsius))
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
