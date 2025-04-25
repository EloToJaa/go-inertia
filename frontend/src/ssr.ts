import { createInertiaApp } from "@inertiajs/svelte";
import createServer from "@inertiajs/svelte/server";
import { mount } from "svelte";

createServer((page) =>
  createInertiaApp({
    page,
    // @ts-expect-error
    resolve: (name) => {
      const pages = import.meta.glob("./pages/**/*.svelte", {
        eager: true,
      });
      return pages[`./pages/${name}.svelte`];
    },
    setup({ el, App, props }) {
      // Svelte 4: return App.render(props)
      // Svelte 5:
      if (el === null) {
        throw new Error("Could not find element with id 'app'");
      }
      mount(App, { target: el, props });
    },
  }),
);
