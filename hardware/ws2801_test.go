package hardware_test

import (
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	. "github.com/declanshanaghy/bbqberry/hardware"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WS2801", func() {
	var (
		strand WS2801
	)

	BeforeEach(func() {
		strand = hardware.NewStrandController()
		hardware.StubSPIBus.Reset()
	})

	Context("sanity checks", func() {
		hwCfg := framework.Config.Hardware

		It("should return correct pixel count", func() {
			numPixels := strand.GetNumPixels()
			Expect(*hwCfg.NumLedPixels).To(Equal(numPixels))
		})
		It("should fail on exceeding max pixel count", func() {
			numPixels := strand.GetNumPixels()
			err := strand.SetPixelColor(numPixels+1, 0)
			Expect(err).Should(HaveOccurred(), "Invalid pixel update was not caught")
		})
		It("should fail on negative pixel number", func() {
			err := strand.SetPixelColor(-1, 0)
			Expect(err).Should(HaveOccurred(), "Negative pixel supdate was not caught")
		})
	})
	Context("setting pixel colors", func() {
		It("should succeed with ints", func() {
			err := strand.SetPixelColor(0, RED)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should succeed with r,g,b", func() {
			err := strand.SetPixelRGB(0, 0x00, 0xFF, 0xFF)
			Expect(err).ToNot(HaveOccurred())
		})
	})
	Context("strand functionality", func() {
		BeforeEach(func() {
			hardware.StubSPIBus.Reset()
		})
		It("update should call TransferAndReceiveData once", func() {
			err := strand.SetPixelColor(0, RED)
			Expect(err).ToNot(HaveOccurred())

			err = strand.Update()

			Expect(err).ToNot(HaveOccurred())
			Expect(1).To(Equal(hardware.StubSPIBus.WriteCallCount))
		})
		It("close should disable all pixels and close the bus", func() {
			err := strand.Close()

			Expect(err).ToNot(HaveOccurred())
			Expect(1).To(Equal(hardware.StubSPIBus.WriteCallCount))
			Expect(1).To(Equal(hardware.StubSPIBus.CloseCallCount))
		})
	})
})
