import { writable } from 'svelte/store';
import { browser } from "$app/environment";

// export const token = writable(keycloak.token);

// // Set up token auto-refresh
// setInterval(async () => {
//     if (keycloak.token) {
//         await keycloak.updateToken(30); // Refresh if token expires in 30s
//         token.set(keycloak.token);
//     }
// }, 10000);



export const player_t = writable("webrtc");

const defaultValue = "grid xl:grid-cols-3 md:grid-cols-2";
const initialValue = browser
  ? window.localStorage.getItem("grid") ?? defaultValue
  : defaultValue;

export const gridSize = writable(initialValue);

gridSize.subscribe((value) => {
  if (browser) {
    window.localStorage.setItem("grid", value);
  }
});
