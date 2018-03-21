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
	start() error
	// tick is called on a time.Ticker period. Returning false will cause the goroutine to exit
	tick() error
	// stop is called when the goroutine is exiting
	stop() error

	// getPeriod will be called by the runner. The time.Duration returned
	// will be used as the period between calls to tick
	getPeriod() time.Duration

	// setPeriod is used to change the period between calls to tick
	setPeriod(time.Duration)

	// GetName returns a human readable name for this background task
	GetName() string
}

// Runnable objects are executed in the background by a runner
type Runnable interface {
	// StartBackground starts the runnable in a goroutine
	StartBackground() error
	// StopBackground stops the background goroutine
	StopBackground() error
	// IsRunning determines if the main loop is executing
	IsRunning() bool
}

type RunnableTicker struct {
	runnable Runnable
	tickable Tickable
}

func newRunnableTicker(tickable Tickable) RunnableTicker {
	r := newRunnable(tickable)
	return RunnableTicker{r, tickable}
}

// runner represents a single background goroutine
type runner struct {
	running  bool
	ch       chan bool
	wg       *sync.WaitGroup
	tickable Tickable
}

func newRunnable(tickable Tickable) Runnable {
	return &runner{
		tickable: tickable,
	}
}

// IsRunning returns the internal state representing if the main loop is running or not.
func (r *runner) IsRunning() bool {
	return r.running
}

// startBackground starts the main loop of the runner resulting the the given
// Tickable being executed on the default Ticker schedule
func (r *runner) StartBackground() error {
	log.Debugf("action=method_entry name=%s", r.tickable.GetName())
	defer log.Debugf("action=method_exit name=%s", r.tickable.GetName())

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
	go r.loop()

	return nil
}

// loop executes the main loop of this runner, calling it's tick method once per period
func (r *runner) loop() {
	log.Debugf("action=method_entry name=%s", r.tickable.GetName())
	defer r.wg.Done()
	defer log.Debugf("action=method_exit name=%s", r.tickable.GetName())

	// Ensure running flag is set
	r.running = true

	// Start the tickable before entering the loop
<<<<<<< Updated upstream
	r.tickable.start()

	for r.running {
		ticker := time.NewTicker(r.tickable.getPeriod())
=======
	if err := o.tickable.start(); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(o.tickable.getPeriod())
	var ticker_err error = nil

	for o.running {
		ticker = time.NewTicker(o.tickable.getPeriod())
>>>>>>> Stashed changes
		select {
		case r.running = <-r.ch:
			log.Debugf("action=rx running=%t", r.running)
		case <-ticker.C:
<<<<<<< Updated upstream
			log.Debugf("action=timeout")
			r.running = r.tickable.tick()
=======
			//log.WithFields(log.Fields{
			//	"name": o.tickable.GetName(),
			//	"period": o.tickable.getPeriod(),
			//}).Debugf("tick")
			tick_err := o.tickable.tick()
			if tick_err != nil {
				log.Error(tick_err)
				o.running = false
			}
>>>>>>> Stashed changes
		}
	}

	// Stop the tickable before exiting
<<<<<<< Updated upstream
	r.tickable.stop()
=======
	if ticker_err != nil {
		if err := o.tickable.stop(); err != nil {
			panic(err)
		}
	}
>>>>>>> Stashed changes

	// Ensure running flag is reset
	r.running = false
}

// StopBackground causes the background goroutine to exit
func (r *runner) StopBackground() error {
	log.Debugf("action=method_entry name=%s", r.tickable.GetName())
	defer log.Debugf("action=method_exit name=%s", r.tickable.GetName())

	if !r.running {
		return errors.New("Cannot execute StopBackground. Not running")
	}

	// Close the run channel which will cause the runner loop to exit
	close(r.ch)

	// Wait at least this amount of time for the loop to exit
	minWait := float64(time.Second.Nanoseconds()) * 0.01
	timeout := math.Max(float64(r.tickable.getPeriod())*1.5, minWait)
	timedOut := waitTimeout(r.wg, time.Duration(int64(timeout)))
	if timedOut {
		return fmt.Errorf("Timed out waiting for background task to exit: name=%s", r.tickable.GetName())
	}

	r.tickable = nil
	return nil
}

// WaitTimeout waits for the WaitGroup for the specified Duration.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")

	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
		log.Debug("action=WaitTimeout status=wg_exited")
	}()
	select {
	case <-c:
		log.Debug("action=WaitTimeout status=exit_success")
		return false // completed normally
	case <-time.After(timeout):
		log.Error("action=WaitTimeout status=exit_failed")
		return true // timed out
	}
}
