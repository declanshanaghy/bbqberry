package influxdb

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/influxdata/influxdb/client"
	v2 "github.com/influxdata/influxdb/client/v2"
)

var defaultTimeout = time.Second

// Cfg holds settings to communicate with influxdb
var Cfg clientConfig

type clientConfig struct {
	client.Config

	Database string
	Host     string
	Port     string
}

func init() {
	LoadConfig()
}

// LoadConfig loads the influxdb connection settings
func LoadConfig() {
	host := os.Getenv("INFLUXDB_HOST")
	if host == "" {
		host = "influxdb"
	}
	port := os.Getenv("INFLUXDB_PORT_HTTP")
	if port == "" {
		port = strconv.Itoa(client.DefaultPort)
	}
	username := os.Getenv("INFLUXDB_USERNAME")
	password := os.Getenv("INFLUXDB_PASSWORD")
	database := os.Getenv("INFLUXDB")
	if database == "" {
		database = "no_name_given"
	}

	addr := net.JoinHostPort(host, port)
	URL, err := client.ParseConnectionString(addr, false)
	if err != nil {
		panic(err)
	}

	cc := client.Config{
		Username: username,
		Password: password,
		URL:      URL,
		Timeout:  defaultTimeout,
	}
	Cfg = clientConfig{
		Database: database,
		Host:     host,
		Port:     port,
	}
	Cfg.Config = cc
}

// NewClient creates a new raw InfluxDB client
func NewClient() (*client.Client, error) {
	c, err := client.NewClient(Cfg.Config)
	if err != nil {
		return nil, fmt.Errorf("Could not create client %s", err)
	}
	duration, v, err := c.Ping()
	log.Debugf("action=influx_ping t=%v version=%s err=%v", duration, v, err)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", Cfg.Config.URL.String(), err.Error())
	}
	return c, nil
}

// ExecuteQuery executes the given query against the database and returns the response or an error
func ExecuteQuery(c *client.Client, query string) (*client.Response, error) {
	log.Info("action=ExecuteQuery q=" + query)
	q := client.Query{
		Command:  query,
		Database: Cfg.Database,
		Chunked:  true,
	}
	response, err := c.Query(q)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// NewHTTPClient creates a new client for reading & writing data to influxDB over HTTP
func NewHTTPClient() (v2.Client, error) {
	addr := Cfg.Config.URL.String()
	c, err := v2.NewHTTPClient(v2.HTTPConfig{
		Addr:     addr,
		Username: Cfg.Config.Username,
		Password: Cfg.Config.Password,
		Timeout:  defaultTimeout,
	})
	if err != nil {
		return nil, err
	}

	log.Infof("action=NewHTTPClient addr=%s username=%s", addr, Cfg.Config.Username)
	return c, nil
}
