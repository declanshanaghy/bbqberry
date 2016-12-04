package backend_test

import (
	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Monitors API", func() {
	It("should return all monitors", func() {
		probe := int32(0)
		params := temperature.GetMonitorsParams{
			Probe: probe,
		}
		m, err := GetTemperatureMonitors(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureMonitors should not have returned an error")
		Expect(m).To(BeNil())
	})
})
