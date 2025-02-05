import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

import inputPlugin from "@macropygia/vite-plugin-glob-input";

export default defineConfig({
	server: {
		port: 5173,
		hmr: {
			protocol: "ws",
			host: "localhost",
			port: 5173,
		},
		proxy: {
			"/api": {
				target: "http://localhost:8080",
				changeOrigin: true,
				rewrite: (path) => path.replace(/^\/api/, ""),
			},
			"/ws": {
				target: "ws://localhost:8080",
				ws: true,
			},
		},
	},
	plugins: [
		solid(),
		inputPlugin({
			patterns: ["./templates/**/*.html"],
		}),
	],
	css: {
		modules: {
			localsConvention: "dashes",
		},
		preprocessorOptions: {
			scss: {
				api: "modern-compiler",
			},
		},
	},
});
