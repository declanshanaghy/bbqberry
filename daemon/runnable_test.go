package daemon

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/Polarishq/middleware/framework/log"
)

type testingTickable struct {
	runner
	startCalls, stopCalls, tickCalls, nTicks int
	tickReturn                               bool
}

func (t *testingTickable) getPeriod() time.Duration {
	return time.Nanosecond
}

func (t *testingTickable) GetName() string {
	return "testingTickable"
}

func (t *testingTickable) start() {
	log.Info("action=start")
	t.startCalls++
}

func (t *testingTickable) stop() {
	log.Info("action=stop")
	t.stopCalls++
}

func (t *testingTickable) tick() bool {
	t.tickCalls++
	// This is very verbose
	//log.Warningf("action=tick tickCalls=%d", t.tickCalls)

	if t.nTicks >= 0 && t.tickCalls > t.nTicks {
		log.Infof("Tick limit reached nTicks=%d tickReturn=%t", t.nTicks, t.tickReturn)
		return t.tickReturn
	}

	return true
}

func (t *testingTickable) StartBackground() error {
	return t.startBackground(t)
}

var _ = Describe("The runner", func() {
	Context("When given a tickable that exits immediately", func() {
		It("it should exit cleanly", func() {
			t := testingTickable{}
			err := t.StartBackground()
			Expect(err).ToNot(HaveOccurred())
			
			Expect(t.running).To(BeTrue(), "Expected the tickable to be running")

			// This should allow at least 1 tickable execution
			time.Sleep(time.Millisecond)

			// Requesting stop should fail because it exits on its own
			err = t.StopBackground()
			Expect(err).To(HaveOccurred())

			Expect(t.startCalls).Should(Equal(1), "Number of calls to start is incorrect")
			Expect(t.tickCalls).Should(Equal(1), "Number of calls to tick is incorrect")
			Expect(t.stopCalls).Should(Equal(1), "Number of calls to stop is incorrect")
		})
	})
	Context("When given a tickable that never exits", func() {
		It("should refuse to start twice and exit when the channel is closed", func() {
			log.Info("It should refuse to start twice and exit when the channel is closed")
			t := testingTickable{
				nTicks: -1,
			}

			actualTicks := 0
			go func() {
				for _ = 0; 5 > t.tickCalls; {
					//log.Warningf("action=TICK_CHECK tickCalls=%d", t.tickCalls)
					time.Sleep(time.Millisecond)
				}
				actualTicks = t.tickCalls
			}()

			err := t.StartBackground()
			Expect(err).ToNot(HaveOccurred())

			// Secondary test. Try to start it twice
			err = t.StartBackground()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Cannot execute StartBackground. Already running"))

			// This should allow at least 1 tickable execution
			time.Sleep(time.Millisecond)

			err = t.StopBackground()
			Expect(err).ToNot(HaveOccurred())

			Expect(t.startCalls).Should(Equal(1), "Number of calls to start is incorrect")
			Expect(t.stopCalls).Should(Equal(1), "Number of calls to stop is incorrect")

			Expect(t.tickCalls).Should(BeNumerically(">=", actualTicks),
				"Number of calls to tick is less than expected")
		})
	})
})
