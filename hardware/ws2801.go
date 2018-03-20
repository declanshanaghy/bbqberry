package hardware

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/kidoman/embd"
)

// WS2801 privdes an interface for communicating with an LED strip which uses the WS2801 chip
type WS2801 interface {
	GetPixels() []byte
	GetNumPixels() int32
	Close() error
	Off() error
	Update() error
	SetPixelRGB(n int32, r uint8, g uint8, b uint8) error
	SetPixelColor(n int32, color int) error
	SetAllPixels(color int) error
}

type ws2801Strand struct {
	bus    embd.SPIBus
	pixels []byte
	data   []byte
}

// newWS2801 creates a new object capable of communicating with a WS2801 LED strip
func newWS2801(nPixels int32, bus embd.SPIBus) WS2801 {
	return &ws2801Strand{
		bus:    bus,
		pixels: make([]uint8, int(nPixels)*3),
		data:   make([]uint8, int(nPixels)*3),
	}
}

func (s *ws2801Strand) GetPixels() []byte {
	return s.pixels
}

func (s *ws2801Strand) GetNumPixels() int32 {
	return int32(len(s.pixels) / 3)
}

func (s *ws2801Strand) Off() error {
	log.Infof("action=Off nPixels=%d", s.GetNumPixels())
	for i := int32(0); i < s.GetNumPixels(); i++ {
		s.SetPixelColor(i, 0)
	}
	return s.Update()
}

func (s *ws2801Strand) Close() error {
	log.Infof("action=Close nPixels=%d", s.GetNumPixels())
	s.Off()
	return s.bus.Close()
}

func (s *ws2801Strand) Update() error {
	copy(s.data, s.pixels)

	pixels := make([]string, s.GetNumPixels())
	for i:=0; i<int(s.GetNumPixels()); i++ {
		pixels[i] = fmt.Sprintf("0x%06x", Color(int(s.data[i]), int(s.data[i * 3 + 1]), int(s.data[i * 3 + 2])))
	}
	//log.Debugf("Update pixels=%v", pixels)

	_, err := s.bus.Write(s.data)
	return err
}

func (s *ws2801Strand) SetPixelColor(n int32, color int) error {
	//log.Debugf("action=SetPixelColor n=%d, color=%#06x", n, color)
	return s.SetPixelRGB(n, uint8(color>>16&0xFF), uint8(color>>8&0xFF), uint8(color&0xFF))
}

func (s *ws2801Strand) SetAllPixels(color int) error {
	//log.Debugf("action=SetAllPixels n=%d, color=%#06x", s.GetNumPixels(),color)
	r := GetRed(color)
	g := GetGreen(color)
	b := GetBlue(color)

	// Set the pixel values
	for i := int32(0); i < s.GetNumPixels(); i++ {
		if err := s.SetPixelRGB(i, r, g, b); err != nil {
			return err
		}
	}

	// Applyt the update
	if err := s.Update(); err != nil {
		return err
	}

	return nil
}

func (s *ws2801Strand) SetPixelRGB(n int32, r uint8, g uint8, b uint8) error {
	if err := s.validatePixel(n); err != nil {
		return err
	}
	base := n * 3
	//log.Debugf("action=SetPixelRGB n=%d base=%d nPixels=%d r=%#02x g=%#02x b=%#02x",
	//	n, base, len(s.pixels), r, g, b)
	s.pixels[base] = r
	s.pixels[base+1] = g
	s.pixels[base+2] = b
	return nil
}

func (s *ws2801Strand) validatePixel(n int32) (err error) {
	if n < int32(0) || n >= s.GetNumPixels() {
		err = fmt.Errorf("action=invalid pixel=%d, max=%d", n, s.GetNumPixels())
	}
	return err
}
