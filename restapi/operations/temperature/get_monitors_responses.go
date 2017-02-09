package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/declanshanaghy/bbqberry/models"
)

/*GetMonitorsOK The currently configured monitor(s) were retrieved successfully

swagger:response getMonitorsOK
*/
type GetMonitorsOK struct {

	/*
	  In: Body
	*/
	Payload models.TemperatureMonitors `json:"body,omitempty"`
}

// NewGetMonitorsOK creates GetMonitorsOK with default headers values
func NewGetMonitorsOK() *GetMonitorsOK {
	return &GetMonitorsOK{}
}

// WithPayload adds the payload to the get monitors o k response
func (o *GetMonitorsOK) WithPayload(payload models.TemperatureMonitors) *GetMonitorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get monitors o k response
func (o *GetMonitorsOK) SetPayload(payload models.TemperatureMonitors) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMonitorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.TemperatureMonitors, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetMonitorsDefault Unexpected error

swagger:response getMonitorsDefault
*/
type GetMonitorsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMonitorsDefault creates GetMonitorsDefault with default headers values
func NewGetMonitorsDefault(code int) *GetMonitorsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMonitorsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get monitors default response
func (o *GetMonitorsDefault) WithStatusCode(code int) *GetMonitorsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get monitors default response
func (o *GetMonitorsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get monitors default response
func (o *GetMonitorsDefault) WithPayload(payload *models.Error) *GetMonitorsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get monitors default response
func (o *GetMonitorsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMonitorsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
