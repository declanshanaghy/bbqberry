package monitors

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

// NewUpdateMonitorParams creates a new UpdateMonitorParams object
// with the default values initialized.
func NewUpdateMonitorParams() *UpdateMonitorParams {
	var ()
	return &UpdateMonitorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMonitorParamsWithTimeout creates a new UpdateMonitorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMonitorParamsWithTimeout(timeout time.Duration) *UpdateMonitorParams {
	var ()
	return &UpdateMonitorParams{

		timeout: timeout,
	}
}

// NewUpdateMonitorParamsWithContext creates a new UpdateMonitorParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMonitorParamsWithContext(ctx context.Context) *UpdateMonitorParams {
	var ()
	return &UpdateMonitorParams{

		Context: ctx,
	}
}

// NewUpdateMonitorParamsWithHTTPClient creates a new UpdateMonitorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMonitorParamsWithHTTPClient(client *http.Client) *UpdateMonitorParams {
	var ()
	return &UpdateMonitorParams{
		HTTPClient: client,
	}
}

/*UpdateMonitorParams contains all the parameters to send to the API endpoint
for the update monitor operation typically these are written to a http.Request
*/
type UpdateMonitorParams struct {

	/*Max
	  The maximium temperature, above which an alert will be generated

	*/
	Max int32
	/*Min
	  The minimum temperature, below which an alert will be generated

	*/
	Min int32
	/*Probe*/
	Probe int32
	/*Scale
	  The temperature scale

	*/
	Scale string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update monitor params
func (o *UpdateMonitorParams) WithTimeout(timeout time.Duration) *UpdateMonitorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update monitor params
func (o *UpdateMonitorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update monitor params
func (o *UpdateMonitorParams) WithContext(ctx context.Context) *UpdateMonitorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update monitor params
func (o *UpdateMonitorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update monitor params
func (o *UpdateMonitorParams) WithHTTPClient(client *http.Client) *UpdateMonitorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update monitor params
func (o *UpdateMonitorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMax adds the max to the update monitor params
func (o *UpdateMonitorParams) WithMax(max int32) *UpdateMonitorParams {
	o.SetMax(max)
	return o
}

// SetMax adds the max to the update monitor params
func (o *UpdateMonitorParams) SetMax(max int32) {
	o.Max = max
}

// WithMin adds the min to the update monitor params
func (o *UpdateMonitorParams) WithMin(min int32) *UpdateMonitorParams {
	o.SetMin(min)
	return o
}

// SetMin adds the min to the update monitor params
func (o *UpdateMonitorParams) SetMin(min int32) {
	o.Min = min
}

// WithProbe adds the probe to the update monitor params
func (o *UpdateMonitorParams) WithProbe(probe int32) *UpdateMonitorParams {
	o.SetProbe(probe)
	return o
}

// SetProbe adds the probe to the update monitor params
func (o *UpdateMonitorParams) SetProbe(probe int32) {
	o.Probe = probe
}

// WithScale adds the scale to the update monitor params
func (o *UpdateMonitorParams) WithScale(scale string) *UpdateMonitorParams {
	o.SetScale(scale)
	return o
}

// SetScale adds the scale to the update monitor params
func (o *UpdateMonitorParams) SetScale(scale string) {
	o.Scale = scale
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMonitorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param max
	qrMax := o.Max
	qMax := swag.FormatInt32(qrMax)
	if qMax != "" {
		if err := r.SetQueryParam("max", qMax); err != nil {
			return err
		}
	}

	// query param min
	qrMin := o.Min
	qMin := swag.FormatInt32(qrMin)
	if qMin != "" {
		if err := r.SetQueryParam("min", qMin); err != nil {
			return err
		}
	}

	// query param probe
	qrProbe := o.Probe
	qProbe := swag.FormatInt32(qrProbe)
	if qProbe != "" {
		if err := r.SetQueryParam("probe", qProbe); err != nil {
			return err
		}
	}

	// query param scale
	qrScale := o.Scale
	qScale := qrScale
	if qScale != "" {
		if err := r.SetQueryParam("scale", qScale); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}