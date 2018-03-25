package hardware

const (
	// RED is a pure red integer representation
	RED = 0xFF0000

	// GREEN is a pure green integer representation
	GREEN = 0x00FF00

	// BLUE is a pure blue integer representation
	BLUE = 0x0000FF
)

func GetRed(color int) uint8 {
	return uint8(color >> 16 & 0xFF)
}

func SetRed(color *int, r uint8) {
	*color = (*color & 0x00FFFF) |  (int(r) << 16)
}

func GetGreen(color int) uint8 {
	return uint8(color >> 8 & 0xFF)
}

func SetGreen(color *int, r uint8) {
	*color = (*color & 0xFF00FF) |  (int(r) << 8)
}

func GetBlue(color int) uint8 {
	return uint8(color & 0xFF)
}

func SetBlue(color *int, r uint8) {
	*color = (*color & 0xFFFF00) |  int(r)
}

func Color(r, g, b int) int {
	return ((int(r) & 0xFF) << 16) | ((int(g) & 0xFF) << 8) | (int(b) & 0xFF)
}
