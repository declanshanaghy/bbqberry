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
	GetTemperatureReading(probe int32) (*models.TemperatureReading, error)
	// GetNumProbes returns the number of configured temperature probes
	GetNumProbes() int32
	// Close closes communication with the underlying hardware
	Close()
}

type temperatureReader struct {
	numProbes int32
	bus       embd.I2CBus
	adc       ADC
	readings  []*models.TemperatureReading
}

// newTemperatureReader constructs a concrete implementation of
// TemperatureReader which can communicate with the underlying hardware
func newTemperatureReader(numProbes int32, bus embd.I2CBus) TemperatureReader {
	o := temperatureReader{
		bus:       	bus,
		adc:       	NewADS1115(bus),
		numProbes:	numProbes,
		readings:	make([]*models.TemperatureReading, numProbes),
	}

	for i := int32(0); i < numProbes; i++ {
		o.readings[i] = &models.TemperatureReading{
			WarningAckd: false,
		}
	}

	return &o
}

func (o *temperatureReader) Close() {
	log.Info("action=Close")
	o.bus.Close()
}

func (o *temperatureReader) GetNumProbes() int32 {
	return o.numProbes
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
	}).Debug("Read probe")

	return int32(a), err
}

func (o *temperatureReader) GetTemperatureReading(probe int32) (*models.TemperatureReading, error) {
	analog, err := o.readProbe(probe)
	if err != nil {
		return nil, err
	}

	hwCfg := framework.Config.Hardware
	physProbe := hwCfg.Probes[probe]
	probeLimits := physProbe.Limits

	vOut := framework.ConvertAnalogToVoltage(analog)
	tempK, tempC, tempF := framework.AdafruitAD8495ThermocoupleVtoKCF(vOut)
	log.Debugf("probe=%d A=%d V=%0.5f K=%d C=%d F=%d minC=%d maxC=%d",
		probe, analog, vOut, tempK, tempC, tempF, probeLimits.MinWarnCelsius, probeLimits.MaxWarnCelsius)

	reading := o.readings[probe]

	if tempC < *probeLimits.MinWarnCelsius {
		reading.Warning = fmt.Sprintf("%d° C exceeds low temperature warning limit of %d° C",
			int32(tempC), int32(*probeLimits.MinWarnCelsius))
	}
	if tempC > *probeLimits.MaxWarnCelsius {
		reading.Warning = fmt.Sprintf("%d° C exceeds high temperature warning limit of %d° C",
			int32(tempC), int32(*probeLimits.MaxWarnCelsius))
	}

	t := strfmt.DateTime(time.Now())

	reading.Probe = &probe
	reading.Updated = &t
	reading.Analog = &analog
	reading.Voltage = &vOut
	reading.Kelvin = &tempK
	reading.Celsius = &tempC
	reading.Fahrenheit = &tempF

	return reading, nil
}
