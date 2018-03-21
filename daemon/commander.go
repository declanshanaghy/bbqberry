package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
	"reflect"
	"github.com/declanshanaghy/bbqberry/restapi/operations/lights"
	"fmt"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	basicTickable

	Options     *framework.CmdOptions

	period      time.Duration
	tempLogger  Runnable
	currentShow *RunnableTicker
	lightShows  map[string]RunnableTicker
	strip       hardware.WS2801
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander() *Commander {
	c := Commander{
		Options: 	framework.NewCmdOptions(),
		strip: 		hardware.NewStrandController(),
		lightShows: getLightShows(1000000000),
	}
	c.runner.tickable = &c
	c.runner.tickable.setPeriod(time.Second)
	return &c
}

func getLightShows(period time.Duration) map[string]RunnableTicker {
	shows := []RunnableTicker {
		newPulser(period),
		newSimpleShifter(period),
		newTemperatureIndicator(),
		newRainbow(period),
	}
	lightShows := make(map[string]RunnableTicker)
	for _, show := range(shows) {
		lightShows[show.tickable.GetName()] = show
	}
	return lightShows
}

// GetName returns a human readable name for this background task
func (o *Commander) GetName() string {
	return "Commander"
}

// Start performs initialization before the first tick
func (o *Commander) start() error {
	o.tempLogger = newTemperatureLoggerRunnable()
	if o.Options.TemperatureLoggerEnabled {
		o.EnableTemperatureLogger()
	}

	show, err := o.getLightShow(o.Options.LightShow)
	if err != nil {
		return err
	}

	p := lights.UpdateGrillLightsParams{
		Name: o.Options.LightShow,
		// The period is in milliseconds so we need to divide
		// the time.Duration by 1000 because it is in nanoseconds
		Period: int64(show.tickable.getPeriod() / 1000),
	}
	return o.UpdateGrillLights(&p)
}

func (o* Commander) getLightShow(name string) (RunnableTicker, error) {
	show, ok := o.lightShows[o.Options.LightShow]
	if ok {
		return show, nil
	} else {
		return show, fmt.Errorf("Unable to find light show with name='%s'", name)
	}
}

func (o *Commander) UpdateGrillLights(params *lights.UpdateGrillLightsParams) error {
	if show, ok := o.lightShows[params.Name]; ok {
		o.changeLightShow(&show, params.Period)
		return nil
	}  else {
		return fmt.Errorf("Invalid light show %s", params.Name)
	}
}

func (o *Commander) changeLightShow(lightShow *RunnableTicker, period int64) error {
	o.DisableLights()

	o.currentShow = lightShow
	o.currentShow.tickable.setPeriod(time.Duration(period) * time.Microsecond)

	if err := o.currentShow.runnable.StartBackground(); err != nil {
		return err
	}

	return nil
}

func (o *Commander) EnableTemperatureLogger() error {
	if ! o.tempLogger.IsRunning() {
		log.WithField("type", reflect.TypeOf(o.tempLogger)).
			Info("Enabling temperature logger")
		if err := o.tempLogger.StartBackground(); err != nil {
			return err
		}
	}

	return nil
}

func (o *Commander) DisableLights() error {
	if o.currentShow != nil && o.currentShow.runnable.IsRunning() {
		log.WithField("type", o.currentShow.tickable.GetName()).
			Info("Shutting down currentShow")
		if err := o.currentShow.runnable.StopBackground(); err != nil {
			return err
		}
	}

	return nil
}

func (o *Commander) DisableTemperatureLogger() error {
	if o.tempLogger.IsRunning() {
		log.WithField("type", reflect.TypeOf(o.tempLogger)).
			Info("Shutting down temperature logger")
		if err := o.tempLogger.StopBackground(); err != nil {
			return err
		}
	}

	return nil
}

// Stop performs cleanup when the goroutine is exiting
func (o *Commander) stop() error {
	if o.tempLogger != nil {
		if err := o.tempLogger.StopBackground(); err != nil {
			return err
		}
	}

	if err := o.DisableLights(); err != nil {
		return err
	}

	log.Info("Clearing all pixels")
	if err := o.strip.Close(); err != nil {
		log.Error(err.Error())
	}

	return nil
}

// Tick executes on a ticker schedule
func (o *Commander) tick() error {
	return nil
}