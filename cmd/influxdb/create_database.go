package main

import (
	"github.com/declanshanaghy/bbqberry/influxdb"
	"fmt"
	"io"
	"encoding/json"
	"github.com/influxdata/influxdb/client"
	"os"
)

func main() {
	//c := cli.New("unknown")
	//c.Host = influxdb.Cfg.Host
	//c.Port, _ = strconv.Atoi(influxdb.Cfg.Port)
	//c.ClientConfig = influxdb.Cfg.Config
	//c.Database = influxdb.Cfg.Database
	//c.Ssl = false
	//c.Format = "json"
	//c.Pretty = true
	//
	//if err := c.ExecuteQuery("CREATE DATABASE " + influxdb.Cfg.Database); err != nil {
	//	panic(err)
	//}
	
	c, err := influxdb.NewClient()
	if err != nil {
		panic(err)
	}

	r, err := influxdb.ExecuteQuery(c, "CREATE DATABASE " + influxdb.Cfg.Database)
	if err != nil {
		panic(err)
	} else {
		writeJSON(r, os.Stdout)
	}
}

func writeJSON(response *client.Response, w io.Writer) {
	var data []byte
	var err error
	data, err = json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Fprintf(w, "Unable to parse json: %s\n", err)
		return
	}
	fmt.Fprintln(w, string(data))
}


