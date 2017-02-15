package temperatures

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	GetTemperatures(params *GetTemperaturesParams) (*GetTemperaturesOK, error)

	SetTransport(transport runtime.ClientTransport)
}
