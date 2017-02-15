package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// HardwareConfig hardware config
// swagger:model HardwareConfig
type HardwareConfig struct {

	// analog max
	// Required: true
	// Minimum: 0
	AnalogMax *int32 `json:"analogMax"`

	// The amount the voltage will increase to reflect a unit increase in analog reading
	// Minimum: 0
	AnalogVoltsPerUnit *float32 `json:"analogVoltsPerUnit,omitempty"`

	// num led pixels
	// Required: true
	// Minimum: 0
	NumLedPixels *int32 `json:"numLedPixels"`

	// probes
	// Required: true
	Probes TemperatureProbes `json:"probes"`

	// vcc
	// Required: true
	Vcc *float32 `json:"vcc"`
}

// Validate validates this hardware config
func (m *HardwareConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalogMax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAnalogVoltsPerUnit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumLedPixels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProbes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVcc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwareConfig) validateAnalogMax(formats strfmt.Registry) error {

	if err := validate.Required("analogMax", "body", m.AnalogMax); err != nil {
		return err
	}

	if err := validate.MinimumInt("analogMax", "body", int64(*m.AnalogMax), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *HardwareConfig) validateAnalogVoltsPerUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.AnalogVoltsPerUnit) { // not required
		return nil
	}

	if err := validate.Minimum("analogVoltsPerUnit", "body", float64(*m.AnalogVoltsPerUnit), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *HardwareConfig) validateNumLedPixels(formats strfmt.Registry) error {

	if err := validate.Required("numLedPixels", "body", m.NumLedPixels); err != nil {
		return err
	}

	if err := validate.MinimumInt("numLedPixels", "body", int64(*m.NumLedPixels), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *HardwareConfig) validateProbes(formats strfmt.Registry) error {

	if err := validate.Required("probes", "body", m.Probes); err != nil {
		return err
	}

	return nil
}

func (m *HardwareConfig) validateVcc(formats strfmt.Registry) error {

	if err := validate.Required("vcc", "body", m.Vcc); err != nil {
		return err
	}

	return nil
}
