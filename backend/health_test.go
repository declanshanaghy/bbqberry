package backend_test

import (
	"os"

	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health API", func() {
	Context("in a correctly configured environment", func() {
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
	Context("in an incorrectly configured environment", func() {
		var origInfluxHost string

		BeforeEach(func() {
			origInfluxHost = os.Getenv("INFLUXDB_HOST")
			os.Setenv("INFLUXDB_HOST", "nonexistent")
			influxdb.LoadConfig()
		})

		AfterEach(func() {
			os.Setenv("INFLUXDB_HOST", origInfluxHost)
			influxdb.LoadConfig()
		})

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
})
