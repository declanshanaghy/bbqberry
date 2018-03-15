package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"os"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	tempLogger   	Runnable
	tempIndicator	Runnable
	lightShow 		Runnable
	strip      	  		hardware.WS2801
	//children			[]*runner
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander() Runnable {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return newRunnable(
		&Commander{
			strip: hardware.NewStrandController(),
		},
	)
}

func (r *Commander) getPeriod() time.Duration {
	return time.Second * 10
}

// GetName returns a human readable name for this background task
func (r *Commander) GetName() string {
	return "commander"
}

//func (r *Commander) startChild(child *runner) {
//	if err := child.StartBackground(); err != nil {
//		log.Error(err.Error())
//	}
//}

// Start performs initialization before the first tick
func (r *Commander) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")

	if v := os.Getenv("LOGGER"); v == "true" {
		r.tempLogger = newTemperatureLogger()
		if err := r.tempLogger.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	if v := os.Getenv("TEMP_INDICATOR"); v == "true" {
		r.tempIndicator = newTemperatureIndicator()
		if err := r.tempIndicator.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	if v := os.Getenv("LIGHT_SHOW"); v == "true" {
		r.lightShow = newLightShow()
		if err := r.lightShow.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

// Stop performs cleanup when the goroutine is exiting
func (r *Commander) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")

	if r.tempLogger != nil {
		if err := r.tempLogger.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	if r.tempIndicator != nil {
		if err := r.tempIndicator.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	if r.lightShow != nil {
		if err := r.lightShow.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	log.Info("Clearing all pixels")
	if err := r.strip.Close(); err != nil {
		log.Error(err.Error())
	}
}

// Tick executes on a ticker schedule
func (r *Commander) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	return true
}
