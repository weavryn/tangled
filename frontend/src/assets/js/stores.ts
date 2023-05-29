import { writable } from "svelte/store";
export const updater = writable(1);
export const watcherOn = writable(false);