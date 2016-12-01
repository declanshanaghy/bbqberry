package hardware

import (
	"github.com/kidoman/embd"
	"github.com/declanshanaghy/bbqberry/framework/log"
)

func NewSPIBus(channel byte) embd.SPIBus {
	log.Infof("action=NewSPIBus channel=%d", channel)
	return &MockSPIBus{}
}