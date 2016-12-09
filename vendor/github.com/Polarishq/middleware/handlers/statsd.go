package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/peterbourgon/g2s"
)

// Stats data structure
type Stats struct {
	handler http.Handler
	statsd  *g2s.Statsd
	prefix  string
}

// NewStatsdHandler a new HTTP handler which publishes statsd metrics
func NewStatsdHandler(prefix string, server string, port int, handler http.Handler) http.Handler {
	if port <= 0 {
		port = 8125
	}

	dest := fmt.Sprintf("%s:%d", server, port)
	log.Infof("sending statsd to %s", dest)

	s, err := g2s.Dial("udp", dest)
	if err != nil {
		panic(fmt.Sprintf("Unable to initialize statsd middleware: %s", err))
	}

	stats := &Stats{
		statsd: s,
		prefix: prefix,
	}

	return stats.Handler(handler)
}

// Handler is a MiddlewareFunc makes Stats implement the Middleware interface.
func (mw *Stats) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		beginning, recorder := mw.Begin(w)

		h.ServeHTTP(recorder, r)

		mw.End(beginning, recorder)
	})
}

// Begin starts a recorder
func (mw *Stats) Begin(w http.ResponseWriter) (time.Time, RecorderResponseWriter) {
	start := time.Now()

	writer := NewRecorderResponseWriter(w, 200)

	return start, writer
}

// EndWithStatus closes the recorder with a specific status
func (mw *Stats) EndWithStatus(start time.Time, status int) {
	end := time.Now()

	responseTime := end.Sub(start)
	statusCode := fmt.Sprintf("%d", status)

	mw.statsd.Counter(1.0, mw.prefix+"status_code."+statusCode, 1)
	mw.statsd.Timing(1.0, mw.prefix+"response_time", responseTime)
}

// End closes the recorder with the recorder status
func (mw *Stats) End(start time.Time, recorder RecorderResponseWriter) {
	mw.EndWithStatus(start, recorder.Status())
}
