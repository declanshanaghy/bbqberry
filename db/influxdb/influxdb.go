package influxdb

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/util"
	clientv1 "github.com/influxdata/influxdb/client"
	"github.com/influxdata/influxdb/client/v2"
)

// Settings holds all pertient connection parameters for InfluxDB
var Settings *influxDBSettings

type influxDBSettings struct {
	clientv1.Config

	Database string
	Host     string
	HTTPPort string
}

func init() {
	LoadConfig()
}

// LoadConfig (re)loads the influxDB config so a connection can be initialized
func LoadConfig() {
	database := os.Getenv("INFLUXDB")
	if database == "" {
		database = framework.DefaultDB
	}

	host := os.Getenv("INFLUXDB_HOST")
	if host == "" {
		host = "influxdb"
	}
	HTTPPort := os.Getenv("INFLUXDB_PORT_HTTP")
	if HTTPPort == "" {
		HTTPPort = strconv.Itoa(clientv1.DefaultPort)
	}
	URL, err := clientv1.ParseConnectionString(net.JoinHostPort(host, HTTPPort), false)
	if err != nil {
		panic(err)
	}

	timeout := util.GetEnvMillisAsDuration("DB_TIMEOUT_MILLIS", 5000)
	username := os.Getenv("INFLUXDB_USERNAME")
	password := os.Getenv("INFLUXDB_PASSWORD")

	Settings = &influxDBSettings{
		Config: clientv1.Config{
			URL:      URL,
			Username: username,
			Password: password,
			Timeout:  timeout,
		},
		Database: database,
		Host:     host,
		HTTPPort: HTTPPort,
	}

	log.WithFields(log.Fields{
		"Database": Settings.Database,
		"Host":     Settings.Host,
		"Port":     Settings.HTTPPort,
		"URL":      Settings.URL.String(),
	}).Infof("Loaded configuration")
}

// NewClientWithTimeout will retry pinging the server until a specified timeout passes
func NewClientWithTimeout(timeout time.Duration) (*clientv1.Client, error) {
	deadline := time.Now().Add(timeout)
	sleep := time.Millisecond * 100
	iterations := 0
	var c *clientv1.Client
	var err error

	hasTimedout := func() bool {
		return time.Now().After(deadline)
	}

	for !hasTimedout() {
		iterations++

		if iterations > 1 {
			log.Errorf("action=sleeping duration=%v", sleep)
			time.Sleep(sleep)
		}

		cfg := Settings.Config

		c, err = clientv1.NewClient(cfg)
		if err != nil {
			log.Errorf("action=create_client err=%v", err)
			if hasTimedout() {
				return nil, fmt.Errorf("Could not create client %s", err)
			}
			continue
		}

		duration, v, err := c.Ping()
		if err != nil {
			log.Errorf("action=ping t=%v version=%s err=%v", duration, v, err)
			if hasTimedout() {
				return nil, fmt.Errorf("failed to connect to %s: %v", cfg.URL.String(), err.Error())
			}
			continue
		}

		// If we made it this far everything is working
		break
	}

	return c, err
}

// ExecuteQuery executes the given query against the database and returns the response or an error
func ExecuteQuery(c *clientv1.Client, query string) (*clientv1.Response, error) {
	log.Infof("action=ExecuteQuery q=%s client=%+v", query, c)
	q := clientv1.Query{
		Command:  query,
		Database: Settings.Database,
		Chunked:  true,
	}
	response, err := c.Query(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// NewHTTPClient creates a new client for reading & writing data to influxDB over HTTP
func NewHTTPClient() (client.Client, error) {
	addr := Settings.Config.URL.String()
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: Settings.Config.Username,
		Password: Settings.Config.Password,
		Timeout:  Settings.Config.Timeout,
	})
	if err != nil {
		return nil, err
	}

	log.Debugf("action=NewHTTPClient addr=%s username=%s", addr, Settings.Config.Username)
	return c, nil
}

// WritePoint writes a single point to the DB with the given name, tags and fields
func WritePoint(name string, tags map[string]string, fields map[string]interface{}) (*client.Point, error) {
	c, err := NewHTTPClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Settings.Database,
		Precision: "s",
	})
	if err != nil {
		return nil, err
	}

	pt, err := client.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		return nil, err
	}

	bp.AddPoint(pt)

	// Write the batch and check for an error
	if err = c.Write(bp); err != nil {
		return nil, err
	}

	log.Debugf("WritePoint=%s", pt.String())
	return pt, nil
}

// Query executes InfluxQL statements against the configured database
func Query(query string) (*client.Response, error) {
	c, err := NewHTTPClient()
	if err != nil {
		return nil, err
	}
	defer c.Close()

	log.Debugf("query=\"%s\"", query)
	q := client.NewQuery(query, Settings.Database, "s")
	return c.Query(q)
}
