package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TemperatureMonitors temperature monitors
// swagger:model TemperatureMonitors
type TemperatureMonitors map[string][]TemperatureMonitor

// Validate validates this temperature monitors
func (m TemperatureMonitors) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", TemperatureMonitors(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}