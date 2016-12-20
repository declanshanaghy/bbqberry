package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/Polarishq/middleware/framework"
	"github.com/Polarishq/middleware/framework/log"
	"github.com/Polarishq/middleware/handlers"
	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/daemon"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations"
	"github.com/declanshanaghy/bbqberry/restapi/operations/health"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	"github.com/go-openapi/swag"
	// Unsure why this is suppressed
	_ "github.com/docker/go-units"
	// Unsure why this is suppressed
	_ "github.com/tylerb/graceful"
)

var commander	*daemon.Commander

func init() {
	commander = daemon.NewCommander()
}

type cmdOptions struct {
	LogFile   string `short:"l" long:"logfile" description:"Specify the log file" default:""`
	Verbose   bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	StaticDir string `short:"s" long:"static" description:"The path to the directory containing static resources" default:""`
}

var cmdOptionsValues cmdOptions

func configureFlags(api *operations.AppAPI) {
	log.Info("action=start")
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "BBQ Berry Server Flags",
			LongDescription:  "BBQ Berry Server Flags",
			Options:          &cmdOptionsValues,
		},
	}
}

func configureAPI(api *operations.AppAPI) http.Handler {
	log.Info("action=start")
	// configure the api here
	api.ServeError = errors.ServeError

	log.SetDebug(cmdOptionsValues.Verbose)
	if cmdOptionsValues.LogFile != "" {
		log.SetOutput(cmdOptionsValues.LogFile)
	}
	api.Logger = log.Infof

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.HealthHealthHandler = health.HealthHandlerFunc(
		func(params health.HealthParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.Health())
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
		globalShutdown()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func globalShutdown() {
	log.Info("action=start")

	if err := commander.Exit(); err != nil {
		log.Error(err.Error())
	}
		
	hardware.Shutdown()

	log.Info("action=done")
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
	log.Infof("action=start scheme=%s", scheme)

	if scheme == "http" {
		go commander.Run()
	}

	log.Infof("action=done scheme=%s", scheme)
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	log.Info("action=start")
	defer log.Info("action=done")
	return handlers.NewPanicHandler(
		handlers.NewLoggingHandler(
			handlers.NewSwaggerUIHandler(cmdOptionsValues.StaticDir, handler)))
}
