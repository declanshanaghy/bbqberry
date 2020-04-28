package backend

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
	"github.com/declanshanaghy/bbqberry/framework"
)

// GetTemperatureProbeReadings reads the current sensor values from the onboard temperature sensors
func GetTemperatureProbeReadings(params *temperatures.GetTemperaturesParams) ([]*models.TemperatureReading, error) {
	readings := make([]*models.TemperatureReading, 0)
	tReader := hardware.NewTemperatureReader()

	for _, probeNumber := range getProbeIndexes(params.Probe) {
		probe := framework.Config.Hardware.Probes[probeNumber]

		if *probe.Enabled {
			reading, err := tReader.GetTemperatureReading(probeNumber)
			if err != nil {
				return nil, err
			}
			readings = append(readings, reading)
		} else {
			readings = append(readings, &models.TemperatureReading{Probe:&probeNumber})
		}
	}

	return readings, nil
}
