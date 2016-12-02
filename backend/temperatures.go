package backend

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	"github.com/go-openapi/strfmt"
)

// GetTemperatureProbeReadings reads the current sensor values from the onboard temperature sensors
func GetTemperatureProbeReadings(params *temperature.GetProbeReadingsParams) (m models.TemperatureReadings,
	err error) {

	tReader := hardware.NewTemperatureReader()

	var probes []int32
	var i int32

	if *params.Probe == 0 {
		probes = make([]int32, tReader.GetNumProbes())
		for i = 0; i < tReader.GetNumProbes(); i++ {
			probes[i] = i + 1
		}
	} else {
		probes = make([]int32, 1)
		probes[0] = *params.Probe
	}

	for i := range probes {
		probeNumber := int32(probes[i])
		reading, err := tReader.GetTemperatureReading(probeNumber)
		if err != nil {
			return nil, err
		}
		
		t := strfmt.DateTime(reading.Time)
		z := models.TemperatureReading{
			Time:       &t,
			Probe:      &reading.Probe,
			Kelvin:     &reading.Kelvin,
			Celsius:    &reading.Celsius,
			Fahrenheit: &reading.Fahrenheit,
		}
		m = append(m, &z)
	}

	return m, err
}
