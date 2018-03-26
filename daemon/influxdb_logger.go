package daemon

import (
	"time"

	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/framework"
)

// influxDBLogger collects and logs temperature metrics
type influxDBLogger struct {
	period	time.Duration
	reader	hardware.TemperatureReader
	probes	*[]int32
}

// newInfluxDBLogger creates a new influxDBLogger instance which can be
// run in the background to collect and log temperature metrics
func newInfluxDBLoggerRunnable() Runnable {
	return newRunnable(newInfluxDBLogger())
}

func newInfluxDBLogger() *influxDBLogger {
	reader := hardware.NewTemperatureReader()
	probes := reader.GetEnabledPobes()

	return &influxDBLogger{
		reader: reader,
		probes: probes,
		period: time.Second,
	}
}

func (o *influxDBLogger) getPeriod() time.Duration {
	return o.period
}

func (o *influxDBLogger) setPeriod(period time.Duration)  {
	o.period = period
}

// GetName returns a human readable name for this background task
func (o *influxDBLogger) GetName() string {
	return "influxDBLogger"
}

// Start performs initialization before the first tick
func (o *influxDBLogger) start() error {
	o.probes = o.reader.GetEnabledPobes()
	log.WithField("probes", len(*o.probes)).Infof("Found enabled probes")

	return o.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (o *influxDBLogger) stop() error {
	return nil
}

// Tick executes on a ticker schedule
func (o *influxDBLogger) tick() error {
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

func (o *influxDBLogger) collectTemperatureMetrics() ([]*models.TemperatureReading, error) {
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

func (o *influxDBLogger) logTemperatureMetrics(readings []*models.TemperatureReading) error {
	log.WithField("numReadings", len(readings)).Debug("logging temperature metrics")

	for _, reading := range readings {
		probe := framework.Constants.Hardware.Probes[*reading.Probe]
		if err := o.writeToInflux(reading, probe); err != nil {
			log.WithField("err", err).Error("Unable to write to InfluxDB")
		}
	}

	return nil
}

func (o *influxDBLogger) writeToInflux(reading *models.TemperatureReading, probe *models.TemperatureProbe) error {
	tags := map[string]string{
		"Probe": fmt.Sprintf("%d", *reading.Probe),
		"Label": *probe.Label,
	}
	fields := map[string]interface{}{
		"Celsius":    *reading.Celsius,
		"Fahrenheit": *reading.Fahrenheit,
		"Kelvin":     *reading.Kelvin,
		"Warning":    reading.Warning,
	}

	log.WithFields(log.Fields{
		"Probe": probe.Label,
		"Fahrenheit": *reading.Fahrenheit,
	}).Debugf("Logging temperature to InfluxDB")

	if _, err := influxdb.WritePoint("temp", tags, fields); err != nil {
		return err
	}
	return nil
}