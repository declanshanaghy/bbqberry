package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"fmt"
	"github.com/declanshanaghy/bbqberry/framework"
	"strconv"
	"math"
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
	return time.Second * 1
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

	avg, err := ti.getAverageTemp()
	if err != nil {
		log.Error(err.Error())
	}

	log.Infof("avg=%0.2f", *avg.Celsius)

	if err := ti.strip.SetAllPixels(0xFF0000); err != nil {
		log.Error(err.Error())
	}
	if err := ti.strip.Update(); err != nil {
		log.Error(err.Error())
	}

	return true
}

func (ti *temperatureIndicator) getAverageTemp() (*models.TemperatureReading, error) {
	log.Debug("action=method_entry numProbes=%d", ti.reader.GetNumProbes())
	defer log.Debug("action=method_exit")

	t := time.Now().Add(ti.getPeriod() * -1)
	t.Format(time.RFC3339)
	ts := "2017-01-24T23:48:00Z"

	tq := 	"SELECT " +
			"	MEAN(Celsius) as Celsius, " +
			"	MEAN(Fahrenheit) as Fahrenheit, " +
			"	MEAN(Kelvin) as Kelvin " +
			"FROM " +
			"	temp " +
			"WHERE " +
			"	Probe='%d' AND " +
			"	time > '%s'"
	q := fmt.Sprintf(tq, framework.Constants.Hardware.AmbientProbe, ts)
	response, err := influxdb.Query(q)
	if err != nil {
		return nil, err
	}

	reading := models.TemperatureReading{}

	toF := func(v interface{}) (float32, error) {
		if s, ok := v.(string); ok {
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return float32(f), nil
			}
		}
		return math.MaxFloat32, err
	}

	if len(response.Results) > 0 {
		r := response.Results[0]
		values := r.Series[0].Values[0]
		log.Infof("values=%+v", values)

		c, err := toF(values[1])
		if err != nil {
			return nil, err
		}

		f, err := toF(values[2])
		if err != nil {
			return nil, err
		}

		k, err := toF(values[3])
		if err != nil {
			return nil, err
		}

		reading.Celsius = &c
		reading.Fahrenheit = &f
		reading.Kelvin = &k
	}

	return &reading, nil
}
