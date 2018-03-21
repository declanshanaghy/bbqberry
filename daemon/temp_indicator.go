package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/models"
)

// temperatureIndicator collects and logs temperature metrics
type temperatureIndicator struct {
<<<<<<< Updated upstream
	period     time.Duration
	reader     hardware.TemperatureReader
	strip      hardware.WS2801
	errorCount int
=======
	basicTickable
	strip      			hardware.WS2801
	errorCount 			int
	probe 				*models.TemperatureProbe
	probeIndex 			int32
>>>>>>> Stashed changes
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
<<<<<<< Updated upstream
func newTemperatureIndicator() Runnable {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return newRunnable(
		&temperatureIndicator{
			reader: hardware.NewTemperatureReader(),
			strip:  hardware.NewStrandController(),
			period: time.Second,
		},
	)
}
=======
func newTemperatureIndicator() RunnableTicker {
	var p *models.TemperatureProbe
	var i int

	// No fail safe if an ambient probe is not found
	for z, probe := range(framework.Constants.Hardware.Probes) {
		if ( *probe.Limits.ProbeType == framework.PROBE_TYPE_AMBIENT && *probe.Enabled ) {
			p = probe
			i = z
		}
	}

	if ( p == nil ) {
		log.Error("Unable to find ambient probe in hardware config")
	}

	t := &temperatureIndicator{
		strip:  	hardware.NewStrandController(),
		probe:  	p,
		probeIndex: int32(i),
	}
	t.Period = time.Second
>>>>>>> Stashed changes

func (r *temperatureIndicator) getPeriod() time.Duration {
	return r.period
}

func (r *temperatureIndicator) setPeriod(period time.Duration)  {
	r.period = period
}

// GetName returns a human readable name for this background task
func (r *temperatureIndicator) GetName() string {
	return "temperatureIndicator"
}

// Start performs initialization before the first tick
<<<<<<< Updated upstream
func (r *temperatureIndicator) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
	r.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (r *temperatureIndicator) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")
}

// Tick executes on a ticker schedule
func (r *temperatureIndicator) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	// Assuming that the ambient probe is #0
	ambientProbeNumber := int32(0)

	avg, err := influxdb.QueryAverageTemperature(r.getPeriod() * 10, ambientProbeNumber)
=======
func (o *temperatureIndicator) start() error {
	return o.tick()
}

// Stop performs cleanup when the goroutine is exiting
func (o *temperatureIndicator) stop() error {
	return nil
}

// Tick executes on a ticker schedule
func (o *temperatureIndicator) tick() error {
	if ( o.probe == nil ) {
		log.Error("Unable to find ambient probe in hardware config")
	}

	avg, err := influxdb.QueryAverageTemperature(o.getPeriod() * 10, o.probeIndex)
>>>>>>> Stashed changes

	if err != nil {
		return err
	}

<<<<<<< Updated upstream
	probe := framework.Constants.Hardware.Probes[ambientProbeNumber]
	min := *probe.TempLimits.MinWarnCelsius
	max := *probe.TempLimits.MaxWarnCelsius

	color := r.getTempColor(*avg.Celsius, min, max)

	if err := r.strip.SetAllPixels(color); err != nil {
		log.Error(err.Error())
=======
	min := o.probe.Limits.MinWarnCelsius
	max := o.probe.Limits.MaxWarnCelsius

	color := getTempColor(*avg.Celsius, *min, *max)

	if err := o.strip.SetAllPixels(color); err != nil {
		return err
>>>>>>> Stashed changes
	}

	return nil
}

<<<<<<< Updated upstream
func (r *temperatureIndicator) getTempColor(temp, min, max int32) int {
=======
func getTempColor(temp, min, max int32) int {
>>>>>>> Stashed changes
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
