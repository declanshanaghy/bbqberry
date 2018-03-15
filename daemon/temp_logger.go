package daemon

import (
	"time"

	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
)

// temperatureLogger collects and logs temperature metrics
type temperatureLogger struct {
	reader hardware.TemperatureReader
}

// newTemperatureLogger creates a new temperatureLogger instance which can be
// run in the background to collect and log temperature metrics
func newTemperatureLogger() Runnable {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return newRunnable(
		&temperatureLogger{
			reader: hardware.NewTemperatureReader(),
		},
	)
}

func (r *temperatureLogger) getPeriod() time.Duration {
	return time.Second * 1
}

// GetName returns a human readable name for this background task
func (r *temperatureLogger) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
func (r *temperatureLogger) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
}

// Stop performs cleanup when the goroutine is exiting
func (r *temperatureLogger) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")
}

// Tick executes on a ticker schedule
func (r *temperatureLogger) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	readings, err := r.collectTemperatureMetrics()
	if err != nil {
		log.Error(err.Error())
	}

	err = r.logTemperatureMetrics(readings)
	if err != nil {
		log.Error(err.Error())
	}

	return true
}

func (r *temperatureLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
	log.Debug("action=method_entry numProbes=%d", r.reader.GetNumProbes())
	defer log.Debug("action=method_exit")

	readings := make([]*models.TemperatureReading, 0)
	for i := int32(0); i < r.reader.GetNumProbes(); i++ {
		log.Debugf("action=iterate probe=%d", i)
		reading := models.TemperatureReading{}
		if err := r.reader.GetTemperatureReading(i, &reading); err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}
	return readings, nil
}

func (r *temperatureLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.Debugf("action=method_entry numReadings=%d", len(readings))
	defer log.Debug("action=method_exit")

	for _, reading := range readings {
		tags := map[string]string{
			"Probe": fmt.Sprintf("%d", *reading.Probe),
		}
		fields := map[string]interface{}{
			"Celsius":    *reading.Celsius,
			"Fahrenheit": *reading.Fahrenheit,
			"Kelvin":     *reading.Kelvin,
			"Warning":    reading.Warning,
		}
		if _, err := influxdb.WritePoint("temp", tags, fields); err != nil {
			return err
		}
	}

	return nil
}
