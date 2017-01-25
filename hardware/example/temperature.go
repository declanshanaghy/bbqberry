// +build ignore

package main

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
)

func main() {
	tReader := hardware.NewTemperatureReader()
	for true {
		reading := models.TemperatureReading{}
		err := tReader.GetTemperatureReading(1, &reading)
		if err != nil {
			panic(err)
		}
		log.Infof("READING: %0.2f\n", *reading.Fahrenheit)
		time.Sleep(1 * time.Second)
	}
}