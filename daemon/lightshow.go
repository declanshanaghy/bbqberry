package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
)

// lightShow displays fancy colors
type lightShow struct {
	performer lightShowPerformance
}

type lightShowPerformance interface {
	init()
	tick()
}

// newLightShow creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newLightShow() Runnable {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
	return newRunnable(
		&lightShow{
			performer:  &simpleShifter{},
		},
	)
}

func (r *lightShow) getPeriod() time.Duration {
	return time.Second
}

// GetName returns a human readable name for this background task
func (r *lightShow) GetName() string {
	return "lightshow"
}

// Start performs initialization before the first tick
func (r *lightShow) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	r.performer.init()
}

// Stop performs cleanup when the goroutine is exiting
func (r *lightShow) stop() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
}

// Tick executes on a ticker schedule
func (r *lightShow) tick() bool {
	log.WithField("action", "method_entry").Info("updating lights")
	defer log.Debug("action=method_exit")

	r.performer.tick()

	return true
}