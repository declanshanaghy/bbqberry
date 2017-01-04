package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
)

// temperatureLogger collects and logs temperature metrics
type temperatureLogger struct {
	runner
	temp hardware.TemperatureArray
}

// NewtemperatureLogger creates a new temperatureLogger instance which can be
// run in the background to collect and log temperature metrics
func newTemperatureLogger() *temperatureLogger {
	log.Debug("action=start")
	defer log.Debug("action=done")
	return &temperatureLogger{
		temp: hardware.NewTemperatureReader(),
	}
}

// StartBackground starts the commander in the background
func (tl *temperatureLogger) StartBackground() error {
	log.Debug("action=start")
	defer log.Debug("action=done")
	return tl.runner.startBackground(tl)
}

func (tl *temperatureLogger) getPeriod() time.Duration {
	return time.Second
}

// GetName returns a human readable name for this background task
func (tl *temperatureLogger) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
func (tl *temperatureLogger) start() {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")
}

// Stop performs cleanup when the goroutine is exiting
func (tl *temperatureLogger) stop() {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")
}

// Tick executes on a ticker schedule
func (tl *temperatureLogger) tick() bool {
	log.Warning("action=Tick")
	defer log.Warning("action=Tick")
	readings := tl.collectTemperatureMetrics()
	tl.logTemperatureMetrics(readings)
	return true
}

func (tl *temperatureLogger) collectTemperatureMetrics() *models.TemperatureReadings {
	log.Infof("action=start numProbes=%d", tl.temp.GetNumProbes())
	readings := models.TemperatureReadings{}
	for i := int32(1); i <= tl.temp.GetNumProbes(); i++ {
		log.Debugf("action=iterate probe=%d", i)
		reading := models.TemperatureReading{}
		if err := tl.temp.GetTemperatureReading(i, &reading); err != nil {
			log.Error(err)
		}
		readings = append(readings, &reading)
	}
	log.Infof("action=done")
	return &readings
}

func (tl *temperatureLogger) logTemperatureMetrics(readings *models.TemperatureReadings) {
	log.Infof("action=start numReadings=%d", len(*readings))
	log.Infof("action=done")
}
