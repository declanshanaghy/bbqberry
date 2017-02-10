package util

import "math"

// Rounds the given number to the requested precision
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// Rounds the given float as per standard math rounding rules
func RoundFloat32ToInt32(v float32) int32 {
	return int32(Round(float64(v), .5, 0))
}
