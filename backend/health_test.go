package backend_test

import (
	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health", func() {
	Describe("Basic test", func() {
		Context("of API", func() {
			It("should return healthy", func() {
				health, err := Health()

				Expect(err).ShouldNot(HaveOccurred(), "Health check returned an error")

				healthy := true
				si := models.ServiceInfo{
					Name:    &framework.ConstantsObj.ServiceName,
					Version: &framework.ConstantsObj.Version,
				}
				h := models.Health{
					Healthy:     &healthy,
					ServiceInfo: &si,
				}

				Expect(h).To(Equal(health))
			})
		})
	})
})
