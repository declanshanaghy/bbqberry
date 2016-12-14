package backend

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/errorcodes"
	"github.com/declanshanaghy/bbqberry/influxdb/util"
	"github.com/declanshanaghy/bbqberry/models"
)

// Health performs all internal health checks to ensure all systems are functioning
func Health() (m models.Health, err error) {
	defer func() {
		log.Infof("service=%s healthy=%t", *m.ServiceInfo.Name, *m.Healthy)
	}()

	healthy := false
	m = models.Health{Healthy: &healthy}

	si := new(models.ServiceInfo)
	si.Name = &framework.Constants.ServiceName
	si.Version = &framework.Constants.Version
	m.ServiceInfo = si

	_, err = util.WriteHealthCheck()
	if err != nil {
		log.Error(err)
		e := new(models.Error)
		code := errorcodes.ErrInfluxWrite
		e.Code = &code
		e.Message = errorcodes.GetText(*e.Code)
		m.Error = e
		return m, nil
	}

	healthy = true
	return m, nil
}
