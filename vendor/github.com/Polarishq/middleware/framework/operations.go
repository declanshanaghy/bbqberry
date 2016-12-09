package framework

import (
	"net/http"

	"github.com/Polarishq/middleware/framework/log"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// IError provides an error reporting interface
type IError interface {
	Code() int32
	Error() string
}

type apiOperation struct {
	Response interface{}
}

// WriteResponse writes the response to the given producer
func (a *apiOperation) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	var err error
	switch t := a.Response.(type) {
	case error:
		errors.ServeError(rw, nil, t)
	default:
		// success sent a 2xx response
		err = producer.Produce(rw, a.Response)
	}

	if err != nil {
		log.Errorf("failed to send response to client error=%s", err.Error())
	}
}

// HandleAPIRequestWithError checks if an error occurred and if so returns a standardized error message
func HandleAPIRequestWithError(response interface{}, e error) middleware.Responder {
	op := apiOperation{}
	if e != nil {
		op.Response = e
	} else {
		op.Response = response
	}

	return &op
}
