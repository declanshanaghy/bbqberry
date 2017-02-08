package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	tempLogger    *temperatureLogger
	tempIndicator *temperatureIndicator
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander() *Commander {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return &Commander{
		tempLogger:    newTemperatureLogger(),
		tempIndicator: newTemperatureIndicator(),
	}
}

// StartBackground starts the commander in the background
func (cmdr *Commander) StartBackground() error {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")

	return cmdr.startBackground(cmdr)
}

func (cmdr *Commander) getPeriod() time.Duration {
	return time.Second * 10
}

// GetName returns a human readable name for this background task
func (cmdr *Commander) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
func (cmdr *Commander) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")

	cmdr.tempLogger.StartBackground()
	cmdr.tempIndicator.StartBackground()
}

// Stop performs cleanup when the goroutine is exiting
func (cmdr *Commander) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")

	cmdr.tempLogger.StopBackground()
	cmdr.tempIndicator.StopBackground()
}

// Tick executes on a ticker schedule
func (cmdr *Commander) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	return true
}
