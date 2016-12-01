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
	if dir == "" {
		// swagger ui path not set up let's try to discover it
		gopath := os.Getenv("GOPATH")
		dir = gopath + "/src" + "/splunk/avanti-container" + AuxPath
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
		
		log.Infof(s.dir)
		http.StripPrefix("", http.FileServer(http.Dir(s.dir))).ServeHTTP(w, r)
		log.Infof("request served by the swagger-ui handler")
		return
	}
	
	s.handler.ServeHTTP(w, r)
}

