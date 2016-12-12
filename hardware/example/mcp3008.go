// +build ignore

// this sample uses the mcp3008 package to interface with the 8-bit ADC and works without code change on bbb and rpi
package main

import (
	"fmt"
	"time"
	"math"

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

// y = mx+b
// b = y - mx
// 10K readings
// const (
// 	x1_temp = 40.0
// 	y1_a	= 952.0
// 	x2_temp = 153.0
// 	y2_a	= 893.0
// )

// 1K readings
const (
	x1_temp = 42.0
	y1_a	= 1020.0
	x2_temp = 183.0
	y2_a	= 1005.0
)

func main() {
	m := (y2_a - y1_a) / (x2_temp - x1_temp)
	b := y2_a - m * x2_temp

	fmt.Printf("(%v, %v) --> (%v, %v)\n", x1_temp, y1_a, x2_temp, y2_a)
	fmt.Printf("m=%v, b=%v\n", m, b)

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

		tempCalc(1023 - avg)
		// fmt.Printf("a=%v, f=%0.2f\n", avg, f)

		// y = mx+b
		// t := m * float64(avg) + b
		// k, c, f := convertVoltToTemp(int(avg))
		// fmt.Printf("a=%v, f=%0.2f, c=%0.2f, k=%0.2f\n", avg, f, c, k)

		time.Sleep(1 * time.Second)
	}
}

func convertVoltToTemp(volt int) (k, c, f float64) {
	// get the Kelvin temperature
	k = math.Log(10240000.0/float64(volt) - 10000)
	k = 1 / (0.001129148 + (0.000234125 * k) + (0.0000000876741 * k * k * k))

	// convert to Celsius and round to 1 decimal place
	c = k - 273.15

	// get the Fahrenheit temperature
	f = (c * 1.8) + 32

	// return all three temperature values
	return
}

func tempCalc(value int) float64 {
    volts := (float64(value) * 5.0) / 1024 // calculate the voltage
    ohms := ((1/volts)*5000)-1000 // calculate the ohms of the thermististor

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

    temp := 1/(a + t1 + t2) // calculate temperature
    tempc := temp - 273.15 - 4 //K to C
    tempf := tempc*9/5 + 32
    // the -4 is error correction for bad python math

    // print out info
    fmt.Printf("%4d/1023 => %5.3f V => %4.1f ohms  => %4.1f K => %4.1f C  => %4.1f F\n", value, volts, ohms, temp, tempc, tempf)

    return tempf	
}
	