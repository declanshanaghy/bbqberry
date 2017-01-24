package daemon

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commander daemon", func() {
	It("should return a sane name", func() {
		cmdr := NewCommander()
		Expect(cmdr.GetName()).ToNot(BeNil())
	})
	It("should start and stop cleanly", func() {
		cmdr := NewCommander()

		err := cmdr.StartBackground()
		Expect(err).ToNot(HaveOccurred())

		// This should allow at least 1 tickable execution
		time.Sleep(time.Millisecond * 1000)

		err = cmdr.StopBackground()
		Expect(err).ToNot(HaveOccurred())
	})
})
