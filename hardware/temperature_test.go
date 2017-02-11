package hardware_test

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hardware package", func() {
	probe := int32(0)
	hwCfg := framework.Constants.Hardware
	lim := hwCfg.Probes[probe].TempLimits
	minC := *lim.MinWarnCelsius
	maxC := *lim.MaxWarnCelsius
	minA := hardware.ConvertCelsiusToAnalog(minC)
	maxA := hardware.ConvertCelsiusToAnalog(maxC)
	_, minF := hardware.ConvertCToKFInt32(float32(minC))
	_, maxF := hardware.ConvertCToKFInt32(float32(maxC))

	log.Infof("minC=%d minA=%d maxC=%d maxA=%d", minC, minA, maxC, maxA)

	Context("conversion functions", func() {
		//It("convert analog to voltage successfully", func() {
		//	v := hardware.ConvertAnalogToVoltage(0)
		//	Expect(v).To(Equal(0))
		//})
		It("should convert celsius to analog successfully", func() {
			a := hardware.ConvertCelsiusToAnalog(400)
			Expect(a).To(Equal(int32(1008)))
		})
	})
	Context("TemperatureReader object", func() {
		It("should return a warning on low temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			actualA := minA - 1
			_, actualF := hardware.ConvertAnalogToCF(actualA)
			hardware.FakeTemps[probe] = actualA
			reader.GetTemperatureReading(probe, &reading)

			msg := fmt.Sprintf("Low temperature limit exceeded: %d °F exceeds limit of %d °F", actualF, minF)
			Expect(reading.Warning).To(Equal(msg))
			Expect(*reading.Celsius).To(BeNumerically("<", minC))
		})
		It("should return no warnings when above the low limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[probe] = minA
			reader.GetTemperatureReading(probe, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(minC))
		})
		It("should return a warning on high temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			actualA := maxA + 2
			_, actualF := hardware.ConvertAnalogToCF(actualA)
			hardware.FakeTemps[probe] = actualA
			reader.GetTemperatureReading(probe, &reading)

			msg := fmt.Sprintf("High temperature limit exceeded: %d °F exceeds limit of %d °F", actualF, maxF)
			Expect(reading.Warning).To(Equal(msg))
			Expect(*reading.Celsius).To(BeNumerically(">", maxC))
		})
		It("should return no warnings when below the high limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[probe] = maxA
			reader.GetTemperatureReading(probe, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(maxC))
		})
	})
})
