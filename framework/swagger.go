package framework

import (
	"net/http"
	"os"
	"strings"
)

const SwaggerUIPath = "/swagger-ui"
const AuxPath = "/static/"

type SwaggerUIHandler struct {
	handler http.Handler
	dir     string
}

func NewSwaggerUIHandler(handler http.Handler, dir string) *SwaggerUIHandler {
	// swagger ui path not set up let's try to discover it
	if dir == "" {
		// swagger ui path not set up let's try to discover it
		gopath := os.Getenv("GOPATH")
		dir = gopath + "/src" + "/github.com/declanshanaghy/bbqberry" + AuxPath
	}
	return &SwaggerUIHandler{handler: handler, dir: dir}
}

const notFoundString = `
swagger-ui is not found.
`

func (s *SwaggerUIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.Index(r.URL.Path, SwaggerUIPath) == 0 {
		if _, err := os.Stat(s.dir); os.IsNotExist(err) {
			nf := notFoundString + " dir=" + s.dir
			http.Error(w, nf, http.StatusNotFound)
		}

		hDir := http.Dir(s.dir)
		srv := http.FileServer(hDir)
		handler := http.StripPrefix("", srv)
		handler.ServeHTTP(w, r)
		return
	}

	s.handler.ServeHTTP(w, r)
}
