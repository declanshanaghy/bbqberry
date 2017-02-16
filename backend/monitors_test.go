package backend_test

import (
	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Monitors backends", func() {
	It("should create a monitor", func() {
		probe := int32(2)
		params := monitors.CreateMonitorParams{
			Probe: probe,
		}
		m, err := backend.CreateMonitor(&params)

		Expect(err).ShouldNot(HaveOccurred(), "CreateMonitor should not have returned an error")
		Expect(m).ToNot(BeNil())

		Expect(*m.Probe).To(Equal(probe))
		Expect(*m.ID).ToNot(BeNil())
	})
	It("should return all monitors", func() {
		probe := int32(0)
		params := monitors.GetMonitorsParams{
			Probe: &probe,
		}
		m, err := backend.GetMonitors(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetMonitors should not have returned an error")
		Expect(m).To(BeNil())
	})
})
