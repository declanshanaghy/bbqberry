package backend

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/errorcodes"
	"github.com/declanshanaghy/bbqberry/models"
)

// Health performs all internal health checks to ensure all systems are functioning
func Health() (*models.Health, error) {
	healthy := false
	m := models.Health{Healthy: &healthy}

	defer func() {
		log.Infof("service=%s healthy=%t", *m.ServiceInfo.Name, *m.Healthy)
	}()

	si := new(models.ServiceInfo)
	si.Name = &framework.Config.ServiceName
	si.Version = &framework.Config.Version
	m.ServiceInfo = si

	tags := map[string]string{"service": framework.Config.ServiceName}
	fields := map[string]interface{}{
		"version": framework.Config.Version,
	}

	_, err := influxdb.WritePoint("Health", tags, fields)
	if err != nil {
		log.Error(err)
		e := new(models.Error)
		code := errorcodes.ErrInfluxWrite
		e.Code = &code
		e.Message = errorcodes.GetText(*e.Code)
		m.Error = e
		return &m, nil
	}

	healthy = true
	return &m, nil
}
