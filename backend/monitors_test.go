package backend_test

import (
	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Monitors API", func() {
	It("should return all monitors", func() {
		probe := int32(0)
		params := monitors.GetMonitorsParams{
			Probe: probe,
		}
		m, err := GetMonitors(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetMonitors should not have returned an error")
		Expect(m).To(BeNil())
	})
})
