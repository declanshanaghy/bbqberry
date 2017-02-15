package hardware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetHardwareHandlerFunc turns a function with the right signature into a get hardware handler
type GetHardwareHandlerFunc func(GetHardwareParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHardwareHandlerFunc) Handle(params GetHardwareParams) middleware.Responder {
	return fn(params)
}

// GetHardwareHandler interface for that can handle valid get hardware params
type GetHardwareHandler interface {
	Handle(GetHardwareParams) middleware.Responder
}

// NewGetHardware creates a new http.Handler for the get hardware operation
func NewGetHardware(ctx *middleware.Context, handler GetHardwareHandler) *GetHardware {
	return &GetHardware{Context: ctx, Handler: handler}
}

/*GetHardware swagger:route GET /hardware Hardware getHardware

Get current configuration settings

*/
type GetHardware struct {
	Context *middleware.Context
	Handler GetHardwareHandler
}

func (o *GetHardware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetHardwareParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
