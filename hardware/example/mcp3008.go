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

		/*
			Voltage divider configuration
					vcc	(3.3v)
					^
					|
					r1	(thermistor)
					|
					|------> vOut
					|
					r2	(1k)
					|
					-
					gnd	(0v)
		*/
		vcc := float32(3.3)
		maxA := float32(1023.0)
		vPerA := vcc / maxA
		r2 := float32(1000.0)
		vOut := analog * vPerA

		r1 := ((vcc * r2) / vOut) - r2
		// log.Infof("A=%0.5f, V=%0.5f, R1=%0.5f", analog, vOut, r1)

		tempK := SteinhartHart(r1)
		tempC, tempF := convertK(tempK)
		log.Infof("A=%0.5f, V=%0.5f, R=%0.5f, K=%0.5f, C=%0.5f, F=%0.5f", analog, vOut, r1, tempK, tempC, tempF)

		time.Sleep(1 * time.Second)
	}
}

func convertK(tempK float32) (tempC float32, tempF float32) {
	tempC = tempK - 273.15 // K to C
	tempF = tempC*1.8 + 32 // Fahrenheit
	return
}

// SteinhartHart calculates temperature from the given analog value using the Steinhart Hart formula
func SteinhartHart(resistance float32) (tempK float32) {
	a := framework.Constants.SteinhartHart.A
	b := framework.Constants.SteinhartHart.B
	c := framework.Constants.SteinhartHart.C
	// Rn := framework.Constants.SteinhartHart.Rn

	v := math.Log(float64(resistance))

	// Steinhart Hart Equation
	// T = 1/(a + b[ln(R)] + (c[ln(R)])^3)
	tempK = float32(1.0 / (a + b*v + c*v*v*v))
	return
}
