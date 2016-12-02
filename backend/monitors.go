package backend

import (
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	"github.com/declanshanaghy/bbqberry/models"
)

func GetTemperatureMonitors(params *temperature.GetMonitorsParams) (m models.TemperatureMonitors,
	err error) {
	return m, err
}
