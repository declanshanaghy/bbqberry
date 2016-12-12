package main

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/Polarishq/middleware/framework/log"
	"time"
)

func main() {
    tReader := hardware.NewTemperatureReader()
    for true {
        reading, err := tReader.GetTemperatureReading(1)
        if err != nil {
            panic(err)
        }
        log.Infof("READING: %+v\n", reading)
        time.Sleep(1 * time.Second)
    }
}