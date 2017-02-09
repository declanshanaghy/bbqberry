package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetMonitorsParams creates a new GetMonitorsParams object
// with the default values initialized.
func NewGetMonitorsParams() GetMonitorsParams {
	var (
		probeDefault = int32(0)
	)
	return GetMonitorsParams{
		Probe: probeDefault,
	}
}

// GetMonitorsParams contains all the bound params for the get monitors operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMonitors
type GetMonitorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The termerature probe for which to retrieve configured monitors (or all probes if the given probe number is 0 or not specified)
	  Required: true
	  Maximum: 7
	  Minimum: 0
	  In: query
	  Default: 0
	*/
	Probe int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetMonitorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qProbe, qhkProbe, _ := qs.GetOK("probe")
	if err := o.bindProbe(qProbe, qhkProbe, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMonitorsParams) bindProbe(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("probe", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("probe", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("probe", "query", "int32", raw)
	}
	o.Probe = value

	if err := o.validateProbe(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetMonitorsParams) validateProbe(formats strfmt.Registry) error {

	if err := validate.MinimumInt("probe", "query", int64(o.Probe), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("probe", "query", int64(o.Probe), 7, false); err != nil {
		return err
	}

	return nil
}
