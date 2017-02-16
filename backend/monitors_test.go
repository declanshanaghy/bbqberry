package backend_test

import (
	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/db/mongodb"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/mgo.v2"
)

var _ = Describe("Monitors backends", func() {

	var (
		session *mgo.Session
		db      *mgo.Database
		err     error
	)

	BeforeEach(func() {
		session, db, err = mongodb.GetSession()
		if err != nil {
			Fail(err.Error())
		}
		if err = db.DropDatabase(); err != nil {
			Fail(err.Error())
		}
	})

	AfterEach(func() {
		db.Logout()
		session.Close()
	})

	It("should create a monitor", func() {
		fake := "fake"
		label := "A probe"
		probe := int32(2)
		min := int32(5)
		max := int32(105)
		scale := "celsius"

		params := monitors.CreateMonitorParams{
			Monitor: &models.TemperatureMonitor{
				ID:    "fake",
				Label: &label,
				Probe: &probe,
				Min:   &min,
				Max:   &max,
				Scale: &scale,
			},
		}
		m, err := backend.CreateMonitor(&params)

		Expect(err).ShouldNot(HaveOccurred(), "CreateMonitor should not have returned an error")
		Expect(m).ToNot(BeNil())

		Expect(*m.Probe).To(Equal(probe))
		Expect(m.Label).To(Equal(label))
		Expect(m.Min).To(Equal(min))
		Expect(m.Max).To(Equal(max))
		Expect(m.Scale).To(Equal(scale))
		Expect(m.ID).ToNot(BeEmpty())
		Expect(m.ID).ToNot(Equal(fake))
		Expect(len(m.ID)).To(Equal(11))
	})
	It("should return all monitors", func() {
		probe := int32(0)
		params := monitors.GetMonitorsParams{
			Probe: &probe,
		}
		m, err := backend.GetMonitors(&params)

		Expect(err).ShouldNot(HaveOccurred(), "GetMonitors should not have returned an error")
		Expect(m).To(BeNil())
	})
})
