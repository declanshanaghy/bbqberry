package framework

import "path"

type CmdOptions struct {
	LogFile   					string	`short:"f" long:"logfile" description:"Specify the log file" default:""`
	Verbose  	 				bool	`short:"v" long:"verbose" description:"Show verbose debug information"`
	ConfDir 					string	`short:"c" long:"conf" description:"The path to the directory containing config resources" default:"./etc"`
	StaticDir 					string	`short:"s" long:"static" description:"The path to the directory containing static resources" default:"./static"`
	TemperatureLoggerEnabled 	bool	`short:"L" long:"logger" description:"Run the temperature logger"`
	LightShow	 				string	`short:"S" long:"show" description:"Run the given lightshow" default:"Temperature" choice:"Rainbow" choice:"Simple Shifter" choice:"Temperature"`
}

func NewCmdOptions() *CmdOptions {
	return &CmdOptions{
		LightShow: "Temperature",
	}
}

func (o* CmdOptions) GetProbesConf() string {
	return path.Join(o.ConfDir, "/probes.json")
}
