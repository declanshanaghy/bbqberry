package temperature

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProbeReadingsHandlerFunc turns a function with the right signature into a get probe readings handler
type GetProbeReadingsHandlerFunc func(GetProbeReadingsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProbeReadingsHandlerFunc) Handle(params GetProbeReadingsParams) middleware.Responder {
	return fn(params)
}

// GetProbeReadingsHandler interface for that can handle valid get probe readings params
type GetProbeReadingsHandler interface {
	Handle(GetProbeReadingsParams) middleware.Responder
}

// NewGetProbeReadings creates a new http.Handler for the get probe readings operation
func NewGetProbeReadings(ctx *middleware.Context, handler GetProbeReadingsHandler) *GetProbeReadings {
	return &GetProbeReadings{Context: ctx, Handler: handler}
}

/*GetProbeReadings swagger:route GET /temperatures/probes Temperature getProbeReadings

Get the current temperature reading from the requested probe(s)

*/
type GetProbeReadings struct {
	Context *middleware.Context
	Handler GetProbeReadingsHandler
}

func (o *GetProbeReadings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetProbeReadingsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
