package config

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	GetConfig(params *GetConfigParams) (*GetConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}
