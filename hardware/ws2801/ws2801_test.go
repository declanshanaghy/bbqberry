package ws2801_test

import (
	"github.com/declanshanaghy/bbqberry/framework_test"
	"github.com/declanshanaghy/bbqberry/hardware"
	. "github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/declanshanaghy/bbqberry/mocks/mock_embd"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WS2801s", func() {
	var (
		t      framework_test.GinkgoTestReporter
		ctrl   *gomock.Controller
		bus    *mock_embd.MockSPIBus
		strand WS2801
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(t)
		bus = mock_embd.NewMockSPIBus(ctrl)

		hardware.Mock = true
		hardware.MockBus = bus

		strand = hardware.NewStrandController()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("Basic test", func() {
		Context("Of sanity", func() {
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
		Context("Setting pixel colors", func() {
			It("should succeed with ints", func() {
				err := strand.SetPixelColor(0, RED)
				Expect(err).ToNot(HaveOccurred())
			})
			It("should succeed with r,g,b", func() {
				err := strand.SetPixelRGB(0, 0x00, 0xFF, 0xFF)
				Expect(err).ToNot(HaveOccurred())
			})
		})
		// Context("of strand update", func() {
		// 	It("should succeeds", func() {
		// 		data := make([]uint8, strand.GetNumPixels()*3)
		// 		bus.EXPECT().TransferAndReceiveData(data)

		// 		err := strand.SetPixelColor(0, RED)
		// 		Expect(err).ToNot(HaveOccurred())
		// 		err = strand.Update()
		// 		Expect(err).ToNot(HaveOccurred())
		// 	})
		// })
	})
})
