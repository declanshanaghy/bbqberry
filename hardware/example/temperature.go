package example

import (
	"github.com/declanshanaghy/bbqberry/hardware"
	"time"
    "fmt"
)

func printTemperature() {
    tReader := hardware.NewTemperatureReader()
    reading, err := tReader.GetTemperatureReading(1)
    if err != nil {
        panic(err)
    }
    fmt.Printf("READING: %v\n", reading)
    time.Sleep(1 * time.Second)
}