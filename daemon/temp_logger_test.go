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

		Expect(len(*readings)).To(Equal(int(temperatureLogger.reader.GetNumProbes())))
		for i, r := range *readings {
			Expect(*r.Probe).To(Equal(int32(i + 1)))
		}
	})
	It("should log temperature readings successfully", func() {
		temperatureLogger := newTemperatureLogger()
		readings := models.TemperatureReadings{}

		for i := int32(0); i < 5; i++ {
			t := strfmt.DateTime(time.Now())
			p := 10 * i
			a := p + 1
			r := p + 3
			v := float32(p + 3)
			k := p + 4
			c := p + 5
			f := p + 6
			readings = append(readings, &models.TemperatureReading{
				Probe:      &i,
				Analog:     &a,
				DateTime:   &t,
				Resistance: &r,
				Voltage:    &v,
				Kelvin:     &k,
				Celsius:    &c,
				Fahrenheit: &f,
			})
		}

		err := temperatureLogger.logTemperatureMetrics(&readings)

		Expect(err).ToNot(HaveOccurred())

	})
})
