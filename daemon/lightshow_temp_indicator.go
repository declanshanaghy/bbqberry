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
	lightShowTickable
	reader      	hardware.TemperatureReader
	errorCount  	int
	probeNumber 	int32
	huePortal		*portal.Portal
	hueGroup		*groups.Group
	hueGroups		*groups.Groups
	initialState	*lights.State
	currentState	*lights.State
	hueUpdInterval 	time.Duration
	hueUpdTime 		time.Time
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newTemperatureIndicator(huePortal *portal.Portal) LightShow {
	t := &temperatureIndicator{
		reader: 		hardware.NewTemperatureReader(),
		huePortal:		huePortal,
		probeNumber:	-1,
		hueUpdTime:		time.Now(),
		hueUpdInterval: time.Second * 10,
	}
	t.strip = hardware.NewGrillLightController()
	t.Period = time.Second

	return newLightShow(t)
}

// GetName returns a human readable name for this background task
func (o *temperatureIndicator) GetName() string {
	return "Temperature"
}

func (o *temperatureIndicator) initializeHue() error {
	if o.huePortal == nil {
		return nil
	}

	o.hueGroups = groups.New(o.huePortal.InternalIPAddress, framework.HUE_KEY)
	allGroups, err := o.hueGroups.GetAllGroups()
	if err != nil {
		return err
	} else {
		log.WithField("group", framework.HUE_ALERT_GROUP).Info("Searching for group")
		for _, g := range allGroups {
			if g.Name == framework.HUE_ALERT_GROUP {
				log.WithFields(log.Fields{
					"group": g,
					"Action": g.Action,
				}).Info("Found hue group")
				o.hueGroup = &g
				o.initialState = &lights.State{
					On: g.Action.On,
					Hue: g.Action.Hue,
					Effect: g.Action.Effect,
					Bri: g.Action.Bri,
					Sat: g.Action.Sat,
					CT: g.Action.CT,
					XY: g.Action.XY,
					Alert: "",
					TransitionTime: g.Action.TransitionTime,
				}
				o.currentState = &lights.State{On: true}
				break
			}
		}
	}

	return nil
}

// Start performs initialization before the first tick
func (o *temperatureIndicator) start() error {
	o.initializeHue()

	probes := framework.Config.GetEnabledProbeIndexes()
	for _, p := range(*probes) {
		probe := *framework.Config.Hardware.Probes[p]
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
	probe := framework.Config.Hardware.Probes[o.probeNumber]
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

	hueGroupName := ""

	if o.hueGroup != nil && time.Now().Sub(o.hueUpdTime) >= 0 {
		hueGroupName = o.hueGroup.Name
		h, s, l := hardware.ColorToPhilipsHueHSB(color)

		o.currentState.Hue = h
		o.currentState.Sat = s
		o.currentState.Bri = l

		if *avg.Celsius > max {
			// If max is exceeded ensure the alert is flashing
			o.currentState.Alert = "lselect"

			log.WithFields(log.Fields{
				"Celsius": *avg.Celsius,
				"color": color.Hex(),
				"name": o.hueGroup.Name,
				"nextHueUpdate": o.hueUpdTime,
				"hueUpdTime": o.hueUpdTime,
				"Alert": o.currentState.Alert,
			}).Info("Updated hue to alert state")

			o.hueGroups.SetGroupState(o.hueGroup.ID, *o.currentState)
		} else if (*o.currentState).Alert == "lselect" && *avg.Celsius < max  {
			//If the current state is an alert and the temp decreased, update hue
			o.currentState.Alert = ""

			log.WithFields(log.Fields{
				"Celsius": *avg.Celsius,
				"color": color.Hex(),
				"name": o.hueGroup.Name,
				"nextHueUpdate": o.hueUpdTime,
				"hueUpdTime": o.hueUpdTime,
				"Alert": o.currentState.Alert,
			}).Info("Cleared hue from alert state")

			o.hueGroups.SetGroupState(o.hueGroup.ID, *o.initialState)
		}
		o.hueUpdTime = time.Now().Add(o.hueUpdInterval)
	}

	log.WithFields(log.Fields{
		"Celsius": *avg.Celsius,
		"Fahrenheit": *avg.Fahrenheit,
		"color": color.Hex(),
		"hueGroupName": hueGroupName,
		"nextHueUpdate": o.hueUpdTime,
	}).Debug("Updated temp indicator")

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
		log.Debugf("%d° C is less than min %d° C...clamping", temp, min)
		temp = min
	}
	if temp > max {
		log.Debugf("%d° C is greater than max %d° C...clamping", temp, max)
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
	}).Debug("Calculating temperature color")

	return color
}