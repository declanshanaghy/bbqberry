package influx

import (
	"fmt"

	"github.com/influxdata/influxdb/client/v2"
)

const (
	host     = "influx"
	portHTTP = 8086
	portUDP  = 8089
	username = "bbqberry"
	password = "piberry"

	// DefaultDB is the default database to read and write data
	DefaultDB = "explore"
)

// GetHTTPClient creates a new client for reading & writing data to influxDB over HTTP
func getHTTPClient() (client.Client, error) {
	return client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%d", host, portHTTP),
		Username: username,
		Password: password,
	})
}

// GetUDPClient creates a new client for reading & writing data to influxDB over UDP
func getUDPClient() (client.Client, error) {
	return client.NewUDPClient(client.UDPConfig{Addr: fmt.Sprintf("%s:%d", host, portUDP)})
}

// GetDefaultClient creates a new client for communication with InfluxDB
// using the default communication protocol (HTTP)
func GetDefaultClient() (client.Client, error) {
	return getHTTPClient()
}
