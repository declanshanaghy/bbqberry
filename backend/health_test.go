package backend_test

import (
	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health API", func() {
	It("should return healthy", func() {
		health, err := Health()

		Expect(err).ShouldNot(HaveOccurred(), "Health check returned an error")

		healthy := true
		si := models.ServiceInfo{
			Name:    &framework.Constants.ServiceName,
			Version: &framework.Constants.Version,
		}
		h := models.Health{
			Healthy:     &healthy,
			ServiceInfo: &si,
		}

		Expect(h).To(Equal(health))
	})
})
