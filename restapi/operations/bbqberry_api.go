package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/declanshanaghy/bbqberry/restapi/operations/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations/health"
	"github.com/declanshanaghy/bbqberry/restapi/operations/lights"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
)

// NewBbqberryAPI creates a new Bbqberry instance
func NewBbqberryAPI(spec *loads.Document) *BbqberryAPI {
	return &BbqberryAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		MonitorsCreateMonitorHandler: monitors.CreateMonitorHandlerFunc(func(params monitors.CreateMonitorParams) middleware.Responder {
			return middleware.NotImplemented("operation MonitorsCreateMonitor has not yet been implemented")
		}),
		HardwareGetHardwareHandler: hardware.GetHardwareHandlerFunc(func(params hardware.GetHardwareParams) middleware.Responder {
			return middleware.NotImplemented("operation HardwareGetHardware has not yet been implemented")
		}),
		MonitorsGetMonitorsHandler: monitors.GetMonitorsHandlerFunc(func(params monitors.GetMonitorsParams) middleware.Responder {
			return middleware.NotImplemented("operation MonitorsGetMonitors has not yet been implemented")
		}),
		TemperaturesGetTemperaturesHandler: temperatures.GetTemperaturesHandlerFunc(func(params temperatures.GetTemperaturesParams) middleware.Responder {
			return middleware.NotImplemented("operation TemperaturesGetTemperatures has not yet been implemented")
		}),
		HealthHealthHandler: health.HealthHandlerFunc(func(params health.HealthParams) middleware.Responder {
			return middleware.NotImplemented("operation HealthHealth has not yet been implemented")
		}),
		LightsUpdateGrillLightsHandler: lights.UpdateGrillLightsHandlerFunc(func(params lights.UpdateGrillLightsParams) middleware.Responder {
			return middleware.NotImplemented("operation LightsUpdateGrillLights has not yet been implemented")
		}),
	}
}

/*BbqberryAPI Rest API definition for BBQ Berry */
type BbqberryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// MonitorsCreateMonitorHandler sets the operation handler for the create monitor operation
	MonitorsCreateMonitorHandler monitors.CreateMonitorHandler
	// HardwareGetHardwareHandler sets the operation handler for the get hardware operation
	HardwareGetHardwareHandler hardware.GetHardwareHandler
	// MonitorsGetMonitorsHandler sets the operation handler for the get monitors operation
	MonitorsGetMonitorsHandler monitors.GetMonitorsHandler
	// TemperaturesGetTemperaturesHandler sets the operation handler for the get temperatures operation
	TemperaturesGetTemperaturesHandler temperatures.GetTemperaturesHandler
	// HealthHealthHandler sets the operation handler for the health operation
	HealthHealthHandler health.HealthHandler
	// LightsUpdateGrillLightsHandler sets the operation handler for the update grill lights operation
	LightsUpdateGrillLightsHandler lights.UpdateGrillLightsHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *BbqberryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *BbqberryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *BbqberryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *BbqberryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *BbqberryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *BbqberryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *BbqberryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the BbqberryAPI
func (o *BbqberryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.MonitorsCreateMonitorHandler == nil {
		unregistered = append(unregistered, "monitors.CreateMonitorHandler")
	}

	if o.HardwareGetHardwareHandler == nil {
		unregistered = append(unregistered, "hardware.GetHardwareHandler")
	}

	if o.MonitorsGetMonitorsHandler == nil {
		unregistered = append(unregistered, "monitors.GetMonitorsHandler")
	}

	if o.TemperaturesGetTemperaturesHandler == nil {
		unregistered = append(unregistered, "temperatures.GetTemperaturesHandler")
	}

	if o.HealthHealthHandler == nil {
		unregistered = append(unregistered, "health.HealthHandler")
	}

	if o.LightsUpdateGrillLightsHandler == nil {
		unregistered = append(unregistered, "lights.UpdateGrillLightsHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *BbqberryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *BbqberryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *BbqberryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *BbqberryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *BbqberryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the bbqberry API
func (o *BbqberryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *BbqberryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/monitors"] = monitors.NewCreateMonitor(o.context, o.MonitorsCreateMonitorHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/hardware"] = hardware.NewGetHardware(o.context, o.HardwareGetHardwareHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/monitors"] = monitors.NewGetMonitors(o.context, o.MonitorsGetMonitorsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/temperatures"] = temperatures.NewGetTemperatures(o.context, o.TemperaturesGetTemperaturesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = health.NewHealth(o.context, o.HealthHealthHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/lights/grill"] = lights.NewUpdateGrillLights(o.context, o.LightsUpdateGrillLightsHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *BbqberryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *BbqberryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
