package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new temperature API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for temperature API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetMonitors gets monitor settings for the requested probe
*/
func (a *Client) GetMonitors(params *GetMonitorsParams) (*GetMonitorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMonitorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMonitors",
		Method:             "GET",
		PathPattern:        "/temperatures/monitors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMonitorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMonitorsOK), nil

}

/*
GetProbeReadings gets the current temperature reading from the requested probe s
*/
func (a *Client) GetProbeReadings(params *GetProbeReadingsParams) (*GetProbeReadingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProbeReadingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProbeReadings",
		Method:             "GET",
		PathPattern:        "/temperatures/probes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProbeReadingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProbeReadingsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
