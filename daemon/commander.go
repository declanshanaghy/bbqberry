package daemon

import (
	"errors"
	"sync"
	"time"

	"github.com/Polarishq/middleware/framework/log"
)

// Commander exposes command, control and state of all background processes
type Commander struct {
	running bool
	ch      chan bool
	wg      *sync.WaitGroup
}

// NewCommander creates a new Commander instance which can be used to query and control background processes
// e.g: Temperature logger, temperature monitor
func NewCommander() *Commander {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	return &Commander{
		running: false,
		ch:      make(chan bool),
		wg:      wg,
	}
}

// Run starts the Commander main loop which maintains control of all child goroutines.
// It should be called in a goroutine itself. Calling Exit will cause this method to exit
func (cmdr *Commander) Run() error {
	if cmdr.running {
		return errors.New("Commander is already running")
	}

	log.Info("action=start")
	defer cmdr.wg.Done()
	defer log.Info("action=done")

	t := time.NewTicker(time.Second * 1)
	cmdr.running = true

	for cmdr.running {
		select {
		case cmdr.running = <-cmdr.ch:
			log.Infof("action=rx running=%t", cmdr.running)
		case <-t.C:
			log.Debugf("action=timeout")
		}
	}

	t.Stop()
	return nil
}

// Exit causes the Run method to return
func (cmdr *Commander) Exit() error {
	log.Info("action=start")
	defer log.Info("action=done")

	if !cmdr.running {
		return errors.New("Commander is not running")
	}

	// Close the run channel which will cause Run to exit
	close(cmdr.ch)

	// Wait for Run to actually exit
	cmdr.wg.Wait()
	return nil
}
