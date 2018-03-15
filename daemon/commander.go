package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	period   		time.Duration
	tempLogger   	Runnable
	tempIndicator	Runnable
	lightShow 		RunnableTicker
	strip      	  	hardware.WS2801
	options			*framework.CmdOptions
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander(options *framework.CmdOptions) *Commander {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	c := Commander{
		strip: hardware.NewStrandController(),
		options: options,
	}
	c.runner.tickable = &c
	return &c
}

func (r *Commander) getPeriod() time.Duration {
	return time.Second * 10
}

func (r *Commander) setPeriod(period time.Duration)  {
	r.period = period
}

// GetName returns a human readable name for this background task
func (r *Commander) GetName() string {
	return "commander"
}

// Start performs initialization before the first tick
func (r *Commander) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")

	r.tempLogger = newTemperatureLogger()
	if r.options.TemperatureLoggerEnabled {
		r.EnableTemperatureLogger()
	}

	r.tempIndicator = newTemperatureIndicator()
	if r.options.TemperatureIndicatorEnabled {
		r.EnableTemperatureIndicator()
	}

	r.lightShow = newSimpleShifter()

	if r.options.LightShowEnabled {
		r.EnableLightShow(r.lightShow.tickable.getPeriod())
	}
}

func (r *Commander) EnableTemperatureLogger() {
	if ! r.tempLogger.IsRunning() {
		if err := r.tempLogger.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) EnableLightShow(period time.Duration) {
	r.lightShow.tickable.setPeriod(period)

	if ! r.lightShow.runnable.IsRunning() {
		if err := r.lightShow.runnable.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) EnableTemperatureIndicator() {
	if ! r.tempIndicator.IsRunning() {
		if err := r.tempIndicator.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) DisableLights() {
	r.disableLightShow()
	r.disableTemperatureIndicator()
}

func (r *Commander) disableLightShow() {
	if ! r.lightShow.runnable.IsRunning() {
		if err := r.lightShow.runnable.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) disableTemperatureIndicator() {
	if ! r.tempIndicator.IsRunning() {
		if err := r.tempIndicator.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) DisableTemperatureLogger() {
	if ! r.tempLogger.IsRunning() {
		if err := r.tempLogger.StopBackground(); err != nil {
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
