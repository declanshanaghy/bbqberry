package daemon

import (
	"time"

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

	It("should return a sane name", func() {
		tl := newTemperatureLogger()
		Expect(tl.GetName()).ToNot(BeNil())
	})
	It("should collect valid temperature readings from all probes", func() {
		temperatureLogger := newTemperatureLogger()
		readings := temperatureLogger.collectTemperatureMetrics()

		Expect(len(*readings)).To(Equal(int(temperatureLogger.reader.GetNumProbes())))
		for i, r := range *readings {
			Expect(*r.Probe).To(Equal(int32(i + 1)))
		}
	})
	It("should start and stop cleanly", func() {
		temperatureLogger := newTemperatureLogger()

		err := temperatureLogger.StartBackground()
		Expect(err).ToNot(HaveOccurred())

		// This should allow at least 1 tickable execution
		time.Sleep(time.Millisecond * 1000)

		err = temperatureLogger.StopBackground()
		Expect(err).ToNot(HaveOccurred())
	})
})
