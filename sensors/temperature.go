package sensors

import (
	"github.com/kidoman/embd"
	"github.com/golang/glog"
	"github.com/kidoman/embd/convertors/mcp3008"
	"math"
)

type TemperatureArray interface {
	/*
	Reads the tempearature from the requested probe and returns the value in Kelvin, Celcius & Fahrenheit
	 */
	GetTemp(probe int) TemperatureReading
}

type TemperatureReading struct {
	K, C, F float64
}

type BBQTemp struct {
	channel byte
	bus embd.SPIBus
	adc *mcp3008.MCP3008
}

func NewTemperature(channel byte) TemperatureArray {
	s := BBQTemp{channel:channel}
	s.init()
	return &s
}

func (s *BBQTemp) init() {
	glog.Info("action=init")
	s.bus = embd.NewSPIBus(embd.SPIMode0, s.channel, 1000000, 8, 0)
	s.adc = mcp3008.New(mcp3008.SingleMode, s.bus)
}

func (s *BBQTemp) Close() {
	glog.Info("action=Close")
	s.bus.Close()
}

func (s *BBQTemp) GetTemp(probe int) TemperatureReading {
	glog.Infof("action=GetTemp probe=%d", probe)
	v, err := s.adc.AnalogValueAt(probe)
	if err != nil {
		panic(err)
	}
	return convertVoltToTemp(v)
}

func convertVoltToTemp(volt int) TemperatureReading {
	// get the Kelvin temperature
	k := math.Log(10240000.0 / float64(volt) - 10000);
	k = 1 / (0.001129148 + (0.000234125 * k) + (0.0000000876741 * k * k * k));

	// convert to Celsius and round to 1 decimal place
	c := k - 273.15;

	// get the Fahrenheit temperature
	f := (c * 1.8) + 32;

	// return all three temperature values
	return TemperatureReading{K: k, C: c, F: f}
}
