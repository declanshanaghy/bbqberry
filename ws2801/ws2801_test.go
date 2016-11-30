package ws2801_test


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/declanshanaghy/bbqberry/ws2801"
)

var _ = Describe("WS2801", func() {
	var (
		strand Strand
	)

	BeforeEach(func() {
		strand := Strand{}
		strand.Init(10)
	})

	AfterEach(func() {
		strand.Close()
	})

	Describe("Basic test", func() {
		Context("of pixel validation", func() {
			It("should fail on invalid range", func() {
				n := strand.GetNumPixels()

				Expect(func() {
					strand.ValidatePixel(n+1)
				}).To(Panic())
			})
		})
	})
})
