import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		cors:false,
		proxy: {
			'/api': {
			  target: 'http://127.0.0.1:1985', // Your backend API server
			  changeOrigin: true,
			  rewrite: (path) => path.replace(/^\/api/, '/api') // Rewrite the URL path
			}
		  },
	},
	plugins: [sveltekit()]
});
