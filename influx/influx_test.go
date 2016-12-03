package influx_test

import (
	. "github.com/declanshanaghy/bbqberry/influx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Influx", func() {
	It("can create the default client", func() {
		Expect(NewDefaultClient()).ToNot(BeNil())
	})
	It("can create an HTTP client", func() {
		Expect(NewHTTPClient()).ToNot(BeNil())
	})
	It("can create a UDP client", func() {
		Expect(NewUDPClient()).ToNot(BeNil())
	})
})
