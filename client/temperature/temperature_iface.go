package temperature

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	GetMonitors(params *GetMonitorsParams) (*GetMonitorsOK, error)
	GetProbeReadings(params *GetProbeReadingsParams) (*GetProbeReadingsOK, error)

	SetTransport(transport runtime.ClientTransport)
}
