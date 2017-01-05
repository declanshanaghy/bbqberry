package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testingTickable struct {
	runner

	// These variables keep track of interactions between this tickable and the runner
	startCalls, stopCalls, tickCalls int

	// These variables control how this tickable behaves

	// max # of times tick can be called before it will indicate that it wants to exit. < 0 means execute forever
	maxTickCalls int

	// duration that each call to tick should sleep
	tickSleep time.Duration
}

func (t *testingTickable) getPeriod() time.Duration {
	return time.Nanosecond
}

func (t *testingTickable) GetName() string {
	return "testingTickable"
}

func (t *testingTickable) start() {
	log.Info("action=method_entry")
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

	if t.tickSleep > 0 {
		log.Infof("action=sleep duration=%d", t.tickSleep)
		time.Sleep(t.tickSleep)
	}

	if t.maxTickCalls >= 0 && t.tickCalls > t.maxTickCalls {
		log.Infof("Tick limit reached nTicks=%d", t.maxTickCalls)
		return false
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
		var t testingTickable

		BeforeEach(func() {
			t = testingTickable{
				maxTickCalls: -1,
			}
		})

		It("should refuse to start twice and timeout when told to stop", func() {
			t.tickSleep = time.Second * 10

			err := t.StartBackground()
			Expect(err).ToNot(HaveOccurred())

			// Secondary test. Try to start it twice
			err = t.StartBackground()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Cannot execute StartBackground. Already running"))

			// This should allow at least 1 tickable execution
			time.Sleep(time.Millisecond)

			err = t.StopBackground()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Timed out waiting for background task to exit: name=testingTickable"))

			Expect(t.startCalls).Should(Equal(1), "Number of calls to start is incorrect")
			Expect(t.tickCalls).Should(Equal(1), "Number of calls to tick is incorrect")
			Expect(t.stopCalls).Should(Equal(0), "Number of calls to stop is incorrect")
		})
		It("should exit when the channel is closed", func() {
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
