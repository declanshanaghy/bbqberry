package daemon

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
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
		started := time.Now()

		temp := hardware.NewTemperatureReader()
		readings := collectTemperatureMetrics(temp)
		
		Expect(len(*readings)).To(Equal(int(temp.GetNumProbes())))		
		for i, r := range *readings {
			Expect(*r.Probe).To(Equal(int32(i+1)))
			Expect(*r.Time).To(BeEquivalentTo(started))
		}
	})
})
