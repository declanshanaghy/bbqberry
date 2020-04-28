package backend

import (
	"github.com/declanshanaghy/bbqberry/restapi/operations/alerts"
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
	return false, nil
}
