package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/declanshanaghy/bbqberry/influx"
	"github.com/declanshanaghy/bbqberry/framework/error"
	"github.com/declanshanaghy/bbqberry/influx/example"
	"fmt"
)

func Health() (m models.Health, err error) {
	defer func() {
		log.Infof("service=%s healthy=%t", *m.ServiceInfo.Name, *m.Healthy)
	}()

	healthy := false
	m = models.Health{ Healthy: &healthy }

	si := new(models.ServiceInfo)
	si.Name = &framework.ConstantsObj.ServiceName
	si.Version = &framework.ConstantsObj.Version
	m.ServiceInfo = si

	client, err := influx.GetDefaultClient()
	if err != nil {
		e := new(models.Error)
		code := error_codes.ErrInfluxUnavailable
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", error_codes.GetText(*e.Code), err)
		m.Error = e
		return m, nil
	}

	tags := map[string]string{"service": *si.Name}
	fields := map[string]interface{}{
		"version":   si.Version,
	}
	
	_, err = influx_example.WriteExamplePoint(client, "health", tags, fields)
	if err != nil {
		e := new(models.Error)
		code := error_codes.ErrInfluxWrite
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", error_codes.GetText(*e.Code), err)
		m.Error = e
		return m, nil
	}

	healthy = true
	return m, nil
}