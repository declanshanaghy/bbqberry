package hardware

import "github.com/lucasb-eyer/go-colorful"


// RED is a pure red integer representation
var RED = colorful.Color{R:1.0, G:0.0, B:0.0}

// GREEN is a pure green integer representation
var GREEN = colorful.Color{R:0.0, G:1.0, B:0.0}

// BLUE is a pure blue integer representation
var BLUE = colorful.Color{R:0.0, G:0.0, B:1.0}

// BLACK is a pure blue integer representation
var BLACK = colorful.Color{R:0.0, G:0.0, B:0.0}


//func GetRed(color int) uint8 {
//	return uint8(color >> 16 & 0xFF)
//}
//
//func SetRed(color *int, r uint8) {
//	*color = (*color & 0x00FFFF) |  (int(r) << 16)
//}
//
//func GetGreen(color int) uint8 {
//	return uint8(color >> 8 & 0xFF)
//}
//
//func SetGreen(color *int, r uint8) {
//	*color = (*color & 0xFF00FF) |  (int(r) << 8)
//}
//
//func GetBlue(color int) uint8 {
//	return uint8(color & 0xFF)
//}
//
//func SetBlue(color *int, r uint8) {
//	*color = (*color & 0xFFFF00) |  int(r)
//}
//
func Color(r, g, b int) colorful.Color {
	return colorful.Color{R: float64(r) / 255.0, G: float64(g) / 255.0, B: float64(b) / 255.0 }
}
