package example

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware/ws2801"
	"github.com/kidoman/embd"
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

func Rainbow(nPixels int, bus embd.SPIBus) {
	strand0 := ws2801.NewWS2801(nPixels, bus)
	strand0.GetNumPixels()
	defer strand0.Close()

	log.Infof("action=Rainbow nPixels=%d", strand0.GetNumPixels())
	n := strand0.GetNumPixels()

	for j := 0; j < 256; j++ {
		for i := 0; i < n; i++ {
			r, g, b := wheel(uint8(((i * 256 / n) + j) % 256))
			strand0.SetPixelRGB(i, r, g, b)
		}

		strand0.Update()
		time.Sleep(5 * time.Millisecond)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)
}

func RedGreenBlueRandom(nPixels int, bus embd.SPIBus) {
	strand0 := ws2801.NewWS2801(nPixels, bus)
	defer strand0.Close()

	log.Infof("action=RedGreenBlueRandom nPixels=%d", strand0.GetNumPixels())

	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGBA(i, color.RGBA{R: 0xFF, G: 0x00, B: 0x00})
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGB(i, 0x00, 0xFF, 0x00)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGB(i, 0x00, 0x00, 0xFF)
	}
	strand0.Update()
	time.Sleep(1 * time.Second)

	rand.Seed(time.Now().UnixNano())
	colors := []int{ws2801.RED, ws2801.GREEN, ws2801.BLUE}
	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelColor(i, colors[rand.Intn(3)])
	}
	strand0.Update()
	time.Sleep(1 * time.Second)
}
