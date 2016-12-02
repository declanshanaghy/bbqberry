package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMonitorsParams creates a new GetMonitorsParams object
// with the default values initialized.
func NewGetMonitorsParams() *GetMonitorsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetMonitorsParams{
		Probe: probeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMonitorsParamsWithTimeout creates a new GetMonitorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMonitorsParamsWithTimeout(timeout time.Duration) *GetMonitorsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetMonitorsParams{
		Probe: probeDefault,

		timeout: timeout,
	}
}

// NewGetMonitorsParamsWithContext creates a new GetMonitorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMonitorsParamsWithContext(ctx context.Context) *GetMonitorsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetMonitorsParams{
		Probe: probeDefault,

		Context: ctx,
	}
}

/*GetMonitorsParams contains all the parameters to send to the API endpoint
for the get monitors operation typically these are written to a http.Request
*/
type GetMonitorsParams struct {

	/*Probe
	  The termerature probe for which to retrieve configured monitors

	*/
	Probe int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get monitors params
func (o *GetMonitorsParams) WithTimeout(timeout time.Duration) *GetMonitorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get monitors params
func (o *GetMonitorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get monitors params
func (o *GetMonitorsParams) WithContext(ctx context.Context) *GetMonitorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get monitors params
func (o *GetMonitorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithProbe adds the probe to the get monitors params
func (o *GetMonitorsParams) WithProbe(probe int32) *GetMonitorsParams {
	o.SetProbe(probe)
	return o
}

// SetProbe adds the probe to the get monitors params
func (o *GetMonitorsParams) SetProbe(probe int32) {
	o.Probe = probe
}

// WriteToRequest writes these params to a swagger request
func (o *GetMonitorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param probe
	if err := r.SetPathParam("probe", swag.FormatInt32(o.Probe)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
