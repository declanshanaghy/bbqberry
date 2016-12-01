package backend

import (
	"github.com/declanshanaghy/bbqberry/models"
	"github.com/declanshanaghy/bbqberry/influx/example"
	"github.com/declanshanaghy/bbqberry/influx"
	"fmt"
)

// Hello World ...
func Hello() (h models.Hello, err error) {
	client := influx.GetClient()
	pt, err := influx_example.WriteExamplePoint(client)

	if err == nil {
		h = models.Hello{}
		message := fmt.Sprintf("Wrote %s:%v at %s", pt.Name(), pt.Fields(), pt.Time())
		h.Message = &message
	}

	return h, err
}