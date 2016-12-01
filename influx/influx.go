package influx

import (
	"log"
	"github.com/influxdata/influxdb/client/v2"
	"fmt"
)

const (
	host = "influx"
	port = 8086
	username = "bbqberry"
	password = "piberry"

	DB = "explore"
)

func GetClient() client.Client {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%d", host, port),
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return c
}
