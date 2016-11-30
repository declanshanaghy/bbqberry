package main

import (
	"flag"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/declanshanaghy/bbqberry/ws2801"
)

func main() {
	flag.Parse()

	ws2801.Rainbow(25)
	//BlinkLED()
}

