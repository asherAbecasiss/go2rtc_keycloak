package keycloak

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func (s *KeycloakServer) MiddlewareAuthKeycloak(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Info().Msgf(" %s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		if r.URL.Path == s.basePath+"/api/ws" {

			token := r.URL.Query().Get("token")

			err := KeycloakApp.VerifyJWT(token)
			if err != nil {

				log.Error().Caller().Msgf("The request requires valid record authorization token to be set.")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == s.basePath+"/api/streams.dot" {
			token := r.Header.Get("Authorization")

			err := s.VerifyJWT(token)
			if err != nil {
				fmt.Println("-->", err)
				log.Error().Caller().Msgf("The request requires valid record authorization token to be set. %s", r.URL.Path)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return

			}

			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		
		err := s.VerifyJWT(token)
		if err != nil {
			fmt.Println("-->", err)
			log.Error().Caller().Msgf("The request requires valid record authorization token to be set. %s", r.URL.Path)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return

		}

		next.ServeHTTP(w, r)

	})
}

func (s *KeycloakServer) VerifyJWT(tokenString string) error {
	// Parse the JWT token and verify using the JWKS
	token, err := jwt.Parse(tokenString, s.jwtKeyFunc)

	
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	// fmt.Println("Token is valid")
	return nil
}

// VerifyJWT verifies the JWT and checks for the realm-management roles
func (s *KeycloakServer) VerifyJWTAdmin(tokenString string) error {
	// Parse the JWT token and verify using the JWKS
	token, err := jwt.Parse(tokenString, s.jwtKeyFunc)

	if err != nil {
		return fmt.Errorf("invalid token: %v", err)
	}

	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("could not extract claims")
	}

	// Check if the resource_access claim exists
	resourceAccess, ok := claims["resource_access"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("resource_access claim not found in token")
	}

	// Check if realm-management exists
	realmManagement, ok := resourceAccess["realm-management"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("realm-management not found in resource_access")
	}

	// Check if roles exist within realm-management
	roles, ok := realmManagement["roles"].([]interface{})
	if !ok {
		return fmt.Errorf("roles not found in realm-management")
	}

	// Convert the roles to a slice of strings
	var roleNames []string
	for _, role := range roles {
		if roleStr, ok := role.(string); ok {
			roleNames = append(roleNames, roleStr)
		}
	}

	// Check for specific roles, e.g., manage-users
	// hasManageUsersRole := contains(roleNames, "manage-users")
	// if hasManageUsersRole {
	// 	fmt.Println("User has manage-users role")
	// } else {
	// 	return fmt.Errorf("user does not have manage-users role")
	//}

	// fmt.Println("Token is valid")
	return nil
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
