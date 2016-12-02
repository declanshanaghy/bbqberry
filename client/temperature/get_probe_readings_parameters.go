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

// NewGetProbeReadingsParams creates a new GetProbeReadingsParams object
// with the default values initialized.
func NewGetProbeReadingsParams() *GetProbeReadingsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetProbeReadingsParams{
		Probe: &probeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProbeReadingsParamsWithTimeout creates a new GetProbeReadingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProbeReadingsParamsWithTimeout(timeout time.Duration) *GetProbeReadingsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetProbeReadingsParams{
		Probe: &probeDefault,

		timeout: timeout,
	}
}

// NewGetProbeReadingsParamsWithContext creates a new GetProbeReadingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProbeReadingsParamsWithContext(ctx context.Context) *GetProbeReadingsParams {
	var (
		probeDefault = int32(0)
	)
	return &GetProbeReadingsParams{
		Probe: &probeDefault,

		Context: ctx,
	}
}

/*GetProbeReadingsParams contains all the parameters to send to the API endpoint
for the get probe readings operation typically these are written to a http.Request
*/
type GetProbeReadingsParams struct {

	/*Probe
	  The termerature probe to read from (or all probes if the given probe number is 0 or not specified)

	*/
	Probe *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get probe readings params
func (o *GetProbeReadingsParams) WithTimeout(timeout time.Duration) *GetProbeReadingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get probe readings params
func (o *GetProbeReadingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get probe readings params
func (o *GetProbeReadingsParams) WithContext(ctx context.Context) *GetProbeReadingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get probe readings params
func (o *GetProbeReadingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithProbe adds the probe to the get probe readings params
func (o *GetProbeReadingsParams) WithProbe(probe *int32) *GetProbeReadingsParams {
	o.SetProbe(probe)
	return o
}

// SetProbe adds the probe to the get probe readings params
func (o *GetProbeReadingsParams) SetProbe(probe *int32) {
	o.Probe = probe
}

// WriteToRequest writes these params to a swagger request
func (o *GetProbeReadingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Probe != nil {

		// query param probe
		var qrProbe int32
		if o.Probe != nil {
			qrProbe = *o.Probe
		}
		qProbe := swag.FormatInt32(qrProbe)
		if qProbe != "" {
			if err := r.SetQueryParam("probe", qProbe); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}