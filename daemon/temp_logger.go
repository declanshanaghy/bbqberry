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
	period	time.Duration
	reader	hardware.TemperatureReader
	probes	*[]int32
}

// newTemperatureLogger creates a new temperatureLogger instance which can be
// run in the background to collect and log temperature metrics
func newTemperatureLoggerRunnable() Runnable {
	return newRunnable(newTemperatureLogger())
}

func newTemperatureLogger() *temperatureLogger {
	reader := hardware.NewTemperatureReader()
	probes := reader.GetEnabledPobes()

	return &temperatureLogger{
		reader: reader,
		probes: probes,
		period: time.Second,
	}
}

func (o *temperatureLogger) getPeriod() time.Duration {
	return o.period
}

func (o *temperatureLogger) setPeriod(period time.Duration)  {
	o.period = period
}

// GetName returns a human readable name for this background task
func (o *temperatureLogger) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
func (o *temperatureLogger) start() error {
	o.probes = o.reader.GetEnabledPobes()
	log.WithField("probes", len(*o.probes)).Infof("Found enabled probes")

	return o.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (o *temperatureLogger) stop() error {
	return nil
}

// Tick executes on a ticker schedule
func (o *temperatureLogger) tick() error {
	readings, err := o.collectTemperatureMetrics()
	if err != nil {
		return err
	}

	err = o.logTemperatureMetrics(readings)
	if err != nil {
		return err
	}

	return nil
}

func (o *temperatureLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
	nProbes := len(*o.probes)
	log.WithField("nProbes", nProbes).Debug("collecting temperature readings")

	readings := make([]*models.TemperatureReading, 0)
	for _, i := range(*o.probes) {
		reading := models.TemperatureReading{}
		if err := o.reader.GetTemperatureReading(i, &reading); err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}
	return readings, nil
}

func (o *temperatureLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.WithField("numReadings", len(readings)).Debug("logging temperature metrics")

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