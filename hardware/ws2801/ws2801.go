package ws2801

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/golang/glog"
	"github.com/kidoman/embd"
)

type WS2801 interface {
	GetNumPixels() int
	Close()
	Off()
	Update() error
	SetPixelRGB(n int, r uint8, g uint8, b uint8) error
	SetPixelRGBA(n int, color color.RGBA) error
	SetPixelColor(n int, color int) error
}

type Strand struct {
	bus    embd.SPIBus
	pixels []uint8
	data   []uint8
}

func NewWS2801(nPixels int, bus embd.SPIBus) WS2801 {
	return &Strand{
		bus:    bus,
		pixels: make([]uint8, nPixels*3),
		data:   make([]uint8, nPixels*3),
	}
}

func (s *Strand) GetNumPixels() int {
	return len(s.pixels) / 3
}

func (s *Strand) Off() {
	glog.Info("action=Off nPixels=%d", s.GetNumPixels())
	for i := 0; i < s.GetNumPixels(); i++ {
		s.SetPixelColor(i, 0)
	}
	s.Update()
}

func (s *Strand) Close() {
	glog.Info("action=Close nPixels=%d", s.GetNumPixels())
	s.Off()
	s.bus.Close()
}

func (s *Strand) Update() error {
	glog.V(3).Infof("action=Update nPixels=%d", s.GetNumPixels())
	copy(s.data, s.pixels)
	return s.bus.TransferAndReceiveData(s.data)
}

func (s *Strand) SetPixelRGBA(n int, color color.RGBA) error {
	glog.V(3).Infof("action=SetPixelRGBA n=%d, color=%#06x", n, color)
	return s.SetPixelRGB(n, color.R, color.G, color.B)
}

func (s *Strand) SetPixelColor(n int, color int) error {
	glog.V(3).Infof("action=SetPixelColor n=%d, color=%#06x", n, color)
	return s.SetPixelRGB(n, uint8(color>>16&0xFF), uint8(color>>8&0xFF), uint8(color&0xFF))
}

func (s *Strand) SetPixelRGB(n int, r uint8, g uint8, b uint8) error {
	if err := s.validatePixel(n); err != nil {
		return err
	}
	base := n * 3
	glog.V(3).Infof("action=SetPixelRGB n=%d base=%d nPixels=%d r=%#02x g=%#02x b=%#02x", n, base, len(s.pixels), r, g, b)
	s.pixels[base] = r
	s.pixels[base+1] = g
	s.pixels[base+2] = b
	return nil
}

func (s *Strand) validatePixel(n int) (err error) {
	if n < 0 || n > s.GetNumPixels() {
		err = errors.New(fmt.Sprintf("action=invalid pixel=%d, max=%d", n, s.GetNumPixels()))
	}
	return err
}
