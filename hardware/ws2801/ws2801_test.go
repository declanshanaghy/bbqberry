package ws2801_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/golang/mock/gomock"
	"github.com/declanshanaghy/bbqberry/mocks/mock_embd"
	"github.com/declanshanaghy/bbqberry/framework_test"
)


var _ = Describe("WS2801", func() {
	var (
		t 		framework_test.GinkgoTestReporter
		ctrl	*gomock.Controller
		bus 	*mock_embd.MockSPIBus
		strand	WS2801
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(t)
		bus = mock_embd.NewMockSPIBus(ctrl)
		strand = NewWS2801(10, bus)
	})

	AfterEach(func() {
	})

	Describe("Basic test", func() {
		Context("of sanity", func() {
			It("should return correct pixel count", func() {
				numPixels := strand.GetNumPixels()
				Expect(10).To(Equal(numPixels))
			})
			//It("should close the bus", func() {
			//	bus.EXPECT().Close()
			//	strand.Close()
			//})
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
	})
})
