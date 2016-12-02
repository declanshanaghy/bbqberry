package framework

import (
	"net/http"
	"runtime/debug"

	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/go-openapi/errors"
)

// PanicHandler is a concreate class which can recover from panics and return a standardized error response
type PanicHandler struct {
	handler http.Handler
}

// NewPanicHandler creates a new PanicHandler object
func NewPanicHandler(handler http.Handler) *PanicHandler {
	return &PanicHandler{handler: handler}
}

func (p *PanicHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("action=panic url=%v r=%v, error=%v", r.RequestURI, r, err)
			debug.PrintStack()
			errors.ServeError(rw, nil, err.(error))
		}
	}()

	p.handler.ServeHTTP(rw, r)
}
