import { createInertiaApp } from "@inertiajs/svelte";
import "./app.css";

createInertiaApp({
  // @ts-expect-error
  resolve: (name) => {
    const pages = import.meta.glob("./pages/**/*.svelte", { eager: true });
    return pages[`./pages/${name}.svelte`];
  },
  setup({ el, App, props }) {
    // @ts-expect-error
    new App({ target: el, props });
  },
});
