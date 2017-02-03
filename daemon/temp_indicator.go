package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
)

// temperatureIndicator collects and logs temperature metrics
type temperatureIndicator struct {
	runner
	reader	hardware.TemperatureReader
	strip	hardware.WS2801
	errorCount	int
}

// newTemperatureIndicator creates a new temperatureIndicator instance which can be
// run in the background to check average temperature and indicate it visually on the LED strip
func newTemperatureIndicator() *temperatureIndicator {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return &temperatureIndicator{
		reader: hardware.NewTemperatureReader(),
		strip: hardware.NewStrandController(),
	}
}

// StartBackground starts the commander in the background
func (ti *temperatureIndicator) StartBackground() error {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	return ti.runner.startBackground(ti)
}

func (ti *temperatureIndicator) getPeriod() time.Duration {
	return time.Second * 10
}

// GetName returns a human readable name for this background task
func (ti *temperatureIndicator) GetName() string {
	return "temperatureIndicator"
}

// Start performs initialization before the first tick
func (ti *temperatureIndicator) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")
}

// Stop performs cleanup when the goroutine is exiting
func (ti *temperatureIndicator) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")
}

// Tick executes on a ticker schedule
func (ti *temperatureIndicator) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	avg, err := framework.QueryAverageTemperature(ti.getPeriod(), framework.Constants.Hardware.AmbientProbeNumber)
	if err != nil {
		log.Error(err.Error())
		return true
	}

	color := getTempColor(*avg.Celsius)
	if err := ti.strip.SetAllPixels(color); err != nil {
		log.Error(err.Error())
	}
	if err := ti.strip.Update(); err != nil {
		log.Error(err.Error())
	}

	return true
}

func getTempColor(temp float32) int {
	// Map the temperature to a color to be displayed on the LED pixels.
	// cold / min = blue	( 0x0000FF ) =
	// hot / max = red ( 0xFF0000 )
	// green LEDs should never be on.
	// Treat each degree above min as a +1 of the red component and -1 of the blue component
	// Therefore:
	// 		avg temp <= min = pure blue
	// 		avg temp >= max = pure red
	// If the max limit is exceeded a visual indicator should be displayed (i.e. flashing)
	min := framework.Constants.Hardware.MinTempWarnCelsius
	max := framework.Constants.Hardware.MaxTempWarnCelsius

	if temp < min {
		log.Warningf("Temp (%0.2f) < min (%0.2f)...clamping", temp, min)
		temp = min
	}
	if temp > max {
		log.Warningf("Temp (%0.2f) > max (%0.2f)...clamping", temp, max)
		temp = max
	}

	rnge := (max - min)

	corrected := temp - min
	scaled := corrected / rnge
	//offset := int(255 * scaled)

	r := int(255 * scaled)
	b := 0xFF - r

	color := r << 16 | b

	//log.Infof("min=%0.2f, max=%0.2f rnge=%0.2f temp=%0.2f, corrected=%0.2f scaled=%0.2f " +
	//		"(r, b) = (%d, %d) = (%x, %x) color=%x", min, max, rnge, temp, corrected, scaled, r, b, r, b, color)

	return color
}

