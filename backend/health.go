package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/influx"
	"github.com/declanshanaghy/bbqberry/framework/error"
	"github.com/declanshanaghy/bbqberry/influx/example"
	"fmt"
)

// Health check endpoint ...
func Health() (h models.Health, err error) {
	defer func() {
		log.Infof("service=%s healthy=%t", *h.ServiceInfo.Name, *h.Healthy)
	}()

	healthy := false
	h = models.Health{ Healthy: &healthy }

	si := new(models.ServiceInfo)
	si.Name = &framework.ConstantsObj.ServiceName
	si.Version = &framework.ConstantsObj.Version
	h.ServiceInfo = si

	client, err := influx.GetDefaultClient()
	if err != nil {
		e := new(models.Error)
		code := error_codes.ERR_INFLUX_UNAVAILABLE
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", error_codes.MESSAGES[*e.Code], err)
		h.Error = e
		return h, nil
	}

	tags := map[string]string{"service": *si.Name}
	fields := map[string]interface{}{
		"version":   si.Version,
	}
	
	_, err = influx_example.WriteExamplePoint(client, "health", tags, fields)
	if err != nil {
		e := new(models.Error)
		code := error_codes.ERR_INFLUX_WRITE_ERROR
		e.Code = &code
		e.Message = fmt.Sprintf("%s %s", error_codes.MESSAGES[*e.Code], err)
		h.Error = e
		return h, nil
	}

	healthy = true
	return h, nil
}