package backend_test

import (
	"fmt"

	"time"

	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Termperatures API", func() {
	hwCfg := framework.Config.Hardware

	It("should return a single temperature reading when given a probe number", func() {
		probe := int32(1)
		params := temperatures.GetTemperaturesParams{
			Probe: &probe,
		}
		m, err := backend.GetTemperatureProbeReadings(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings "+
			"should not have returned an error")
		Expect(m).To(HaveLen(1), "Incorrect number of readings returned")
	})
	It("should return all temperature readings when not given a probe number", func() {
		started := time.Now()

		params := temperatures.GetTemperaturesParams{}
		m, err := backend.GetTemperatureProbeReadings(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings "+
			"should not have returned an error")
		Expect(m).To(HaveLen(len(hwCfg.Probes)), "Incorrect number of readings returneds")

		for i, reading := range m {
			Expect(int32(i)).To(
				Equal(*reading.Probe),
				fmt.Sprintf("Probe %d has incorrect number", i))
			dt, err := strfmt.ParseDateTime(reading.DateTime.String())
			Expect(err).ToNot(HaveOccurred())
			Expect(time.Time(dt)).Should(BeTemporally("~", started, time.Second))
		}
	})
	It("should return all temperature readings when given a negative probe number", func() {
		started := time.Now()

		probe := int32(-1)
		params := temperatures.GetTemperaturesParams{
			Probe: &probe,
		}
		m, err := backend.GetTemperatureProbeReadings(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetTemperatureProbeReadings "+
			"should not have returned an error")
		Expect(m).To(HaveLen(len(hwCfg.Probes)), "Incorrect number of readings returneds")

		for i, reading := range m {
			Expect(int32(i)).To(
				Equal(*reading.Probe),
				fmt.Sprintf("Probe %d has incorrect number", i))
			dt, err := strfmt.ParseDateTime(reading.DateTime.String())
			Expect(err).ToNot(HaveOccurred())
			Expect(time.Time(dt)).Should(BeTemporally("~", started, time.Second))
		}
	})
})
