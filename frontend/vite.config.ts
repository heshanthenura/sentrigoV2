import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import tailwindcss from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  base: "/static/",
  plugins: [react(), tailwindcss() as any],
  server: {
    proxy: {
      "/api": "http://192.168.1.101:8080",
    },
  },
});
