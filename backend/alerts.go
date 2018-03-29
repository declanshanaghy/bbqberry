package backend

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations/alerts"
	"github.com/Polarishq/middleware/framework/log"
)

// AlertsManager provides methods to operate on monitors within the backend database
type AlertsManager struct {
}

// NewAlertsManager creates a manager which can operate on temperature monitors
func NewAlertsManager() (*AlertsManager) {
	return &AlertsManager{
	}
}

// ClearAlert clears the warning alert from a probe
func (m *AlertsManager) ClearAlert(params *alerts.UpdateAlertParams) (bool, error) {
	t := hardware.NewTemperatureReader()
	reading, err := t.GetTemperatureReading(int32(params.Probe))
	if err != nil {
		return false, err
	}

	reading.WarningAckd = true

	log.Infof("ClearAlert %s", *reading)

	return true, nil
}
