package backend

import (
	"fmt"

	"github.com/declanshanaghy/bbqberry/db/mongodb"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	"gopkg.in/mgo.v2/bson"
)

const collection = "monitors"

// CreateMonitor creates a new temperature monitor
func CreateMonitor(params *monitors.CreateMonitorParams) (*models.TemperatureMonitor, error) {
	session, db, err := mongodb.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	defer db.Logout()

	c := db.C(collection)

	result := new(models.TemperatureMonitor)

	err = c.Find(bson.M{"probe": params.Monitor.Probe}).One(result)
	if err == nil && result.Probe != nil {
		return nil, fmt.Errorf("Monitor already exists for the requested probe")
	}

	if err = c.Insert(params.Monitor); err != nil {
		return nil, err
	}

	err = c.Find(bson.M{"probe": params.Monitor.Probe}).One(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMonitors reads all configured temperature monitors
func GetMonitors(params *monitors.GetMonitorsParams) (*models.TemperatureMonitors, error) {
	return nil, nil
}
