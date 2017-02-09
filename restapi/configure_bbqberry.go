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
	bbqframework "github.com/declanshanaghy/bbqberry/framework"
	bbqhandlers "github.com/declanshanaghy/bbqberry/framework/handlers"
	"github.com/declanshanaghy/bbqberry/backend"
	"github.com/declanshanaghy/bbqberry/daemon"
	"github.com/declanshanaghy/bbqberry/hardware"
	"github.com/declanshanaghy/bbqberry/restapi/operations"
	"github.com/declanshanaghy/bbqberry/restapi/operations/health"
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperature"
	"github.com/declanshanaghy/bbqberry/restapi/operations/config"
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

func configureFlags(api *operations.BbqberryAPI) {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "BBQ Berry Server Flags",
			LongDescription:  "BBQ Berry Server Flags",
			Options:          &cmdOptionsValues,
		},
	}
}

func configureAPI(api *operations.BbqberryAPI) http.Handler {
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
	api.ConfigGetHardwareConfigHandler = config.GetHardwareConfigHandlerFunc(
		func(params config.GetHardwareConfigParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(bbqframework.Constants.Hardware, nil)
		})
	api.TemperatureGetProbeReadingsHandler = temperature.GetProbeReadingsHandlerFunc(
		func(params temperature.GetProbeReadingsParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.GetTemperatureProbeReadings(&params))
		})
	api.TemperatureGetMonitorsHandler = temperature.GetMonitorsHandlerFunc(
		func(params temperature.GetMonitorsParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.GetTemperatureMonitors(&params))
		})
	
	globalMiddleware := setupGlobalMiddleware(api.Serve(setupMiddlewares))
	
	globalStartup()
	api.ServerShutdown = func() {
		globalShutdown()
	}
	
	return globalMiddleware
}

func globalStartup() {
	log.Info("action=method_entry")
	defer log.Info("action=method_exit")
	
	hardware.Startup()
	
	if ( ! commander.IsRunning() ) {
		if err := commander.StartBackground(); err != nil {
			panic(err)
		}
	}
	
	return
}

func globalShutdown() {
	log.Info("action=method_entry")
	defer log.Info("action=method_exit")
	
	if ( commander.IsRunning() ) {
		if err := commander.StopBackground(); err != nil {
			panic(err)
		}
	}
	
	hardware.Shutdown()
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
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	log.Debug("action=method_entry")
	defer log.Debug("action=method_exit")
	
	return handlers.NewPanicHandler(
		handlers.NewLoggingHandler(
			bbqhandlers.NewStaticHandler(cmdOptionsValues.StaticDir, handler)))
}
