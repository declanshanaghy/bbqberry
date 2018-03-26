package daemon

import (
	"time"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
)

// simpleShifter displays fancy colors
type simpleShifter struct {
	basicTickable
	curled  int32
	lastled int32
	strip   hardware.WS2801
}

// newSimpleShifter creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newSimpleShifter(period time.Duration) RunnableTicker {
	t := &simpleShifter{
		strip: hardware.NewStrandController(),
	}
	t.Period = period

	return newRunnableTicker(t)
}

// GetName returns a human readable name for this background task
func (o *simpleShifter) GetName() string {
	return "Simple Shifter"
}

// Start performs initialization before the first tick
func (o *simpleShifter) start() error {
	log.WithFields(log.Fields{
		"name": o.GetName(),
		"period": o.getPeriod(),
	}).Info("Starting tickable execution")
	return o.strip.Off()
}

// Stop performs cleanup when the goroutine is exiting
func (o *simpleShifter) stop() error {
	log.WithFields(log.Fields{
		"name": o.GetName(),
	}).Info("Stopping tickable execution")
	return o.strip.Off()
}

func (o *simpleShifter) tick() error {
	log.WithFields(log.Fields{
		"curled":  o.curled,
		"lastled": o.lastled,
		"action":  "method_entry",
	}).Debug("simpleShifter updating lights")

	if err := o.strip.SetPixelColor(o.lastled, hardware.BLACK); err != nil {
		log.Error(err.Error())
	}
	if err := o.strip.SetPixelColor(o.curled, hardware.GREEN); err != nil {
		log.Error(err.Error())
	}
	if err := o.strip.Update(); err != nil {
		return err
	}

	o.lastled = o.curled
	o.curled++

	if o.curled >= o.strip.GetNumPixels() {
		o.curled = 0
	}

	return nil
}