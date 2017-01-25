package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
)

// temperatureIndicator collects and logs temperature metrics
type temperatureIndicator struct {
	runner
	reader	hardware.TemperatureReader
	strip	hardware.WS2801
	errorCount	int
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newTemperatureIndicator() *temperatureIndicator {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return &temperatureIndicator{
		reader: hardware.NewTemperatureReader(),
		strip: hardware.NewStrandController(),
	}
}

// StartBackground starts the commander in the background
func (ti *temperatureIndicator) StartBackground() error {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return ti.runner.startBackground(ti)
}

func (ti *temperatureIndicator) getPeriod() time.Duration {
	return time.Second * 10
}

// GetName returns a human readable name for this background task
func (ti *temperatureIndicator) GetName() string {
	return "temperatureIndicator"
}

// Start performs initialization before the first tick
func (ti *temperatureIndicator) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
}

// Stop performs cleanup when the goroutine is exiting
func (ti *temperatureIndicator) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")
}

// Tick executes on a ticker schedule
func (ti *temperatureIndicator) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	avg, err := framework.QueryAverageTemperature(ti.getPeriod(), framework.Constants.Hardware.AmbientProbeNumber)
	if err != nil {
		log.Error(err.Error())
		return true
	}

	log.Infof("avg=%0.2f", *avg.Fahrenheit)

	if err := ti.strip.SetAllPixels(0xFF0000); err != nil {
		log.Error(err.Error())
	}
	if err := ti.strip.Update(); err != nil {
		log.Error(err.Error())
	}

	return true
}

