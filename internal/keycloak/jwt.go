package keycloak

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

type JWKSManager struct {
	JWTJWKS        string
	jwtHTTPClient  *http.Client
	jwtKeyFunc     keyfunc.Keyfunc
	jwtLastRefresh time.Time
}

func (m *JWKSManager) GetKeyFunc() (jwt.Keyfunc, error) {
	// Fetch the JWKS if it hasn't been fetched recently
	if m.jwtKeyFunc == nil || time.Since(m.jwtLastRefresh) > time.Hour {
		// Fetch JWKS from the URL
		res, err := m.jwtHTTPClient.Get(m.JWTJWKS)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get JWKS key function")
			os.Exit(0)

		}

		defer res.Body.Close()

		// Decode the response body
		var raw json.RawMessage
		err = json.NewDecoder(res.Body).Decode(&raw)
		if err != nil {
			return nil, err
		}

		// Parse the JWKS from the raw JSON
		tmp, err := keyfunc.NewJWKSetJSON(raw)
		if err != nil {
			return nil, err
		}

		// Store the parsed JWKS and refresh time
		m.jwtKeyFunc = tmp
		m.jwtLastRefresh = time.Now()
	}

	// Return the key function for JWT validation
	return m.jwtKeyFunc.Keyfunc, nil
}
