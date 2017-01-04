// +build ignore

// this sample uses the mcp3008 package to interface with the 8-bit ADC and works without code change on bbb and rpi
package main

import (
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	channel = 1
	speed   = 1000000
	bpw     = 8
	delay   = 0
)

func main() {
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, channel, speed, bpw, delay)
	defer spiBus.Close()

	adc := mcp3008.New(mcp3008.SingleMode, spiBus)
	readings := [10000]int{}

	for true {
		for i := range readings {
			v, err := adc.AnalogValueAt(0)
			if err != nil {
				panic(err)
			}
			readings[i] = v
		}
		tot := 0
		for _, v := range readings {
			tot += v
		}
		analog := float32(tot) / float32(len(readings))

		time.Sleep(1 * time.Second)
	}
}
