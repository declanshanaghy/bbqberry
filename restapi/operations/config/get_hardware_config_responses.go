package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/declanshanaghy/bbqberry/models"
)

/*GetHardwareConfigOK The config was retrieved successfully

swagger:response getHardwareConfigOK
*/
type GetHardwareConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.HardwareConfig `json:"body,omitempty"`
}

// NewGetHardwareConfigOK creates GetHardwareConfigOK with default headers values
func NewGetHardwareConfigOK() *GetHardwareConfigOK {
	return &GetHardwareConfigOK{}
}

// WithPayload adds the payload to the get hardware config o k response
func (o *GetHardwareConfigOK) WithPayload(payload *models.HardwareConfig) *GetHardwareConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hardware config o k response
func (o *GetHardwareConfigOK) SetPayload(payload *models.HardwareConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHardwareConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetHardwareConfigDefault Unexpected error

swagger:response getHardwareConfigDefault
*/
type GetHardwareConfigDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHardwareConfigDefault creates GetHardwareConfigDefault with default headers values
func NewGetHardwareConfigDefault(code int) *GetHardwareConfigDefault {
	if code <= 0 {
		code = 500
	}

	return &GetHardwareConfigDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get hardware config default response
func (o *GetHardwareConfigDefault) WithStatusCode(code int) *GetHardwareConfigDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get hardware config default response
func (o *GetHardwareConfigDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get hardware config default response
func (o *GetHardwareConfigDefault) WithPayload(payload *models.Error) *GetHardwareConfigDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get hardware config default response
func (o *GetHardwareConfigDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHardwareConfigDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
