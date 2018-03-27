package hardware_test

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
)

var _ = Describe("Temperature", func() {
	probe := int32(0)
	hwCfg := framework.Constants.Hardware
	lim := hwCfg.Probes[probe].Limits
	minC := *lim.MinWarnCelsius
	maxC := *lim.MaxWarnCelsius
	minA := framework.ConvertCelsiusToAnalog(minC)
	maxA := framework.ConvertCelsiusToAnalog(maxC)
	_, minF := framework.ConvertCToKFInt32(float32(minC))
	_, maxF := framework.ConvertCToKFInt32(float32(maxC))

	log.Infof("minC=%d minA=%d maxC=%d maxA=%d", minC, minA, maxC, maxA)

	Context("conversion functions", func() {
		It("convert analog to voltage successfully", func() {
			v := framework.ConvertAnalogToVoltage(0)
			Expect(v).To(Equal(float32(0.0)))
		})
		It("should convert celsius to analog successfully", func() {
			a := framework.ConvertCelsiusToAnalog(400)
			Expect(a).To(Equal(int32(1008)))
		})
	})
	Context("TemperatureReader object", func() {
		It("should return a warning on low temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			actualA := minA - 1
			_, actualF := framework.ConvertAnalogToCF(actualA)
			stubembd.SetFakeTemp(probe, actualA)
			reader.GetTemperatureReading(probe, &reading)

			msg := fmt.Sprintf("%d° F exceeds low temperature limit of %d° F", actualF, minF)
			Expect(reading.Warning).To(Equal(msg))
			Expect(*reading.Celsius).To(BeNumerically("<", minC))
		})
		It("should return no warnings when above the low limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			stubembd.SetFakeTemp(probe, minA)
			reader.GetTemperatureReading(probe, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(minC))
		})
		It("should return a warning on high temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			actualA := maxA + 2
			_, actualF := framework.ConvertAnalogToCF(actualA)
			stubembd.SetFakeTemp(probe, actualA)
			reader.GetTemperatureReading(probe, &reading)

			msg := fmt.Sprintf("%d° F exceeds high temperature limit of %d° F", actualF, maxF)
			Expect(reading.Warning).To(Equal(msg))
			Expect(*reading.Celsius).To(BeNumerically(">", maxC))
		})
		It("should return no warnings when below the high limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			stubembd.SetFakeTemp(probe, maxA)
			reader.GetTemperatureReading(probe, &reading)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(Equal(maxC))
		})
	})
})
