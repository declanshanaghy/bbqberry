package backend

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
)

var _ = Describe("Monitors backends", func() {
	var mgr *MonitorsManager
	var err error

	expectCollectionCount := func(n int, label string) {
		l, err := mgr.GetCollection().Count()
		Expect(err).ShouldNot(HaveOccurred(),
			fmt.Sprintf("Monitors collection count should not have returned an error during %s", label))
		Expect(l).To(Equal(n), fmt.Sprintf("Unexpected collection count during %s", label))

		indb := make([]map[string]interface{}, 0)
		err = mgr.GetCollection().Find(nil).All(&indb)
		Expect(err).ShouldNot(HaveOccurred(),
			fmt.Sprintf("Find All should not have returned an error during %s", label))
		Expect(indb).To(HaveLen(n),
			fmt.Sprintf("Unexpected number of documents found in collection during %s", label))
	}

	BeforeEach(func() {
		collectionName := fmt.Sprintf("monitors_test%d", config.GinkgoConfig.ParallelNode)
		mgr, err = newMonitorsManagerForCollection(collectionName)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		c := mgr.GetCollection()
		c.DropCollection()
		Expect(err).ToNot(HaveOccurred())

		log.WithField("Name", c.FullName).Info("Collection dropped")
		expectCollectionCount(0, fmt.Sprintf("BeforeEach %s", mgr.GetCollection().FullName))

		mgr.Close()
	})

	It("should create a monitor", func() {
		probe := int32(2)
		fake := "fake"
		label := "test create"
		min := int32(5)
		max := int32(105)
		scale := "celsius"

		params := monitors.CreateMonitorParams{
			Monitor: &models.TemperatureMonitor{
				ID:    fake,
				Label: &label,
				Probe: &probe,
				Min:   &min,
				Max:   &max,
				Scale: &scale,
			},
		}
		monitor, err := mgr.CreateMonitor(&params)

		Expect(err).ShouldNot(HaveOccurred(), "CreateMonitor should not have returned an error")
		Expect(monitor).ToNot(BeNil())

		Expect(*monitor.Probe).To(Equal(probe))
		Expect(*monitor.Label).To(Equal(label))
		Expect(*monitor.Min).To(Equal(min))
		Expect(*monitor.Max).To(Equal(max))
		Expect(*monitor.Scale).To(Equal(scale))
		Expect(monitor.ID).ToNot(Equal(fake))
	})
	It("should return all monitors when not given a probe number", func() {
		scale := "celsius"

		created := make([]*models.TemperatureMonitor, 0, len(framework.Constants.Hardware.Probes))
		for i, settings := range framework.Constants.Hardware.Probes {
			probe := int32(i)
			label := fmt.Sprintf("test get all - probe %d", i)
			monitor := models.TemperatureMonitor{
				Label: &label,
				Probe: &probe,
				Min:   settings.Limits.MinWarnCelsius,
				Max:   settings.Limits.MaxWarnCelsius,
				Scale: &scale,
			}
			mgr.GetCollection().Insert(monitor)
			created = append(created, &monitor)
		}
		expectCollectionCount(len(created), fmt.Sprintf("create all test setup %s",
			mgr.GetCollection().FullName))

		params := monitors.GetMonitorsParams{}
		monitors, err := mgr.GetMonitors(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetMonitors should not have returned an error")
		Expect(monitors).To(HaveLen(len(created)))
	})
})
