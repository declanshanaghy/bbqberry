package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/lucasb-eyer/go-colorful"
	"fmt"
	"github.com/heatxsink/go-hue/groups"
	"github.com/heatxsink/go-hue/lights"
	"github.com/heatxsink/go-hue/portal"
	"github.com/declanshanaghy/bbqberry/db/influxdb"
)

// temperatureIndicator collects and logs temperature metrics
type temperatureIndicator struct {
	basicTickable
	reader      	hardware.TemperatureReader
	strip       	hardware.WS2801
	errorCount  	int
	probeNumber 	int32
	huePortal		*portal.Portal
	hueGroup		*groups.Group
	hueGroups		*groups.Groups
	initialState	*lights.State
	currentState	*lights.State
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newTemperatureIndicator(huePortal *portal.Portal) RunnableTicker {
	t := &temperatureIndicator{
		reader: 		hardware.NewTemperatureReader(),
		strip:  		hardware.NewStrandController(),
		huePortal:		huePortal,
		probeNumber:	-1,
	}
	t.Period = time.Second * 5

	return newRunnableTicker(t)
}

// GetName returns a human readable name for this background task
func (o *temperatureIndicator) GetName() string {
	return "Temperature"
}

// Start performs initialization before the first tick
func (o *temperatureIndicator) start() error {
	o.hueGroups = groups.New(o.huePortal.InternalIPAddress, framework.HUE_KEY)
	allGroups, err := o.hueGroups.GetAllGroups()
	if err != nil {
		return err
	} else {
		log.WithField("group", framework.HUE_ALERT_GROUP).Info("Searching for group")
		for _, g := range allGroups {
			if g.Name == framework.HUE_ALERT_GROUP {
				log.WithField("group", g).Info("Found hue group")
				o.hueGroup = &g
				o.initialState = &o.hueGroup.Action
				o.currentState = &lights.State{On: true}
				break
			}
		}
	}

	probes := o.reader.GetEnabledPobes()
	for _, p := range(*probes) {
		probe := *framework.Constants.Hardware.Probes[p]
		log.WithFields(log.Fields{
			"p": p,
			"enabled": *probe.Enabled,
			"type": *probe.Limits.ProbeType,
		}).Info("Checking probe")
		if *probe.Limits.ProbeType == "ambient" {
			o.probeNumber = p
			break
		}
	}

	if o.probeNumber == -1 {
		return fmt.Errorf("Unable to find an embient probe to monitor")
	}

	return nil
}

// Stop performs cleanup when the goroutine is exiting
func (o *temperatureIndicator) stop() error {
	if o.initialState != nil {
		log.WithField("state", *o.initialState).Info("Resetting hue initial state")
		_, err := o.hueGroups.SetGroupState(o.hueGroup.ID, *o.initialState)
		if err != nil {
			return err
		}
	}
	return nil
}

// Tick executes on a ticker schedule
func (o *temperatureIndicator) tick() error {
	probe := framework.Constants.Hardware.Probes[o.probeNumber]
	min := *probe.Limits.MinWarnCelsius
	max := *probe.Limits.MaxWarnCelsius

	avg, err := influxdb.QueryAverageTemperature(o.getPeriod() * 10, o.probeNumber)
	if err != nil {
		log.Error(err)
		return nil
	}

	color := getTempColor(*avg.Celsius, min, max)
	if err := o.strip.SetAllPixels(color); err != nil {
		return err
	}

	h, s, l := hardware.ColorToHue(color)
	o.currentState.Hue = h
	o.currentState.Sat = s
	o.currentState.Bri = l

	r, g, b := color.RGB255()
	log.WithFields(log.Fields{
		"temp": *avg.Fahrenheit,
		"color": color.Hex(),
		"(r, g, b)": fmt.Sprintf("(%02x, %02x, %02x)", r, g, b),
		"(h, s, l)": fmt.Sprintf("(%d, %d, %d)", h, s, l),
	}).Debugf("Mapped color to hue")

	o.hueGroups.SetGroupState(o.hueGroup.ID, *o.currentState)
	return nil
}

func getTempColor(temp, min, max int32) colorful.Color {
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

	color := colorful.Color{R: float64(red) / 255, G: 0.0, B: float64(blu) / 255}

	log.WithFields(log.Fields{
		"min": min,
		"max": max,
		"rnge": rnge,
		"temp": temp,
		"corrected": corrected,
		"scaled": scaled,
		"red": fmt.Sprintf("0x%02x", red),
		"blu": fmt.Sprintf("0x%02x", blu),
		"colorHex": color.Hex(),
	}).Debugf("Calculating temperature color")

	return color
}