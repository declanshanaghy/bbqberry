package hardware

import (
	"fmt"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/kidoman/embd"
	"github.com/lucasb-eyer/go-colorful"
)

// WS2801 privdes an interface for communicating with an LED strip which uses the WS2801 chip
type WS2801 interface {
	GetPixels() []byte
	GetColors() ([]*colorful.Color, error)
	GetNumPixels() int32
	Close() error
	Off() error
	Update() error
	SetPixelRGB(n int32, r uint8, g uint8, b uint8) error
	SetPixelColor(n int32, color colorful.Color) error
	SetAllPixels(color colorful.Color) error
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

func (o *ws2801Strand) GetPixels() []byte {
	return o.pixels
}

func (o *ws2801Strand) GetColors() ([]*colorful.Color, error) {
	colors := make([]*colorful.Color, o.GetNumPixels())

	for i := int32(0); i < o.GetNumPixels(); i++ {
		c, err := o.GetPixelColor(i)
		if err != nil {
			return nil, err
		}
		colors[i] = c
	}
	
	return colors, nil
}

func (o *ws2801Strand) GetNumPixels() int32 {
	return int32(len(o.pixels) / 3)
}

func (o *ws2801Strand) Off() error {
	log.Infof("action=Off nPixels=%d", o.GetNumPixels())
	for i := int32(0); i < o.GetNumPixels(); i++ {
		o.SetPixelColor(i, BLACK)
	}
	return o.Update()
}

func (o *ws2801Strand) Close() error {
	log.Infof("action=Close nPixels=%d", o.GetNumPixels())
	o.Off()
	return o.bus.Close()
}

func (o *ws2801Strand) Update() error {
	copy(o.data, o.pixels)
	_, err := o.bus.Write(o.data)
	return err
}

func (o *ws2801Strand) SetPixelColor(n int32, color colorful.Color) error {
	if err := o.validatePixel(n); err != nil {
		return err
	}
	r, g, b := color.RGB255()
	return o.SetPixelRGB(n, r, g, b)
}

func (o *ws2801Strand) GetPixelColor(n int32) (*colorful.Color, error) {
	if err := o.validatePixel(n); err != nil {
		return nil, err
	}
	base := n * 3
	r := int(o.pixels[base])
	g := int(o.pixels[base+1])
	b := int(o.pixels[base+2])

	c := Color(r, g, b)
	return &c, nil
}

func (o *ws2801Strand) SetAllPixels(color colorful.Color) error {
	//log.Debugf("action=SetAllPixels n=%d, color=%#06x", o.GetNumPixels(), color)
	r, g, b := color.RGB255()

	for i := int32(0); i < o.GetNumPixels(); i++ {
		if err := o.SetPixelRGB(i, r, g, b); err != nil {
			return err
		}
	}

	// Applyt the update
	if err := o.Update(); err != nil {
		return err
	}

	return nil
}

func (o *ws2801Strand) SetPixelRGB(n int32, r uint8, g uint8, b uint8) error {
	if err := o.validatePixel(n); err != nil {
		return err
	}
	base := n * 3
	o.pixels[base] = r
	o.pixels[base+1] = g
	o.pixels[base+2] = b
	return nil
}

func (o *ws2801Strand) validatePixel(n int32) (err error) {
	if n < int32(0) || n >= o.GetNumPixels() {
		err = fmt.Errorf("action=invalid pixel=%d, max=%d", n, o.GetNumPixels())
	}
	return err
}