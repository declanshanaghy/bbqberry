package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TemperatureMonitor temperature monitor
// swagger:model TemperatureMonitor
type TemperatureMonitor struct {

	// The maximium temperature, below which an alert will be generated
	// Required: true
	Max *float32 `json:"max"`

	// The minimum temperature, below which an alert will be generated
	// Required: true
	Min *float32 `json:"min"`

	// probe
	// Required: true
	// Maximum: 3
	// Minimum: 0
	Probe *int32 `json:"probe"`

	// The temperature scale
	// Required: true
	Scale *string `json:"scale"`
}

// Validate validates this temperature monitor
func (m *TemperatureMonitor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProbe(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TemperatureMonitor) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *TemperatureMonitor) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

func (m *TemperatureMonitor) validateProbe(formats strfmt.Registry) error {

	if err := validate.Required("probe", "body", m.Probe); err != nil {
		return err
	}

	if err := validate.MinimumInt("probe", "body", int64(*m.Probe), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("probe", "body", int64(*m.Probe), 3, false); err != nil {
		return err
	}

	return nil
}

func (m *TemperatureMonitor) validateScale(formats strfmt.Registry) error {

	if err := validate.Required("scale", "body", m.Scale); err != nil {
		return err
	}

	return nil
}