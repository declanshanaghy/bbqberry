package daemon

import (
	"time"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
)

// pulser displays a single colored pulse on the strand
type pulser struct {
	basicTickable
	color			int
	mask			int
	diff			uint8
	strip   		hardware.WS2801
}

// newSimpleShifter creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newPulser(period time.Duration) RunnableTicker {
	t := &pulser{
		strip: hardware.NewStrandController(),
		mask: 0xFFFFFF,
		diff: 1,

	}
	t.Period = period

	return newRunnableTicker(t)
}

// GetName returns a human readable name for this background task
func (o *pulser) GetName() string {
	return "Pulser"
}

// Start performs initialization before the first tick
func (o *pulser) start() {
	o.strip.SetAllPixels(o.color);
}

// Stop performs cleanup when the goroutine is exiting
func (o *pulser) stop() {
	o.strip.Off()
}

func (o *pulser) tick() bool {
	r := hardware.GetRed(o.color)
	hardware.SetRed(&o.color, r + o.diff)

	g := hardware.GetGreen(o.color)
	hardware.SetGreen(&o.color, g + o.diff)

	b := hardware.GetBlue(o.color)
	hardware.SetBlue(&o.color, b + o.diff)

	log.Debugf("color=0x%06x", o.color)
	o.strip.SetAllPixels(o.color);

	return true
}