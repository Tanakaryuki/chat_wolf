import react from "@vitejs/plugin-react-swc";
import { visualizer } from "rollup-plugin-visualizer";
import { defineConfig, splitVendorChunkPlugin } from "vite";
import pluginRewriteAll from "vite-plugin-rewrite-all";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		react(),
		tsconfigPaths(),
		pluginRewriteAll(),
		splitVendorChunkPlugin(),
		visualizer(),
	],
	envDir: "..",
	build: {
		assetsInlineLimit: 4096,
		cssCodeSplit: true,
		sourcemap: false,
		minify: "esbuild",
		rollupOptions: {
			output: {
				chunkFileNames: "assets/js/vendor/[hash].js",
				entryFileNames: "assets/js/[name].js",
			},
		},
	},
});
