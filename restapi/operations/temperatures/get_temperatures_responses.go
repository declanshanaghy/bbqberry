package temperatures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/declanshanaghy/bbqberry/models"
)

/*GetTemperaturesOK Temperature was read successfully

swagger:response getTemperaturesOK
*/
type GetTemperaturesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.TemperatureReading `json:"body,omitempty"`
}

// NewGetTemperaturesOK creates GetTemperaturesOK with default headers values
func NewGetTemperaturesOK() *GetTemperaturesOK {
	return &GetTemperaturesOK{}
}

// WithPayload adds the payload to the get temperatures o k response
func (o *GetTemperaturesOK) WithPayload(payload []*models.TemperatureReading) *GetTemperaturesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get temperatures o k response
func (o *GetTemperaturesOK) SetPayload(payload []*models.TemperatureReading) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTemperaturesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.TemperatureReading, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTemperaturesDefault Unexpected error

swagger:response getTemperaturesDefault
*/
type GetTemperaturesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTemperaturesDefault creates GetTemperaturesDefault with default headers values
func NewGetTemperaturesDefault(code int) *GetTemperaturesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTemperaturesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get temperatures default response
func (o *GetTemperaturesDefault) WithStatusCode(code int) *GetTemperaturesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get temperatures default response
func (o *GetTemperaturesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get temperatures default response
func (o *GetTemperaturesDefault) WithPayload(payload *models.Error) *GetTemperaturesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get temperatures default response
func (o *GetTemperaturesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTemperaturesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
