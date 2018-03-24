package stubembd

import (
	"github.com/Polarishq/middleware/framework/log"
	"math/rand"
	"github.com/declanshanaghy/bbqberry/framework"
)

// fakeTemps can be o.t to return specific analog readings during tests
var fakeTemps = make(map[int32]int32, 0)

func init() {
	resetFakeTemps()
}

func resetFakeTemps() {
	hwCfg := framework.Constants.Hardware

	if framework.Constants.Stub {
		nProbes := int32(len(hwCfg.Probes))
		for probe := int32(0); probe < nProbes; probe++ {
			limit := framework.Constants.Hardware.Probes[probe].Limits
			min := framework.ConvertCelsiusToAnalog(-15)
			max := framework.ConvertCelsiusToAnalog(*limit.MaxAbsCelsius)
			analog := int32(rand.Intn(int(min)) + int(max-min))
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
	}).Infof("Set fake probe value")
}

func getFakeTemp(probe int32) int32 {
	limit := framework.Constants.Hardware.Probes[probe].Limits
	min := framework.ConvertCelsiusToAnalog(-15)
	max := framework.ConvertCelsiusToAnalog(*limit.MaxAbsCelsius)
	analog := fakeTemps[probe]
	celcius, _ := framework.ConvertAnalogToCF(analog)

	log.WithFields(log.Fields{
		"probe":   probe,
		"analog":  fakeTemps[probe],
		"celcius": celcius,
	}).Infof("Got fake probe value")


	if analog >= max {
	   analog = min
	}
	celcius, _ = framework.ConvertAnalogToCF(analog)
	analog2 := framework.ConvertCelsiusToAnalog(celcius + 1)

	fakeTemps[probe] = analog

	log.WithFields(log.Fields{
		"probe": probe,
		"analog2": analog2,
		"celcius": celcius,
	}).Infof("Probe updated")

	return analog
}