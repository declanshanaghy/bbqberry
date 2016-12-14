package ws2801_test

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	. "github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/declanshanaghy/bbqberry/stubs/stubembd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WS2801", func() {
	var (
		bus    *stubembd.StubSPIBus
		strand WS2801
	)

	BeforeEach(func() {
		bus = stubembd.NewStubSPIBus()
		hardware.StubBus = bus
		strand = hardware.NewStrandController()
	})

	Context("sanity checks", func() {
		It("should return correct pixel count", func() {
			numPixels := strand.GetNumPixels()
			Expect(hardware.HardwareConfig.NumLEDPixels).To(Equal(numPixels))
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
		It("update should call TransferAndReceiveData once", func() {
			err := strand.SetPixelColor(0, RED)
			Expect(err).ToNot(HaveOccurred())

			err = strand.Update()

			Expect(err).ToNot(HaveOccurred())
			Expect(1).To(Equal(bus.TransferAndReceiveDataCallCount))
		})
		It("close should disable all pixels and close the bus", func() {
			err := strand.Close()

			Expect(err).ToNot(HaveOccurred())
			Expect(1).To(Equal(bus.TransferAndReceiveDataCallCount))
			Expect(1).To(Equal(bus.CloseCallCount))
		})
	})
})
