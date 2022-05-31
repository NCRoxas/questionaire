import { defineConfig } from "vite";
import { VitePWA } from "vite-plugin-pwa";
import pluginRewriteAll from "vite-plugin-rewrite-all";
import vue from "@vitejs/plugin-vue";

// process.env.VITE_BACKEND = "http://localhost:5000/api/v1";

// https://vitejs.dev/config/
export default defineConfig({
	resolve: { alias: { "@": "/src" } },
	server: { port: 8000 },
	preview: { port: 8000 },
	build: { outDir: "../dist" },
	plugins: [
		vue(),
		pluginRewriteAll(),
		VitePWA({
			registerType: "autoUpdate",
			includeAssets: ["images/**/*"],
			manifest: {
				name: "Questionaire",
				short_name: "Questionaire",
				description: "Online learning platform",
				theme_color: "#ffffff",
				icons: [
					{
						src: "/images/icons/manifest-icon-192.maskable.png",
						sizes: "192x192",
						type: "image/png",
						purpose: "any"
					},
					{
						src: "/images/icons/manifest-icon-192.maskable.png",
						sizes: "192x192",
						type: "image/png",
						purpose: "maskable"
					},
					{
						src: "/images/icons/manifest-icon-512.maskable.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "any"
					},
					{
						src: "/images/icons/manifest-icon-512.maskable.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "maskable"
					}
				]
			}
		})
	]
});
