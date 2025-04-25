import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import laravel from "laravel-vite-plugin";
import { sveltePreprocess } from "svelte-preprocess";

export default defineConfig({
  plugins: [
    laravel({
      input: ["src/app.ts", "src/app.css"],
      ssr: "src/ssr.ts", // Enable SSR
      publicDirectory: "public",
      buildDirectory: "bootstrap",
      refresh: true,
    }),
    svelte({ preprocess: sveltePreprocess({ typescript: true }) }),
  ],
  build: {
    ssr: true, // Enable SSR
    outDir: "bootstrap",
    rollupOptions: {
      input: "src/ssr.ts",
      output: {
        entryFileNames: "assets/[name].js",
        chunkFileNames: "assets/[name].js",
        assetFileNames: "assets/[name][extname]",
        manualChunks: undefined, // Disable automatic chunk splitting
      },
    },
  },
});
