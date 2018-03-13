package backend

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/mongodb"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	"github.com/go-openapi/errors"
	"gopkg.in/mgo.v2/bson"
)

const defaultCollection = "monitors"

// MonitorsManager provides methods to operate on monitors within the backend database
type MonitorsManager struct {
	*mongodb.ResourceManager
}

// NewMonitorsManager creates a manager which can operate on temperature monitors
func NewMonitorsManager() (*MonitorsManager, error) {
	return newMonitorsManagerForCollection(defaultCollection)
}

// newMonitorsManagerForCollection creates a manager which can operate on temperature monitors
// pass in a collection to specify a non default collection name to use within the database
// this functionality is provided to support unitttest parallelization by using different collection names
func newMonitorsManagerForCollection(collection string) (*MonitorsManager, error) {
	if collection != defaultCollection {
		log.WithField("CollectionName", collection).Warning(
			"MonitorsManager custom setup completed")
	}

	rm, err := mongodb.NewResourceManager(collection)
	if err != nil {
		return nil, err
	}

	m := MonitorsManager{
		ResourceManager: rm,
	}
	return &m, nil
}

// CreateMonitor creates a new temperature monitor
func (m *MonitorsManager) CreateMonitor(params *monitors.CreateMonitorParams) (*models.TemperatureMonitor, error) {
	c := m.GetCollection()

	result := new(models.TemperatureMonitor)

	err := c.Find(bson.M{"probe": params.Monitor.Probe}).One(result)
	if err == nil && result.Probe != nil {
		return nil, errors.New(400, "Monitor already exists for probe %d", *params.Monitor.Probe)
	}

	monitor := params.Monitor
	monitor.ID = ""
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
func (m *MonitorsManager) GetMonitors(params *monitors.GetMonitorsParams) ([]*models.TemperatureMonitor, error) {

	c := m.GetCollection()

	log.Error("hello")

	result := make([]*models.TemperatureMonitor, 0)
	if params.Probe != nil {
		if err := c.Find(bson.M{"probe": params.Probe}).All(&result); err != nil {
			return nil, err
		}
	} else {
		if err := c.Find(nil).All(&result); err != nil {
			return nil, err
		}
	}

	return result, nil
}
