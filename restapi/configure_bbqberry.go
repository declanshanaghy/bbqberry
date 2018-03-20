package restapi

import (
	"crypto/tls"
	"net/http"
	
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/tylerb/graceful"
	
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
	"github.com/declanshanaghy/bbqberry/restapi/operations/temperatures"
	"github.com/declanshanaghy/bbqberry/restapi/operations/monitors"
	opshardware "github.com/declanshanaghy/bbqberry/restapi/operations/hardware"
	"github.com/go-openapi/swag"
	// Unsure why this is suppressed
	_ "github.com/docker/go-units"
	// Unsure why this is suppressed
	_ "github.com/tylerb/graceful"
	"os"
	"syscall"
	"os/signal"
	"github.com/declanshanaghy/bbqberry/restapi/operations/lights"
	"sync"
)


var shutdown			sync.Mutex
var commander			*daemon.Commander
var cmdOptionsValues	bbqframework.CmdOptions


func configureFlags(api *operations.BbqberryAPI) {
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

	commander = daemon.NewCommander(&cmdOptionsValues)

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
	api.HardwareGetHardwareHandler = opshardware.GetHardwareHandlerFunc(
		func(params opshardware.GetHardwareParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(bbqframework.Constants.Hardware, nil)
		})
	api.TemperaturesGetTemperaturesHandler = temperatures.GetTemperaturesHandlerFunc(
		func(params temperatures.GetTemperaturesParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(backend.GetTemperatureProbeReadings(&params))
		})
	api.MonitorsCreateMonitorHandler = monitors.CreateMonitorHandlerFunc(
		func(params monitors.CreateMonitorParams) middleware.Responder {
			mgr, err := backend.NewMonitorsManager()
			if err != nil {
				return framework.HandleAPIRequestWithError(nil, err)
			}
			defer mgr.Close()

			return framework.HandleAPIRequestWithError(mgr.CreateMonitor(&params))
		})
	api.MonitorsGetMonitorsHandler = monitors.GetMonitorsHandlerFunc(
		func(params monitors.GetMonitorsParams) middleware.Responder {
			mgr, err := backend.NewMonitorsManager()
			if err != nil {
				return framework.HandleAPIRequestWithError(nil, err)
			}
			defer mgr.Close()

			return framework.HandleAPIRequestWithError(mgr.GetMonitors(&params))
		})
	api.LightsUpdateGrillLightsHandler = lights.UpdateGrillLightsHandlerFunc(
		func(params lights.UpdateGrillLightsParams) middleware.Responder {
			return framework.HandleAPIRequestWithError(commander.UpdateGrillLights(&params))
		})

	globalMiddleware := setupGlobalMiddleware(api.Serve(setupMiddlewares))
	
	globalStartup()
	api.ServerShutdown = func() {
		globalShutdown()
	}
	
	return globalMiddleware
}

func registerSignals() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGKILL,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		log.WithField("signal", s).Info("Received signal")
		globalShutdown()
		os.Exit(42)
	}()
}

func setupHardware() {
	hardware.Startup()

	if err := commander.StartBackground(); err != nil {
		panic(err)
	}
}

func globalStartup() {
	registerSignals()
	setupHardware()
}

func globalShutdown() {
	shutdown.Lock()
	defer shutdown.Unlock()

	if commander.IsRunning() {
		if err := commander.StopBackground(); err != nil {
			panic(err)
		}
	}
	
	hardware.Shutdown()
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	_ = tlsConfig
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
	_ = s
	_ = scheme
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handlers.NewPanicHandler(
		handlers.NewLoggingHandler(
			bbqhandlers.NewStaticHandler(cmdOptionsValues.StaticDir, handler)))
}
