package influx

import (
	"github.com/influxdata/influxdb/client/v2"
	"fmt"
)

const (
	host = "influx"
	port_http = 8086
	port_udp = 8089
	username = "bbqberry"
	password = "piberry"

	DB = "explore"
)

func GetHttpClient() (client.Client, error) {
	return client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%d", host, port_http),
		Username: username,
		Password: password,
	})
}

func GetUdpClient() (client.Client, error) {
	return client.NewUDPClient(client.UDPConfig{Addr:fmt.Sprintf("%s:%d", host, port_udp)})
}

func GetDefaultClient() (client.Client, error) {
	return GetHttpClient()
}