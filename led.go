package main

import (
	"github.com/kidoman/embd"
	"os"
	"os/signal"
	"fmt"
	"time"
)

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
		fmt.Println("Feck Interrupt")
		led.Off()
		signal.Stop(interrupt)
	}()

	for {
		select {
		case <-time.After(250 * time.Millisecond):
			fmt.Println("On")
			if err := led.Toggle(); err != nil {
				panic(err)
			}
		case <-interrupt:
			return
		}
		select {
		case <-time.After(250 * time.Millisecond):
			fmt.Println("Off")
			if err := led.Off(); err != nil {
				panic(err)
			}
		case <-interrupt:
			return
		}
	}
}
