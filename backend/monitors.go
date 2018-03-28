package backend

import (
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	"fmt"
	//"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework"
)

// MonitorsManager provides methods to operate on monitors within the backend database
type MonitorsManager struct {
}

// NewMonitorsManager creates a manager which can operate on temperature monitors
func NewMonitorsManager() (*MonitorsManager) {
	return &MonitorsManager{
	}
}

// UpdateMonitor updates the temperature monitor settings for a given probe
func (m *MonitorsManager) UpdateMonitor(params *monitors.UpdateMonitorParams) (bool, error) {
	min := *params.Monitor.Min
	max := *params.Monitor.Max

	if min > max {
		return false, fmt.Errorf("Minimum %d cannot be greater then Maximum %d", min, max)
	}

	probe := framework.Config.Hardware.Probes[*params.Monitor.Probe]

	if (params.Monitor.Label != "" ) {
		probe.Label = &params.Monitor.Label
	}

	probe.Limits.MinWarnCelsius = &min
	probe.Limits.MaxWarnCelsius = &max

	framework.Config.Save()
	return true, nil
}
