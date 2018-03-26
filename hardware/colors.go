package hardware

import (
	"github.com/lucasb-eyer/go-colorful"
)


// RED is a pure red integer representation
var RED = colorful.Color{R:1.0, G:0.0, B:0.0}

// GREEN is a pure green integer representation
var GREEN = colorful.Color{R:0.0, G:1.0, B:0.0}

// BLUE is a pure blue integer representation
var BLUE = colorful.Color{R:0.0, G:0.0, B:1.0}

// BLACK is a pure blue integer representation
var BLACK = colorful.Color{R:0.0, G:0.0, B:0.0}


func Color(r, g, b int) colorful.Color {
	return colorful.Color{R: float64(r) / 255.0, G: float64(g) / 255.0, B: float64(b) / 255.0 }
}

func ColorToHue(color colorful.Color) (uint16, uint8, uint8) {
	m := 65535.0 / 360

	h, s, l := color.Hsl()

	hi := uint16(h * m)
	si := uint8(254 * s)
	li := uint8(254 * l)

	//log.WithFields(log.Fields{
	//	"color": color.Hex(),
	//	"(h, s, l)": fmt.Sprintf("(%0.2f, %0.2f, %0.2f)", h, s, l),
	//	"(hi, si, li)": fmt.Sprintf("(%d, %d, %d)", hi, si, li),
	//}).Debug("Mapped color to hue")

	return hi, si, li
}

