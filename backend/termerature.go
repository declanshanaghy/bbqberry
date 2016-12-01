package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/restapi/operations/sensors"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/go-openapi/strfmt"
)

func GetTemperature(params *sensors.GetTemperatureParams) (m models.Temperature, err error) {
	bus := hardware.NewSPIBus(1)
	sTemp := hardware.NewTemperature(bus)
	reading := sTemp.GetTemp(params.Probe)
	
	m.Time = strfmt.DateTime(reading.Time)
	m.Probe = &reading.Probe
	m.Kelvin = &reading.Kelvin
	m.Celcuis = &reading.Celcius
	m.Fahrenheit = &reading.Fahrenheit

	return m, err
}