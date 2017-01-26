package framework

import (
	"time"
	"fmt"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"strconv"
	"math"
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/Polarishq/middleware/framework/log"
	"errors"
)


// QueryAverageTemperature retrieves the average temperature over the given period for the requested probe from InfluxDB
func QueryAverageTemperature(period time.Duration, probe int32) (*models.TemperatureReading, error) {
	log.Debug("action=method_entry period=%v probe=%d", period, probe)
	defer log.Debug("action=method_exit")

	t := time.Now().Add(period * -1)
	ts := t.Format(time.RFC3339)

	tq := `SELECT MEAN(Celsius) as Celsius, MEAN(Fahrenheit) as Fahrenheit, MEAN(Kelvin) as Kelvin FROM temp WHERE Probe='%d' AND time > '%s'`
	q := fmt.Sprintf(tq, probe, ts)

	response, err := influxdb.Query(q)
	if err != nil {
		return nil, err
	}

	toF := func(v interface{}) (float32, error) {
		s := fmt.Sprintf("%v", v)
		f, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return float32(f), nil
		}
		return math.MaxFloat32, err
	}

	if len(response.Results) > 0 && len(response.Results[0].Series) > 0 {
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

		reading := models.TemperatureReading{
			Probe: &probe,
			Celsius: &c,
			Fahrenheit: &f,
			Kelvin: &k,
		}

		return &reading, nil
	} else {
		return nil, errors.New("No results returned from InfluxDB")
	}
}
