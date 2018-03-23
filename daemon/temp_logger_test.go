package daemon

import (
	"time"

	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature daemon", func() {
	It("should start and stop cleanly", func() {
		temperatureLogger := newTemperatureLoggerRunnable()

		err := temperatureLogger.StartBackground()
		Expect(err).ToNot(HaveOccurred())

		// This should allow at least 1 tickable execution
		time.Sleep(time.Millisecond * 1000)

		err = temperatureLogger.StopBackground()
		Expect(err).ToNot(HaveOccurred())
	})
	It("should collect valid temperature readings from all probes", func() {
		temperatureLogger := newTemperatureLogger()
		temperatureLogger.period = time.Nanosecond

		readings, err := temperatureLogger.collectTemperatureMetrics()

		Expect(err).ToNot(HaveOccurred())

		Expect(len(readings)).To(Equal(len(*temperatureLogger.reader.GetEnabledPobes())))
		for i, r := range readings {
			Expect(*r.Probe).To(Equal(int32(i)))
		}
	})
	It("should log temperature readings successfully", func() {
		temperatureLogger := &temperatureLogger{
			reader: hardware.NewTemperatureReader(),
			period: time.Nanosecond,
		}
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
