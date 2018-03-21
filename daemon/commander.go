package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
)

// Commander is the main controller of all background goroutines
type Commander struct {
<<<<<<< Updated upstream
	runner
	period   		time.Duration
	tempLogger   	Runnable
	tempIndicator	Runnable
	lightShow 		RunnableTicker
	strip      	  	hardware.WS2801
	options			*framework.CmdOptions
=======
	basicTickable
	runner

	Options     *framework.CmdOptions

	period      time.Duration
	tempLogger  Runnable
	currentShow *RunnableTicker
	lightShows  map[string]RunnableTicker
	strip       hardware.WS2801
>>>>>>> Stashed changes
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
<<<<<<< Updated upstream
func NewCommander(options *framework.CmdOptions) *Commander {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	c := Commander{
		strip: hardware.NewStrandController(),
		options: options,
=======
func NewCommander() *Commander {
	c := Commander{
		Options:    &framework.CmdOptions{},
		strip:      hardware.NewStrandController(),
		lightShows: initializeLightShows(1000000000),
>>>>>>> Stashed changes
	}
	c.runner.tickable = &c
	return &c
}

<<<<<<< Updated upstream
func (r *Commander) getPeriod() time.Duration {
	return time.Second * 10
}

func (r *Commander) setPeriod(period time.Duration)  {
	r.period = period
=======
func initializeLightShows(period time.Duration) map[string]RunnableTicker {
	log.WithField("period", period).Infof("Iniitializing light shows")

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
>>>>>>> Stashed changes
}

// GetName returns a human readable name for this background task
func (r *Commander) GetName() string {
	return "commander"
}

// Start performs initialization before the first tick
<<<<<<< Updated upstream
func (r *Commander) start() {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_entry")

	r.tempLogger = newTemperatureLogger()
	if r.options.TemperatureLoggerEnabled {
		r.EnableTemperatureLogger()
	}

	r.tempIndicator = newTemperatureIndicator()
	if r.options.TemperatureIndicatorEnabled {
		r.EnableTemperatureIndicator()
	}

	r.lightShow = newSimpleShifter()

	if r.options.LightShowEnabled {
		r.EnableLightShow(r.lightShow.tickable.getPeriod())
	}
}

func (r *Commander) EnableTemperatureLogger() {
	if ! r.tempLogger.IsRunning() {
		if err := r.tempLogger.StartBackground(); err != nil {
			log.Error(err.Error())
		}
=======
func (o *Commander) start() error {
	o.tempLogger = newTemperatureLogger()
	if o.Options.TemperatureLoggerEnabled {
		o.EnableTemperatureLogger()
	}

	show, err := o.GetLightShow(o.Options.LightShow)
	if ( err != nil ) {
		return err
	} else {
		p := lights.UpdateGrillLightsParams{
			Name: o.Options.LightShow,
			// The period is in milliseconds so we need to divide
			// the time.Duration by 1000 because it is in nanoseconds
			Period: int64(show.tickable.getPeriod() / 1000),
		}
		if err := o.UpdateGrillLights(&p); err != nil {
			return err
		}
	}

	return nil
}

func (o *Commander) UpdateGrillLights(params *lights.UpdateGrillLightsParams) error {
	show, err := o.GetLightShow(params.Name)
	if ( err != nil ) {
		return err
	} else {
		o.changeLightShow(&show, params.Period)
		return nil
	}
}

func (o *Commander) GetLightShow(name string) (show RunnableTicker, err error) {
	if show, ok := o.lightShows[name]; ok {
		log.WithFields(log.Fields{
			"show": show,
			"name": name,
		}).Infof("Found light show")
		return show, nil
	} else {
		return show, fmt.Errorf("Invalid light show name='%s'", name)
>>>>>>> Stashed changes
	}
}

func (r *Commander) EnableLightShow(period time.Duration) {
	r.lightShow.tickable.setPeriod(period)

	if ! r.lightShow.runnable.IsRunning() {
		if err := r.lightShow.runnable.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) EnableTemperatureIndicator() {
	if ! r.tempIndicator.IsRunning() {
		if err := r.tempIndicator.StartBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

func (r *Commander) DisableLights() {
	r.disableLightShow()
	r.disableTemperatureIndicator()
}

func (r *Commander) disableLightShow() {
	if ! r.lightShow.runnable.IsRunning() {
		if err := r.lightShow.runnable.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}
}

<<<<<<< Updated upstream
func (r *Commander) disableTemperatureIndicator() {
	if ! r.tempIndicator.IsRunning() {
		if err := r.tempIndicator.StopBackground(); err != nil {
			log.Error(err.Error())
=======
func (o *Commander) DisableLights() error {
	if o.currentShow != nil && o.currentShow.runnable.IsRunning() {
		log.WithField("name", o.currentShow.tickable.GetName()).
			Info("Shutting down current light show")
		if err := o.currentShow.runnable.StopBackground(); err != nil {
			return err
>>>>>>> Stashed changes
		}
	}

	return nil
}

<<<<<<< Updated upstream
func (r *Commander) DisableTemperatureLogger() {
	if ! r.tempLogger.IsRunning() {
		if err := r.tempLogger.StopBackground(); err != nil {
			log.Error(err.Error())
=======
func (o *Commander) DisableTemperatureLogger() error {
	if o.tempLogger.IsRunning() {
		log.WithField("type", reflect.TypeOf(o.tempLogger)).
			Info("Shutting down temperature logger")
		if err := o.tempLogger.StopBackground(); err != nil {
			return err
>>>>>>> Stashed changes
		}
	}

	return nil
}

// Stop performs cleanup when the goroutine is exiting
<<<<<<< Updated upstream
func (r *Commander) stop() {
	log.Debug("action=stop")
	defer log.Debug("action=stop")

	if r.tempLogger != nil {
		if err := r.tempLogger.StopBackground(); err != nil {
			log.Error(err.Error())
		}
	}

	log.Info("Clearing all pixels")
	if err := r.strip.Close(); err != nil {
		log.Error(err.Error())
=======
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
		return err
>>>>>>> Stashed changes
	}

	return nil
}

// Tick executes on a ticker schedule
<<<<<<< Updated upstream
func (r *Commander) tick() bool {
	log.Debug("action=tick")
	defer log.Debug("action=tick")

	return true
=======
func (o *Commander) tick() error {
	return nil
>>>>>>> Stashed changes
}
