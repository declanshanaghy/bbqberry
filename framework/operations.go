package framework

import (
	"net/http"

	"github.com/declanshanaghy/bbqberry/framework/log"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// Closer provides an interface to receive a callback when the service is shutting down
type Closer interface {
	Close()
}

// IError provides an interface for obtaining error codes and messages
type IError interface {
	Code() int32
	Error() string
}

type apiOperation struct {
	Response interface{}
}

// WriteResponse writes the pending HTTP response to the given producer
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

// HandleAPIRequestWithError evaluates the given response and error object,
// if an error has occurred a standardized HTTP error response is returned in JSON format,
// otherwise the given response is returned.
func HandleAPIRequestWithError(response interface{}, e error) middleware.Responder {
	op := apiOperation{}
	if e != nil {
		op.Response = e
	} else {
		op.Response = response
	}

	return &op
}
