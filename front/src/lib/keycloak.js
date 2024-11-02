
import Keycloak from "keycloak-js";
import * as jwt_decode from 'jwt-decode';
// Initialize Keycloak
export const keycloak = new Keycloak({
	url: "http://localhost:8080/",
        realm: "go2rtc",
        clientId: "go2rtc",
});

// Initialize Keycloak and check login status
export async function initializeKeycloak() {
    try {
        const authenticated = await keycloak.init({
            onLoad: 'check-sso', // or 'login-required'
            pkceMethod: 'S256'   // PKCE recommended for public clients
        });

        if (!authenticated) {
            await keycloak.login();
        }

        return authenticated;
    } catch (error) {
        console.error('Failed to initialize Keycloak:', error);
    }
}

export function getUserInfo() {
    if (keycloak.token) {
        // Decode the JWT token
        const decodedToken = jwt_decode.jwtDecode(keycloak.token);
        console.log(decodedToken.email);

        // Extract the username (usually in the `preferred_username` or `name` field)
        const username = decodedToken.email 
        
        return username;
    }
    return null;
}