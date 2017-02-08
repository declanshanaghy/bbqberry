package hardware

import (
	"github.com/declanshanaghy/bbqberry/framework"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Temperature package", func() {
	Context("global functions", func() {
		precision := 0.001

		It("should convert celsius properly", func() {
			// Format is [c] = [k, f]
			conversions := map[float32][2]float32{
				-273.15: {0, -459.67},
				0.0:     {273.15, 32.0},
				100.0:   {373.15, 212.0},
				250.0:   {523.15, 482.0},
			}

			for c, answer := range conversions {
				k, f := convertCToKF(c)
				Expect(k).To(BeNumerically("~", answer[0], precision))
				Expect(f).To(BeNumerically("~", answer[1], precision))
			}
		})

		It("should convert analog voltage properly", func() {
			hwCfg := framework.Constants.Hardware

			// Format is [v] = [k, c, f]
			conversions := map[float32][3]float32{
				0.0:       {23, -250, -418},
				hwCfg.VCC: {683, 410, 770},
			}

			for v, answer := range conversions {
				k, c, f := adafruitAD8495ThermocoupleVtoKCF(v)
				Expect(k).To(BeNumerically("~", answer[0], precision))
				Expect(c).To(BeNumerically("~", answer[1], precision))
				Expect(f).To(BeNumerically("~", answer[2], precision))
			}
		})
	})
})
