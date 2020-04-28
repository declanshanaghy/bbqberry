package framework

import (
	"os"

	"github.com/declanshanaghy/bbqberry/models"
	"encoding/json"
	"io/ioutil"
	"github.com/Polarishq/middleware/framework/log"
)

// DefaultDB is the default database name that should be used if an override is not provided
const DefaultDB = "bbqberry"

var Enabled = true
var Disabled = false

const HUE_KEY = "5EPXHOGHzm7TGha3IFumdF2bTLdcwuae-21iQguC"
const HUE_ALERT_GROUP = "Patio"
const AWS_DEFAULT_REGION = "us-east-1"

type config struct {
	ServiceName string
	Version     string
	Stub        bool
	Hardware    models.HardwareConfig
	options 	*CmdOptions
}

// Config contains information about the running service
var Config config

func init() {
	stub := false
	if os.Getenv("STUB") != "" {
		stub = true
	}

	/******************************** BEGIN PSEUDO CONSTANTS *********************************************/
	// Electrical Config
	vcc := float32(3.3)
	analogMax := int32(26453)

	// Accessories
	nPixels := int32(25)

	hwCfg := models.HardwareConfig{
		NumLedPixels: &nPixels,
		Vcc:          &vcc,
		AnalogMax:    &analogMax,
	}

	Config = config{
		ServiceName: "bbqberry",
		Version:     "v1",
		Stub:        stub,
		Hardware:    hwCfg,
	}
}

func (o*config) Initialize(options *CmdOptions) error {
	o.options = options

	text, err := ioutil.ReadFile(options.GetProbesConf())
	if err != nil {
		return err
	}

	err = json.Unmarshal(text, &o.Hardware.Probes)
	if err != nil {
		return err
	}

	return nil
}

func (o*config) Save() error {
	t, err := json.MarshalIndent(o.Hardware.Probes, "  ", "  ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(o.options.GetProbesConf(), t, 0644); err != nil {
		return err
	}

	return nil
}

func (o *config) GetEnabledProbeIndexes() *[]int32 {
	enabled := make([]int32, 0)

	for i, probe := range o.Hardware.Probes {
		if *probe.Enabled {
			enabled = append(enabled, int32(i))
		}
	}

	log.WithField("enabled", enabled).Info("GetEnabledProbeIndexes")

	return &enabled
}

func (o *config) GetEnabledProbes() []*models.TemperatureProbe {
	enabled := make([]*models.TemperatureProbe, 0)

	for _, probe := range o.Hardware.Probes {
		if *probe.Enabled {
			enabled = append(enabled, probe)
		}
	}

	return enabled
}

