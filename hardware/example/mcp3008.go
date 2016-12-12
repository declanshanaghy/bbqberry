// +build ignore

// this sample uses the mcp3008 package to interface with the 8-bit ADC and works without code change on bbb and rpi
package main

import (
	"time"
	"math"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
	"github.com/Polarishq/middleware/framework/log"
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
		for i, _ := range readings {
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

		calculateTemperature(avg)
		// fmt.Printf("a=%v, f=%0.2f\n", avg, f)

		time.Sleep(1 * time.Second)
	}
}

func calculateTemperature(value int) (float64, float64, float64) {
	// iBBQ probe is 100.8K at 25c
	
    volts := (float64(value) * 3.3) / 1024 // calculate the voltage
    ohms := ((1/volts)*3300)-1000 // calculate the ohms of the thermististor

    lnohm := math.Log1p(ohms) // take ln(ohms)

    // a, b, & c values from http://www.thermistor.com/calculators.php
    // using curve R (-6.2%/C @ 25C) Mil Ratio X
    // a =  0.002197222470870
    // b =  0.000161097632222
    // c =  0.000000125008328

    a :=  0.000570569668444 
    b :=  0.000239344111326 
    c :=  0.000000047282773 

    // Steinhart Hart Equation
    // T = 1/(a + b[ln(ohm)] + c[ln(ohm)]^3)

    t1 := (b*lnohm) // b[ln(ohm)]
    c2 := c*lnohm // c[ln(ohm)]
    t2 := math.Pow(c2,3) // c[ln(ohm)]^3

    tempk := 1/(a + t1 + t2) // calculate temperature
    tempc := tempk - 273.15 - 4 //K to C
    tempf := tempc*9/5 + 32
    // the -4 is error correction for bad python math

    // print out info
    log.Infof("%4d/1023 => %5.3f V => %4.1f ohms  => %4.1f K => %4.1f C  => %4.1f F\n", value, volts, ohms, tempk, tempc, tempf)

    return tempk, tempc, tempf	
}
	