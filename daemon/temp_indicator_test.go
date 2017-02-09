package daemon

import (
	"github.com/declanshanaghy/bbqberry/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature indicator", func() {
	lim := framework.Constants.Hardware.Probes[0].TempLimits
	min := float32(*lim.MinWarnCelsius)
	max := float32(*lim.MaxWarnCelsius)

	It("should map min temp to pure blue", func() {
		color := getTempColor(*lim.MinWarnCelsius, *lim.MinWarnCelsius, *lim.MaxWarnCelsius)
		Expect(color).To(Equal(0x0000FF))
	})
	It("should map max temp to pure red", func() {
		color := getTempColor(*lim.MaxWarnCelsius, *lim.MinWarnCelsius, *lim.MaxWarnCelsius)
		Expect(color).To(Equal(0xFF0000))
	})
	It("should map mid temp to equal red & blue", func() {
		color := getTempColor(int32((max-min)*0.5+min), int32(min), int32(max))
		Expect(color).To(Equal(0x7F0080))
	})
	It("should map 1/4 between min and max to 1/4 between blue and red", func() {
		color := getTempColor(int32((max-min)*0.25+min), int32(min), int32(max))
		Expect(color).To(Equal(0x3F00C0))
	})
	It("should map 3/4 between min and max to 3/4 between blue and red", func() {
		color := getTempColor(int32((max-min)*0.75+min), int32(min), int32(max))
		Expect(color).To(Equal(0xbf0040))
	})
})
