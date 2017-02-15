package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
)

// CreateMonitor reads all configured temperature monitors
func CreateMonitor(params *monitors.CreateMonitorParams) (m models.TemperatureMonitors,
		err error) {
	return m, err
}

// GetMonitors reads all configured temperature monitors
func GetMonitors(params *monitors.GetMonitorsParams) (m models.TemperatureMonitors,
	err error) {
	return m, err
}
