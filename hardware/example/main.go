package example

import (
	"sync"
	"time"

	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/kidoman/embd"
)

func reader(w *sync.WaitGroup, temp chan<- *hardware.TemperatureReading, t hardware.TemperatureArray) {
	// When the sending channel is closed a panic will occur, this is the signal to exit
	defer func() {
		if r := recover(); r != nil {
			log.Infof("action=recover r=%v", r)
			if err, ok := r.(error); !ok {
				log.Errorf("pkg: %v err=%v ok=%v", r, err, ok)
			}
		}
		log.Infof("action=done")
		w.Done()
	}()

	ticker := time.NewTicker(time.Second * 1)
	for tick := range ticker.C {
		t, _ := t.GetTemperatureReading(1)
		log.Debugf("action=tx time=%s temp=%v", tick, t)
		temp <- t
	}
}

func processor(w *sync.WaitGroup, temp <-chan *hardware.TemperatureReading) {
	loop := true
	for loop {
		select {
		case t, more := <-temp:
			if more {
				log.Debugf("action=rx temp=%v", t)
			} else {
				log.Infof("action=quit")
				loop = false
			}
		case <-time.After(time.Second * 1):
			continue
		}
	}
	log.Infof("action=done")
	w.Done()
}

type closer struct {
	w     *sync.WaitGroup
	bus0  embd.SPIBus
	bus1  embd.SPIBus
	cTemp chan *hardware.TemperatureReading
}

func (c *closer) Close() {
	c.bus0.Close()
	c.bus1.Close()
	close(c.cTemp)
	c.w.Wait()
}

// ReaderWriterExample provides an example of how to read temperature values from the hardware and communicate them
// to a separate goroutine for processing
func ReaderWriterExample() framework.Closer {
	var w sync.WaitGroup
	w.Add(2)

	bus0 := embd.NewSPIBus(embd.SPIMode0, 0, 1000000, 8, 0)
	go func() {
		Rainbow(25, bus0)
	}()

	bus1 := embd.NewSPIBus(embd.SPIMode0, 1, 1000000, 8, 0)
	sTemp := hardware.NewTemperatureArray(1, bus1)
	cTemp := make(chan *hardware.TemperatureReading, 1)

	go reader(&w, cTemp, sTemp)
	go processor(&w, cTemp)

	return &closer{w: &w, bus0: bus0, bus1: bus1, cTemp: cTemp}
}
