package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TemperatureProbe temperature probe
// swagger:model TemperatureProbe
type TemperatureProbe struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// label
	// Required: true
	Label *string `json:"label"`

	// limits
	// Required: true
	Limits *TemperatureLimits `json:"limits"`
}

// Validate validates this temperature probe
func (m *TemperatureProbe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLimits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemperatureProbe) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *TemperatureProbe) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *TemperatureProbe) validateLimits(formats strfmt.Registry) error {

	if err := validate.Required("limits", "body", m.Limits); err != nil {
		return err
	}

	if m.Limits != nil {

		if err := m.Limits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limits")
			}
			return err
		}
	}

	return nil
}
