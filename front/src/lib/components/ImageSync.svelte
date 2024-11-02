<script>
    import { Spinner } from "flowbite-svelte";
      import { keycloak } from '$lib/keycloak.js'
    export let url,size;
  
    let image = null;
    $: image;
  
    const preload = async (src) => {
  
      const resp = await fetch(src, {
        headers: {
          Authorization: keycloak.token,
        },
      });
      if (!resp.ok) {
        const text = await resp.text();
        console.log(text);
        throw Error(text);
      }
  
      const blob = await resp.blob();
  
      return new Promise(function (resolve) {
        let reader = new FileReader();
        reader.readAsDataURL(blob);
        reader.onload = () => resolve(reader.result);
        reader.onerror = (error) => reject("Error: ", error);
      });
    };
  </script>
  
  {#await preload(url)}
    <Spinner color="purple" size="6" />
  {:then base64}
    <img src={base64} alt="" class="{size} rounded object-cover block" />
    {:catch error}
    <p class="text-[13px] font-medium text-red-600">RTC: {error.message}</p>
    <!-- 	<code>{base64}</code> -->
  {/await}
  