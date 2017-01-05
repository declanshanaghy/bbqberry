package handlers

import (
	"fmt"
	"net/http"
	"time"

	"os"
	"strconv"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/ooyala/go-dogstatsd"
)

// Stats data structure
type Stats struct {
	client *dogstatsd.Client
}

// NewStatsdHandler creates a new HTTP handler which publishes statsd metrics
// To override the statsd server set the STATSD_HOST environment variable
// To override the statsd port set the STATSD_PORT environment variable
func NewStatsdHandler(prefix string, handler http.Handler) http.Handler {
	stats := newStatsd(prefix)
	return stats.Handler(handler)
}

func newStatsd(prefix string) *Stats {
	var client *dogstatsd.Client
	var err error
	var port int
	var server string

	if os.Getenv("DISABLE_METRICS") == "" {
		if os.Getenv("STATSD_PORT") != "" {
			port, err = strconv.Atoi(os.Getenv("STATSD_PORT"))
			if err != nil {
				panic(err)
			}
		} else {
			port = 8125
		}
		if os.Getenv("STATSD_HOST") != "" {
			server = os.Getenv("STATSD_HOST")
		} else {
			server = "statsd"
		}

		dest := fmt.Sprintf("%s:%d", server, port)
		log.Info("dest=", dest)

		// Create a new dog statsd client
		client, err = dogstatsd.New(dest)
		if err != nil {
			panic(err)
		}

		// Prefix every metric with the app name
		client.Namespace = prefix + "."

		// TODO: Setup additional container info in tags
		// c.Tags = append(c.Tags, "us-east-1a")
	}

	return &Stats{
		client: client,
	}
}

// Handler is a MiddlewareFunc makes Stats implement the Middleware interface.
func (mw *Stats) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		beginning, recorder := mw.Begin(w)
		h.ServeHTTP(recorder, req)
		mw.End(beginning, recorder, req)
	})
}

// Begin starts a recorder
func (mw *Stats) Begin(w http.ResponseWriter) (time.Time, RecorderResponseWriter) {
	start := time.Now()
	writer := NewRecorderResponseWriter(w, 0)
	return start, writer
}

// EndWithStatus closes the recorder with a specific status
func (mw *Stats) EndWithStatus(start time.Time, status int, req *http.Request) {
	end := time.Now()
	responseTime := end.Sub(start)

	if mw.client != nil {
		mw.recordStatus(status)
		mw.recordResponseTime(responseTime, req)
	}
}

func (mw *Stats) recordStatus(status int) {
	// Log http.status metric tagged with the code that is being returned
	tags := []string{
		fmt.Sprintf("code:%d", status),
	}
	metric := "http.status"
	if err := mw.client.Count(metric, 1, tags, 1.0); err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("metric=%s, value=%d, tags=%v", metric, status, tags)
	}
}

func (mw *Stats) recordResponseTime(responseTime time.Duration, req *http.Request) {
	// Log http.response_time metric tagged with the URL that was requested
	metric := "http.response_time_ms"
	value := responseTime.Seconds() * 1000
	tags := extractRequestTags(req)
	if err := mw.client.Histogram(metric, value, tags, 1.0); err != nil {
		log.Error(err.Error())
	} else {
		log.Infof("metric=%s, value=%f, tags=%v", metric, value, tags)
	}
}

func extractRequestTags(req *http.Request) []string {
	tags := make([]string, 1, 2)
	tags[0] = fmt.Sprintf("path:%s", req.URL.Path)

	if req.URL.RawQuery != "" {
		tags = append(tags, fmt.Sprintf("query:%s", req.URL.RawQuery))
	}
	return tags
}

// End closes the recorder with the recorder status
func (mw *Stats) End(start time.Time, recorder RecorderResponseWriter, req *http.Request) {
	mw.EndWithStatus(start, recorder.Status(), req)
}
