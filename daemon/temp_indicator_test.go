package daemon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/declanshanaghy/bbqberry/framework"
)

var _ = Describe("Temperature indicator", func() {
	It("should map min temp to pure blue", func() {
		color := getTempColor(framework.Constants.Hardware.MinTempWarnCelsius)
		Expect(color).To(Equal(0x0000FF))
	})
	It("should map max temp to pure red", func() {
		color := getTempColor(framework.Constants.Hardware.MaxTempWarnCelsius)
		Expect(color).To(Equal(0xFF0000))
	})
	It("should map mid temp to equal red & blue", func() {
		min := framework.Constants.Hardware.MinTempWarnCelsius
		max := framework.Constants.Hardware.MaxTempWarnCelsius
		color := getTempColor((max - min) * 0.5 + min)
		Expect(color).To(Equal(0x7F0080))
	})
	It("should map 1/4 between min and max to 1/4 between blue and red", func() {
		min := framework.Constants.Hardware.MinTempWarnCelsius
		max := framework.Constants.Hardware.MaxTempWarnCelsius
		color := getTempColor((max - min) * 0.25 + min)
		Expect(color).To(Equal(0x3F00C0))
	})
	It("should map 3/4 between min and max to 3/4 between blue and red", func() {
		min := framework.Constants.Hardware.MinTempWarnCelsius
		max := framework.Constants.Hardware.MaxTempWarnCelsius
		color := getTempColor((max - min) * 0.75 + min)
		Expect(color).To(Equal(0xBF0040))
	})
})
