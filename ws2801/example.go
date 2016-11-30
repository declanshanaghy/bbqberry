package ws2801

import (
	"github.com/golang/glog"
	"image/color"
	"time"
	"math/rand"
)

/*
  Ported from example python code:

 def colorwipe(self, pixels, c, delay):
	for i in range(len(pixels)):
		self.setpixelcolor(pixels, i, c)
		self.writestrip(pixels)
		time.sleep(delay)

 def Wheel(self, WheelPos):
	if (WheelPos < 85):
   		return self.Color(WheelPos * 3, 255 - WheelPos * 3, 0)
	elif (WheelPos < 170):
   		WheelPos -= 85;
   		return self.Color(255 - WheelPos * 3, 0, WheelPos * 3)
	else:
		WheelPos -= 170;
		return self.Color(0, WheelPos * 3, 255 - WheelPos * 3)

 def rainbowCycle(self, pixels, wait):
	for j in range(256): # one cycle of all 256 colors in the wheel
    	   for i in range(len(pixels)):
 # tricky math! we use each pixel as a fraction of the full 96-color wheel
 # (thats the i / strip.numPixels() part)
 # Then add in j which makes the colors go around per pixel
 # the % 96 is to make the wheel cycle around
      		self.setpixelcolor(pixels, i, self.Wheel( ((i * 256 / len(pixels)) + j) % 256) )
	   self.writestrip(pixels)
	   time.sleep(wait)

 def cls(self, pixels):
          for i in range(len(pixels)):
                self.setpixelcolor(pixels, i, self.Color(0,0,0))
                self.writestrip(pixels)
 */
func wheel(wp uint8) (r, g, b uint8) {
	glog.V(4).Infof("action=wheel wp=%d", wp)
	if wp < 85 {
		r = wp * 3
		g = 255 - wp * 3
		b = 0
	} else if wp < 170 {
		wp -= 85
		r = 255 - wp * 3
		g = 0
		b = wp * 3
	} else {
		wp -= 170
		r = 0
		g = wp * 3
		b = 255 - wp * 3
	}
	return
}

func Rainbow(nPixels int) {
	strand0 := NewWS2801(nPixels, 0)
	strand0.GetNumPixels()
	defer strand0.Close()

	glog.Infof("action=Rainbow nPixels=%d", strand0.GetNumPixels())
	n := strand0.GetNumPixels()

	for j := 0; j < 256; j++ {
		glog.V(4).Infof("action=outer j=%d", j)

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

func RedGreenBlueRandom(nPixels int) {
	strand0 := NewWS2801(nPixels, 0)
	defer strand0.Close()

	glog.Infof("action=RedGreenBlueRandom nPixels=%d", strand0.GetNumPixels())

	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelRGBA(i, color.RGBA{R:0xFF, G:0x00, B:0x00})
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
	colors := []int{RED, GREEN, BLUE}
	for i := 0; i < strand0.GetNumPixels(); i++ {
		strand0.SetPixelColor(i, colors[rand.Intn(3)])
	}
	strand0.Update()
	time.Sleep(1 * time.Second)
}