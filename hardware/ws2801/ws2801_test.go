package ws2801_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/golang/mock/gomock"
	"github.com/declanshanaghy/bbqberry/mocks"
	"fmt"
)

type GinkgoTestReporter struct {}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args))
}

var _ = Describe("WS2801", func() {
	var (
		t 		GinkgoTestReporter
		ctrl	*gomock.Controller
		bus 	*mock_embd.MockSPIBus
		strand	WS2801
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(t)
		bus := mock_embd.NewMockSPIBus(ctrl)
		strand = NewWS2801(10, bus)
	})

	AfterEach(func() {
		strand.Close()
		bus.Close()
	})

	Describe("Basic test", func() {
		Context("of sanity", func() {
			It("should return correct pixel count", func() {
				Expect(func() {
					strand.GetNumPixels()
				}).To(Equal(10))
			})
			//It("should fail on invalid range", func() {
			//	Expect(func() {
			//		strand.ValidatePixel(n+1)
			//	}).To(Panic())
			//})
		})
	})
})
