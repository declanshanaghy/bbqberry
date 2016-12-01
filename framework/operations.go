package framework

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/Polarishq/middleware/framework/log"
	"runtime/debug"
)

type Closer interface {
	Close()
}

type IError interface {
	Code() int32
	Error() string
}

type apiOperation struct {
	Response interface{}
}

func (a *apiOperation) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	var err error
	switch t := a.Response.(type) {
	case error:
		log.Error(t)
		errors.ServeError(rw, nil, t)
	default:
		// success sent a 2xx response
		err = producer.Produce(rw, a.Response)
	}

	if err != nil {
		log.Errorf("failed to send response to client error=%s", err.Error())
	}
}

func HandleApiRequestWithError(response interface{}, e error) middleware.Responder {
	op := apiOperation{}
	if e != nil {
		op.Response = e
	} else {
		op.Response = response
	}

	return &op
}

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