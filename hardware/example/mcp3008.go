// +build ignore

// this sample uses the mcp3008 package to interface with the 8-bit ADC and works without code change on bbb and rpi
package main

import (
	"math"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
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
	readings := [1000]int{}

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
		avg := tot / len(readings)

		tempK, tempC, tempF, voltage, resistance := SteinhartHart(int32(avg))
		log.Infof("A=%v, V=%v, R=%v, K=%v, C=%v, F=%v", avg, voltage, resistance, tempK, tempC, tempF)

		time.Sleep(1 * time.Second)
	}
}

// SteinhartHart calculates temperature from the given analog value using the Steinhart Hart formula
func SteinhartHart(analog int32) (tempK float32, tempC float32, tempF float32, voltage float32, resistance int32) {
	// iBBQ probe is 100.8K at 25c

	volts := (float64(analog) * 3.3) / 1024 // calculate the voltage
	voltage = float32(volts)
	ohms := ((1 / volts) * 3300) - 1000 // calculate the resistance of the thermististor
	resistance = int32(ohms)

	lnohm := math.Log1p(ohms) // take ln(ohms)

	a := framework.Constants.SteinhartHart.A
	b := framework.Constants.SteinhartHart.B
	c := framework.Constants.SteinhartHart.C

	// Steinhart Hart Equation
	// T = 1/(a + b[ln(ohm)] + c[ln(ohm)]^3)
	t1 := (b * lnohm)     // b[ln(ohm)]
	c2 := c * lnohm       // c[ln(ohm)]
	t2 := math.Pow(c2, 3) // c[ln(ohm)]^3

	tempK = float32(1 / (a + t1 + t2)) // Calculate temperature in Kelvin
	tempC = tempK - 273.15 - 4         // K to C (the -4 is error correction for bad python math)
	tempF = tempC*9/5 + 32             // Fahrenheit

	return
}
