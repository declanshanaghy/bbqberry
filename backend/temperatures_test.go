package backend_test

import (
	"fmt"

	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Termperatures", func() {
	Describe("Basic test", func() {
		Context("of temperature API", func() {
			It("should return a single temeprature reading", func() {
				probe := int32(1)
				params := temperature.GetProbeReadingsParams{
					Probe: &probe,
				}
				m, err := GetTemperatureProbeReadings(&params)

				Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings should not have returned an error")
				Expect(m).To(HaveLen(1), "Incorrect number of readings returned")
			})
			It("should return all temeprature readings", func() {
				probe := int32(0)
				params := temperature.GetProbeReadingsParams{
					Probe: &probe,
				}
				m, err := GetTemperatureProbeReadings(&params)

				Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings should not have returned an error")
				Expect(m).To(
					HaveLen(int(hardware.HardwareConfig.NumTemperatureProbes)),
					"Incorrect number of readings returneds")

				for i, reading := range m {
					Expect(int32(i+1)).To(
						Equal(*reading.Probe),
						fmt.Sprintf("Probe %d has incorrect number", i))
				}
			})
		})
	})
})
