package backend_test

import (
	// "github.com/declanshanaghy/bbqberry/mocks/mock_embd"
	// "github.com/golang/mock/gomock"
	// "github.com/declanshanaghy/bbqberry/framework_test"

	"fmt"

	. "github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Termperatures API", func() {
	// var (
	// 	t      framework_test.GinkgoTestReporter
	// 	ctrl   *gomock.Controller
	// 	bus    *mock_embd.MockSPIBus
	// )

	// BeforeEach(func() {
	// 	ctrl = gomock.NewController(t)
	// 	bus = mock_embd.NewMockSPIBus(ctrl)

	// 	hardware.Mock = true
	// 	hardware.MockBus = bus
	// })

	// AfterEach(func() {
	// 	ctrl.Finish()
	// })

	It("should return a single temperature reading when given a probe number", func() {
		// FFS Mocking not working
		// no matching expected call: []interface {}.%!v(MISSING)(%!v(MISSING))
		// var data [3]uint8
		// data[0] = 1
		// data[1] = uint8(1)<<7 | uint8(1)<<4
		// data[2] = 0
		// bus.EXPECT().TransferAndReceiveData(data[:])

		probe := int32(1)
		params := temperature.GetProbeReadingsParams{
			Probe: &probe,
		}
		m, err := GetTemperatureProbeReadings(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings should not have returned an error")
		Expect(m).To(HaveLen(1), "Incorrect number of readings returned")
	})
	It("should return all temperature readings when not given a probe number", func() {
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
