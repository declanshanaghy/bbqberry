package daemon

import (
	"time"

	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	"github.com/go-openapi/strfmt"
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
	It("should start and stop cleanly", func() {
		temperatureLogger := newTemperatureLogger()

		err := temperatureLogger.StartBackground()
		Expect(err).ToNot(HaveOccurred())

		// This should allow at least 1 tickable execution
		time.Sleep(time.Millisecond * 1000)

		err = temperatureLogger.StopBackground()
		Expect(err).ToNot(HaveOccurred())
	})
	It("should collect valid temperature readings from all probes", func() {
		temperatureLogger := newTemperatureLogger()
		readings, err := temperatureLogger.collectTemperatureMetrics()

		Expect(err).ToNot(HaveOccurred())

		Expect(len(readings)).To(Equal(int(temperatureLogger.reader.GetNumProbes())))
		for i, r := range readings {
			Expect(*r.Probe).To(Equal(int32(i)))
		}
	})
	It("should log temperature readings successfully", func() {
		temperatureLogger := newTemperatureLogger()
		readings := make([]*models.TemperatureReading, 0)

		for i := int32(0); i < 5; i++ {
			t := strfmt.DateTime(time.Now())
			p := 10 * i
			a := p + 1
			v := float32(p + 2)
			k := p + 3
			c := p + 4
			f := p + 5
			readings = append(readings, &models.TemperatureReading{
				Probe:      &i,
				Analog:     &a,
				DateTime:   &t,
				Voltage:    &v,
				Kelvin:     &k,
				Celsius:    &c,
				Fahrenheit: &f,
			})
		}

		err := temperatureLogger.logTemperatureMetrics(readings)

		Expect(err).ToNot(HaveOccurred())
	})
})
