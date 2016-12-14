package util

import (
	// "github.com/influxdata/influxdb/client"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/influxdb"
	"github.com/influxdata/influxdb/client/v2"
)

// WriteHealthCheck writes a health check metric for the service into influxdb
func WriteHealthCheck() (*client.Point, error) {
	c, err := influxdb.NewHTTPClient()
	if err != nil {
		return nil, err
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influxdb.Cfg.Database,
		Precision: "s",
	})
	if err != nil {
		return nil, err
	}

	tags := map[string]string{"service": framework.Constants.ServiceName}
	fields := map[string]interface{}{
		"version": framework.Constants.Version,
	}

	pt, err := client.NewPoint("health", tags, fields, time.Now())
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
