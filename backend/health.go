package backend

import (
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/influx"
	"github.com/declanshanaghy/bbqberry/framework/error"
	"fmt"
	"github.com/declanshanaghy/bbqberry/influx/example"
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

	client, err := influx.GetDefaultClient()
	if err != nil {
		e := new(models.Error)
		code := errorcodes.ErrInfluxUnavailable
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", errorcodes.GetText(*e.Code), err)
		m.Error = e
		return m, nil
	}

	tags := map[string]string{"service": *si.Name}
	fields := map[string]interface{}{
		"version": si.Version,
	}

	_, err = influxexample.WriteExamplePoint(client, "health", tags, fields)
	if err != nil {
		e := new(models.Error)
		code := errorcodes.ErrInfluxWrite
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", errorcodes.GetText(*e.Code), err)
		m.Error = e
		return m, nil
	}

	healthy = true
	return m, nil
}
