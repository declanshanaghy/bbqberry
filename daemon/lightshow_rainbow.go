package daemon

import (
	"time"
	"github.com/declanshanaghy/bbqberry/hardware"
)

// rainbow displays a single colored pulse on the strand
type rainbow struct {
	basicTickable
	j		int
	strip   hardware.WS2801
}

// newSimpleShifter creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newRainbow(period time.Duration) RunnableTicker {
	t := &rainbow{
		strip: hardware.NewStrandController(),
	}
	t.Period = time.Millisecond * 1

	return newRunnableTicker(t)
}

// GetName returns a human readable name for this background task
func (o *rainbow) GetName() string {
	return "Rainbow"
}

// Start performs initialization before the first tick
func (o *rainbow) start() error {
	return o.strip.Off()
}

// Stop performs cleanup when the goroutine is exiting
func (o *rainbow) stop() error {
	return o.strip.Off()
}
/*
# Define the wheel function to interpolate between different hues.
def wheel(pos):
    if pos < 85:
        return Adafruit_WS2801.RGB_to_color(pos * 3, 255 - pos * 3, 0)
    elif pos < 170:
        pos -= 85
        return Adafruit_WS2801.RGB_to_color(255 - pos * 3, 0, pos * 3)
    else:
        pos -= 170
        return Adafruit_WS2801.RGB_to_color(0, pos * 3, 255 - pos * 3)
*/
func (o *rainbow) wheel(pos int) int {
	if pos < 85 {
		return hardware.Color(pos * 3, 255 - pos * 3, 0)
	} else if pos < 170 {
		pos -= 85
		return hardware.Color(255 - pos * 3, 0, pos * 3)
	} else {
		pos -= 170
	}

	return hardware.Color(0, pos * 3, 255 - pos * 3)
}
/*
def rainbow_colors(pixels, wait=0.1):
    for j in range(256):  # one cycle of all 256 colors in the wheel
        for i in range(pixels.count()):
            pixels.set_pixel(i, wheel(((256 // pixels.count() + j)) % 256))
        pixels.show()
        if wait > 0:
            time.sleep(wait)
 */
func (o *rainbow) tick() error {
	nPixels := int(o.strip.GetNumPixels())

	// This "w" cycles all pixels through the same color
	//w := ((256 / nPixels + o.j)) % 256

	for i := 0; i < nPixels; i++ {
		// This "w" cycles each pixel individually
		w := ((i * 256 / nPixels) + o.j) % 256
		color := o.wheel(w)
		o.strip.SetPixelColor(int32(i), color)
	}

	//log.Infof("j=%d, w=%d, pixels=%v", o.j, w, o.strip.GetPixels())
	if err := o.strip.Update(); err != nil {
		return err
	}

	o.j++
	if o.j == 256 {
		o.j = 0
	}

	return nil
}

