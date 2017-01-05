package handlers

import (
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/Polarishq/middleware/framework/log"
)

// LoggingHandler provides a middleware handler which logs all requests and responses
type LoggingHandler struct {
	handler         http.Handler
	logResponseBody bool
}

// NewLoggingHandler creates a middleware handler which logs all requests and responses
func NewLoggingHandler(handler http.Handler) *LoggingHandler {
	return &LoggingHandler{handler: handler, logResponseBody: false}
}

// NewLoggingHandlerWithResponseBody creates a middleware handler which logs all requests and responses,
// with an option to enable logging of the response body
func NewLoggingHandlerWithResponseBody(handler http.Handler, captureBody bool) *LoggingHandler {
	return &LoggingHandler{handler: handler, logResponseBody: captureBody}
}

type loggingResponseWriter struct {
	headers     http.Header
	w           http.ResponseWriter
	data        []byte
	code        int
	captureBody bool
}

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lw.captureBody {
		lw.data = append(lw.data, b...)
	}
	return lw.w.Write(b)
}

func (lw *loggingResponseWriter) WriteHeader(code int) {
	lw.headers = lw.Header()
	lw.code = code
	lw.w.WriteHeader(code)
}

func (lw *loggingResponseWriter) Header() http.Header {
	return lw.w.Header()
}

func stringify(r *http.Request) string {
	dump, _ := httputil.DumpRequest(r, true)
	return strings.Replace(string(dump), "\n", " ", -1)
}

func (l *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	log.Debugf("LoggingHandler: request: %+v string_request %+v", r, stringify(r))
	lwr := loggingResponseWriter{w: w, captureBody: l.logResponseBody}
	l.handler.ServeHTTP(&lwr, r)
	endTime := time.Now()
	// if the code is 0, it means that an outter handler will write the code
	log.Debugf("LoggingHandler: response code=%d header='%v' string_body='%+v' time=%v",
		lwr.code, lwr.headers, string(lwr.data), endTime.Sub(startTime))
}
