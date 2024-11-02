package keycloak

import (
	"fmt"
	"net/http"

	"github.com/AlexxIT/go2rtc/internal/app"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog"
)

var (
	jwksURL = "http://localhost:8080/realms/go2rtc/protocol/openid-connect/certs"
)

type KeycloakServer struct {
	jwtKeyFunc jwt.Keyfunc
	basePath   string
}

var log zerolog.Logger

func Init() {
	var cfg struct {
		Mod struct {
			BasePath string `yaml:"base_path"`
		} `yaml:"api"`
	}

	// load config from YAML
	app.LoadConfig(&cfg)

	log = app.GetLogger("keycloak")

	manager := &JWKSManager{
		JWTJWKS:       jwksURL,
		jwtHTTPClient: &http.Client{},
	}

	// Fetch the key function from the JWKS endpoint
	keyFunc, err := manager.GetKeyFunc()
	if err != nil {
		fmt.Printf("Failed to get JWKS key function: %v\n", err)
		return

	}

	KeycloakApp = KeycloakServer{

		jwtKeyFunc: keyFunc,
		basePath:   cfg.Mod.BasePath,
	}

	log.Info().Str("url", jwksURL).Msg("keycloak init")

}

var KeycloakApp KeycloakServer
