package log

import (
	"os"

	"github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"path"
	"runtime"
	"strings"
)

type contextHook struct{}

func (hook contextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook contextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 5, 5)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 2)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") &&
			!strings.Contains(name, "framework/log") {
			file, line := fu.FileLine(pc[i] - 2)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}

func init() {
	logrus.SetFormatter(&prefixed.TextFormatter{TimestampFormat: "Jan 02 03:04:05.000"})
	logrus.AddHook(contextHook{})
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, v ...interface{}) {
	logrus.Warningf(format, v...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

// Error logs a message at level Error on the standard logger.
func Error(v ...interface{}) {
	logrus.Error(v...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(v ...interface{}) {
	logrus.Warning(v...)
}

// Info logs a message at level Info on the standard logger.
func Info(v ...interface{}) {
	logrus.Info(v...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(v ...interface{}) {
	logrus.Debug(v...)
}

// there is no fatal on purpose - log and panic instead

// EnableJSONOutput enables JSON output on the logger
func EnableJSONOutput() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(contextHook{})
}

// EnableTextOutput enables plain text output on the logger
func EnableTextOutput() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "Jan 02 03:04:05.000"})
}

// SetOutput sets the standard logger output.
func SetOutput(name string) {
	out, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.SetOutput(os.Stderr)
	}
	logrus.SetOutput(out)
}

// SetDebug sets the log level to debug
func SetDebug(on bool) {
	if on {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.AddHook(contextHook{})
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// SetWarn sets the log level to warn
func SetWarn() {
	logrus.SetLevel(logrus.WarnLevel)

}
