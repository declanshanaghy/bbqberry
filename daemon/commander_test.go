package daemon

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Commander daemon", func() {
	It("should exit when told", func() {
		cmdrChan := mach(chan bool, 1)
		cmdrChan <- false
		StartCommander()
		readings := collectTemperatureMetrics(temp)

		Expect(len(*readings)).To(Equal(int(temp.GetNumProbes())))
		for i, r := range *readings {
			Expect(*r.Probe).To(Equal(int32(i + 1)))
		}
	})
})
