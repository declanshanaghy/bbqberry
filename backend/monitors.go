package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
)

// GetTemperatureMonitors reads all configured temperature monitors
func GetTemperatureMonitors(params *temperature.GetMonitorsParams) (m models.TemperatureMonitors,
	err error) {
	return m, err
}
