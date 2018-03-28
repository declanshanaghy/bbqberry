package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// TemperatureMonitor temperature monitor
// swagger:model TemperatureMonitor
type TemperatureMonitor struct {

	// label
	Label string `json:"label,omitempty"`

	// The maximium temperature, above which an alert will be generated
	// Required: true
	Max *int32 `json:"max"`

	// The minimum temperature, below which an alert will be generated
	// Required: true
	Min *int32 `json:"min"`

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

var temperatureMonitorTypeScalePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["celsius"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		temperatureMonitorTypeScalePropEnum = append(temperatureMonitorTypeScalePropEnum, v)
	}
}

const (
	// TemperatureMonitorScaleCelsius captures enum value "celsius"
	TemperatureMonitorScaleCelsius string = "celsius"
)

// prop value enum
func (m *TemperatureMonitor) validateScaleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, temperatureMonitorTypeScalePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TemperatureMonitor) validateScale(formats strfmt.Registry) error {

	if err := validate.Required("scale", "body", m.Scale); err != nil {
		return err
	}

	// value enum
	if err := m.validateScaleEnum("scale", "body", *m.Scale); err != nil {
		return err
	}

	return nil
}
