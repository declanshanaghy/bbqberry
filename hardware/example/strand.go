// + build ignore

package main

import (
	"math/rand"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
)

func wheel(wp uint8) (r, g, b uint8) {
	log.Debugf("action=wheel wp=%d", wp)
	if wp < 85 {
		r = wp * 3
		g = 255 - wp*3
		b = 0
	} else if wp < 170 {
		wp -= 85
		r = 255 - wp*3
		g = 0
		b = wp * 3
	} else {
		wp -= 170
		r = 0
		g = wp * 3
		b = 255 - wp*3
	}
	return
}

// Rainbow cycles the LED strand through the rainbow colors
func Rainbow(strand0 hardware.WS2801) {
	log.Infof("action=Rainbow nPixels=%d", strand0.GetNumPixels())
	n := strand0.GetNumPixels()

	for j := int32(0); j < 256; j++ {
		for i := int32(0); i < n; i++ {
			r, g, b := wheel(uint8(((i * 256 / n) + j) % 256))
			strand0.SetPixelRGB(i, r, g, b)
		}

		strand0.Update()
		time.Sleep(5 * time.Millisecond)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)
}

// RedGreenBlueRandom sets the LED colors to all Red, all Green,
// all Blue, then all random assignments of Red, Green or Blue
func RedGreenBlueRandom(strand0 hardware.WS2801) {

	log.Infof("action=RedGreenBlueRandom nPixels=%d", strand0.GetNumPixels())

	for i := int32(0); i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGB(i, 0xFF, 0x00, 0x00)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	for i := int32(0); i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGB(i, 0x00, 0xFF, 0x00)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	for i := int32(0); i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGB(i, 0x00, 0x00, 0xFF)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	colors := []int{hardware.RED, hardware.GREEN, hardware.BLUE}
	for i := int32(0); i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelColor(i, colors[rand.Intn(3)])
	}
	strand0.Update()
	time.Sleep(1 * time.Second)
}

func main() {
	strand := hardware.NewGrillLightController()
	defer strand.Close()

	RedGreenBlueRandom(strand)
	Rainbow(strand)
}
