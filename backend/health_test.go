package backend

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/framework"
)

func TestHealth(t *testing.T) {
	health, err := Health()
	assert.Nil(t, err, "Standard health check returned an error")
	assert.Equal(t, *health.Healthy, true, "Standard health check returned unhealthy")

	si := new(models.ServiceInfo)
	si.Name = &framework.ConstantsObj.ServiceName
	si.Version = &framework.ConstantsObj.Version

	assert.Equal(t, health.ServiceInfo, si, "Standard health check returned unexpected ServiceInfo")
}