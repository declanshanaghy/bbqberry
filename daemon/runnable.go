package daemon

import (
	"errors"
	"sync"
	"time"

	"fmt"
	"math"

	"github.com/Polarishq/middleware/framework/log"
)

// Tickable objects are executed in the background by a runner
type Tickable interface {
	// start is called when the goroutine is starting up, before the first tick
	start()
	// tick is called on a time.Ticker period. Returning false will cause the goroutine to exit
	tick() bool
	// stop is called when the goroutine is exiting
	stop()

	// getPeriod will be called by the runner. The time.Duration returned will be used as the period between calls to tick
	getPeriod() time.Duration

	// GetName is used for a human to identify background tasks
	GetName() string
}

// runner represents a single background goroutine
type runner struct {
	running  bool
	ch       chan bool
	wg       *sync.WaitGroup
	tickable Tickable
}

// StartBackground starts the main loop of the runner resulting the the given
// Tickable being executed on the default Ticker schedule
func (r *runner) startBackground(tickable Tickable) error {
	log.Debug("action=start")
	defer log.Debug("action=done")

	if r.running {
		return errors.New("Cannot execute StartBackground. Already running")
	}

	// Initialize control structures
	wg := &sync.WaitGroup{}
	wg.Add(1)
	r.running = true
	r.ch = make(chan bool)
	r.wg = wg

	// Launch background goroutine
	r.tickable = tickable
	go r.loop(tickable)

	return nil
}

// loop executes the main loop of this runner, calling the Tickable every second
func (r *runner) loop(tickable Tickable) {
	log.Info("action=start")
	defer r.wg.Done()
	defer log.Info("action=done")

	// Ensure running flag is set
	r.running = true

	// Start the tickable before entering the loop
	tickable.start()

	ticker := time.NewTicker(tickable.getPeriod())
	for r.running {
		select {
		case r.running = <-r.ch:
			log.Infof("action=rx running=%t", r.running)
		case <-ticker.C:
			log.Debugf("action=timeout")
			r.running = tickable.tick()
		}
	}

	// Stop the tickable before exiting
	tickable.stop()

	// Ensure running flag is reset
	r.running = false
}

// StopBackground causes the background goroutine to exit
func (r *runner) StopBackground() error {
	log.Debug("action=start")
	defer log.Debug("action=done")

	if !r.running {
		return errors.New("Cannot execute StopBackground. Not running")
	}

	// Close the run channel which will cause the runner loop to exit
	close(r.ch)

	// Wait at least 1 second for the loop to exit
	timeout := math.Max(float64(r.tickable.getPeriod())*1.5, float64(time.Second.Nanoseconds()))
	timedOut := waitTimeout(r.wg, time.Duration(int64(timeout)))
	if timedOut {
		return errors.New(fmt.Sprintf("Timed out waiting for background task to exit: name=%s", r.tickable.GetName()))
	}

	r.tickable = nil
	return nil
}

// WaitTimeout waits for the WaitGroup for the specified Duration.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	log.Infof("action=WaitTimeout timeout=%d", timeout)

	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
		log.Info("action=WaitTimeout status=wg_exited")
	}()
	select {
	case <-c:
		log.Info("action=WaitTimeout status=exit_success")
		return false // completed normally
	case <-time.After(timeout):
		log.Info("action=WaitTimeout status=exit_failed")
		return true // timed out
	}
}
