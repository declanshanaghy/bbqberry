package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/declanshanaghy/bbqberry/models"
)

// GetProbeReadingsReader is a Reader for the GetProbeReadings structure.
type GetProbeReadingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProbeReadingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProbeReadingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetProbeReadingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProbeReadingsOK creates a GetProbeReadingsOK with default headers values
func NewGetProbeReadingsOK() *GetProbeReadingsOK {
	return &GetProbeReadingsOK{}
}

/*GetProbeReadingsOK handles this case with default header values.

Temperature was read successfully
*/
type GetProbeReadingsOK struct {
	Payload *models.TemperatureReading
}

func (o *GetProbeReadingsOK) Error() string {
	return fmt.Sprintf("[GET /temperatures/probes][%d] getProbeReadingsOK  %+v", 200, o.Payload)
}

func (o *GetProbeReadingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemperatureReading)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProbeReadingsDefault creates a GetProbeReadingsDefault with default headers values
func NewGetProbeReadingsDefault(code int) *GetProbeReadingsDefault {
	return &GetProbeReadingsDefault{
		_statusCode: code,
	}
}

/*GetProbeReadingsDefault handles this case with default header values.

Unexpected error
*/
type GetProbeReadingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get probe readings default response
func (o *GetProbeReadingsDefault) Code() int {
	return o._statusCode
}

func (o *GetProbeReadingsDefault) Error() string {
	return fmt.Sprintf("[GET /temperatures/probes][%d] getProbeReadings default  %+v", o._statusCode, o.Payload)
}

func (o *GetProbeReadingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
