package hardware_test

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hardware package", func() {
	Context("conversion functions", func() {
		It("convert celsius to analog successfully", func() {
			hwCfg := framework.Constants.Hardware
			lim := hwCfg.Probes[0].TempLimits
			min := *lim.MinAbsCelsius
			max := *lim.MaxAbsCelsius

			l := hardware.ConvertCelsiusToAnalog(min)
			log.Infof("MIN=%d, l=%d", min, l)
			h := hardware.ConvertCelsiusToAnalog(max)
			log.Infof("MAX=%d h=%d", max, h)

			Expect(l).To(Equal(int32(310)))
			Expect(h).To(Equal(int32(1008)))
		})
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
		hwCfg := framework.Constants.Hardware
		lim := hwCfg.Probes[0].TempLimits
		min := *lim.MinWarnCelsius
		max := *lim.MaxWarnCelsius
		lowAnalogBoundary := hardware.ConvertCelsiusToAnalog(min)
		highAnalogBoundary := hardware.ConvertCelsiusToAnalog(max)

		It("should return a warning on low temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = lowAnalogBoundary - 1
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal("Low temperature limit exceeded: " +
				"-41 °F exceeds limit of -40 °F"))
			Expect(*reading.Celsius).To(BeNumerically("<", min))
		})
		It("should return no warnings when above the low limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = lowAnalogBoundary
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(min))
		})
		It("should return a warning on high temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = highAnalogBoundary
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal("High temperature limit exceeded: " +
				"609 °F exceeds limit of 608 °F"))
			Expect(*reading.Celsius).To(BeNumerically(">", max))
		})
		It("should return no warnings when below the high limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = highAnalogBoundary - 1
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(max))
		})
	})
})
