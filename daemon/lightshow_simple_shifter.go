package daemon

import (
	"time"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"

)

// simpleShifter displays fancy colors
type simpleShifter struct {
	curled  int32
	lastled int32
	strip   hardware.WS2801
	period	time.Duration
}

// newLightShow creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newSimpleShifter() RunnableTicker {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")

	t := &simpleShifter{
		strip: hardware.NewStrandController(),
		period: time.Second,
	}

	return newRunnableTicker(t)
}

func (r *simpleShifter) getPeriod() time.Duration {
	return r.period
}

func (r *simpleShifter) setPeriod(period time.Duration)  {
	r.period = period
}

// GetName returns a human readable name for this background task
func (r *simpleShifter) GetName() string {
	return "lightshow"
}

// Start performs initialization before the first tick
<<<<<<< Updated upstream
func (r *simpleShifter) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	r.strip.Off()
}

// Stop performs cleanup when the goroutine is exiting
func (r *simpleShifter) stop() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
}

func (r *simpleShifter) tick() bool {
=======
func (o *simpleShifter) start() error {
	return o.strip.Off()
}

// Stop performs cleanup when the goroutine is exiting
func (o *simpleShifter) stop() error {
	return o.strip.Off()
}

func (o *simpleShifter) tick() error {
>>>>>>> Stashed changes
	log.WithFields(log.Fields{
		"curled":  r.curled,
		"lastled": r.lastled,
		"action":  "method_entry",
	}).Info("simpleShifter updating lights")
	defer log.Debug("action=method_exit")

<<<<<<< Updated upstream
	if err := r.strip.SetPixelColor(r.lastled, 0x000000); err != nil {
		log.Error(err.Error())
	}
	if err := r.strip.SetPixelColor(r.curled, 0xFF0000); err != nil {
		log.Error(err.Error())
=======
	if err := o.strip.SetPixelColor(o.lastled, 0x000000); err != nil {
		return err
	}
	if err := o.strip.SetPixelColor(o.curled, 0xFF0000); err != nil {
		return err
>>>>>>> Stashed changes
	}
	r.strip.Update()

	r.lastled = r.curled
	r.curled++

	if r.curled >= r.strip.GetNumPixels() {
		r.curled = 0
	}

	return nil
}
