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
	period time.Duration
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
			period: time.Second * 1,
		},
	)
}

func (r *temperatureLogger) getPeriod() time.Duration {
	return r.period
}

func (r *temperatureLogger) setPeriod(period time.Duration)  {
	r.period = period
}

// GetName returns a human readable name for this background task
func (r *temperatureLogger) GetName() string {
	return "temperatureLogger"
}

// Start performs initialization before the first tick
<<<<<<< Updated upstream
func (r *temperatureLogger) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
	r.tick()
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
=======
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
>>>>>>> Stashed changes
	if err != nil {
		log.Error(err.Error())
	}

<<<<<<< Updated upstream
	err = r.logTemperatureMetrics(readings)
	if err != nil {
		log.Error(err.Error())
=======
	if err := o.logTemperatureMetrics(readings); err != nil {
		return err
>>>>>>> Stashed changes
	}

	return nil
}

<<<<<<< Updated upstream
func (r *temperatureLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
	log.Debug("action=method_entry numProbes=%d", r.reader.GetNumProbes())
	defer log.Debug("action=method_exit")

	readings := make([]*models.TemperatureReading, 0)
	for i := int32(0); i < r.reader.GetNumProbes(); i++ {
=======
func (o *temperatureLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
	nProbes := len(*o.probes)

	log.WithField("numProbes", nProbes).Info("collecting temperature metrics")
	readings := make([]*models.TemperatureReading, nProbes)

	for _, i := range(*o.probes) {
>>>>>>> Stashed changes
		log.Debugf("action=iterate probe=%d", i)
		reading := models.TemperatureReading{}
		if err := r.reader.GetTemperatureReading(i, &reading); err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}
	return readings, nil
}

<<<<<<< Updated upstream
func (r *temperatureLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.Debugf("action=method_entry numReadings=%d", len(readings))
	defer log.Debug("action=method_exit")
=======
func (o *temperatureLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.WithField("numReadings", len(readings)).Info("logging temperature metrics")
>>>>>>> Stashed changes

	for _, reading := range readings {
		log.WithField("reading", reading).Info("logging temperature reading")
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
