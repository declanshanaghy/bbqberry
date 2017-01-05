package framework

import (
	"fmt"
	"net/http"
)

type jwtContextKey string

// JWTKey is used to store and retrieve the token from the context
var JWTKey jwtContextKey

func init() {
	JWTKey = "polarisjwt"
}

// PolarisJWT is used by all Polaris services to infer auth information
type PolarisJWT struct {
	TenantID string
	UserID   string
	Scope    string
}

// GetPolarisJWT retrieves the JWT from the http request
func GetPolarisJWT(httpRequest *http.Request) (*PolarisJWT, error) {
	ctx := httpRequest.Context()
	jwt := ctx.Value(JWTKey).(PolarisJWT)
	if &jwt != nil {
		return &jwt, nil
	}
	return nil, fmt.Errorf("Error retrieving JWT from HTTP request")
}
