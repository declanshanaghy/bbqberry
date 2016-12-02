package framework

import (
	"net/http"
	"runtime/debug"

	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/go-openapi/errors"
)

type PanicHandler struct {
	handler http.Handler
}

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
