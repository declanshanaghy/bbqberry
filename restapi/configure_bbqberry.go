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
	"time"
)


var runner				*daemon.Runnable
var commander			*daemon.Commander
var cmdOptionsValues	bbqframework.CmdOptions

func init() {
	commander = daemon.NewCommander()
}

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

	commander.Options = &cmdOptionsValues

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
<<<<<<< Updated upstream
	api.LightsEnableShifterHandler = lights.EnableShifterHandlerFunc(
		func(params lights.EnableShifterParams) middleware.Responder {
			p := time.Duration(params.Period) * time.Millisecond
			commander.EnableLightShow(p)
			return framework.HandleAPIRequestWithError(true, nil)
=======
	api.LightsUpdateGrillLightsHandler = lights.UpdateGrillLightsHandlerFunc(
		func(params lights.UpdateGrillLightsParams) middleware.Responder {
			err := commander.UpdateGrillLights(&params)
			if ( err != nil ) {
				return framework.HandleAPIRequestWithError(false, err)
			} else {
				return framework.HandleAPIRequestWithError(false, err)
			}
>>>>>>> Stashed changes
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
		syscall.SIGUSR1,
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
	log.Info("action=method_entry")
	defer log.Info("action=method_exit")

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
