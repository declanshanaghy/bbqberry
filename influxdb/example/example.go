package example

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"github.com/influxdata/influxdb/client/v2"
)

// WriteExamplePoint writes a single point to the influxDB
func WriteExamplePoint(c client.Client, name string, tags map[string]string,
	fields map[string]interface{}) (*client.Point, error) {

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxdb.Settings.Database,
		Precision: "s",
	})

	if err != nil {
		log.Error(err)
	}

	pt, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		log.Error(err)
	}

	bp.AddPoint(pt)

	// Write the batch and check for an error
	if err = c.Write(bp); err != nil {
		return nil, err
	}

	return pt, nil
}
