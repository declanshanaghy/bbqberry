package stubembd

import (
	"github.com/Polarishq/middleware/framework/log"
	"math/rand"
	"github.com/declanshanaghy/bbqberry/framework"
	"time"
	"sync"
)

// fakeTemps can be o.t to return specific analog readings during tests
var fakeTemps = make(map[int32]int32, 0)
var mux = &sync.Mutex{}

func init() {
	resetFakeTemps()
}

func resetFakeTemps() {
	rand.Seed( time.Now().UTC().UnixNano())

	if framework.Config.Stub {
		nProbes := int32(4)
		for probe := int32(0); probe < nProbes; probe++ {
			min := framework.ConvertCelsiusToAnalog(50)
			analog := min
			celcius, _ := framework.ConvertAnalogToCF(analog)
			fakeTemps[probe] = analog
			log.WithFields(log.Fields{
				"probe": probe,
				"analog": analog,
				"celcius": celcius,
			}).Infof("Probe initialized")
		}
	}
}

func SetFakeTemp(probe int32, analog int32) {
	fakeTemps[probe] = analog

	celcius, _ := framework.ConvertAnalogToCF(analog)
	log.WithFields(log.Fields{
		"probe":   probe,
		"analog":  analog,
		"celcius": celcius,
	}).Info("Set fake probe value")
}

func getFakeTemp(probe int32) int32 {
	limit := framework.Config.Hardware.Probes[probe].Limits
	min := framework.ConvertCelsiusToAnalog(*limit.MinWarnCelsius - 15)
	max := framework.ConvertCelsiusToAnalog(*limit.MaxAbsCelsius)

	mux.Lock()
	defer mux.Unlock()

	analog := fakeTemps[probe]
	celcius, _ := framework.ConvertAnalogToCF(analog)

	log.WithFields(log.Fields{
		"probe":   probe,
		"analog":  fakeTemps[probe],
		"celcius": celcius,
	}).Debug("Got fake probe value")

	if analog >= max {
		analog = min //int32(rand.Intn(int(min)) + int(max-min))
	}
	celcius, _ = framework.ConvertAnalogToCF(analog)
	analog2 := framework.ConvertCelsiusToAnalog(celcius + 1)

	fakeTemps[probe] = analog2

	log.WithFields(log.Fields{
		"probe": probe,
		"analog": analog,
		"analog2": analog2,
		"celcius": celcius,
	}).Debug("Probe updated")

	return analog
}