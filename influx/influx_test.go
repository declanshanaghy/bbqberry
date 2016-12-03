package influx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/declanshanaghy/bbqberry/influx"
)

var _ = Describe("Influx", func() {
	It("can create the default connection", func() {
		Expect(GetDefaultClient()).ToNot(BeNil())
	})
})
