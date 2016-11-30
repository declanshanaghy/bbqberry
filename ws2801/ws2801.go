package ws2801

import (
	"github.com/kidoman/embd"
	"github.com/golang/glog"
	"image/color"
	"fmt"
)

type WS2801 interface {
	GetNumPixels() int
	Close()
	Off()
	Update() error
	SetPixelRGB(n int, r uint8, g uint8, b uint8)
	SetPixelRGBA(n int, color color.RGBA)
	SetPixelColor(n int, color int)
}

type Strand struct {
	channel byte
	bus embd.SPIBus
	pixels []uint8
	data []uint8
}

func NewWS2801(nPixels int, channel byte) WS2801 {
	s := Strand{channel:channel}
	s.init(nPixels)
	return &s
}

func (s *Strand) init(nPixels int) {
	glog.Info("action=Init nPixels=%s", nPixels)

	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	s.bus = embd.NewSPIBus(embd.SPIMode0, s.channel, 1000000, 8, 0)

	s.pixels = make([]uint8, nPixels * 3)
	s.data = make([]uint8, nPixels * 3)
}

func (s *Strand) GetNumPixels() int {
	return len(s.pixels) / 3
}

func (s *Strand) Off() {
	glog.Info("action=Off nPixels=%s", s.GetNumPixels())
	for i := 0; i < s.GetNumPixels(); i++ {
		s.SetPixelColor(i, 0)
	}
	s.Update()
}

func (s *Strand) Close() {
	glog.Info("action=Close nPixels=%s", s.GetNumPixels())
	s.Off()
	s.bus.Close()
	embd.CloseSPI()
}

func (s *Strand) Update() error {
	glog.V(3).Infof("action=Update nPixels=%d", s.GetNumPixels())
	copy(s.data, s.pixels)
	return s.bus.TransferAndReceiveData(s.data)
}

func (s *Strand) SetPixelRGBA(n int, color color.RGBA) {
	glog.V(3).Infof("action=SetPixelRGBA n=%d, color=%#06x", n, color)
	s.SetPixelRGB(n, color.R, color.G, color.B)
}

func (s *Strand) SetPixelColor(n int, color int) {
	glog.V(3).Infof("action=SetPixelColor n=%d, color=%#06x", n, color)
	s.SetPixelRGB(n, uint8(color>>16 & 0xFF), uint8(color>>8 & 0xFF), uint8(color & 0xFF))
}

func (s *Strand) SetPixelRGB(n int, r uint8, g uint8, b uint8) {
	s.ValidatePixel(n)
	base := n * 3
	glog.V(3).Infof("action=SetPixelRGB n=%d base=%d nPixels=%d r=%#02x g=%#02x b=%#02x", n, base, len(s.pixels), r, g, b)
	s.pixels[base] = r
	s.pixels[base+1] = g
	s.pixels[base+2] = b
}

func (s *Strand) ValidatePixel(n int) {
	if n > s.GetNumPixels() {
		msg := fmt.Sprintf("action=invalid pixel=%d, max=%d", n, s.GetNumPixels())
		glog.Error(msg)
		panic(msg)
	}
}