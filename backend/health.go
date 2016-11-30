package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/Polarishq/middleware/framework/log"
)

// Health check endpoint ...
func Health() (models.Health, error) {
	healthy := false
	h := models.Health{ Healthy: &healthy }

	si := new(models.ServiceInfo)
	si.Name = &framework.ConstantsObj.ServiceName
	si.Version = &framework.ConstantsObj.Version
	h.ServiceInfo = si

	ecode := int32(900)
	e := new(models.Error)
	e.Code = &ecode
	e.Message = "Skeleton implementation, insert your own functional checks here"
	h.Error = e

	log.Infof("service=%s healthy=%t", *h.ServiceInfo.Name, *h.Healthy)

	return h, nil
}