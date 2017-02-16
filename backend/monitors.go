package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	"github.com/declanshanaghy/bbqberry/db/mongodb"
)

const collection = "monitors"

// CreateMonitor creates a new temperature monitor
func CreateMonitor(params *monitors.CreateMonitorParams) (*models.TemperatureMonitor, error) {
	session, db, err := mongodb.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := db.C(collection)
	err = c.Insert(&params)
	if err != nil {
		return nil, err
	}

	ID := "bollox"
	m := models.TemperatureMonitor{
		Probe: &params.Probe,
		ID: &ID,
	}

	return &m, nil
}

// GetMonitors reads all configured temperature monitors
func GetMonitors(params *monitors.GetMonitorsParams) (*models.TemperatureMonitors, error) {
	return nil, nil
}
