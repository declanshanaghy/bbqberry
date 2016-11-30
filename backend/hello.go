package backend

import "github.com/declanshanaghy/bbqberry/models"

// Hello World ...
func Hello() (models.Hello, error) {
	h := models.Hello{}
	message := "A dinner fit for a king"
	h.Message = &message

	return h, nil
}