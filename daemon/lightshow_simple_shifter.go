package daemon

import (
	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
)

type simpleShifter struct {
	strip   hardware.WS2801
	curled  int32
	lastled int32
}

func (r *simpleShifter) init() {
	r.strip = hardware.NewStrandController()
}

func (r *simpleShifter) tick() {
	log.WithFields(log.Fields{
		"curled":  r.curled,
		"lastled": r.lastled,
		"action":  "method_entry",
	}).Info("simpleShifter updating lights")
	defer log.Debug("action=method_exit")

	if err := r.strip.SetPixelColor(r.lastled, 0x000000); err != nil {
		log.Error(err.Error())
	}
	if err := r.strip.SetPixelColor(r.curled, 0xFF0000); err != nil {
		log.Error(err.Error())
	}

	r.lastled = r.curled
	r.curled++

	if r.curled >= r.strip.GetNumPixels()+100 {
		r.curled = 0
		r.lastled = 0
	}
}
