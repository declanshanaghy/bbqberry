package monitors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/declanshanaghy/bbqberry/models"
)

/*CreateMonitorOK The currently configured monitor(s) were retrieved successfully

swagger:response createMonitorOK
*/
type CreateMonitorOK struct {

	/*
	  In: Body
	*/
	Payload models.TemperatureMonitors `json:"body,omitempty"`
}

// NewCreateMonitorOK creates CreateMonitorOK with default headers values
func NewCreateMonitorOK() *CreateMonitorOK {
	return &CreateMonitorOK{}
}

// WithPayload adds the payload to the create monitor o k response
func (o *CreateMonitorOK) WithPayload(payload models.TemperatureMonitors) *CreateMonitorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create monitor o k response
func (o *CreateMonitorOK) SetPayload(payload models.TemperatureMonitors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMonitorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.TemperatureMonitors, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreateMonitorDefault Unexpected error

swagger:response createMonitorDefault
*/
type CreateMonitorDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateMonitorDefault creates CreateMonitorDefault with default headers values
func NewCreateMonitorDefault(code int) *CreateMonitorDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateMonitorDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create monitor default response
func (o *CreateMonitorDefault) WithStatusCode(code int) *CreateMonitorDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create monitor default response
func (o *CreateMonitorDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create monitor default response
func (o *CreateMonitorDefault) WithPayload(payload *models.Error) *CreateMonitorDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create monitor default response
func (o *CreateMonitorDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateMonitorDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}