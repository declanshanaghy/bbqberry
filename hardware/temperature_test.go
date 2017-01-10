package hardware_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"fmt"
	"github.com/declanshanaghy/bbqberry/framework"
)

var _ = Describe("TemperatureReader Object", func() {
	Context("limit tests", func() {
		hwCfg := framework.Constants.Hardware
		
		// Analog reading corresponding to the low threshold
		lowBoundary := int32(318)
		
		// Analog reading corresponding to the high threshold
		highBoundary := int32(735)
		
		It("should return a warning on low temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = lowBoundary - 1
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal("Low temperature limit exceeded: actual=-45.04 °C < threshold=-45.00 °C"))
			Expect(*reading.Celsius).To(BeNumerically("<", hwCfg.MinTempWarnCelsius))
		})
		It("should return no warnings when above the low limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = lowBoundary
			reader.GetTemperatureReading(1, &reading)
			fmt.Printf("Warning is %s\n", reading.Warning)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(BeNumerically(">", hwCfg.MinTempWarnCelsius))
		})
		It("should return a warning on high temp reading", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = highBoundary + 1
			reader.GetTemperatureReading(1, &reading)

			Expect(reading.Warning).To(Equal("High temperature limit exceeded: actual=225.02 °C > threshold=225.00 °C"))
			Expect(*reading.Celsius).To(BeNumerically(">", hwCfg.MaxTempWarnCelsius))
		})
		It("should return no warnings when below the high limit", func() {
			reader := hardware.NewTemperatureReader()
			reading := models.TemperatureReading{}

			hardware.FakeTemps[1] = highBoundary
			reader.GetTemperatureReading(1, &reading)
			fmt.Printf("Warning is %s\n", reading.Warning)

			Expect(reading.Warning).To(Equal(""))
			Expect(*reading.Celsius).To(BeNumerically("<", hwCfg.MaxTempWarnCelsius))
		})
	})
})
