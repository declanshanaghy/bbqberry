package daemon

import (
	"errors"
	"sync"
	"time"

	"fmt"
	"math"

	"github.com/Polarishq/middleware/framework/log"
)

// tickable objects are executed in the background by a runner
type tickable interface {
	// start is called when the goroutine is starting up, before the first tick
	start() error
	// tick is called on a time.Ticker period. Returning false will cause the goroutine to exit
	tick() error
	// stop is called when the goroutine is exiting
	stop() error

	// getPeriod will be called by the runneo. The time.Duration returned
	// will be used as the period between calls to tick
	getPeriod() time.Duration

	// setPeriod is used to change the period between calls to tick
	setPeriod(time.Duration)

	// GetName returns a human readable name for this background task
	GetName() string
}

// basicTickable provides common functionality for tickable implementations
type basicTickable struct {
	Period	time.Duration
}

func (o *basicTickable) getPeriod() time.Duration {
	return o.Period
}

func (o *basicTickable) setPeriod(period time.Duration)  {
	o.Period = period
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
	tickable tickable
}

func newRunnableTicker(tickable tickable) RunnableTicker {
	r := newRunnable(tickable)
	return RunnableTicker{r, tickable}
}

// runner represents a single background goroutine
type runner struct {
	running  bool
	ch       chan bool
	wg       *sync.WaitGroup
	tickable tickable
}

func newRunnable(tickable tickable) Runnable {
	return &runner{
		tickable: tickable,
	}
}

// IsRunning returns the internal state representing if the main loop is running or not.
func (o *runner) IsRunning() bool {
	return o.running
}

// startBackground starts the main loop of the runner resulting the the given
// tickable being executed on the default Ticker schedule
func (o *runner) StartBackground() error {
	log.WithField("name", o.tickable.GetName()).Infof("Starting in background")

	if o.running {
		return errors.New("Cannot execute StartBackground. Already running")
	}

	// Initialize control structures
	wg := &sync.WaitGroup{}
	wg.Add(1)
	o.running = true
	o.ch = make(chan bool)
	o.wg = wg

	// Launch background goroutine
	go o.loop()

	return nil
}

// loop executes the main loop of this runner, calling it's tick method once per period
func (o *runner) loop() {
	log.WithField("name", o.tickable.GetName()).Infof("Starting loop")
	defer o.wg.Done()
	defer log.WithField("name", o.tickable.GetName()).Infof("Exiting loop")

	// Ensure running flag is set
	o.running = true

	// Start the tickable before entering the loop
	if err := o.tickable.start(); err != nil {
		panic(err)
	}

	ticker := time.NewTicker(o.tickable.getPeriod())
	var tickerErr error = nil

	for o.running {
		ticker = time.NewTicker(o.tickable.getPeriod())
		select {
		case o.running = <-o.ch:
			log.WithFields(log.Fields{
				"name": o.tickable.GetName(),
				"period": o.tickable.getPeriod(),
			}).Debug("idle")
		case <-ticker.C:
			//log.WithFields(log.Fields{
			//	"name": o.tickable.GetName(),
			//	"period": o.tickable.getPeriod(),
			//}).Debugf("tick")
			tickerErr = o.tickable.tick()
			if tickerErr != nil {
				log.Error(tickerErr)
				o.running = false
			}
		}
	}

	// Stop the tickable before exiting
	if tickerErr == nil {
		if err := o.tickable.stop(); err != nil {
			log.Error(err)
		}
	}

	// Ensure running flag is reset
	o.running = false
}

// StopBackground causes the background goroutine to exit
func (o *runner) StopBackground() error {
	if !o.running {
		return fmt.Errorf("Cannot execute StopBackground on %s. Not running", o.tickable.GetName())
	}

	log.WithField("name", o.tickable.GetName()).Infof("Stopping background routine")
	defer log.WithField("name", o.tickable.GetName()).Infof("Stopping background routine succeeded")

	// Close the run channel which will cause the runner loop to exit
	close(o.ch)

	// Wait at least this amount of time for the loop to exit
	minWait := float64(time.Second.Nanoseconds()) * 0.01
	timeout := math.Max(float64(o.tickable.getPeriod())*50, minWait)
	timedOut := waitTimeout(o.wg, time.Duration(int64(timeout)))
	if timedOut {
		return fmt.Errorf("Timed out waiting for background task to exit: name=%s", o.tickable.GetName())
	}

	return nil
}

// WaitTimeout waits for the WaitGroup for the specified Duration.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
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