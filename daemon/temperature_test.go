package daemon

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature daemon", func() {
	var (
		bus *stubembd.StubSPIBus
	)

	BeforeEach(func() {
		bus = stubembd.NewStubSPIBus()
		hardware.StubBus = bus
	})

	It("should collect valid temperature readings from all probes", func() {
		temperatureLogger := newTemperatureLogger()
		readings := temperatureLogger.collectTemperatureMetrics()

		Expect(len(*readings)).To(Equal(int(temperatureLogger.temp.GetNumProbes())))
		for i, r := range *readings {
			Expect(*r.Probe).To(Equal(int32(i + 1)))
		}
	})
})
