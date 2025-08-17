import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "node:path";

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 5173,
    strictPort: true,
    proxy: {
      // Dev-only: forward API calls to Go backend
      "/api": "http://localhost:8080",
    },
  },
  build: {
    // IMPORTANT: output the built app into the Go server's static dir
    outDir: resolve(__dirname, "../server/public"),
    emptyOutDir: true,
  },
});
