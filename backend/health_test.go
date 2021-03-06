package backend_test

import (
	"os"

	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/errorcodes"
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
				Name:    &framework.Config.ServiceName,
				Version: &framework.Config.Version,
			}
			h := models.Health{
				Healthy:     &healthy,
				ServiceInfo: &si,
			}

			Expect(*health).To(Equal(h))
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

		It("should fail writing to InfluxDB", func() {
			health, err := Health()

			Expect(err).ShouldNot(HaveOccurred(), "Health check returned an error")

			healthy := false
			si := models.ServiceInfo{
				Name:    &framework.Config.ServiceName,
				Version: &framework.Config.Version,
			}
			h := models.Health{
				Healthy:     &healthy,
				ServiceInfo: &si,
			}

			code := errorcodes.ErrInfluxWrite
			e := models.Error{
				Code:    &code,
				Message: errorcodes.GetText(code),
			}
			h.Error = &e

			Expect(*health).To(Equal(h))
		})
	})
})
