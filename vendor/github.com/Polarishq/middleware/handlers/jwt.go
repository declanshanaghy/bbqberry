package handlers

import (
	"context"
	"net/http"

	"github.com/Polarishq/middleware/framework"
)

// JWTHandler helps with unmarshalling a JWT token into meaningful structs
type JWTHandler struct {
	handler http.Handler
}

// NewJWTHandler is a constructor
func NewJWTHandler(handler http.Handler) *JWTHandler {
	return &JWTHandler{handler: handler}
}

func (j *JWTHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dummyJWT := framework.PolarisJWT{
		TenantID: "00000000-0000-0000-0000-000000000000",
		UserID:   "00000000-0000-0000-0000-000000000000",
		Scope:    "*",
	}
	// TODO: actually pull this information from a JWT token in the request, duh
	newCtx := context.WithValue(r.Context(), framework.JWTKey, dummyJWT)
	r = r.WithContext(newCtx)
	j.handler.ServeHTTP(w, r)
}
