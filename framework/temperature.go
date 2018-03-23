package framework

import (
	"github.com/declanshanaghy/bbqberry/util"
)


// AdafruitAD8495ThermocoupleVtoKCF converts the voltage read from the Adafruit Thermocouple breakout board
// to temperatures in Kelvin, Celsius and Fahrenheit
func AdafruitAD8495ThermocoupleVtoKCF(v float32) (tempK int32, tempC int32, tempF int32) {
	// https://www.adafruit.com/product/1778
	// Analog Output K-Type Thermocouple Amplifier - AD8495 Breakout
	// PRODUCT ID: 1778
	// Temperature = (Vout - 1.25) / 0.005 V
	// e.g:
	// v = 1.5VDC
	// The temperature is (1.5 - 1.25) / 0.005 = 50°C

	fTempC := (v - 1.25) / 0.005
	fTempK, fTempF := ConvertCToKF(fTempC)
	tempC = util.RoundFloat32ToInt32(fTempC)
	tempF = util.RoundFloat32ToInt32(fTempF)
	tempK = util.RoundFloat32ToInt32(fTempK)
	return
}

// ConvertVoltageToTemperature converts the given voltage value to its corresponding temperature values
//func ConvertVoltageToTemperature(v float32) (tempK int32, tempC int32, tempF int32) {
//	return AdafruitAD8495ThermocoupleVtoKCF(v)
//}

// ConvertAnalogToVoltage converts an analog reading to its corresponding voltage value
func ConvertAnalogToVoltage(analog int32) float32 {
	hwCfg := Constants.Hardware
	vcc := *hwCfg.Vcc
	// volts per analog unit = VCC / Analog max
	amax := float32(*hwCfg.AnalogMax)
	avpu := vcc / amax
	return float32(analog) * avpu
}

// ConvertVoltageToAnalog converts the given voltage to its corresponding analog value
func ConvertVoltageToAnalog(v float32) (a int32) {
	hwCfg := Constants.Hardware
	vcc := *hwCfg.Vcc
	amax := float32(*hwCfg.AnalogMax)
	// volts per analog unit = VCC / Analog max
	avpu := vcc / amax
	// Therefore:
	// 	analog = volts / avpu
	a = util.RoundFloat32ToInt32(v / avpu)
	return
}

// ConvertCelsiusToVoltage converts a celsius temperature to its corresponding voltage
func ConvertCelsiusToVoltage(c int32) (v float32) {
	// According to AdafruitAD8495ThermocoupleVtoKCF
	// 	c = (v - 1.25 ) / 0.005
	// Therefore:
	v = float32(c)*0.005 + 1.25
	return
}

// ConvertCelsiusToAnalog converts the given celsius to its corresponding analog value
func ConvertCelsiusToAnalog(c int32) (a int32) {
	v := ConvertCelsiusToVoltage(c)
	return ConvertVoltageToAnalog(v)
}

// ConvertAnalogToCF converts the given analog value to its corresponding celsius & fahrenheit value
func ConvertAnalogToCF(a int32) (int32, int32) {
	v := ConvertAnalogToVoltage(a)
	_, c, f := AdafruitAD8495ThermocoupleVtoKCF(v)
	return c, f
}

// ConvertCToKF converts a celsius temperature to kelvin and fahrenheit
func ConvertCToKF(tempC float32) (tempK, tempF float32) {
	tempK = tempC + 273.15 // C to K
	tempF = tempC*1.8 + 32 // C to F
	return
}

// ConvertCToKFInt32 converts a celsius temperature to kelvin and fahrenheit
func ConvertCToKFInt32(tempC float32) (tempK, tempF int32) {
	k, f := ConvertCToKF(tempC)
	return int32(k), int32(f)
}
