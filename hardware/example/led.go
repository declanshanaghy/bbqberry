package example

import (
	"os"
	"os/signal"
	"time"

	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/kidoman/embd"
)

// BlinkLED blinks the on board led continuously until a SIGKILL is received
func BlinkLED() {
	if err := embd.InitLED(); err != nil {
		panic(err)
	}
	defer embd.CloseLED()

	led, err := embd.NewLED(0)
	if err != nil {
		panic(err)
	}
	defer func() {
		led.Off()
		led.Close()
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer func() {
		led.Off()
		signal.Stop(interrupt)
	}()

	for {
		select {
		case <-time.After(250 * time.Millisecond):
			log.Infof("action=on led=0")
			if err := led.Toggle(); err != nil {
				panic(err)
			}
		case <-interrupt:
			return
		}
		select {
		case <-time.After(250 * time.Millisecond):
			log.Infof("action=off led=0")
			if err := led.Off(); err != nil {
				panic(err)
			}
		case <-interrupt:
			return
		}
	}
}
