package influxdb_test

import (
	. "github.com/declanshanaghy/bbqberry/influxdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Influx", func() {
	It("can create the default client", func() {
		Expect(NewHTTPClient()).ToNot(BeNil())
	})
})
