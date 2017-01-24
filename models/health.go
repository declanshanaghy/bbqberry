package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Health health
// swagger:model Health
type Health struct {

	// error
	Error *Error `json:"error,omitempty"`

	// Flag indicating whether or not ALL internal checks passed
	// Required: true
	Healthy *bool `json:"healthy"`

	// service info
	// Required: true
	ServiceInfo *ServiceInfo `json:"service_info"`
}

// Validate validates this health
func (m *Health) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Health) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *Health) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *Health) validateServiceInfo(formats strfmt.Registry) error {

	if err := validate.Required("service_info", "body", m.ServiceInfo); err != nil {
		return err
	}

	if m.ServiceInfo != nil {

		if err := m.ServiceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_info")
			}
			return err
		}
	}

	return nil
}
