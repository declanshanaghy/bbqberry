package hardware_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/Polarishq/middleware/framework/log"
)

var _ = Describe("Colors", func() {
	Context("RGB to hue", func() {
		It("should convert red successfully", func() {
			// Any of the following should be the same
			c := colorful.Color{0.313725, 0.478431, 0.721569}
			c, err := colorful.Hex("#517AB8")
			if err != nil{
				log.Error(err)
			}
			c = colorful.Hsv(216.0, 0.56, 0.722)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
			c = colorful.Xyz(0.189165, 0.190837, 0.480248)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
			c = colorful.Xyy(0.219895, 0.221839, 0.190837)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
			c = colorful.Lab(0.507850, 0.040585,-0.370945)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
			c = colorful.Luv(0.507849,-0.194172,-0.567924)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
			c = colorful.Hcl(276.2440, 0.373160, 0.507849)
			log.Infof("RGB values: %v, %v, %v", c.R, c.G, c.B)
		})
	})
})
