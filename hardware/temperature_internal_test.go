package hardware

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/declanshanaghy/bbqberry/framework"
)

var _ = Describe("Temperature package", func() {
	Context("global functions", func() {
		precision := 0.001
		
		It("should convert celsius properly", func() {
			// Format is [c] = [k, f]
			conversions := map[float32][2]float32 {
				-273.15: [2]float32{0, -459.67},
				0.0: [2]float32{273.15, 32.0},
				100.0: [2]float32{373.15, 212.0},
				250.0: [2]float32{523.15, 482.0},
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
			conversions := map[float32][3]float32 {
				0.0: [3]float32{23.15, -250.0, -418.0},
				hwCfg.VCC: [3]float32{683.15, 410.0, 770.0},
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
