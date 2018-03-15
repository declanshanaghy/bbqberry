package framework

type CmdOptions struct {
	LogFile   					string	`short:"f" long:"logfile" description:"Specify the log file" default:""`
	Verbose  	 				bool	`short:"v" long:"verbose" description:"Show verbose debug information"`
	StaticDir 					string	`short:"s" long:"static" description:"The path to the directory containing static resources" default:""`
	TemperatureLoggerEnabled 	bool	`short:"L" long:"logger" description:"Run the temperature logger"`
	TemperatureIndicatorEnabled bool	`short:"I" long:"indicator" description:"Run the temperature indicator"`
	LightShowEnabled 			bool	`short:"S" long:"show" description:"Run the lightshow"`
}
