package monitors

import (
	"github.com/go-openapi/runtime"
)

// Interface for client methods
type ClientInterface interface {
	CreateMonitor(params *CreateMonitorParams) (*CreateMonitorOK, error)
	GetMonitors(params *GetMonitorsParams) (*GetMonitorsOK, error)

	SetTransport(transport runtime.ClientTransport)
}