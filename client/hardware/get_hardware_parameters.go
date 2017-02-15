package hardware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetHardwareParams creates a new GetHardwareParams object
// with the default values initialized.
func NewGetHardwareParams() *GetHardwareParams {

	return &GetHardwareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHardwareParamsWithTimeout creates a new GetHardwareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHardwareParamsWithTimeout(timeout time.Duration) *GetHardwareParams {

	return &GetHardwareParams{

		timeout: timeout,
	}
}

// NewGetHardwareParamsWithContext creates a new GetHardwareParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHardwareParamsWithContext(ctx context.Context) *GetHardwareParams {

	return &GetHardwareParams{

		Context: ctx,
	}
}

/*GetHardwareParams contains all the parameters to send to the API endpoint
for the get hardware operation typically these are written to a http.Request
*/
type GetHardwareParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hardware params
func (o *GetHardwareParams) WithTimeout(timeout time.Duration) *GetHardwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hardware params
func (o *GetHardwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hardware params
func (o *GetHardwareParams) WithContext(ctx context.Context) *GetHardwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hardware params
func (o *GetHardwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetHardwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}