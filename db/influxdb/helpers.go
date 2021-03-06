package influxdb

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/models"
)

// QueryAverageTemperature retrieves the average temperature over the given period for the requested probe from InfluxDB
func QueryAverageTemperature(period time.Duration, probe int32) (*models.TemperatureReading, error) {
	log.Debug("action=method_entry period=%s probe=%d", period, probe)
	defer log.Debug("action=method_exit")

	t := time.Now().Add(period * -1)
	ts := t.Format(time.RFC3339)

	tq := `SELECT MEAN(Celsius) as Celsius, MEAN(Fahrenheit) as Fahrenheit, MEAN(Kelvin) as Kelvin FROM temp WHERE Probe='%d' AND time > '%s'`
	q := fmt.Sprintf(tq, probe, ts)

	response, err := Query(q)
	if err != nil {
		return nil, err
	}

	toF := func(v interface{}) (int32, error) {
		s := fmt.Sprintf("%v", v)
		f, err := strconv.ParseFloat(s, 32)
		if err == nil {
			return int32(f), nil
		}
		return math.MaxInt32, err
	}

	if len(response.Results) > 0 && len(response.Results[0].Series) > 0 {
		r := response.Results[0]
		values := r.Series[0].Values[0]

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
			Probe:      &probe,
			Celsius:    &c,
			Fahrenheit: &f,
			Kelvin:     &k,
		}

		return &reading, nil
	}

	return nil, errors.New("No results returned from InfluxDB")
}
