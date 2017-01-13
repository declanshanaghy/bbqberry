package restapi

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server Configuration", func() {
	It("global setup and shutdown should start the commander", func() {
		globalStartup()
		Expect(commander.IsRunning()).To(BeTrue())
		
		globalShutdown()
		Expect(commander.IsRunning()).To(BeFalse())
	})
})
