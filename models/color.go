package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Color color
// swagger:model Color
type Color struct {

	// Color in hex representation
	// Required: true
	Hex *string `json:"hex"`
}

// Validate validates this color
func (m *Color) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHex(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Color) validateHex(formats strfmt.Registry) error {

	if err := validate.Required("hex", "body", m.Hex); err != nil {
		return err
	}

	return nil
}
