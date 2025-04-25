import { createInertiaApp } from "@inertiajs/svelte";
import { mount } from "svelte";
import "./app.css";

createInertiaApp({
  // @ts-expect-error
  resolve: (name) => {
    const pages = import.meta.glob("./pages/**/*.svelte", { eager: true });
    return pages[`./pages/${name}.svelte`];
  },
  setup({ el, App, props }) {
    if (el === null) {
      throw new Error("Could not find element with id 'app'");
    }
    mount(App, { target: el, props });
  },
});
