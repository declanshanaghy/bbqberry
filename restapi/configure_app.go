package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/framework"
	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations"
	"github.com/declanshanaghy/bbqberry/restapi/operations/example"
	"github.com/declanshanaghy/bbqberry/restapi/operations/health"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	"github.com/go-openapi/swag"
)

type CmdOptions struct {
	LogFile   string `short:"l" long:"logfile" description:"Specify the log file" default:""`
	Verbose   bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	StaticDir string `short:"s" long:"static" description:"The path to the static dirs" default:""`
}

var CmdOptionsValues CmdOptions // export for testing

func configureFlags(api *operations.AppAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		swag.CommandLineOptionsGroup{
			ShortDescription: "BBQ Berry Server Flags",
			LongDescription:  "BBQ Berry Server Flags",
			Options:          &CmdOptionsValues,
		},
	}

}
func configureAPI(api *operations.AppAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	log.SetDebug(CmdOptionsValues.Verbose)
	if CmdOptionsValues.LogFile != "" {
		log.SetOutput(CmdOptionsValues.LogFile)
	}
	api.Logger = log.Infof

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.HealthHealthHandler = health.HealthHandlerFunc(func(params health.HealthParams) middleware.Responder {
		return framework.HandleAPIRequestWithError(backend.Health())
	})
	api.ExampleHelloHandler = example.HelloHandlerFunc(
		func(params example.HelloParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.Hello())
		})
	api.TemperatureGetProbeReadingsHandler = temperature.GetProbeReadingsHandlerFunc(
		func(params temperature.GetProbeReadingsParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.GetTemperatureProbeReadings(&params))
		})
	api.TemperatureGetMonitorsHandler = temperature.GetMonitorsHandlerFunc(
		func(params temperature.GetMonitorsParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.GetTemperatureMonitors(&params))
		})

	hardware.Startup()
	api.ServerShutdown = func() {
		hardware.Shutdown()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return framework.NewPanicHandler(framework.NewSwaggerUIHandler(handler, CmdOptionsValues.StaticDir))
}
