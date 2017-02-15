package hardware

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	GetHardware(params *GetHardwareParams) (*GetHardwareOK, error)

	SetTransport(transport runtime.ClientTransport)
}
