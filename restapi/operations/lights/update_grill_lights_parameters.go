package lights

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

// NewUpdateGrillLightsParams creates a new UpdateGrillLightsParams object
// with the default values initialized.
func NewUpdateGrillLightsParams() UpdateGrillLightsParams {
	var (
		periodDefault = int64(500000)
	)
	return UpdateGrillLightsParams{
		Period: periodDefault,
	}
}

// UpdateGrillLightsParams contains all the bound params for the update grill lights operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateGrillLights
type UpdateGrillLightsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The light show to enable
	  Required: true
	  In: query
	*/
	Name string
	/*The time period between updates in microseconds
	  Required: true
	  Minimum: 1
	  In: query
	  Default: 500000
	*/
	Period int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateGrillLightsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qName, qhkName, _ := qs.GetOK("name")
	if err := o.bindName(qName, qhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	qPeriod, qhkPeriod, _ := qs.GetOK("period")
	if err := o.bindPeriod(qPeriod, qhkPeriod, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateGrillLightsParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("name", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("name", "query", raw); err != nil {
		return err
	}

	o.Name = raw

	if err := o.validateName(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateGrillLightsParams) validateName(formats strfmt.Registry) error {

	if err := validate.Enum("name", "query", o.Name, []interface{}{"Simple Shifter", "Rainbow", "Temperature"}); err != nil {
		return err
	}

	return nil
}

func (o *UpdateGrillLightsParams) bindPeriod(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("period", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("period", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("period", "query", "int64", raw)
	}
	o.Period = value

	if err := o.validatePeriod(formats); err != nil {
		return err
	}

	return nil
}

func (o *UpdateGrillLightsParams) validatePeriod(formats strfmt.Registry) error {

	if err := validate.MinimumInt("period", "query", int64(o.Period), 1, false); err != nil {
		return err
	}

	return nil
}
