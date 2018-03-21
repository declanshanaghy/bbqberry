package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
)

// temperatureIndicator collects and logs temperature metrics
type temperatureIndicator struct {
	basicTickable
	reader     hardware.TemperatureReader
	strip      hardware.WS2801
	errorCount int
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newTemperatureIndicator() RunnableTicker {
	t := &temperatureIndicator{
		reader: hardware.NewTemperatureReader(),
		strip:  hardware.NewStrandController(),
	}
	t.Period = time.Second

	return newRunnableTicker(t)
}

// GetName returns a human readable name for this background task
func (o *temperatureIndicator) GetName() string {
	return "Temperature"
}

// Start performs initialization before the first tick
func (o *temperatureIndicator) start() error {
	return o.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (o *temperatureIndicator) stop() error {
	return nil
}

// Tick executes on a ticker schedule
func (o *temperatureIndicator) tick() error {
	// Assuming that the ambient probe is #0
	ambientProbeNumber := int32(0)
	probe := framework.Constants.Hardware.Probes[ambientProbeNumber]
	min := *probe.Limits.MinWarnCelsius
	max := *probe.Limits.MaxWarnCelsius

	avg, err := influxdb.QueryAverageTemperature(o.getPeriod() * 10, ambientProbeNumber)
	if err != nil {
		log.Error(err)
		return nil
	}

	color := getTempColor(*avg.Celsius, min, max)

	if err := o.strip.SetAllPixels(color); err != nil {
		return err
	}

	return nil
}

func getTempColor(temp, min, max int32) int {
	// Map the temperature to a color to be displayed on the LED pixels.
	// cold / min = blue	( 0x0000FF ) =
	// hot / max = red ( 0xFF0000 )
	// green LEDs should never be on.
	// Treat each degree above min as a +1 of the red component and -1 of the blue component
	// Therefore:
	// 		avg temp <= min = pure blue
	// 		avg temp >= max = pure red
	// If the max limit is exceeded a visual indicator should be displayed (i.e. flashing)

	if temp < min {
		log.Warningf("Temp (%d) < min (%d)...clamping", temp, min)
		temp = min
	}
	if temp > max {
		log.Warningf("Temp (%d) > max (%d)...clamping", temp, max)
		temp = max
	}

	rnge := max - min
	corrected := temp - min
	scaled := float32(corrected) / float32(rnge)

	red := int(255 * scaled)
	blu := 0xFF - red

	color := red<<16 | blu

	log.WithFields(log.Fields{
		"min": min,
		"max": max,
		"rnge": rnge,
		"temp": temp,
		"corrected": corrected,
		"scaled": scaled,
		"color": color,
	}).Debugf("Calculating temperature color")

	return color
}