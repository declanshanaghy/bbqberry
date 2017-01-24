package influxdb

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/influxdata/influxdb/client"
	clientV2 "github.com/influxdata/influxdb/client/v2"
)

var defaultTimeout = time.Second

// Settings holds all pertient connection parameters for InfluxDB
var Settings *influxDBSettings

type influxDBSettings struct {
	client.Config

	Database string
	Host     string
	Port     string
}

func init() {
	LoadConfig()
}

func LoadConfig() {
	database := os.Getenv("INFLUXDB")
	if database == "" {
		database = "no_name_given"
	}
	
	host := os.Getenv("INFLUXDB_HOST")
	if host == "" {
		host = "influxdb"
	}
	port := os.Getenv("INFLUXDB_PORT_HTTP")
	if port == "" {
		port = strconv.Itoa(client.DefaultPort)
	}
	URL, err := client.ParseConnectionString(net.JoinHostPort(host, port), false)
	if err != nil {
		panic(err)
	}
	
	username := os.Getenv("INFLUXDB_USERNAME")
	password := os.Getenv("INFLUXDB_PASSWORD")
	
	Settings = &influxDBSettings{
		Config: client.Config{
			URL: URL,
			Username: username,
			Password: password,
		},
		Database: database,
		Host: host,
		Port: port,
	}
	log.Infof("action=LoadConfig influxDBSettings=%+v URL=%s", Settings, Settings.URL.String())
}

// NewClientWithTimeout will retry pinging the server until a specified timeout passes
func NewClientWithTimeout(timeout time.Duration) (*client.Client, error) {
	deadline := time.Now().Add(timeout)
	sleep := time.Millisecond * 100
	iterations := 0
	var c *client.Client
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
		
		c, err = client.NewClient(cfg)
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
func ExecuteQuery(c *client.Client, query string) (*client.Response, error) {
	log.Infof("action=ExecuteQuery q=%s client=%+v", query, c)
	q := client.Query{
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
func NewHTTPClient() (clientV2.Client, error) {
	addr := Settings.Config.URL.String()
	c, err := clientV2.NewHTTPClient(clientV2.HTTPConfig{
		Addr:     addr,
		Username: Settings.Config.Username,
		Password: Settings.Config.Password,
		Timeout:  defaultTimeout,
	})
	if err != nil {
		return nil, err
	}

	log.Infof("action=NewHTTPClient addr=%s username=%s", addr, Settings.Config.Username)
	return c, nil
}

// WritePoint writes a single point to the DB with the given, name, tags and fields
func WritePoint(name string, tags map[string]string, fields map[string]interface{}) (*clientV2.Point, error) {
	c, err := NewHTTPClient()
	if err != nil {
		return nil, err
	}

	// Create a new point batch
	bp, err := clientV2.NewBatchPoints(clientV2.BatchPointsConfig{
		Database:  Settings.Database,
		Precision: "s",
	})
	if err != nil {
		return nil, err
	}

	pt, err := clientV2.NewPoint(name, tags, fields, time.Now())
	if err != nil {
		return nil, err
	}

	bp.AddPoint(pt)

	// Write the batch and check for an error
	if err = c.Write(bp); err != nil {
		return nil, err
	}

	log.Debugf("Point=%v pt=%s", pt.Name(), pt.String())
	return pt, nil
}