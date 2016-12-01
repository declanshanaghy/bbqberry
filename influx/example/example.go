package influx_example

import (
	"time"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/declanshanaghy/bbqberry/influx"
)

func WriteExamplePoint(c client.Client, name string, tags map[string]string,
	fields map[string]interface{}) (*client.Point, error) {
	
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influx.DB,
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
	} else {
		return pt, nil
	}
}
