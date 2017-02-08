package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/declanshanaghy/bbqberry/models"
)

// GetConfigReader is a Reader for the GetConfig structure.
type GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/*GetConfigOK handles this case with default header values.

The config was retrieved successfully
*/
type GetConfigOK struct {
	Payload *models.Config
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigOK  %+v", 200, o.Payload)
}

func (o *GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigDefault creates a GetConfigDefault with default headers values
func NewGetConfigDefault(code int) *GetConfigDefault {
	return &GetConfigDefault{
		_statusCode: code,
	}
}

/*GetConfigDefault handles this case with default header values.

Unexpected error
*/
type GetConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get config default response
func (o *GetConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigDefault) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
