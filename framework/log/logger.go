package log

import (
	"os"

	logrus "github.com/Sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"path"
	"runtime"
	"strings"
)

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
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
	logrus.AddHook(ContextHook{})
}

func Debugf(format string, v ...interface{}) {
	logrus.Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	logrus.Infof(format, v...)
}

func Warningf(format string, v ...interface{}) {
	logrus.Warningf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func Error(v ...interface{}) {
	logrus.Error(v...)
}

func Warning(v ...interface{}) {
	logrus.Warning(v...)
}

func Info(v ...interface{}) {
	logrus.Info(v...)
}

func Debug(v ...interface{}) {
	logrus.Debug(v...)
}

// there is no fatal on purpose - log and panic instead

func EnableJsonOutput() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.AddHook(ContextHook{})
}

func EnableTextOutput() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "Jan 02 03:04:05.000"})
}

func SetOutput(name string) {
	out, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.SetOutput(os.Stderr)
	}
	logrus.SetOutput(out)
}

func SetDebug(on bool) {
	if on {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.AddHook(ContextHook{})
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func SetWarn() {
	logrus.SetLevel(logrus.WarnLevel)

}
