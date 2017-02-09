package config

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	GetHardwareConfig(params *GetHardwareConfigParams) (*GetHardwareConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}
