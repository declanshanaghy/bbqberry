package hardware

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/kidoman/embd"
)

// WS2801 privdes an interface for communicating with an LED strip which uses the WS2801 chip
type WS2801 interface {
	GetNumPixels() int
	Close() error
	Off() error
	Update() error
	SetPixelRGB(n int, r uint8, g uint8, b uint8) error
	SetPixelColor(n int, color int) error
}

type ws2801Strand struct {
	bus    embd.SPIBus
	pixels []uint8
	data   []uint8
}

// newWS2801 creates a new object capable of communicating with a WS2801 LED strip
func newWS2801(nPixels int, bus embd.SPIBus) WS2801 {
	return &ws2801Strand{
		bus:    bus,
		pixels: make([]uint8, nPixels*3),
		data:   make([]uint8, nPixels*3),
	}
}

func (s *ws2801Strand) GetNumPixels() int {
	return len(s.pixels) / 3
}

func (s *ws2801Strand) Off() error {
	log.Info("action=Off nPixels=%d", s.GetNumPixels())
	for i := 0; i < s.GetNumPixels(); i++ {
		s.SetPixelColor(i, 0)
	}
	return s.Update()
}

func (s *ws2801Strand) Close() error {
	log.Info("action=Close nPixels=%d", s.GetNumPixels())
	s.Off()
	return s.bus.Close()
}

func (s *ws2801Strand) Update() error {
	log.Debugf("action=Update nPixels=%d", s.GetNumPixels())
	copy(s.data, s.pixels)
	return s.bus.TransferAndReceiveData(s.data)
}

func (s *ws2801Strand) SetPixelColor(n int, color int) error {
	log.Debugf("action=SetPixelColor n=%d, color=%#06x", n, color)
	return s.SetPixelRGB(n, uint8(color>>16&0xFF), uint8(color>>8&0xFF), uint8(color&0xFF))
}

func (s *ws2801Strand) SetPixelRGB(n int, r uint8, g uint8, b uint8) error {
	if err := s.validatePixel(n); err != nil {
		return err
	}
	base := n * 3
	log.Debugf("action=SetPixelRGB n=%d base=%d nPixels=%d r=%#02x g=%#02x b=%#02x", n, base, len(s.pixels), r, g, b)
	s.pixels[base] = r
	s.pixels[base+1] = g
	s.pixels[base+2] = b
	return nil
}

func (s *ws2801Strand) validatePixel(n int) (err error) {
	if n < 0 || n > s.GetNumPixels() {
		err = fmt.Errorf("action=invalid pixel=%d, max=%d", n, s.GetNumPixels())
	}
	return err
}