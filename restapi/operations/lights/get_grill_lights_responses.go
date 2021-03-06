package lights

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/declanshanaghy/bbqberry/models"
)

// GetGrillLightsOKCode is the HTTP code returned for type GetGrillLightsOK
const GetGrillLightsOKCode int = 200

/*GetGrillLightsOK Pixels were read successfully

swagger:response getGrillLightsOK
*/
type GetGrillLightsOK struct {

	/*
	  In: Body
	*/
	Payload GetGrillLightsOKBody `json:"body,omitempty"`
}

// NewGetGrillLightsOK creates GetGrillLightsOK with default headers values
func NewGetGrillLightsOK() *GetGrillLightsOK {
	return &GetGrillLightsOK{}
}

// WithPayload adds the payload to the get grill lights o k response
func (o *GetGrillLightsOK) WithPayload(payload GetGrillLightsOKBody) *GetGrillLightsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get grill lights o k response
func (o *GetGrillLightsOK) SetPayload(payload GetGrillLightsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGrillLightsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetGrillLightsDefault Unexpected error

swagger:response getGrillLightsDefault
*/
type GetGrillLightsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetGrillLightsDefault creates GetGrillLightsDefault with default headers values
func NewGetGrillLightsDefault(code int) *GetGrillLightsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetGrillLightsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get grill lights default response
func (o *GetGrillLightsDefault) WithStatusCode(code int) *GetGrillLightsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get grill lights default response
func (o *GetGrillLightsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get grill lights default response
func (o *GetGrillLightsDefault) WithPayload(payload *models.Error) *GetGrillLightsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get grill lights default response
func (o *GetGrillLightsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetGrillLightsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
