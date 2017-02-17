package backend

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
)

// GetTemperatureProbeReadings reads the current sensor values from the onboard temperature sensors
func GetTemperatureProbeReadings(params *temperatures.GetTemperaturesParams) (*models.TemperatureReadings, error) {
	m := models.TemperatureReadings{}
	tReader := hardware.NewTemperatureReader()

	for _, probeNumber := range getProbeIndexes(params.Probe) {
		reading := models.TemperatureReading{}
		err := tReader.GetTemperatureReading(probeNumber, &reading)
		if err != nil {
			return nil, err
		}

		m = append(m, &reading)
	}

	return &m, nil
}
