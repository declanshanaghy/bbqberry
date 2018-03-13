package backend

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
)

// GetTemperatureProbeReadings reads the current sensor values from the onboard temperature sensors
func GetTemperatureProbeReadings(params *temperatures.GetTemperaturesParams) ([]*models.TemperatureReading, error) {
	m := make([]*models.TemperatureReading, 0)
	tReader := hardware.NewTemperatureReader()

	for _, probeNumber := range getProbeIndexes(params.Probe) {
		reading := models.TemperatureReading{}
		err := tReader.GetTemperatureReading(probeNumber, &reading)
		if err != nil {
			return nil, err
		}

		m = append(m, &reading)
	}

	return m, nil
}
