package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"fmt"
)

// temperatureLogger collects and logs temperature metrics
type temperatureLogger struct {
	runner
	reader hardware.TemperatureReader
}

// newTemperatureLogger creates a new temperatureLogger instance which can be
// run in the background to collect and log temperature metrics
func newTemperatureLogger() *temperatureLogger {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return &temperatureLogger{
		reader: hardware.NewTemperatureReader(),
	}
}

// StartBackground starts the commander in the background
func (tl *temperatureLogger) StartBackground() error {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return tl.runner.startBackground(tl)
}

func (tl *temperatureLogger) getPeriod() time.Duration {
	return time.Second * 1
}

// GetName returns a human readable name for this background task
func (tl *temperatureLogger) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
func (tl *temperatureLogger) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
}

// Stop performs cleanup when the goroutine is exiting
func (tl *temperatureLogger) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")
}

// Tick executes on a ticker schedule
func (tl *temperatureLogger) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	readings, err := tl.collectTemperatureMetrics()
	if err != nil {
		log.Error(err.Error())
	}

	err = tl.logTemperatureMetrics(readings)
	if err != nil {
		log.Error(err.Error())
	}

	return true
}

func (tl *temperatureLogger) collectTemperatureMetrics() (*models.TemperatureReadings, error) {
	log.Debug("action=method_entry numProbes=%d", tl.reader.GetNumProbes())
	defer log.Debug("action=method_exit")

	readings := models.TemperatureReadings{}
	for i := int32(1); i <= tl.reader.GetNumProbes(); i++ {
		log.Debugf("action=iterate probe=%d", i)
		reading := models.TemperatureReading{}
		if err := tl.reader.GetTemperatureReading(i, &reading); err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}
	return &readings, nil
}

func (tl *temperatureLogger) logTemperatureMetrics(readings *models.TemperatureReadings) error {
	log.Debugf("action=method_entry numReadings=%d", len(*readings))
	defer log.Debug("action=method_exit")

	for _, reading := range *readings {
		tags := map[string]string{
			"Probe": fmt.Sprintf("%d", *reading.Probe),
		}
		fields := map[string]interface{}{
			"Celsius": *reading.Celsius,
			"Fahrenheit": *reading.Fahrenheit,
			"Kelvin": *reading.Kelvin,
			"Warning": reading.Warning,
		}
		if _, err := influxdb.WritePoint("temp", tags, fields); err != nil {
			return err
		}
	}

	return nil
}
