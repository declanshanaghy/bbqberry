package daemon

import (
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/framework"
	"reflect"
	"github.com/declanshanaghy/bbqberry/restapi/operations/lights"
	"fmt"
	"github.com/declanshanaghy/bbqberry/restapi/operations/system"
	"os/exec"
	"github.com/heatxsink/go-hue/portal"
	"github.com/declanshanaghy/bbqberry/models"
)

// Commander is the main controller of all background goroutines
type Commander struct {
	runner
	basicTickable

	Options     *framework.CmdOptions

	period      time.Duration
	tempLoggers [2]Runnable
	currentShow *RunnableTicker
	lightShows  map[string]RunnableTicker
	strip       hardware.WS2801
	huePortal   *portal.Portal
}

// NewCommander creates a Commander instance which can be
// used to query and control all background processes.
// e.g: Temperature logger, temperature monitor
func NewCommander() *Commander {
	c := Commander{
		Options: 	framework.NewCmdOptions(),
		strip: 		hardware.NewStrandController(),
	}
	c.runner.tickable = &c
	c.runner.tickable.setPeriod(time.Second)
	return &c
}

// GetName returns a human readable name for this background task
func (o *Commander) GetName() string {
	return "Commander"
}

func (o *Commander) initializeHue() error {
	pp, err := portal.GetPortal()
	if err != nil {
		return err
	}
	o.huePortal = &pp[0]

	return nil
}

// Start performs initialization before the first tick
func (o *Commander) start() (error) {
	err := o.initializeHue()
	if err != nil {
		log.WithField("err", err).Error("Unable to initialize hue, functionality will be disabled")
	}

	o.initializeLightShows()

	o.tempLoggers = [2]Runnable {
		newInfluxDBLoggerRunnable(),
		newDynamoDBLoggerRunnable(),
	}
	if o.Options.TemperatureLoggerEnabled {
		o.EnableTemperatureLogging()
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

	_, err = o.UpdateGrillLights(&p)

	log.Info("Commander up and running")

	return err
}

// Stop performs cleanup when the goroutine is exiting
func (o *Commander) stop() error {
	if err := o.DisableTemperatureLogging(); err != nil {
		log.Error(err.Error())
	}

	if err := o.DisableLights(); err != nil {
		log.Error(err.Error())
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

func (o* Commander) initializeLightShows() {
	shows := []RunnableTicker {
		newSimpleShifter(time.Second),
		newTemperatureIndicator(o.huePortal),
		newRainbow(time.Millisecond * 100),
	}
	lightShows := make(map[string]RunnableTicker)
	for _, show := range(shows) {
		lightShows[show.tickable.GetName()] = show
	}

	o.lightShows = lightShows
}

func (o* Commander) getLightShow(name string) (RunnableTicker, error) {
	show, ok := o.lightShows[o.Options.LightShow]
	if ok {
		return show, nil
	} else {
		return show, fmt.Errorf("Unable to find light show with name='%s'", name)
	}
}

func (o *Commander) UpdateGrillLights(params *lights.UpdateGrillLightsParams) (bool, error) {
	if show, ok := o.lightShows[params.Name]; ok {
		o.changeLightShow(&show, params.Period)
		return true, nil
	}  else {
		return false, fmt.Errorf("Invalid light show %s", params.Name)
	}
}

func (o *Commander) ShutdownSystem(params *system.ShutdownParams) (*models.Shutdown, error) {
	tNow := time.Now().Local()
	tShdn := tNow.Add(time.Minute * time.Duration(1))

	//Shutdown uses time format of hh:mm
	tShutdownAbbr := fmt.Sprintf("%04d/%02d/%02d %02d:%02d", tShdn.Year(), tShdn.Month(),
		tShdn.Day(), tShdn.Hour(), tShdn.Minute())
	tShutdown := fmt.Sprintf("%02d:%02d", tShdn.Hour(), tShdn.Minute())

	tActual, err := time.ParseInLocation("2006/01/02 15:04" , tShutdownAbbr, time.Local)
	if err != nil {
		return nil, err
	}

	tDiff := tActual.Sub(tNow).Round(time.Second)

	log.WithFields(log.Fields{
		"tNow": tNow,
		"tShdn": tShdn,
		"tActual": tActual,
		"tDiff": tDiff,
		"tShutdown": tShutdown,
	}).Info("Shutting down at specified time")

	if ! framework.Config.Stub {
		_, err = exec.Command("/usr/bin/sudo", "/sbin/shutdown", "-h", tShutdown,
			"BBQBerry initiated shutdown").Output()
		if err != nil {
			return nil, err
		}
	}

	t := tActual.String()
	m := fmt.Sprintf("Will shutdown in %s", tDiff)
	shdn := models.Shutdown{
		ShutdownTime: &t,
		Message: &m,
	}
	return &shdn, nil
}

func (o *Commander) changeLightShow(lightShow *RunnableTicker, period int64) error {
	o.DisableLights()

	o.currentShow = lightShow
	o.currentShow.tickable.setPeriod(time.Duration(period) * time.Microsecond)


	log.WithFields(log.Fields{
		"name": o.currentShow.tickable.GetName(),
	}).Info("Change light show")
	if err := o.currentShow.runnable.StartBackground(); err != nil {
		return err
	}

	return nil
}

func (o *Commander) EnableTemperatureLogging() error {
	for _, tempLogger := range(o.tempLoggers) {
		if ! tempLogger.IsRunning() {
			log.WithField("type", reflect.TypeOf(tempLogger)).
				Info("Enabling temperature logger")
			if err := tempLogger.StartBackground(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *Commander) DisableTemperatureLogging() error {
	for _, tempLogger := range(o.tempLoggers) {
		log.WithFields(log.Fields{
			"running": tempLogger.IsRunning(),
		}).Info("Disabling temperature logger")

		if tempLogger.IsRunning() {
			if err := tempLogger.StopBackground(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *Commander) DisableLights() error {
	if o.currentShow != nil {
		log.WithFields(log.Fields{
			"name": o.currentShow.tickable.GetName(),
			"running": o.currentShow.runnable.IsRunning(),
		}).Info("Disabling lights")

		if o.currentShow.runnable.IsRunning() {
			if err := o.currentShow.runnable.StopBackground(); err != nil {
				return err
			}
		}
	}

	return nil
}

