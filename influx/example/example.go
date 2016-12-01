package influx_example

import (
	"time"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/declanshanaghy/bbqberry/influx"
	"math/rand"
)

func WriteExamplePoint(c client.Client) (*client.Point, error) {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influx.DB,
		Precision: "s",
	})

	if err != nil {
		log.Error(err)
	}

	rand.Seed(time.Now().UnixNano())
	r := 100.0
	i := float64(int(rand.Float64() * 100.0)) + rand.Float64()
	r -= i
	s := float64(int(rand.Float64() * 100.0) % int(r-1)) + rand.Float64()
	r -= s
	u := r

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   i,
		"system": s,
		"user":   u,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Error(err)
	}

	bp.AddPoint(pt)

	// Write the batch
	if err = c.Write(bp); err != nil {
		return nil, err
	} else {
		return pt, nil
	}
}
