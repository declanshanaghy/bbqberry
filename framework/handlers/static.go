package handlers

import (
	"net/http"
	"strings"
	"os"
	"github.com/Polarishq/middleware/framework/log"
	"fmt"
	"errors"
)

// staticPath is the static URL prefix from which all static assets are served
const staticPath = "/static"

// StaticHandler provides a middleware handler which serves the swagger UI
type StaticHandler struct {
	handler 	http.Handler
	fileServer  http.Handler
}

// NewStaticHandler creates a new middleware handler which serves the swagger UI
func NewStaticHandler(dir string, handler http.Handler) *StaticHandler {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		panic(errors.New(fmt.Sprintf(dir+" does not exist")))
	}


	return &StaticHandler{
		handler: handler,
		fileServer: http.FileServer(http.Dir(dir)),
	}
}

func (s *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idx := strings.Index(r.URL.Path, staticPath)
	if idx == 0 {
		http.StripPrefix(staticPath, s.fileServer).ServeHTTP(w, r)
		log.Debug("request served by StaticHandler")
		return
	}

	s.handler.ServeHTTP(w, r)
}

