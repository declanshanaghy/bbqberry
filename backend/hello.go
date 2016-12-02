package backend

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/declanshanaghy/bbqberry/influx"
	"github.com/declanshanaghy/bbqberry/influx/example"
	"github.com/declanshanaghy/bbqberry/models"
)

// Hello World ...
func Hello() (m models.Hello, err error) {
	client, err := influx.GetDefaultClient()
	if err != nil {
		return m, err
	}

	tags := map[string]string{"cpu": "total"}
	rand.Seed(time.Now().UnixNano())
	r := 100.0
	i := float64(int(rand.Float64()*100.0)) + rand.Float64()
	r -= i
	s := float64(int(rand.Float64()*100.0)%int(r-1)) + rand.Float64()
	r -= s
	u := r
	fields := map[string]interface{}{
		"idle":   i,
		"system": s,
		"user":   u,
	}

	pt, err := influx_example.WriteExamplePoint(client, "cpu_usage", tags, fields)
	if err == nil {
		m = models.Hello{}
		message := fmt.Sprintf("Wrote %s:%v at %s", pt.Name(), pt.Fields(), pt.Time())
		m.Message = &message
	}

	return m, err
}
