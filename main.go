package main

import (
	"flag"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/kidoman/embd"
	"github.com/declanshanaghy/bbqberry/samples"
	"github.com/declanshanaghy/bbqberry/sensors"
	"time"
	"github.com/golang/glog"
	"fmt"
	"sync"
)

func reader(w *sync.WaitGroup, temp chan<- sensors.TemperatureReading, t sensors.TemperatureArray) {
	// When the sending channel is closed a panic will occur, this is the signal to exit
	defer func() {
		if r := recover(); r != nil {
			glog.Infof("action=recover r=%v", r)
			if err, ok := r.(error); !ok {
				glog.Errorf("pkg: %v err=%v ok=%v", r, err, ok)
			}
		}
		glog.Infof("action=done")
		w.Done()
	}()

	ticker := time.NewTicker(time.Second * 1)
	for tick := range ticker.C {
		t := t.GetTemp(0)
		glog.Infof("action=tx time=%s temp=%v", tick, t)
		temp <- t
	}
}

func processor(w *sync.WaitGroup, temp <-chan sensors.TemperatureReading) {
	loop := true
	for loop {
		select {
		case t, more := <-temp:
			if more {
				glog.Infof("action=rx temp=%v", t)
			} else {
				glog.Infof("action=quit")
				loop = false
			}
		case <-time.After(time.Second * 1):
			continue
		}
	}
	glog.Infof("action=done")
	w.Done()
}

func main() {
	flag.Parse()

	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	go func() {
		samples.Rainbow(25)
	}()

	var w sync.WaitGroup
	w.Add(3)

	sTemp := sensors.NewTemperature(1)
	cTemp := make(chan sensors.TemperatureReading, 1)

	go reader(&w, cTemp, sTemp)
	go processor(&w, cTemp)

	var input string
	fmt.Scanln(&input)

	close(cTemp)

	w.Wait()
}

