import { error } from "@sveltejs/kit";
/** @type {import('./$types').PageLoad} */

export const ssr = false;

export async function load(event) {
  return {
    cameraName: event.params.slug,
  };
}
