package temperatures

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTemperaturesHandlerFunc turns a function with the right signature into a get temperatures handler
type GetTemperaturesHandlerFunc func(GetTemperaturesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTemperaturesHandlerFunc) Handle(params GetTemperaturesParams) middleware.Responder {
	return fn(params)
}

// GetTemperaturesHandler interface for that can handle valid get temperatures params
type GetTemperaturesHandler interface {
	Handle(GetTemperaturesParams) middleware.Responder
}

// NewGetTemperatures creates a new http.Handler for the get temperatures operation
func NewGetTemperatures(ctx *middleware.Context, handler GetTemperaturesHandler) *GetTemperatures {
	return &GetTemperatures{Context: ctx, Handler: handler}
}

/*GetTemperatures swagger:route GET /temperatures Temperatures getTemperatures

Get the current temperature reading from the requested probe(s)

*/
type GetTemperatures struct {
	Context *middleware.Context
	Handler GetTemperaturesHandler
}

func (o *GetTemperatures) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTemperaturesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
