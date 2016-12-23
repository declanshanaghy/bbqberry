package daemon

import (
	"github.com/Polarishq/middleware/framework/log"
	"time"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	ticks int
}

// NewCommander creates a new Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander() *Commander {
	log.Debug("action=start")
	defer log.Debug("action=done")
	return &Commander{}
}

// StartBackground starts the commander in the background
func (cmdr *Commander) StartBackground() error {
	log.Debug("action=start")
	defer log.Debug("action=done")
	return cmdr.startBackground(cmdr)
}

func (cmdr *Commander) getPeriod() time.Duration {
	return time.Second
}

// Start performs initialization before the first tick
func (cmdr *Commander) start() {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")
}

// Stop performs cleanup when the goroutine is exiting
func (cmdr *Commander) stop() {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")
}

// Tick executes on a ticker schedule
func (cmdr *Commander) tick() bool {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")

	cmdr.ticks++
	if cmdr.ticks >= 5 {
		return false
	}

	return true
}
